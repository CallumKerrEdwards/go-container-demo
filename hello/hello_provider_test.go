package hello

import "testing"

func TestHelloWorldProvider(t *testing.T) {
    //given
    expected := "This is a receipt 📃"
    
    //when
    actual := HelloWorldProvider()
    
    //then
    if actual != expected {
        t.Errorf("Result '%s' was not the expected result '%s'", actual, expected)
    }
}