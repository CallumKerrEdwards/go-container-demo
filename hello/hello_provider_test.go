package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldProvider(t *testing.T) {
	//given
	expected := "This is a receipt 📃"

	//when
	actual := HelloWorldProvider()

	//then
	assert.Equal(t, expected, actual, "A notice of receipt")
}
