package robot

import (
	"os"
	"testing"
)

func TestNewTextMessage(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	m := NewTextMessage()
	m.SetContent("test")
	_, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTextMessage_AtAll(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	m := NewTextMessage()
	m.SetContent("test").AtAll()
	_, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
