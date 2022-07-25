package hamming_numbers_test

import (
	"testing"

	. "CodeWars/hamming_numbers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHammingNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HammingNumbers Suite")
}

func dotest(n int, exp uint) {
	Expect(Hammer(n)).To(Equal(exp))
}

var _ = Describe("Sample tests", func() {
	It("Should handle basic cases", func() {
		dotest(1, 1)
		dotest(2, 2)
		dotest(3, 3)
		dotest(4, 4)
		dotest(5, 5)
		dotest(6, 6)
		dotest(7, 8)
		dotest(8, 9)
		dotest(9, 10)
		dotest(10, 12)
		dotest(11, 15)
		dotest(12, 16)
		dotest(13, 18)
		dotest(14, 20)
		dotest(15, 24)
		dotest(16, 25)
		dotest(17, 27)
		dotest(18, 30)
		dotest(19, 32)
		dotest(1691, 2125764000)
	})
})
