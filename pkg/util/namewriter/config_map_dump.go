package namewriter

import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func UpdateBucketrData(user, bucketName string) {

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err)
	// }

	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Error(err)
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Error(err)
		panic(err)
	}

	configMapData := make(map[string]string, 0)

	configMapData[user+"-bucket-name"] = bucketName

	configMap :=
		&apiv1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigMap",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bucket-data",
				Namespace: "ibm-cosi-driver",
			},
			Data: configMapData,
		}

	if cm, err := clientset.CoreV1().ConfigMaps("ibm-cosi-driver").Get(context.TODO(), "bucket-data", metav1.GetOptions{}); errors.IsNotFound(err) {

		fmt.Println(cm)
		_, err = clientset.CoreV1().ConfigMaps("ibm-cosi-driver").Create(context.TODO(), configMap, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}

	} else {
		_, err = clientset.CoreV1().ConfigMaps("ibm-cosi-driver").Update(context.TODO(), configMap, metav1.UpdateOptions{})

		if err != nil {
			panic(err)
		}

	}

}
