package main

import (
	"github.com/vmware-tanzu/kubeapps/cmd/apprepository-controller/pkg/apis/apprepository/v1alpha1"
	"github.com/vmware-tanzu/kubeapps/pkg/chart"
	log "k8s.io/klog/v2"
)

const repoName = "foo-repo"

func main() {

	agent := ""
	target := &chart.Details{
		AppRepositoryResourceName: repoName,
		ChartName:                 "nginx",
		ReleaseName:               "foo",
		Version:                   "17.1.0",
		TarballURL:                "https://charts.bitnami.com/bitnami/redis-17.1.0.tgz",
	}
	// httpClient := httpclient.newHTTPClient(repoURL, []chart.Details{target}, agent)
	chUtils := chart.NewChartClient(agent)
	chUtils.Init(&v1alpha1.AppRepository{}, nil, nil)
	c, err := chUtils.GetChart(target, "https://charts.bitnami.com/bitnami")
	if err != nil {
		panic(err)
	}
	log.Infof("%s ", c.Name(), err)

}
