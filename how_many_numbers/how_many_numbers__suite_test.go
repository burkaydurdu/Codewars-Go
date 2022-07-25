package how_many_numbers_test

import (
	"testing"

	. "CodeWars/how_many_numbers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHowManyNumbers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HowManyNumbers Suite")
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(FindAll(10, 3)).To(Equal([]int{8, 118, 334}))
		Expect(FindAll(27, 3)).To(Equal([]int{1, 999, 999}))
		Expect(FindAll(84, 4)).To(BeEmpty())
		Expect(FindAll(35, 6)).To(Equal([]int{123, 116999, 566666}))
		Expect(FindAll(89, 7)).To(Equal([]int{123, 116999, 566666}))
	})
})
