package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("archive REST api", func() {

	It("returns \"hello\" from /archive", func() {
		hostname := os.Getenv("KUBEARCHIVE_HOSTNAME")
		if len(hostname) == 0 {
			hostname = "localhost:8080"
		}

		testURL := fmt.Sprintf("http://%s/archive", hostname)
		resp, err := http.Get(testURL)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		result := map[string]string{}
		err = json.Unmarshal(body, &result)
		Expect(err).NotTo(HaveOccurred())

		Expect(result["message"]).To(Equal("hello"))

	})
})
