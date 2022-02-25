package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show secret list",
	Long:  `Show secret list`,
	RunE:  showSecretList,
}

func showSecretList(cmd *cobra.Command, args []string) (err error) {

	client, _ := NewClient()

	n, _ := cmd.Flags().GetString("namespace")
	secrets, err := client.CoreV1().Secrets(n).List(context.TODO(), meta_v1.ListOptions{})

	if err != nil {
		return
	}

	for _, secret := range secrets.Items {
		fmt.Println(secret.GetName())
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	namespace, _, _ := kubeConfig.Namespace()

	listCmd.Flags().StringP("namespace", "n", namespace, "Namespace")
}
