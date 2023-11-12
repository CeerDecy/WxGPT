package gptclient

import (
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	client := DefaultClient()
	response, err := client.GetResponse("你好")
	if err != nil {
		log.Println(err)
	}
	log.Println("resp: ", response)
}
