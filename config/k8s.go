package config

import (
	"log"

	"github.com/anfimovoleh/k8s-spec-generator/clients/k8s"
	"github.com/caarlos0/env"
)

type K8S struct {
	Host string `env:"K8S_GEN_API_HOST,required"`
	Port int    `env:"K8S_GEN_API_PORT,required"`
}

func (c *configImpl) K8S() k8s.Client {
	if c.k8sClient != nil {
		return c.k8sClient
	}

	k8sConfig := &K8S{}
	if err := env.Parse(k8sConfig); err != nil {
		log.Fatalln("failed to read kubernetes configuration:", err)
	}

	c.k8sClient = k8s.New("")
	return c.k8sClient
}
