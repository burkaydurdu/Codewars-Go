package matrix_determinant_test

import (
	"testing"

	. "CodeWars/matrix_determinant"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMatrixDeterminant(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MatrixDeterminant Suite")
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(Determinant([][]int{{1}})).To(Equal(1))
		Expect(Determinant([][]int{{1, 3}, {2, 5}})).To(Equal(-1))
		Expect(Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}})).To(Equal(-20))
	})
})
