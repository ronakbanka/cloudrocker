package focker_test

import (
	"github.com/hatofmonkeys/cloudfocker/focker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Focker", func() {
	var (
		testfocker *focker.Focker
		buffer     *gbytes.Buffer
	)
	BeforeEach(func() {
		testfocker = focker.NewFocker()
		buffer = gbytes.NewBuffer()
	})

	Describe("Displaying the docker version", func() {
		It("should tell Docker to output its version", func() {
			testfocker.DockerVersion(buffer)
			Eventually(buffer).Should(gbytes.Say(`Checking Docker version`))
			Eventually(buffer).Should(gbytes.Say(`Client API version: `))
			Eventually(buffer).Should(gbytes.Say(`Go version \(client\): go`))
		})
	})

	/*
	  Describe("Bootstrapping the base image", func() {
	  	It("should download and tag the lucid64 filesystem", func() {
	  		Expect(true).To(Equal(true))
	  		})
	  	})

	  Describe("Writing a dockerfile", func() {
	  	It("should output a valid dockerfile", func() {
	  		Expect(true).To(Equal(true))
	  		})
	  	})

	  Describe("Building a docker container", func() {
	  	It("should output a built container tag", func() {
	  		Expect(true).To(Equal(true))
	  		})
	  	})

	  Describe("Running the docker container", func() {
	  	It("should output a valid URL for the running application", func() {
	  		Expect(true).To(Equal(true))
	  		})
	  	})
	*/
})