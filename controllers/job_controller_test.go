package controllers

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Job controller", func() {

	var job *batchv1.Job

	BeforeEach(func() {
		By("creating a test Job")
		job = &batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "default",
				Name:      "test-job",
			},
			Spec: batchv1.JobSpec{
				Template: v1.PodTemplateSpec{
					Spec: v1.PodSpec{
						Containers: []v1.Container{
							{
								Name:  "test",
								Image: "ubi9",
							},
						},
						RestartPolicy: v1.RestartPolicyNever,
					},
				},
			},
		}
		err := k8sClient.Create(ctx, job)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		By("deleting the test Job")
		err := k8sClient.Delete(ctx, job)
		Expect(err).NotTo(HaveOccurred())
	})

	It("adds an annotation to a Job", func() {
		Eventually(func() bool {
			By("getting the test job")
			err := k8sClient.Get(ctx, client.ObjectKeyFromObject(job), job)
			if err != nil {
				Expect(client.IgnoreNotFound(err)).NotTo(HaveOccurred())
				return false
			}
			By("checking the annotations")
			annotations := job.GetAnnotations()
			if annotations == nil {
				return false
			}
			fmt.Printf("found annotations %s", annotations)
			val, exists := annotations[ArchivedAnnotation]
			return exists && val == "true"
		}, 10*time.Second, 500*time.Millisecond).Should(BeTrue())
	})
})
