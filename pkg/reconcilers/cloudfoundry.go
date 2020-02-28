package reconcilers

import (
	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	"github.com/teddyking/circle/pkg/repositories"
	ctrl "sigs.k8s.io/controller-runtime"
)

//go:generate counterfeiter . RuntimeRepo

type RuntimeRepo interface {
	GetCloudFoundry(req ctrl.Request) (*runtimev1alpha1.CloudFoundry, error)
}

type CloudFoundryReconciler struct {
	RuntimeRepo RuntimeRepo
}

func (r *CloudFoundryReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_, err := r.RuntimeRepo.GetCloudFoundry(req)
	if err != nil {
		return ctrl.Result{}, repositories.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}
