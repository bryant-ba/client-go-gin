package service

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	. "main.go/src/lib"
)

type Deployment struct {
	Name                string
	Replicas            int32
	AvailableReplicas   int32
	UnavailableReplicas int32
	Images              string
	Labels              map[string]string
}

func ListDeployment(g *gin.Context) {
	ns := g.Query("ns")

	dps, err := K8sClient.AppsV1().Deployments(ns).List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		g.Error(err)
	}
	ret := make([]*Deployment, 0)
	for _, item := range dps.Items {
		ret = append(ret, &Deployment{
			Name:                item.Name,
			Replicas:            item.Status.Replicas,
			AvailableReplicas:   item.Status.AvailableReplicas,
			UnavailableReplicas: item.Status.UnavailableReplicas,
			Images:              item.Spec.Template.Spec.Containers[0].Image,
			Labels:              item.GetLabels(),
		})
	}
	g.JSON(200, ret)
	return
}

func GetPodByDep(ns string, dep *v1.Deployment) []*Pod {
	ctx := context.Background()
	listopt := metaV1.ListOptions{
		LabelSelector: "",
	}
	list, err := K8sClient.CoreV1().Pods(ns).List(ctx, listopt)
	if err != nil {
		panic(err.Error())
	}
	pods := make([]*Pod, len(list.Items))
	for i, pod := range list.Items {
		pods[i] = &Pod{
			Namespace:  pod.Namespace,
			Name:       pod.Name, //获取 pod名称
			Status:     string(pod.Status.Phase),
			Images:     pod.Spec.Containers[0].Image,
			NodeName:   pod.Spec.NodeName, //所属节点
			Labels:     pod.Labels,
			CreateTime: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}
	return pods
}
