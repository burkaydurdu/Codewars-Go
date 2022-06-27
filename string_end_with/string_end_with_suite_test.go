package string_end_with_test

import (
	"testing"

	. "CodeWars/string_end_with"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStringEndWith(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringEndWith Suite")
}

var _ = Describe("Example test:", func() {
	It("Should work for fixed tests", func() {
		Expect(Solution("", "")).To(Equal(true))
		Expect(Solution(" ", "")).To(Equal(true))
		Expect(Solution("abc", "c")).To(Equal(true))
		Expect(Solution("banana", "ana")).To(Equal(true))
		Expect(Solution("a", "z")).To(Equal(false))
		Expect(Solution("", "t")).To(Equal(false))
		Expect(Solution("$a = $b + 1", "+1")).To(Equal(false))
		Expect(Solution("    ", "   ")).To(Equal(true))
		Expect(Solution("1oo", "100")).To(Equal(false))
	})
})
