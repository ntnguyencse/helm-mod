package main

import (
	"log"
	"os"

	"github.com/ntnguyencse/helm/pkg/action"
	"github.com/ntnguyencse/helm/pkg/cli"
	"github.com/ntnguyencse/helm/pkg/kube"
	"github.com/spf13/cobra"
)

const LKaaSManagedFieldsManager = "helm-l-kaas-automation"

func ApplyHelmWrapper(kubeconfig string, chartpath string, debugFlag bool, options []string) error {

	cli.NewCustomEnvs(kubeconfig, debugFlag)
	log.SetFlags(log.Lshortfile)
	kube.ManagedFieldsManager = LKaaSManagedFieldsManager
	actionConfig := new(action.Configuration)
	cmd, err := newRootCmd(actionConfig, os.Stdout, os.Args[1:])
	if err != nil {
		warning("%+v", err)
		os.Exit(1)
	}

	// run when each command's execute method is called
	cobra.OnInitialize(func() {
		helmDriver := os.Getenv("HELM_DRIVER")
		if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), helmDriver, debug); err != nil {
			log.Fatal(err)
		}
		if helmDriver == "memory" {
			loadReleasesInMemory(actionConfig)
		}
	})


	if err := cmd.Execute(); err != nil {
		debug("%+v", err)
		switch e := err.(type) {
		case pluginError:
			os.Exit(e.code)
		default:
			os.Exit(1)
		}
	}
	return nil
}
