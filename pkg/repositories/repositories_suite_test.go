package repositories_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	runtimev1alpha1 "github.com/teddyking/circle/apis/runtime/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	runtimev1alpha1.AddToScheme(scheme.Scheme)
	RunSpecs(t, "Repositories Suite")
}
