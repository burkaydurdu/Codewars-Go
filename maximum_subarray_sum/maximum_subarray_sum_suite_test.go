package maximum_subarray_sum_test

import (
	"testing"

	. "CodeWars/maximum_subarray_sum"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMaximumSubarraySum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaximumSubarraySum Suite")
}

var _ = Describe("Tests", func() {
	It("Basic tests", func() {
		Expect(MaximumSubarraySum([]int{})).To(Equal(0))
		Expect(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})).To(Equal(6))
		Expect(MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})).To(Equal(0))
	})
})
