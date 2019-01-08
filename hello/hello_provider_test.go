package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageProvider(t *testing.T) {
	//given
	expected := "This is a receipt 📃"

	//when
	actual := MessageProvider()

	//then
	assert.Equal(t, expected, actual, "A notice of receipt")
}
