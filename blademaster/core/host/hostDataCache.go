package host

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnHostDataCachePacket(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InHostDataCachePacket
	if !p.PraseInHostDataCachePacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error gamecache packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a gamecache packet but not in server !")
		return
	}
	//找到对应房间
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(),
		uPtr.GetUserChannelServerID(),
		uPtr.GetUserRoomID())
	if rm == nil ||
		rm.Id <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a gamecache packet but not in room !")
		return
	}
	if rm.HostUserID != uPtr.Userid {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a gamecache packet but is not host !")
		return
	}
	//保存缓存
	rm.RoomSaveCache(pkt.PageNum, pkt.Cache)
	DebugInfo(2, "Recv a gamecache packet seq", pkt.PageNum, "from User", uPtr.UserName)
}
