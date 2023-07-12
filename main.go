package main

import (
	"fmt"

	helmwrapper "github.com/ntnguyencse/helm/cmd/helm"
)

func main() {
	kubePath := "/home/ubuntu/config"
	calicoChart := "https://docs.tigera.io/calico/charts"

	helmArgs := []string{"install", calicoChart, "--kubeconfig", kubePath, "--dry-run"}
	err := helmwrapper.ApplyHelmWrapper(kubePath, calicoChart, true, helmArgs, []string{})
	if err != nil {
		fmt.Println("error:", err)
	}

}
