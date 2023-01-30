package service

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"main.go/src/lib"
)

type Pod struct {
	Namespace  string
	Status     string
	Images     string
	NodeName   string
	CreateTime string
	Labels     map[string]string
	Name       string
}

func ListAllPod(g *gin.Context) {
	ns := g.Query("ns")

	pods, err := lib.K8sClient.CoreV1().Pods(ns).List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		g.Error(err)
	}
	ret := make([]*Pod, 0)
	for _, item := range pods.Items {
		ret = append(ret, &Pod{
			Namespace:  item.Namespace,
			Name:       item.Name,
			Status:     string(item.Status.Phase),
			Labels:     item.Labels,
			NodeName:   item.Spec.NodeName,
			Images:     item.Spec.Containers[0].Image,
			CreateTime: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	g.JSON(200, ret)
	return
}
