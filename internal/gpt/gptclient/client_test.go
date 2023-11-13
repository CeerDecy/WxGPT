package gptclient

import (
	"fmt"
	"log"
	"reflect"
	"sync"
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

func TestStreamClient(t *testing.T) {
	client := DefaultClient()
	response, err := client.GetStreamResponse("怎么求导")
	if err != nil {
		log.Println(err)
	}
	content := make([]byte, 0)
	lock := sync.Mutex{}
	sign := make(chan struct{})
	go func() {
		for {
			recv, err := response.Recv()
			if err != nil {
				log.Println(err)
				sign <- struct{}{}
				break
			}
			for _, v := range recv.Choices {
				lock.Lock()
				content = append(content, v.Delta.Content...)
				lock.Unlock()
			}
		}
	}()
	log.Println("准备输出")
	for {
		select {
		case <-sign:
			return
		default:
			lock.Lock()
			fmt.Print(string(content))
			content = make([]byte, 0)
			lock.Unlock()
		}
	}
}
