package repositories

import (
	"context"

	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Runtime struct {
	KubeClient client.Client
}

func (r *Runtime) GetCloudFoundry(req ctrl.Request) (*runtimev1alpha1.CloudFoundry, error) {
	cf := &runtimev1alpha1.CloudFoundry{}
	if err := r.KubeClient.Get(context.Background(), types.NamespacedName{Name: req.Name}, cf); err != nil {
		return &runtimev1alpha1.CloudFoundry{}, ErrCloudFoundryRuntimeNotFound{Name: req.Name}
	}

	return cf, nil
}
