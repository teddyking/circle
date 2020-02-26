package reconcilers

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

type CloudFoundry struct{}

func NewCloudFoundryReconciler() *CloudFoundry {
	return &CloudFoundry{}
}

func (r *CloudFoundry) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
