/*
Kubernetesvapi package to initialise connection to a minikube cluster
Add a conifg map to the minikube cluster
Read a config map from minikube cluster
Update a config map in minikube cluster
*/
package kubernetesvapi

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//Function to initialise a connection to minikube cluster and return an object for further operations like create,read and update
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

//Function to update a config map
func UpdateConfigMap(clientSet *kubernetes.Clientset, data map[string]string, name string, namespace string) {
	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}
	_, err := clientSet.CoreV1().ConfigMaps(namespace).Update(context.Background(), cm, metav1.UpdateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

}

//Function to create a config map
func CreateConfigMap(clientSet *kubernetes.Clientset, data map[string]string, name string, namespace string) {
	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}
	_, err := clientSet.CoreV1().ConfigMaps(namespace).Create(context.Background(), cm, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}

}

//Function to read a config map
func ReadConfigMap(clientSet *kubernetes.Clientset, name string, namespace string) (string, error) {
	cm, err := clientSet.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	JSONFromCM := cm.Data["configuration"]
	return JSONFromCM, nil
}
