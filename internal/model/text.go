package model

import "time"

type TextReceive struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   uint64 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MediaId      uint64 `xml:"MediaId"`
	MsgDataId    uint64 `xml:"MsgDataId"`
	Idx          uint64 `xml:"Idx"`
}

// <xml>
// <ToUserName><![CDATA[toUser]]></ToUserName>
// <FromUserName><![CDATA[fromUser]]></FromUserName>
// <CreateTime>12345678</CreateTime>
// <MsgType><![CDATA[text]]></MsgType>
// <Content><![CDATA[你好]]></Content>
// </xml>
type TextResponse struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   uint64 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}

func DefaultTextResp(openid string, token string, content string) *TextResponse {
	return &TextResponse{
		ToUserName:   openid,
		FromUserName: token,
		CreateTime:   uint64(time.Now().Unix()),
		MsgType:      "text",
		Content:      content,
	}
}
