package email_resource_test

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Out", func() {
	BeforeEach(func() {
		Run("go", "build", "-o", "../bin/out", "../actions/out")
	})

	It("should print the current time as a version and exit 0", func() {
		output, err := RunWithStdinAllowError("some input data", "../bin/out", "some", "arguments")
		Expect(err).NotTo(HaveOccurred())

		var outdata struct {
			Version struct {
				Time time.Time
			}
		}
		Expect(json.Unmarshal([]byte(output), &outdata)).To(Succeed())
		Expect(outdata.Version.Time).To(BeTemporally("~", time.Now(), 5*time.Second))

		var untyped map[string]interface{}
		Expect(json.Unmarshal([]byte(output), &untyped)).To(Succeed())
		Expect(untyped).To(HaveKey("version"))
	})
})