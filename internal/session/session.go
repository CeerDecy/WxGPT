package session

import (
	"log"
	"sync"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/sashabaranov/go-openai"
	_ "github.com/sashabaranov/go-openai"
)

var ChatSession cmap.ConcurrentMap

func init() {
	ChatSession = cmap.New()
}

type Session struct {
	stream  *openai.ChatCompletionStream
	buf     []byte
	Content []byte
	Done    bool
	Sign    chan struct{}
	Lock    sync.Mutex
}

func NewSession(stream *openai.ChatCompletionStream) *Session {
	return &Session{
		stream:  stream,
		buf:     make([]byte, 0),
		Content: make([]byte, 0),
		Done:    false,
		Sign:    make(chan struct{}, 1),
		Lock:    sync.Mutex{},
	}
}

func (s *Session) ReadResp() (res []byte) {
	s.Lock.Lock()
	res = make([]byte, len(s.buf))
	copy(res, s.buf)
	s.buf = make([]byte, 0)
	s.Lock.Unlock()
	return res
}

func (s *Session) ReadStream() {
	go func() {
		for {
			recv, err := s.stream.Recv()
			if err != nil {
				s.stream.Close()
				s.Done = true
				s.Sign <- struct{}{}
				return
			}
			for _, v := range recv.Choices {
				s.Lock.Lock()
				s.buf = append(s.buf, v.Delta.Content...)
				s.Content = append(s.Content, v.Delta.Content...)
				log.Printf("%s", string(s.buf))
				s.Lock.Unlock()
			}
		}
	}()
}
