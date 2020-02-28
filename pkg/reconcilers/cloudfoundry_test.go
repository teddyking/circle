package reconcilers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/circle/pkg/reconcilers"
	"github.com/teddyking/circle/pkg/repositories"

	"errors"

	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	"github.com/teddyking/circle/pkg/reconcilers/reconcilersfakes"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ = Describe("Cloudfoundry", func() {
	var (
		fakeRuntimeRepo *reconcilersfakes.FakeRuntimeRepo
		reconciler      *CloudFoundryReconciler
		reconcileErr    error
	)

	BeforeEach(func() {
		fakeRuntimeRepo = &reconcilersfakes.FakeRuntimeRepo{}
		reconciler = &CloudFoundryReconciler{RuntimeRepo: fakeRuntimeRepo}
	})

	JustBeforeEach(func() {
		_, reconcileErr = reconciler.Reconcile(ctrl.Request{})
	})

	It("fetches the CloudFoundry", func() {
		Expect(fakeRuntimeRepo.GetCloudFoundryCallCount()).To(Equal(1))
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
			Expect(reconcileErr).To(HaveOccurred())
		})
	})
})
