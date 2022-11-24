package main

import (
	"fmt"

	"github.com/anfimovoleh/k8s-spec-generator/config"
)

func main() {
	var (
		cfg       = config.New()
		k8sClient = cfg.K8S()
	)

	fmt.Println(k8sClient.Namespaces())
}
