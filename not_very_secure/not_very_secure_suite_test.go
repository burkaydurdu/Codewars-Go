package not_very_secure_test

import (
	"testing"

	. "CodeWars/not_very_secure"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNotVerySecure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NotVerySecure Suite")
}

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(Alphanumeric(".*?")).To(Equal(false))
		Expect(Alphanumeric("a")).To(Equal(true))
		Expect(Alphanumeric("Mazinkaiser")).To(Equal(true))
		Expect(Alphanumeric("hello world_")).To(Equal(false))
		Expect(Alphanumeric("PassW0rd")).To(Equal(true))
		Expect(Alphanumeric("     ")).To(Equal(false))
		Expect(Alphanumeric("")).To(Equal(false))
		Expect(Alphanumeric("\n\t\n")).To(Equal(false))
		Expect(Alphanumeric("ciao\n$$_")).To(Equal(false))
		Expect(Alphanumeric("__ * __")).To(Equal(false))
		Expect(Alphanumeric("&)))(((")).To(Equal(false))
		Expect(Alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
	})
})
