package main

import (
	"fmt"

	helmwrapper "github.com/ntnguyencse/helm/cmd/helm"
)

func main() {
	kubePath := "/home/ubuntu/ntnguyen-helm/helm/test/kubeconfig/cluster-test"
	chartName2 := "chartname21"
	ciliumPath := "https://github.com/prometheus-community/helm-charts/releases/download/kube-prometheus-stack-46.7.0/kube-prometheus-stack-46.7.0.tgz"
	ciliumHelmArgs := []string{"install", chartName2, ciliumPath, "--kubeconfig", kubePath, "--debug", "--v", "5"}
	err := helmwrapper.ApplyHelmWrapper(kubePath, ciliumPath, true, true, ciliumHelmArgs, []string{})
	if err != nil {
		fmt.Println("error: 2", err)
	}
}