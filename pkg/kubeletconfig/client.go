package kubeletconfig

import (
	"context"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/openshift/rosa/pkg/ocm"
	"github.com/openshift/rosa/pkg/rosa"
)

type KubeletConfigClient interface {
	CanCreateKubeletConfig(clusterId string) (bool, error)
	GetKubeletConfig(clusterId string) (*cmv1.KubeletConfig, error)
	CreateKubeletConfig(clusterId string, options KubeletConfigOptions) (*cmv1.KubeletConfig, error)
	UpdateKubeletConfig(clusterId string, options KubeletConfigOptions) (*cmv1.KubeletConfig, error)
	DeleteKubeletConfig(clusterId string, options KubeletConfigOptions) (*cmv1.KubeletConfig, error)
}

type KubeletConfigClientImpl struct {
	ocm *ocm.Client
}

func NewKubeletConfigClient(_ context.Context, runtime *rosa.Runtime) KubeletConfigClient {
	return &KubeletConfigClientImpl{ocm: runtime.OCMClient}
}

func (k KubeletConfigClientImpl) CanCreateKubeletConfig(clusterId string) (bool, error) {
	panic("implement me")
}

func (k KubeletConfigClientImpl) GetKubeletConfig(clusterId string) (*cmv1.KubeletConfig, error) {
	panic("implement me")
}

func (k KubeletConfigClientImpl) CreateKubeletConfig(clusterId string, options KubeletConfigOptions) (*cmv1.KubeletConfig, error) {
	panic("implement me")
}

func (k KubeletConfigClientImpl) UpdateKubeletConfig(clusterId string, options KubeletConfigOptions) (*cmv1.KubeletConfig, error) {
	panic("implement me")
}

func (k KubeletConfigClientImpl) DeleteKubeletConfig(clusterId string, options KubeletConfigOptions) (*cmv1.KubeletConfig, error) {
	panic("implement me")
}
