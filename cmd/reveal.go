package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// revealCmd represents the reveal command
var (
	secret string

	revealCmd = &cobra.Command{
		Use:   "reveal",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples1`,
		RunE:  revealSecretData,
	}
)

func revealSecretData(cmd *cobra.Command, args []string) (err error) {

	client, _ := NewClient()

	n, _ := cmd.Flags().GetString("namespace")
	s, _ := cmd.Flags().GetString("secret")
	secret, err := client.CoreV1().Secrets(n).Get(context.TODO(), s, meta_v1.GetOptions{})

	secretList := secret.Data

	for key, value := range secretList {
		fmt.Println(key, ":", string(value))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(revealCmd)

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	namespace, _, _ := kubeConfig.Namespace()

	revealCmd.Flags().StringP("namespace", "n", namespace, "Namespace")
	revealCmd.Flags().StringP("secret", "s", secret, "Secret name")
}
