package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

type podsOptions struct {
	namespace  string
	output     string
	kubeconfig string
}

// rootCmd is a pointer to &cobra.Command{}; rootCmd is of type *cobra.Command
var rootCmd = &cobra.Command{
	Use:   "devtool",
	Short: "A CLI tool for DevOps tasks",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("devtool v0.0.1")
	},
}

func newPodsCmd() *cobra.Command {
	opts := &podsOptions{}

	cmd := &cobra.Command{
		Use:   "pods",
		Short: "List pods in a namespace",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Listing pods in namespace: %s\n", opts.namespace)
			fmt.Printf("Output format: %s\n", opts.output)
			fmt.Printf("Using kubeconfig: %s\n", opts.kubeconfig)

			loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
			loadingRules.ExplicitPath = opts.kubeconfig

			configOverrides := &clientcmd.ConfigOverrides{}
			kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

			restConfig, err := kubeConfig.ClientConfig()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error loading kubeconfig: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Succesfully loaded kubeconfig, host: %s\n", restConfig.Host)
		},
	}

	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "default", "Kubernetes namespace")
	cmd.Flags().StringVarP(&opts.output, "output", "o", "table", "Output format (table|json|yaml)")
	cmd.Flags().StringVar(&opts.kubeconfig, "kubeconfig", "", "Path to kubeconfig file (default: $KUBECONFIG or ~/.kube/config)")

	return cmd
}

func main() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(newPodsCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
