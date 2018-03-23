package main

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("spinnaker").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, element := range pods.Items {
		if strings.Contains(element.ObjectMeta.Name, "spin-clouddriver") {
			fmt.Printf("Found the desired clouddriver pod\n")
			fmt.Printf(element.ObjectMeta.Name)
			err = clientset.CoreV1().Pods("spinnaker").Delete(element.ObjectMeta.Name, &metav1.DeleteOptions{})
			if errors.IsNotFound(err) {
				fmt.Printf("\nPod not found\n")
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				fmt.Printf("\nError deleting pod %v\n", statusError.ErrStatus.Message)
			} else if err != nil {
				panic(err.Error())
			} else {
				fmt.Printf("\nDeleted pod\n")
			}
		}
	}
}
