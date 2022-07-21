package sum_of_intervals_test

import (
	"testing"

	. "CodeWars/sum_of_intervals"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSumOfIntervals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumOfIntervals Suite")
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(SumOfIntervals([][2]int{{1, 5}})).To(Equal(4))
		Expect(SumOfIntervals([][2]int{{1, 5}, {5, 6}, {1, 5}})).To(Equal(5))
		Expect(SumOfIntervals([][2]int{{1, 5}, {6, 10}})).To(Equal(8))
		Expect(SumOfIntervals([][2]int{{1, 5}, {1, 5}})).To(Equal(4))
		Expect(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})).To(Equal(7))
		Expect(SumOfIntervals([][2]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}})).To(Equal(19))
	})
})
