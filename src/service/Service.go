package service

import (
	"context"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	. "main.go/src/lib"
)

type Service struct {
	Name       string
	Type       string
	ClusterIP  string
	ExternalIp []string
	Ports      []string
	Select     map[string]string
}

func ListService(g *gin.Context) {
	ns := g.Query("ns")
	svc, err := K8sClient.CoreV1().Services(ns).List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		g.Error(err)
		return
	}
	ret := make([]*Service, 0)
	for _, item := range svc.Items {
		ret = append(ret, &Service{
			Name:       "",
			Type:       "",
			ClusterIP:  "",
			ExternalIp: nil,
			Ports:      nil,
			Select:     nil,
		})
	}
	g.JSON(200, svc)
}
