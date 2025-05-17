package cmd

import (
	"github.com/davidemaggi/koncierge/cmd/context"
	"github.com/davidemaggi/koncierge/cmd/forward"
	"github.com/davidemaggi/koncierge/cmd/namespace"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "koncierge",
	Short: "Your faithful assistant to interact with your Kubernetes cluster",
	Long: `Koncierge is here to be your assistant and your guide managing your k8s cluster.
You can manage your Contexts, Namespace, Kubeconfigs and port forwards.

You are lazy, I'm lazy lets get lazy together and let Koncierge do the dirty job`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize service container once
		container.Init(config.IsVerbose)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.koncierge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&config.IsVerbose, "verbose", false, "Log All information, not just the bad things")

	startFile := ""

	if home := homedir.HomeDir(); home != "" {
		startFile = filepath.Join(home, ".kube", "config")
	}

	rootCmd.PersistentFlags().StringVarP(&config.KubeConfigFile, "kubeconfig", "f", startFile, "The Kubeconfig file to use")

	startCtx := k8s.GetCurrentContextAsStringFromConfig(config.KubeConfigFile)
	rootCmd.PersistentFlags().StringVarP(&config.KubeContext, "context", "c", startCtx, "The preferred Context to use")

	rootCmd.AddCommand(forward.FwdCmd)

	rootCmd.AddCommand(namespace.NsCmd)
	rootCmd.AddCommand(ConfigCmd)
	rootCmd.AddCommand(context.CtxCmd)
	rootCmd.AddCommand(InfoCmd)

}
