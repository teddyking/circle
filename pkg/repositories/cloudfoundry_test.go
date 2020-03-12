package repositories_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/circle/pkg/repositories"

	"context"

	cfv1alpha1 "github.com/teddyking/circle/apis/cf/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("CloudFoundry", func() {
	var (
		fakeClient client.Client
		repo       *CloudFoundry
	)

	BeforeEach(func() {
		fakeClient = fake.NewFakeClient()
		repo = &CloudFoundry{KubeClient: fakeClient}
	})

	Describe("CreateOrg", func() {
		var cfOrg cfv1alpha1.Org

		BeforeEach(func() {
			cfOrg = cfv1alpha1.Org{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-org-1",
				},
				Spec: cfv1alpha1.OrgSpec{
					Guid: "my-guid-1",
				},
			}
		})

		When("the Org doesn't exist", func() {
			It("creates the Org", func() {
				Expect(repo.CreateOrg("my-org-1", "my-guid-1")).To(Succeed())

				var createdOrg cfv1alpha1.Org
				Expect(fakeClient.Get(context.Background(), types.NamespacedName{Name: "my-org-1"}, &createdOrg)).To(Succeed())
				Expect(createdOrg.ObjectMeta.Name).To(Equal("my-org-1"))
				Expect(createdOrg.Spec.Guid).To(Equal("my-guid-1"))
			})
		})

		When("the Org exists", func() {
			BeforeEach(func() {
				Expect(fakeClient.Create(context.Background(), &cfOrg)).To(Succeed())
			})

			It("doesn't error", func() {
				Expect(repo.CreateOrg("my-org-1", "my-guid-1")).To(Succeed())
			})
		})
	})
})
