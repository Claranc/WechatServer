package wx

import (
	"encoding/xml"
	"strconv"
	"time"
)

type Base struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	MsgType      CDATAText
	CreateTime   CDATAText
}

func (b *Base) InitBaseData(w *WeixinClient, msgtype string) {

	b.FromUserName = value2CDATA(w.Message["ToUserName"].(string))
	b.ToUserName = value2CDATA(w.Message["FromUserName"].(string))
	b.CreateTime = value2CDATA(strconv.FormatInt(time.Now().Unix(), 10))
	b.MsgType = value2CDATA(msgtype)
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}


//文字类型
type TextMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content CDATAText
	MediaId CDATAText
}

//图片类型
type ImageMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Image []ImageInfo //xml嵌套的话就使用切片类型
}

type ImageInfo struct {
	MediaId CDATAText
}
