package weird_string_case_test

import (
	"testing"

	. "CodeWars/weird_string_case"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWeirdStringCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WeirdStringCase Suite")
}

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values", func() {
		Expect(ToWeirdCase("abc def")).To(Equal("AbC DeF"))
		Expect(ToWeirdCase("ABC")).To(Equal("AbC"))
		Expect(ToWeirdCase("This is a test Looks like you passed")).To(Equal("ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"))
	})
})
