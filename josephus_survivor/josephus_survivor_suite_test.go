package josephus_survivor_test

import (
	"testing"

	. "CodeWars/josephus_survivor"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJosephusSurvivor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JosephusSurvivor Suite")
}

var _ = Describe("Sample Tests", func() {
	It("should handle basic tests", func() {
		dotest(7, 3, 4)
		dotest(11, 19, 10)
		dotest(40, 3, 28)
		dotest(14, 2, 13)
		dotest(100, 1, 100)
	})
})

func dotest(n, k, e int) {
	Expect(JosephusSurvivor(n, k)).To(Equal(e))
}
