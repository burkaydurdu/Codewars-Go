package build_a_pile_of_cubes_test

import (
	"testing"

	. "CodeWars/build_a_pile_of_cubes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuildAPileOfCubes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuildAPileOfCubes Suite")
}

func dotest(n int, exp int) {
	var ans = FindNb(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(4183059834009, 2022)
		dotest(24723578342962, -1)
	})
})
