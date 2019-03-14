package tgo

import (
	"bytes"
	"encoding/binary"
	"github.com/tgo-team/tgo-talk/tgo/packets"
)

type Client struct {
	ClientID uint64
}

func NewClient(clientID uint64) *Client {
	return &Client{ClientID: clientID}
}
func (c *Client) MarshalBinary() (data []byte, err error) {
	var body bytes.Buffer
	body.Write(packets.EncodeUint64(c.ClientID))
	return body.Bytes(), nil
}

func (c *Client) UnmarshalBinary(data []byte) error {
	c.ClientID = binary.BigEndian.Uint64(data[:8])
	return nil
}

type Storage interface {
	// ------ 消息操作 -----
	SaveMsg(msgContext *MsgContext) error // 保存消息
	StorageMsgChan() chan *MsgContext     // 读取消息
	// ------ 管道操作 -----
	SaveChannel(c *Channel) error                   // 保存管道
	GetChannel(channelID uint64) (*Channel, error)  // 获取管道
	AddClient(c *Client) error                 // 添加消费者
	Bind(clientID uint64, channelID uint64) error // 绑定消费者和通道的关系
	GetClientIDs(channelID uint64) ([]uint64,error) // 获取所属管道所有的客户端
}
