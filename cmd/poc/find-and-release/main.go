package main

import (
	"github.com/vmware-tanzu/kubeapps/cmd/apprepository-controller/pkg/apis/apprepository/v1alpha1"
	"github.com/vmware-tanzu/kubeapps/cmd/kubeapps-apis/plugins/helm/packages/v1alpha1/utils"
	agent "github.com/vmware-tanzu/kubeapps/cmd/kubeapps-apis/plugins/pkg/agent"
	"github.com/vmware-tanzu/kubeapps/pkg/chart"
	"github.com/vmware-tanzu/kubeapps/pkg/chart/models"
	"github.com/vmware-tanzu/kubeapps/pkg/dbutils"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	log "k8s.io/klog/v2"
)

// k port-forward  svc/kubeapps-postgresql 5432:5432 &
func main() {
	chartToInstall := getChart()
	releaseChart("redis1", *chartToInstall, chartToInstall.ChartVersions[0])
}

func getChart() *models.Chart {

	var dbConfig = dbutils.Config{URL: "localhost:5432", Database: "assets", Username: "postgres", Password: "PS7WGeU17M"}

	manager, err := utils.NewPGManager(dbConfig, "kubeapps")
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = manager.Init()
	if err != nil {
		log.Fatalf("%s", err)
	}
	charts, err := manager.GetPaginatedChartListWithFilters(utils.ChartQuery{ChartName: "redis"}, 0, 10)
	if err != nil {
		panic(err)
	}
	log.Infof("size: %d", len(charts))
	return charts[0]
}

func releaseChart(releaseName string, chartToInstall models.Chart, chartVersion models.ChartVersion) {

	target := &chart.Details{
		AppRepositoryResourceName: "dummy",
		ChartName:                 chartToInstall.Name,
		ReleaseName:               releaseName,
		Version:                   chartVersion.Version,
		TarballURL:                chartVersion.URLs[0],
	}

	// httpClient := httpclient.newHTTPClient(repoURL, []chart.Details{target}, agent)
	chUtils := chart.NewChartClient("dummy")
	chUtils.Init(&v1alpha1.AppRepository{}, nil, nil)
	helmChart, err := chUtils.GetChart(target, "https://charts.bitnami.com/bitnami")
	if err != nil {
		panic(err)
	}
	log.Infof("%s ", helmChart.Name())

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

	release, err := agent.CreateRelease(actionConfig, releaseName, "default", "", helmChart, nil, 20)
	if err != nil {
		panic(err)
	}
	log.Infof(release.Name)

}
