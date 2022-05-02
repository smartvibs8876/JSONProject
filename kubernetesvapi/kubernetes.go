package kubernetesvapi

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func InitConnection() *kubernetes.Clientset {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	clientSet := kubernetes.NewForConfigOrDie(config)

	_, err = clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return clientSet
}

func CreateConfigMap(clientSet *kubernetes.Clientset, data map[string]string, name string) {
	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		Data: data,
	}
	err := clientSet.CoreV1().ConfigMaps("default").Delete(context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = clientSet.CoreV1().ConfigMaps("default").Create(context.Background(), cm, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}

func ReadConfigMap(clientSet *kubernetes.Clientset, name string) map[string]string {
	cm, err := clientSet.CoreV1().ConfigMaps("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	JSONFromCM := cm.Data
	return JSONFromCM
}
