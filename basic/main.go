package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type podsOptions struct {
	namespace string
	output    string
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
		},
	}

	cmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "default", "Kubernetes namespace")
	cmd.Flags().StringVarP(&opts.output, "output", "o", "table", "Output format (table|json|yaml)")

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
