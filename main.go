package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here
	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	namespace, _, _ := kubeConfig.Namespace()

	// Command line arguments
	n := flag.String("namespace", namespace, "namespace")
	s := flag.String("secret", "default-secret", "secret name")
	flag.Parse()

	client, _ := newClient()

	secret, err := client.CoreV1().Secrets(*n).Get(context.TODO(), *s, meta_v1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	secretList := secret.Data

	for key, value := range secretList {
		fmt.Println(key, ":", string(value))
	}

}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}
