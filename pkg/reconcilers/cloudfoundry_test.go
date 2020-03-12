package reconcilers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/circle/pkg/reconcilers"

	"errors"

	"github.com/cloudfoundry-community/go-cfclient"
	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	"github.com/teddyking/circle/pkg/reconcilers/reconcilersfakes"
	"github.com/teddyking/circle/pkg/repositories"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ = Describe("CloudFoundry Reconcile", func() {
	var (
		fakeRuntimeRepo     *reconcilersfakes.FakeRuntimeRepo
		fakeCFRepo          *reconcilersfakes.FakeCFRepo
		fakeCFClientCreator CFClientCreator
		fakeCFClient        *reconcilersfakes.FakeCFClient
		cfclientConfig      *cfclient.Config
		reconciler          *CloudFoundryReconciler
		reconcileErr        error
	)

	BeforeEach(func() {
		fakeRuntimeRepo = &reconcilersfakes.FakeRuntimeRepo{}
		fakeCFRepo = &reconcilersfakes.FakeCFRepo{}
		fakeCFClient = &reconcilersfakes.FakeCFClient{}

		fakeCFClientCreator = func(config *cfclient.Config) (CFClient, error) {
			cfclientConfig = config
			return fakeCFClient, nil
		}

		cf := &runtimev1alpha1.CloudFoundry{
			ObjectMeta: metav1.ObjectMeta{
				Name: "my-cf-1",
			},
			Spec: runtimev1alpha1.CloudFoundrySpec{
				API: "https://api.sys.my-cf-1.example.com",
			},
		}

		orgs := []cfclient.Org{
			cfclient.Org{Name: "my-org-1", Guid: "my-guid-1"},
			cfclient.Org{Name: "my-org-2", Guid: "my-guid-2"},
		}

		fakeRuntimeRepo.GetCloudFoundryReturns(cf, nil)
		fakeCFClient.ListOrgsReturns(orgs, nil)
	})

	JustBeforeEach(func() {
		reconciler = &CloudFoundryReconciler{RuntimeRepo: fakeRuntimeRepo, CFRepo: fakeCFRepo, CFClientCreator: fakeCFClientCreator}

		req := ctrl.Request{NamespacedName: types.NamespacedName{Name: "my-cf-1"}}
		_, reconcileErr = reconciler.Reconcile(req)
	})

	It("fetches the CloudFoundry", func() {
		Expect(fakeRuntimeRepo.GetCloudFoundryCallCount()).To(Equal(1))
	})

	It("creates a cfclient with the correct configuration", func() {
		Expect(cfclientConfig).To(Equal(&cfclient.Config{
			ApiAddress: "https://api.sys.my-cf-1.example.com",
		}))
	})

	It("fetches all Orgs in the CloudFoundry", func() {
		Expect(fakeCFClient.ListOrgsCallCount()).To(Equal(1))
	})

	It("creates Org resources for each of the fetched Orgs", func() {
		Expect(fakeCFRepo.CreateOrgCallCount()).To(Equal(2))

		name1, guid1 := fakeCFRepo.CreateOrgArgsForCall(0)
		Expect(name1).To(Equal("my-org-1"))
		Expect(guid1).To(Equal("my-guid-1"))

		name2, guid2 := fakeCFRepo.CreateOrgArgsForCall(1)
		Expect(name2).To(Equal("my-org-2"))
		Expect(guid2).To(Equal("my-guid-2"))
	})

	When("fetching the CloudFoundry doesn't exist", func() {
		BeforeEach(func() {
			fakeRuntimeRepo.GetCloudFoundryReturns(&runtimev1alpha1.CloudFoundry{}, repositories.ErrCloudFoundryRuntimeNotFound{})
		})

		It("doesn't return an error", func() {
			Expect(reconcileErr).NotTo(HaveOccurred())
		})
	})

	When("fetching the CloudFoundry returns an error", func() {
		BeforeEach(func() {
			fakeRuntimeRepo.GetCloudFoundryReturns(&runtimev1alpha1.CloudFoundry{}, errors.New("error-fetching-cf"))
		})

		It("returns the error", func() {
			Expect(reconcileErr).To(MatchError("error-fetching-cf"))
		})
	})

	When("creating the cfclient returns an error", func() {
		BeforeEach(func() {
			fakeCFClientCreator = func(config *cfclient.Config) (CFClient, error) {
				return &cfclient.Client{}, errors.New("error-creating-cfclient")
			}
		})

		It("returns the error", func() {
			Expect(reconcileErr).To(MatchError("error-creating-cfclient"))
		})
	})

	When("fetching the Orgs in the CloudFoundry returns an error", func() {
		BeforeEach(func() {
			fakeCFClient.ListOrgsReturns([]cfclient.Org{}, errors.New("error-fetching-cf-orgs"))
		})

		It("returns the error", func() {
			Expect(reconcileErr).To(MatchError("error-fetching-cf-orgs"))
		})
	})

	When("creating an Org returns an error", func() {
		BeforeEach(func() {
			fakeCFRepo.CreateOrgReturns(errors.New("error-creating-cf-org"))
		})

		It("returns the error", func() {
			Expect(reconcileErr).To(MatchError("error-creating-cf-org"))
		})
	})
})
