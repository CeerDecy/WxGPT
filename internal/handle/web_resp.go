package handle

import (
	"io"

	"github.com/gin-gonic/gin"

	"WxGPT/internal/session"
)

func StreamWeb(ctx *gin.Context) {
	sid, _ := ctx.GetQuery("sid")
	streamRaw, ok := session.ChatSession.Get(sid)
	if !ok {
		ctx.String(200, "not found session id")
		return
	}
	//fmt.Println(streamRaw)
	sess := streamRaw.(*session.Session)
	if sess.Done {
		ctx.String(200, string(sess.Content))
	} else {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
		ctx.Stream(func(w io.Writer) bool {
			res := sess.ReadResp()
			_, err := w.Write(res)
			select {
			case <-sess.Sign:
				return false
			default:
				return err == nil
			}
		})
	}

}
