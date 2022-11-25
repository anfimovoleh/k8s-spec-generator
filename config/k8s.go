package config

import (
	"log"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/anfimovoleh/k8s-spec-generator/clients/k8s"
)

func (c *configImpl) K8S(kubeConfig string) k8s.Client {
	if c.k8sClient != nil {
		return c.k8sClient
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatalln("failed to build kubeconfig:", err)
	}

	c.k8sClient, err = k8s.New(config)
	if err != nil {
		log.Fatalln("failed to kubernetes initialize client:", err)
	}

	return c.k8sClient
}
