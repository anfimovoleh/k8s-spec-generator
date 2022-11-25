package config

import "github.com/anfimovoleh/k8s-spec-generator/clients/k8s"

type Config interface {
	K8S(kubeConfig string) k8s.Client
}

type configImpl struct {
	k8sClient k8s.Client
}

func New() Config {
	return &configImpl{}
}
