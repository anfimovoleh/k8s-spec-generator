package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/anfimovoleh/k8s-spec-generator/config"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	var (
		cfg       = config.New()
		k8sClient = cfg.K8S(*kubeconfig)
	)

	fmt.Println(k8sClient.Namespaces())
}
