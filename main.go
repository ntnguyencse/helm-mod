package main

import (
	"fmt"

	helmwrapper "github.com/ntnguyencse/helm-mod/cmd/helm"
)

func main() {
	kubePath := "/home/ubuntu/ntnguyen-helm/helm/test/kubeconfig/cluster-test"
	chartName2 := "cilium"
	ciliumPath := "https://github.com/ntnguyencse/helm-mod/raw/main/test/charts/cilium-1.13.0.tar.gz"
	ciliumHelmArgs := []string{"install", chartName2, ciliumPath, "--namespace", "cilium-system", "--kubeconfig", kubePath}
	err := helmwrapper.ApplyHelmWrapper(kubePath, ciliumPath, true, true, ciliumHelmArgs, []string{})
	if err != nil {
		fmt.Println("error: 2", err)
	}
}
