package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	rootCmd = &cobra.Command{
		Use:   "revs",
		Short: "revs is command that reveal Kubernetes secret resources",
		Long:  `revs is command that reveal Kubernetes secret resources without base64 decode`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("revs called")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func NewClient() (kubernetes.Interface, error) {
	kubeconfig := getEnv("KUBECONFIG", clientcmd.RecommendedHomeFile)
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
