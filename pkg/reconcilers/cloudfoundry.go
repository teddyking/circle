package reconcilers

import (
	"github.com/cloudfoundry-community/go-cfclient"
	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	"github.com/teddyking/circle/pkg/repositories"
	ctrl "sigs.k8s.io/controller-runtime"
)

//go:generate counterfeiter . RuntimeRepo
//go:generate counterfeiter . CFRepo
//go:generate counterfeiter . CFClient

type RuntimeRepo interface {
	//TODO: change param to types.NamespacedName
	GetCloudFoundry(req ctrl.Request) (*runtimev1alpha1.CloudFoundry, error)
}

type CFRepo interface {
	CreateOrg(name, guid string) error
}

type CFClient interface {
	ListOrgs() ([]cfclient.Org, error)
}

type CFClientCreator func(config *cfclient.Config) (CFClient, error)

type CloudFoundryReconciler struct {
	RuntimeRepo     RuntimeRepo
	CFRepo          CFRepo
	CFClientCreator CFClientCreator
}

func (r *CloudFoundryReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	cf, err := r.RuntimeRepo.GetCloudFoundry(req)
	if err != nil {
		return ctrl.Result{}, repositories.IgnoreNotFound(err)
	}

	cfClient, err := r.CFClientCreator(&cfclient.Config{ApiAddress: cf.Spec.API})
	if err != nil {
		return ctrl.Result{}, err
	}

	orgs, err := cfClient.ListOrgs()
	if err != nil {
		return ctrl.Result{}, err
	}

	for _, org := range orgs {
		if err := r.CFRepo.CreateOrg(org.Name, org.Guid); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}
