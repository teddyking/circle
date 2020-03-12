package repositories

import (
	"context"
	"fmt"

	cfv1alpha1 "github.com/teddyking/circle/apis/cf/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type CloudFoundry struct {
	KubeClient client.Client
}

func (r *CloudFoundry) CreateOrg(name, guid string) error {
	org := &cfv1alpha1.Org{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: cfv1alpha1.OrgSpec{
			Guid: guid,
		},
	}

	result, err := controllerutil.CreateOrUpdate(context.Background(), r.KubeClient, org, func() error {
		return nil
	})

	fmt.Printf("RESULT: %s", result)

	return err
}
