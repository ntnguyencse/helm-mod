package helm

import (
	"fmt"
	"log"
	"os"

	"github.com/ntnguyencse/helm/pkg/action"
	"github.com/ntnguyencse/helm/pkg/cli"
	"github.com/ntnguyencse/helm/pkg/kube"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	chartUrl       = "https:/github.com...."
	valuesFilePath = "...."
	chartsPath     = "..."
)
const LKaaSManagedFieldsManager = "helm-l-kaas-automation"

func test11() {
	args := []string{
		"--values=valuesPathFile",
		"install",
	}
	_ = args

}

// Simulate install (dry-run)
func SimulateInstall() {

}
func ApplyHelmWrapper(kubeconfig string, chartpath string, debugFlag bool, helmArgs, options []string) error {
	fmt.Println("Begining of ApplyHelmWrapper")
	_, _, outbuff, _ := genericclioptions.NewTestIOStreams()
	// Create new Custom Envs with kubeconfig and debug flag
	cli.NewCustomEnvs(kubeconfig, debugFlag)
	log.SetFlags(log.Lshortfile)
	// Store args of Helm command
	// var helmArgs []string
	kube.ManagedFieldsManager = LKaaSManagedFieldsManager
	actionConfig := new(action.Configuration)
	cmd, err := newRootCmd(actionConfig, outbuff, helmArgs)
	if err != nil {
		warning("%+v", err)
		log.Println(err, "Error: Error when create newRootCmd")
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
			log.Println(err, e, "Plugin Error: Error when perform")
			return err
		default:
			log.Println(err, e, "Error: ")
			return err
		}
	}
	return nil
}
