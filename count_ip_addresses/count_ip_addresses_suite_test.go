package count_ip_addresses_test

import (
	"testing"

	. "CodeWars/count_ip_addresses"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCountIpAddresses(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountIpAddresses Suite")
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		Expect(IpsBetween("10.0.0.0", "10.0.0.50")).To(Equal(50))
		Expect(IpsBetween("20.0.0.10", "20.0.1.0")).To(Equal(246))
	})
})
