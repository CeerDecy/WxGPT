package model

// ReceiveMsg
// <ToUserName><![CDATA[toUser]]></ToUserName>
// <FromUserName><![CDATA[fromUser]]></FromUserName>
// <CreateTime>1348831860</CreateTime>
// <MsgType><![CDATA[image]]></MsgType>
// <Content><![CDATA[this is a test]]></Content>
// <MsgId>1234567890123456</MsgId>
// <MsgDataId>xxxx</MsgDataId>
// <Idx>xxxx</Idx>
type ReceiveMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   uint64 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MediaId      uint64 `xml:"MediaId"`
	MsgDataId    uint64 `xml:"MsgDataId"`
	Idx          uint64 `xml:"Idx"`
}
