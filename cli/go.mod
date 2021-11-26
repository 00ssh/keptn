module github.com/keptn/keptn/cli

go 1.16

require (
	github.com/alecthomas/jsonschema v0.0.0-20211022214203-8b29eab41725
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/cloudevents/sdk-go/v2 v2.6.1
	github.com/docker/docker-credential-helpers v0.6.4
	github.com/go-test/deep v1.0.8
	github.com/google/uuid v1.3.0
	github.com/hashicorp/go-version v1.2.0
	github.com/keptn/go-utils v0.11.0
	github.com/keptn/kubernetes-utils v0.10.1-0.20211102080304-e59377afdc8b
	github.com/mattn/go-shellwords v1.0.12
	github.com/mitchellh/mapstructure v1.4.2
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gotest.tools v2.2.0+incompatible
	helm.sh/helm/v3 v3.7.0
	k8s.io/api v0.22.3
	k8s.io/apimachinery v0.22.3
	k8s.io/cli-runtime v0.22.3
	k8s.io/client-go v0.22.3
	k8s.io/kubectl v0.22.1
)

// required as per https://github.com/helm/helm/issues/9354
replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
