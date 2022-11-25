package k8s

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client interface {
	Namespaces() ([]string, error)
}

type clientImpl struct {
	config    *rest.Config
	clientSet *kubernetes.Clientset
}

func New(config *rest.Config) (Client, error) {
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		config:    config,
		clientSet: clientSet,
	}, nil
}

func (c clientImpl) Namespaces() ([]string, error) {
	list, err := c.clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespacesNames []string
	for _, item := range list.Items {
		namespacesNames = append(namespacesNames, item.Name)
	}

	return namespacesNames, nil
}
