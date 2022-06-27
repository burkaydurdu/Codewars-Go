package vowel_count_test

import (
	"testing"

	. "CodeWars/vowel_count"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVowelCount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VowelCount Suite")
}

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(GetCount("abracadabra")).To(Equal(5))
		Expect(GetCount("rkrkrk")).To(Equal(0))
	})
})
