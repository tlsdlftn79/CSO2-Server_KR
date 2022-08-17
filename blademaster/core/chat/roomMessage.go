package chat

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnChatRoomMessage(p *InChatPacket, client net.Conn) {
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent RoomMessage but not in server !")
		return
	}
	if uPtr.GetUserRoomID() <= 0 {
		DebugInfo(2, "Error : User", string(uPtr.IngameName), "sent RoomMessage but not in room !")
		return
	}
	//找到对应房间
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(), uPtr.GetUserChannelID(), uPtr.GetUserRoomID())
	if rm == nil || rm.Id <= 0 {
		DebugInfo(2, "Error : User", string(uPtr.IngameName), "sent RoomMessage but not in room !")
		return
	}
	if uPtr.CurrentIsIngame {
		DebugInfo(2, "Error : User", string(uPtr.IngameName), "sent RoomMessage but ingame !")
		return
	}
	//发送数据
	msg := BuildChatMessage(uPtr, p, ChatRoom)
	rm.RoomMutex.Lock()
	defer rm.RoomMutex.Unlock()
	for _, v := range rm.Users {
		if !v.CurrentIsIngame {
			SendPacket(BytesCombine(BuildHeader(v.CurrentSequence, PacketTypeChat), msg), v.CurrentConnection)
		}
	}
	DebugInfo(1, "User", string(uPtr.IngameName), "say <", string(p.Message), "> in room id", rm.Id)
}
