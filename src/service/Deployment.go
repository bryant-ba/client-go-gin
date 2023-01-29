package service

import (
	"context"
	"github.com/gin-gonic/gin"
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
			Labels:              item.Labels,
		})
	}
	g.JSON(200, ret)
	return
}
