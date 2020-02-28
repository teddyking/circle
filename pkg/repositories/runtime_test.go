package repositories_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/circle/pkg/repositories"

	"context"

	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("Runtime", func() {
	var (
		fakeClient client.Client
		cfRuntime  *runtimev1alpha1.CloudFoundry
		req        ctrl.Request
		repo       *Runtime
	)

	BeforeEach(func() {
		cfRuntime = &runtimev1alpha1.CloudFoundry{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-cf-1",
			},
			Spec: runtimev1alpha1.CloudFoundrySpec{},
		}

		req = ctrl.Request{NamespacedName: types.NamespacedName{Name: "my-cf-1"}}

		fakeClient = fake.NewFakeClient()
		repo = &Runtime{KubeClient: fakeClient}
	})

	Describe("GetCloudFoundry", func() {
		When("the CloudFoundry exists", func() {
			BeforeEach(func() {
				Expect(fakeClient.Create(context.Background(), cfRuntime)).To(Succeed())
			})

			It("fetches the CloudFoundry from the k8s API", func() {
				cf, err := repo.GetCloudFoundry(req)
				Expect(err).NotTo(HaveOccurred())

				Expect(cf.ObjectMeta.Name).To(Equal("my-cf-1"))
			})
		})

		When("the CloudFoundry doesn't exist", func() {
			It("returns an ErrCloudFoundryRuntimeNotFound", func() {
				_, err := repo.GetCloudFoundry(req)
				Expect(err).To(HaveOccurred())

				Expect(err.Error()).To(Equal("CloudFoundry runtime 'my-cf-1' not found"))
			})
		})
	})
})
