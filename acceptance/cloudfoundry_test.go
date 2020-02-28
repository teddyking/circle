package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"path/filepath"

	"github.com/teddyking/circle/test/helpers"
)

var _ = Describe("CloudFoundry", func() {
	var (
		kubectl *helpers.Kubectl
	)

	BeforeEach(func() {
		var err error

		kubectl, err = helpers.NewKubectlForContext("kind-kind", GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Expect(kubectl.Apply(pathToCRDs)).To(Succeed())
	})

	AfterEach(func() {
		Expect(kubectl.Delete(pathToCRDs)).To(Succeed())
	})

	Describe("Create", func() {
		var (
			testCloudFoundryYML = filepath.Join("..", "test", "assets", "cloudfoundry.test.yml")
		)

		BeforeEach(func() {
			Expect(kubectl.Apply(testCloudFoundryYML)).To(Succeed())
		})

		AfterEach(func() {
			Expect(kubectl.Delete(testCloudFoundryYML)).To(Succeed())
		})

		XIt("creates Org resources in k8s for each Org resource in the CloudFoundry", func() {
			Eventually(func() error {
				return kubectl.Get("org", "acceptance-org")
			}).Should(Succeed())
		})
	})
})
