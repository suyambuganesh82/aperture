package controlplane

import (
	"context"
	"errors"

	"go.uber.org/fx"
	"gopkg.in/yaml.v3"

	policiesv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	policysyncv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/sync/v1"
	"github.com/fluxninja/aperture/v2/pkg/alerts"
	"github.com/fluxninja/aperture/v2/pkg/config"
	"github.com/fluxninja/aperture/v2/pkg/log"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/circuitfactory"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/crwatcher"
	"github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/resources/classifier/compiler"
	"github.com/fluxninja/aperture/v2/pkg/status"
	"github.com/fluxninja/aperture/v2/pkg/webhooks/policyvalidator"
)

// FxOut is the output of the controlplane module.
type FxOut struct {
	fx.Out
	Validator policyvalidator.PolicySpecValidator `group:"policy-validators"`
}

// FxIn is the input for the AddAgentInfoAttribute function.
type FxIn struct {
	fx.In
	Unmarshaller config.Unmarshaller
}

// providePolicyValidator provides classification Policy Custom Resource validator
//
// Note: This validator must be registered to be accessible.
func providePolicyValidator(in FxIn) (FxOut, error) {
	var config crwatcher.CRWatcherConfig
	err := in.Unmarshaller.UnmarshalKey(crwatcher.ConfigKey, &config)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal Kubernetes watcher config")
		return FxOut{}, nil
	}

	if !config.Enabled {
		log.Info().Msg("Kubernetes watcher is disabled")
		return FxOut{}, nil
	}

	return FxOut{
		Validator: &PolicySpecValidator{},
	}, nil
}

// PolicySpecValidator Policy implementation of PolicySpecValidator interface.
type PolicySpecValidator struct{}

// ValidateSpec checks the validity of a Policy spec
//
// returns:
// * true, "", nil when Policy is valid
// * false, message, nil when Policy is invalid
// and
// * false, "", err on other errors.
//
// ValidateSpec checks the syntax, validity of extractors, and validity of
// rego modules (by attempting to compile them).
func (v *PolicySpecValidator) ValidateSpec(
	ctx context.Context,
	name string,
	yamlSrc []byte,
) (bool, string, error) {
	_, _, err := ValidateAndCompile(ctx, name, yamlSrc)
	if err != nil {
		// there is no need to handle validator errors. just return validation result.
		return false, err.Error(), nil
	}
	return true, "", nil
}

// ValidateAndCompile checks the validity of a single Policy and compiles it.
func ValidateAndCompile(ctx context.Context, name string, yamlSrc []byte) (*circuitfactory.Circuit, *policiesv1.Policy, error) {
	if len(yamlSrc) == 0 {
		return nil, nil, errors.New("empty policy")
	}

	var yamlRaw map[string]any
	err := yaml.Unmarshal(yamlSrc, &yamlRaw)
	if err != nil {
		return nil, nil, err
	}
	// TODO: this is a hack to make unmarshal work. We should configure unmarshaller to not fail on unknown fields.
	delete(yamlRaw, "metadata")
	yamlSrc, err = yaml.Marshal(yamlRaw)
	if err != nil {
		return nil, nil, err
	}

	policy := &policiesv1.Policy{}
	err = config.UnmarshalYAML(yamlSrc, policy)
	if err != nil {
		return nil, nil, err
	}

	alerter := alerts.NewSimpleAlerter(100)
	registry := status.NewRegistry(log.GetGlobalLogger(), alerter)
	circuit, err := CompilePolicy(policy, registry)
	if err != nil {
		return nil, nil, err
	}

	if policy.GetResources() != nil {
		classifiers := policy.GetResources().GetFlowControl().GetClassifiers()
		for _, c := range classifiers {
			_, err = compiler.CompileRuleset(ctx, name, &policysyncv1.ClassifierWrapper{
				Classifier: c,
				ClassifierAttributes: &policysyncv1.ClassifierAttributes{
					PolicyName:      "dummy",
					PolicyHash:      "dummy",
					ClassifierIndex: 0,
				},
			})
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return circuit, policy, nil
}
