package host

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnHostDataUsersPacket(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InHostDataUsersPacket
	if !p.PraseInHostDataUsersPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error gameusers packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a gameusers packet but not in server !")
		return
	}
	//找到对应房间
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(),
		uPtr.GetUserChannelServerID(),
		uPtr.GetUserRoomID())
	if rm == nil ||
		rm.Id <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a gameusers packet but not in room !")
		return
	}
	if rm.HostUserID != uPtr.Userid {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a gameusers packet but is not host !")
		return
	}
	//保存
	DebugInfo(2, "Recv a gameusers packet seq", pkt.PageNum, "from User", uPtr.UserName)
}
