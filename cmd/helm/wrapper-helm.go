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
func ApplyHelmWrapper(kubeconfig string, chartpath string, debugFlag bool, dryRun bool, helmArgs, options []string) error {

	fmt.Println("Begining of ApplyHelmWrapper")

	// valueFilePath := "test"
	_, _, outbuff, _ := genericclioptions.NewTestIOStreams()
	// Create new Custom Envs with kubeconfig and debug flag
	customSettings := cli.NewCustomEnvs(kubeconfig, debugFlag)
	log.SetFlags(log.Lshortfile)

	kube.ManagedFieldsManager = LKaaSManagedFieldsManager
	actionConfig := new(action.Configuration)

	cobra.OnInitialize(func() {
		helmDriver := os.Getenv("HELM_DRIVER")
		if err := actionConfig.Init(customSettings.RESTClientGetter(), customSettings.Namespace(), helmDriver, debug); err != nil {
			log.Fatal(err)
		}
		if helmDriver == "memory" {
			loadReleasesInMemory(actionConfig)
		}
	})
	cmd, err := newRootCmd(actionConfig, outbuff, helmArgs)
	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "path to the kubeconfig file")
	// Create Install flag
	installCmd := newInstallCmd(actionConfig, outbuff)
	// installCmd.Flags().BoolVar(&dryRun, "dry-run", false, "simulate an install")

	// Add install flags
	// actionInstall := new(action.Install)
	// actionInstall.DryRun = true
	// actionInstall.ClientOnly = true
	//Value Options "-f value-files"
	// installOption := values.Options{
	// 	ValueFiles: []string{valueFilePath},
	// }
	// Flags Set of command
	// installFlags := cmd.Flags()
	// installFlags.
	// addInstallFlags(installCmd, installFlags, actionInstall, &installOption)
	// Add Helm install command
	cmd.AddCommand(installCmd)
	fmt.Println("Created root cmd")
	if err != nil {
		warning("%+v", err)
		log.Println(err, "Error: Error when create newRootCmd")
	}

	// run when each command's execute method is called

	cmd.SetArgs(helmArgs)
	// installCmd.SetArgs(helmArgs)
	cmd.DebugFlags()
	installCmd.DebugFlags()
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error when execute")
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
	fmt.Println("Not Error when execute")
	fmt.Println("Output:", string(outbuff.Bytes()))
	return nil
}