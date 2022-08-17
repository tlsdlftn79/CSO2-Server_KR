package chat

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnChat(p *PacketData, client net.Conn) {
	var pkt InChatPacket
	if p.PraseInChatPacket(&pkt) {

		switch pkt.Type {
		case ChatDirectMessage:
			OnChatDirectMessage(&pkt, client)
		case ChatChannel:
			OnChatChannelMessage(&pkt, client)
		case ChatRoom:
			OnChatRoomMessage(&pkt, client)
		case ChatIngameGlobal:
			OnChatGlobalMessage(&pkt, client)
		case ChatIngameTeam:
			OnChatTeamMessage(&pkt, client)
		default:
			DebugInfo(2, "Unknown chat packet", pkt.Type, "from", client.RemoteAddr().String())
		}
	} else {
		DebugInfo(2, "Error : Recived a illegal chat packet from", client.RemoteAddr().String())
	}
}

func BuildChatMessage(sender *User, p *InChatPacket, chattype uint8) []byte {
	temp := make([]byte, 10+len(sender.IngameName)+int(p.MessageLen))
	offset := 0
	WriteUint8(&temp, chattype, &offset)
	WriteUint8(&temp, sender.Gm, &offset)
	WriteString(&temp, []byte(sender.IngameName), &offset)

	if sender.IsVIP() {
		WriteUint8(&temp, 1, &offset)
	} else {
		WriteUint8(&temp, 0, &offset)
	}
	WriteUint8(&temp, sender.VipLevel, &offset)

	WriteLongString(&temp, p.Message, &offset)
	return temp[:offset]
}
