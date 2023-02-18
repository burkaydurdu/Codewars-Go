package explosive_sum

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExplosiveSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Example")
}

func dotest(n uint64, exp uint64) {
	Expect(ExpSum(n)).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		dotest(1, 1)
		dotest(2, 2)
		dotest(3, 3)
		dotest(4, 5)
		dotest(5, 7)
		dotest(10, 42)
	})
})
