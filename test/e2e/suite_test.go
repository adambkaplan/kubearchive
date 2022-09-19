package e2e

import (
	"context"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/adambkaplan/kubearchive/test/e2e/framework"
)

var (
	kubeClient  client.Client
	testContext context.Context
	cancel      context.CancelFunc
)

func TestE2E(t *testing.T) {
	suite := os.Getenv("KUBEARCHIVE_TEST_SUITE")
	if suite != "e2e" {
		t.Skip("skipping e2e test because KUBEARCHIVE_TEST_SUITE=e2e was not set")
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Suite")
}

var _ = BeforeSuite(func() {
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	rest, err := ctrl.GetConfig()
	Expect(err).NotTo(HaveOccurred())

	kubeClient, err = client.New(rest, client.Options{
		Scheme: scheme,
	})
	Expect(err).NotTo(HaveOccurred())
	testContext, cancel = context.WithCancel(context.Background())
})

var _ = BeforeEach(func() {
	By("waiting for project deployment")
	err := wait.PollImmediate(10*time.Second, 5*time.Minute, func() (bool, error) {
		return framework.IsDeploymentReady(testContext, kubeClient, "kubearchive", "kubearchive-controller-manager")
	})
	Expect(err).NotTo(HaveOccurred())

})

var _ = AfterSuite(func() {
	cancel()
})
