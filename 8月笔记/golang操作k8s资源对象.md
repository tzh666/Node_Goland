## golang操作k8s资源对象

```go
package main

import (
	"context"
	"fmt"
	"k8s/utils"
	"log"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 1.获取namespace列表  (指定包的版本号,go get k8s.io/client-go@v0.20.4)

// 2.获取deployment列表  (context.TODO(),不知道传什么就传这个就行了)

// 3.创建deployment

// 4.修改deployment

// 5.获取service列表

// 6.创建service

// 7.修改service

// 8.删除deployment

// 9.删除service

func main() {

	// k8s 配置文件
	kubeconfig := "etc/config"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// 1.获取namespace列表
	namespaceClient := clientset.CoreV1().Namespaces()
	namespaceResult, err := namespaceClient.List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	namespaces := []string{}
	for _, namespace := range namespaceResult.Items {
		namespaces = append(namespaces, namespace.Name)
		fmt.Println(namespace.Name)
	}

	// 2.获取deployment列表
	for _, namespace := range namespaces {

		// 获取deployment列表
		deploymentClient := clientset.AppsV1().Deployments(namespace)

		deployResult, err := deploymentClient.List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		} else {
			for _, deployment := range deployResult.Items {
				fmt.Println("deployment 名字: ", deployment.Name)
				fmt.Println()
			}
		}
	}

	// 3.创建deployment
	fmt.Println("创建Deployment")
	deploymentClient := clientset.AppsV1().Deployments("default")
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "test-nginx-dev",
			Labels: map[string]string{
				"source": "cmdb",
				"app":    "nginx",
				"env":    "test",
			},
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: utils.Int32Ptr(2),
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"source": "cmdb",
					"app":    "nginx",
					"env":    "test",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"source": "cmdb",
						"app":    "nginx",
						"env":    "test",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http-nginx",
									ContainerPort: 80,                  // 暴露的端口
									Protocol:      coreV1.ProtocolSCTP, // TCP协议
								},
							},
						},
					},
				},
			},
		},
	}
	// deployment, err = deploymentClient.Create(context.TODO(), deployment, metaV1.CreateOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(deployment.Status)
	// }

	// 4.修改deployment,这里修改得对应上面自定义的deployment资源
	// 4.1、获取deployment
	deployment, err = deploymentClient.Get(context.TODO(), "test-nginx-dev", metaV1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	} else {
		deployment.Spec.Replicas = utils.Int32Ptr(2)
		// update deplyment
		_, err := deploymentClient.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
		if err != nil {
			log.Fatal(err)
		}
	}

	// 5.获取service列表
	for _, namespace := range namespaces {
		serveiceClient := clientset.CoreV1().Services(namespace)
		serveiceResult, err := serveiceClient.List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		} else {
			for _, service := range serveiceResult.Items {
				fmt.Println("service 名字: ", service.Name)
				fmt.Println("service 信息: ", service.Namespace, service.Labels, service.Spec.ClusterIP, service.Spec.ClusterIPs, service.Spec.ExternalIPs)
				fmt.Println()
			}
		}
	}
	// 6.创建service
	serveiceClient := clientset.CoreV1().Services("default")
	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "test-nginx-dev",
			Labels: map[string]string{
				"source": "cmdb",
				"app":    "nginx",
				"env":    "test",
			},
		},
		Spec: coreV1.ServiceSpec{
			Selector: map[string]string{
				"source": "cmdb",
				"app":    "nginx",
				"env":    "test",
			},
			Type: coreV1.ServiceTypeNodePort,
			Ports: []coreV1.ServicePort{
				{
					Name:     "http-nginx",
					Port:     80,
					Protocol: coreV1.ProtocolTCP,
				},
			},
		},
	}

	// 6.1、创建service
	_, err = serveiceClient.Create(context.TODO(), service, metaV1.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// 7.修改service
	service, err = serveiceClient.Get(context.TODO(), "test-nginx-dev", metaV1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	if service.Spec.Type == coreV1.ServiceTypeNodePort {
		service.Spec.Selector = map[string]string{
			"source": "cmdb1",
			"app":    "nginx1",
			"env":    "test1",
		}
	} else {
		service.Spec.Selector = map[string]string{
			"source": "cmdb2",
			"app":    "nginx2",
			"env":    "test2",
		}
	}
	_, err = serveiceClient.Update(context.TODO(), service, metaV1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}
	// 8.删除deployment
	// deploymentClient.Delete(context.TODO(), "test-nginx-dev", metaV1.DeleteOptions{})

	// 9.删除service
	// serveiceClient.Delete(context.TODO(), "test-nginx-dev", metaV1.DeleteOptions{})
}
```

