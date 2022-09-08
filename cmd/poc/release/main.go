package main

import (
	"github.com/vmware-tanzu/kubeapps/cmd/apprepository-controller/pkg/apis/apprepository/v1alpha1"
	"github.com/vmware-tanzu/kubeapps/cmd/kubeapps-apis/plugins/pkg/agent"
	"github.com/vmware-tanzu/kubeapps/pkg/chart"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	log "k8s.io/klog/v2"
)

func main() {

	releasename := "redis1"
	target := &chart.Details{
		AppRepositoryResourceName: "test",
		ChartName:                 "nginx",
		ReleaseName:               releasename,
		Version:                   "17.1.0",
		TarballURL:                "https://charts.bitnami.com/bitnami/redis-17.1.0.tgz",
	}
	chUtils := chart.NewChartClient("my-user-agent")
	chUtils.Init(&v1alpha1.AppRepository{
		Spec: v1alpha1.AppRepositorySpec{
			TLSInsecureSkipVerify: true,
		},
	}, nil, nil)
	chart, err := chUtils.GetChart(target, "https://charts.bitnami.com/bitnami")

	if err != nil {
		panic(err)
	}
	var settings = cli.New()

	actionConfig := new(action.Configuration)
	err = actionConfig.Init(
		settings.RESTClientGetter(),
		settings.Namespace(),
		"secret",
		func(format string, v ...interface{}) {
			log.Infof(format, v...)
		},
	)
	if err != nil {
		panic(err)
	}

	release, err := agent.CreateRelease(actionConfig, releasename, "default", "", chart, nil, 20)
	if err != nil {
		panic(err)
	}
	log.Infof(release.Name)

}
