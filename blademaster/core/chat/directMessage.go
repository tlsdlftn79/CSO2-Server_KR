package chat

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnChatDirectMessage(p *InChatPacket, client net.Conn) {
	//检索数据包
	if len(p.Destination) <= 0 || p.DestinationLen <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a null destination directchat packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent DirectMessage but not in server !")
		return
	}
	if CompareBytes([]byte(uPtr.IngameName), p.Destination) {
		DebugInfo(2, "Error : User", uPtr.UserName, "sent DirectMessage to self !")
		return
	}
	reciver := GetUserFromIngameName(p.Destination)
	if reciver == nil ||
		reciver.Userid <= 0 {
		DebugInfo(2, "Error : User", uPtr.UserName, "sent DirectMessage but reciver not in server !")
		return
	}
	//发送数据
	SendPacket(BuildDirectMessage(uPtr, reciver, 0, p), uPtr.CurrentConnection)
	SendPacket(BuildDirectMessage(uPtr, reciver, 1, p), reciver.CurrentConnection)
	DebugInfo(1, "User", uPtr.IngameName, "say <", string(p.Message), "> to User", reciver.UserName)
}

func BuildDirectMessage(sender *User, reciver *User, isReciver uint8, p *InChatPacket) []byte {
	if isReciver != 0 { //发送给接收者
		temp := make([]byte, 10+len(sender.IngameName)+int(p.MessageLen))
		offset := 0
		WriteUint8(&temp, ChatDirectMessage, &offset)
		WriteUint8(&temp, sender.Gm, &offset)
		WriteUint8(&temp, 1, &offset)
		WriteString(&temp, []byte(sender.IngameName), &offset)

		if sender.IsVIP() {
			WriteUint8(&temp, 1, &offset)
		} else {
			WriteUint8(&temp, 0, &offset)
		}
		WriteUint8(&temp, sender.VipLevel, &offset)

		WriteLongString(&temp, p.Message, &offset)
		return BytesCombine(BuildHeader(reciver.CurrentSequence, PacketTypeChat), temp[:offset])
	}
	temp := make([]byte, 10+len(reciver.IngameName)+int(p.MessageLen))
	offset := 0
	WriteUint8(&temp, ChatDirectMessage, &offset)
	WriteUint8(&temp, reciver.Gm, &offset)
	WriteUint8(&temp, 0, &offset)
	WriteString(&temp, []byte(reciver.IngameName), &offset)

	if sender.IsVIP() {
		WriteUint8(&temp, 1, &offset)
	} else {
		WriteUint8(&temp, 0, &offset)
	}
	WriteUint8(&temp, sender.VipLevel, &offset)

	WriteLongString(&temp, p.Message, &offset)
	return BytesCombine(BuildHeader(sender.CurrentSequence, PacketTypeChat), temp[:offset])

}
