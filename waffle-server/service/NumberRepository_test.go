package service_test

import (
	. "github.com/1ambda/gokit-waffle/waffle-server/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NumberRepository", func() {
	Describe("Submission", func() {
		Describe("NewSubmission", func() {
			It("should return valid submission instance", func() {
				s := NewSubmission("1ambda", 3)
				Expect(s).NotTo(Equal(nil))
			})
		})
	})
})
