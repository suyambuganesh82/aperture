local creator = import '../../grafana/dashboard_group.libsonnet';
local utils = import '../common/utils.libsonnet';
local blueprint = import './elasticsearch.libsonnet';

local policy = blueprint.policy;
local config = blueprint.config;

function(params, metadata={}) {
  // make sure param object contains fields that are in config
  local extra_keys = std.setDiff(std.objectFields(params), std.objectFields(config)),
  assert std.length(extra_keys) == 0 : 'Unknown keys in params: ' + extra_keys,

  local c = std.mergePatch(config, params),
  local metadataWrapper = metadata { values: std.toString(params) },

  local updated_cfg = utils.add_kubelet_overload_confirmations(c).updated_cfg,

  local infraMeters = if std.objectHas(c.policy.resources, 'infra_meters') then c.policy.resources.infra_meters else {},
  assert !std.objectHas(infraMeters, 'elasticsearch') : 'An infra meter with name elasticsearch already exists. Please choose a different name.',
  local config_with_elasticsearch_infra_meter = updated_cfg {
    policy+: {
      resources+: {
        infra_meters+: {
          elasticsearch: {
            agent_group: if std.objectHas(updated_cfg.policy.elasticsearch, 'agent_group') then updated_cfg.policy.elasticsearch.agent_group else 'default',
            per_agent_group: false,
            receivers: {
              elasticsearch: std.prune(updated_cfg.policy.elasticsearch {
                agent_group: null,
                metrics+: {
                  'elasticsearch.node.operations.current': {
                    enabled: true,
                  },
                  'elasticsearch.node.operations.completed': {
                    enabled: true,
                  },
                  'elasticsearch.node.operations.time': {
                    enabled: true,
                  },
                  'elasticsearch.node.operations.get.completed': {
                    enabled: true,
                  },
                  'elasticsearch.node.operations.get.time': {
                    enabled: true,
                  },
                  'jvm.memory.heap.utilization': {
                    enabled: true,
                  },
                },
              }),
            },
          },
        },
      },
    },
  },

  local p = policy(config_with_elasticsearch_infra_meter, metadataWrapper),
  local d = creator(p.policyResource, config_with_elasticsearch_infra_meter),

  policies: {
    [std.format('%s-cr.yaml', config_with_elasticsearch_infra_meter.policy.policy_name)]: p.policyResource,
    [std.format('%s.yaml', config_with_elasticsearch_infra_meter.policy.policy_name)]: p.policyDef { metadata: metadataWrapper },
  },
  dashboards: {
    [std.format('%s.json', config_with_elasticsearch_infra_meter.policy.policy_name)]: d.mainDashboard,
    [std.format('signals-%s.json', config_with_elasticsearch_infra_meter.policy.policy_name)]: d.signalsDashboard,
  },
}
