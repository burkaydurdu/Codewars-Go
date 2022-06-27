package multiplication_table_test

import (
	"testing"

	. "CodeWars/multiplication_table"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMultiplicationTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultiplicationTable Suite")
}

var _ = Describe("Sample Tests", func() {
	It("Sample Tests", func() {
		Expect(MultiplicationTable(1)).To(Equal([][]int{{1}}))
		Expect(MultiplicationTable(2)).To(Equal([][]int{{1, 2}, {2, 4}}))
		Expect(MultiplicationTable(3)).To(Equal([][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}))
	})
})
