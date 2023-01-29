package service

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	. "main.go/src/lib"
)

type Namespace struct {
	Name   string
	Status string
	Labels map[string]string
}

func ListNamespace(g *gin.Context) {
	ns, err := K8sClient.CoreV1().Namespaces().List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		g.Error(err)
		return
	}
	ret := make([]*Namespace, 0)
	for _, item := range ns.Items {
		ret = append(ret, &Namespace{
			Name:   item.Name,
			Status: string(item.Status.Phase),
			Labels: item.Labels,
		})
	}
	g.JSON(200, ret)
	return
}
