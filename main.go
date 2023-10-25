package main

import (
	"fmt"

	helmwrapper "github.com/ntnguyencse/helm-mod/cmd/helm"
	// "k8s.io/client-go/discovery"
	// "k8s.io/client-go/tools/clientcmd"
)

func main() {
	// kubePath := "/home/ubuntu/ntnguyen-helm/helm/test/kubeconfig/cluster-test"
	kubePath := "/home/ubuntu/helm-mod/helm-mod/cluster-test"
	fmt.Println("kubePath: ", kubePath)
	chartName2 := "prometheus"
	// ciliumPath := "https://github.com/ntnguyencse/helm-mod/raw/main/test/charts/cilium-1.13.0.tar.gz"
	ciliumPath := "https://github.com/prometheus-community/helm-charts/releases/download/kube-prometheus-stack-48.1.1/kube-prometheus-stack-48.1.1.tgz"
	ciliumHelmArgs := []string{"install", chartName2, ciliumPath, "--namespace", "cilium-system", "--kubeconfig", kubePath}
	err := helmwrapper.ApplyHelmWrapper(kubePath, ciliumPath, true, false, ciliumHelmArgs, []string{})
	if err != nil {
		fmt.Println("error: 2", err)
	}
	// config, err := clientcmd.BuildConfigFromFlags("", kubePath)
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// 	// return nil, err
	// }

	// discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	// if err != nil {
	// 	fmt.Println("error in discoveryClient", err)
	// 	// return nil, err
	// }

	// information, err := discoveryClient.ServerVersion()
	// if err != nil {
	// 	fmt.Println("Error while fetching server version information", err)
	// 	// return nil, err
	// }

	// fmt.Println("Version", information)
	// // return information, err
}
