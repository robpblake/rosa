package kubeletconfig

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"
)

var _ = Describe("Create KubeletConfig Options", func() {

	Context("Create New Options", func() {

		It("Initialises defaults correctly", func() {
			o := NewCreateOptions()
			Expect(o.podPidsLimit).To(Equal(0))
		})
	})

	Context("Add Flags to Options", func() {
		It("Adds --pods-pids-limits flag", func() {
			o := NewCreateOptions()
			flags := &pflag.FlagSet{}
			o.AddFlags(flags)

			Expect(flags.Lookup("pod-pids-limit")).NotTo(BeNil())
		})
	})

	Context("Validates passed flags", func() {
		It("Fails if pod-pids-limit is less than minimum", func() {

		})
	})

})
