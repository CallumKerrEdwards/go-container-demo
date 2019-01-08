package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCustomerName(t *testing.T) {
	//given
	givenName := "Test Name"

	//when
	cus := Customer{givenName}

	//then
	assert.Equal(t, givenName, cus.Name, "A customer should have a name")
}
