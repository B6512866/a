package entity

import (
	"testing"
	"github.com/onsi/gomega"
)

func TestAllfield(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	number := uint(1)

	customer := Customer{
		Name: "Jessada",
		Email: "guy@gmail.com",
		CustomerID: "L1234567",
		Age: 18,
		Nember: &number,
	}

	ok, err := customer.Validate()

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}