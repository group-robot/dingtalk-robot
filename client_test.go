package robot

import (
	"os"
	"testing"
)

func TestClient_Send(t *testing.T) {
	c := NewClient(os.Getenv("webhok"), os.Getenv("secret"))
	response, err := c.Send(&TextMessage{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
