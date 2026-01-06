package entity

import (
	"testing"
	"github.com/onsi/gomega"
)

func TestNameField(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	number := uint(1)

	customer := Customer{
		Name: "",
		Email: "guy@gmail.com",
		CustomerID: "L1234567",
		Age: 18,
		Nember: &number,
	}
	
	ok, err := customer.Validate()

	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err.Error()).To(gomega.Equal("Name Is Required"))
}