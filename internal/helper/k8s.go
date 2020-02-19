package helper

import (
	"fmt"
	"os"

	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func getClient() dynamic.Interface {
	var kubeconfig = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	var master string

	// creates the connection
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return client
}

func CreateProxy(proxy *projectcontour.HTTPProxy) error {
	client := getClient()
	converter := NewUnstructuredConverter()
	convertedProxy, err := converter.Convert(proxy)
	if err != nil {
		panic(err)
	}
	_, err = client.Resource(HTTPProxyGVR).Namespace(Namespace).Create(convertedProxy, metav1.CreateOptions{})
	return err
}

func RemoveAllProxies() error {
	client := getClient()
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	if err := client.Resource(HTTPProxyGVR).Namespace(Namespace).DeleteCollection(deleteOptions, metav1.ListOptions{}); err != nil {
		return err
	}
	return nil
}
