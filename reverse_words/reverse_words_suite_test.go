package reverse_words_test

import (
	"testing"

	. "CodeWars/reverse_words"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReverseWords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReverseWords Suite")
}

var _ = Describe("ReverseWords", func() {
	It("should work for sample test cases", func() {
		Expect(ReverseWords("The quick brown fox jumps over the lazy dog.")).To(Equal("ehT kciuq nworb xof spmuj revo eht yzal .god"))
		Expect(ReverseWords("apple")).To(Equal("elppa"))
		Expect(ReverseWords("a b c d")).To(Equal("a b c d"))
		Expect(ReverseWords("double  spaced  words")).To(Equal("elbuod  decaps  sdrow"))
		Expect(ReverseWords("stressed desserts")).To(Equal("desserts stressed"))
		Expect(ReverseWords("hello hello")).To(Equal("olleh olleh"))
	})
})
