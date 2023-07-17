package main

import (
	"fmt"

	helmwrapper "github.com/ntnguyencse/helm/cmd/helm"
)

func main() {
	kubePath := "/home/ubuntu/config"
	calicoChartPackagedPath := "https://github.com/jenkinsci/helm-charts/releases/download/jenkins-4.4.1/jenkins-4.4.1.tgz"
	valueFile := "/home/ubuntu/ntnguyen-helm/helm/test/values.yaml"
	chartName := "chartname"
	helmArgs := []string{"install", "-f", valueFile, chartName, calicoChartPackagedPath, "--kubeconfig", kubePath, "--dry-run", "--debug", "--v", "5"}
	err := helmwrapper.ApplyHelmWrapper(kubePath, calicoChartPackagedPath, true, true, helmArgs, []string{})
	if err != nil {
		fmt.Println("error:", err)
	}

}
