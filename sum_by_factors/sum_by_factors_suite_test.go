package sum_by_factors_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "CodeWars/sum_by_factors"
)

func TestSumByFactors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumByFactors Suite")
}

func dotest(lst []int, exp string) {
	var ans = SumOfDivided(lst)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("SumOfDivided", func() {
	It("Basic tests", func() {
		lst1 := []int{12, 15}
		dotest(lst1, "(2 12)(3 27)(5 15)")

		lst2 := []int{15, 21, 24, 30, 45}
		dotest(lst2, "(2 54)(3 135)(5 90)(7 21)")

		lst3 := []int{100, 100, 107, 110, 116, 118, 123, 126, 158, 204}
		dotest(lst3, "(2 1032)(3 453)(5 310)(7 126)(11 110)(17 204)(29 116)(41 123)(59 118)(79 158)(107 107)")
	})
})
