---
sidebar_label: Agent
hide_title: true
keywords:
  - aperturectl
  - aperturectl_install_agent
---

<!-- markdownlint-disable -->

## aperturectl install agent

Install Aperture Agent

### Synopsis

Use this command to install Aperture Agent on your Kubernetes cluster.
Refer https://artifacthub.io/packages/helm/aperture/aperture-agent#parameters for list of configurable parameters for preparing values file.

```
aperturectl install agent [flags]
```

### Examples

```
aperturectl install agent --values-file=values.yaml

aperturectl install agent --values-file=values.yaml --namespace=aperture
```

### Options

```
  -h, --help   help for agent
```

### Options inherited from parent commands

```
      --dry-run              If set to true, only the manifests will be generated and no installation will be performed
      --kube-config string   Path to the Kubernetes cluster config. Defaults to '~/.kube/config'
      --namespace string     Namespace in which the component will be installed. Defaults to 'default' namespace (default "default")
      --values-file string   Values YAML file containing parameters to customize the installation
      --version string       Version of the Aperture (default "latest")
```

### SEE ALSO

- [aperturectl install](/reference/aperturectl/install/install.md) - Install Aperture components
