package gptclient

import (
	"log"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	client := DefaultClient()
	response, err := client.GetResponse("怎么求导")
	if err != nil {
		log.Println(reflect.TypeOf(err))
		if err.Error() == `Post "https://proxy.geekai.co/v1/chat/completions": context deadline exceeded` {
			log.Println("链接超时")
		}
	}

	log.Println("resp: ", response)
}
