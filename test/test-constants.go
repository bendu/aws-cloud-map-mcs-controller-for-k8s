package test

import (
	"github.com/aws/aws-cloud-map-mcs-controller-for-k8s/pkg/model"
)

const (
	NsName        = "ns-name"
	NsId          = "ns-id"
	SvcName       = "svc-name"
	SvcId         = "svc-id"
	EndptId1      = "endpoint-id-1"
	EndptId2      = "endpoint-id-2"
	EndptIp1      = "192.168.0.1"
	EndptIp2      = "192.168.0.2"
	EndptPort1    = 1
	EndptPortStr1 = "1"
	EndptPort2    = 2
	EndptPortStr2 = "2"
	OpId1         = "operation-id-1"
	OpId2         = "operation-id-2"
	OpStart       = 1
)

func GetTestHttpNamespace() *model.Namespace {
	return &model.Namespace{
		Id:   NsId,
		Name: NsName,
		Type: model.HttpNamespaceType,
	}
}

func GetTestDnsNamespace() *model.Namespace {
	return &model.Namespace{
		Id:   NsId,
		Name: NsName,
		Type: model.DnsPrivateNamespaceType,
	}
}

func GetTestService() *model.Service {
	return &model.Service{
		Namespace: NsName,
		Name:      SvcName,
		Endpoints: []*model.Endpoint{GetTestEndpoint(), GetTestEndpoint2()},
	}
}

func GetTestEndpoint() *model.Endpoint {
	return &model.Endpoint{
		Id:         EndptId1,
		IP:         EndptIp1,
		Port:       EndptPort1,
		Attributes: make(map[string]string, 0),
	}
}

func GetTestEndpoint2() *model.Endpoint {
	return &model.Endpoint{
		Id:         EndptId2,
		IP:         EndptIp2,
		Port:       EndptPort2,
		Attributes: make(map[string]string, 0),
	}
}
