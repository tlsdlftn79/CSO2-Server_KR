package room

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	connecting = 0
	joinFailed = 2004
)

func OnRoomFeedback(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InFeedbackPacket
	if !p.PraseRoomFeedbackPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error Feedback packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "send Feedback packet but not in server !")
		return
	}

	switch pkt.ErrorCode {
	case connecting:
		DebugInfo(2, "User", uPtr.UserName, "try to join a match ...")
	case joinFailed:
		DebugInfo(2, "User", uPtr.UserName, "try to join a match but failed !")
	default:
		DebugInfo(2, "User", uPtr.UserName, "send unkown feedback packet code", pkt.ErrorCode)

	}
}
