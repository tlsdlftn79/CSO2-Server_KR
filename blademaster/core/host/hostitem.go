package host

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnHostItemUsing(p *PacketData, client net.Conn) {
	//检查数据包
	var pkt InHostItemUsingPacket
	if !p.PraseInHostItemUsingPacket(&pkt) {
		DebugInfo(2, "Error : Cannot prase a ItemUsing packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : A host send ItemUsing but not in server!")
		return
	}
	dest := GetUserFromID(pkt.UserID)
	if dest == nil ||
		dest.Userid <= 0 {
		DebugInfo(2, "Error : A host send ItemUsing but dest user is null!")
		return
	}
	//找到玩家的房间
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(),
		uPtr.GetUserChannelID(),
		uPtr.GetUserRoomID())
	if rm == nil ||
		rm.Id <= 0 {
		DebugInfo(2, "Error : User", uPtr.UserName, "try to send ItemUsing but in a null room !")
		return
	}
	//是不是房主
	if rm.HostUserID != uPtr.Userid {
		DebugInfo(2, "Error : User", uPtr.UserName, "try to send ItemUsing but isn't host !")
		return
	}
	//发送用户背包数据
	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeHost), BuildItemUsing(pkt.UserID, pkt.ItemID))
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Send User", dest.UserName, "ItemUsed packet to host", uPtr.UserName)
}

func BuildItemUsing(uid uint32, itemid uint32) []byte {
	buf := make([]byte, 10)
	offset := 0
	WriteUint8(&buf, ItemUsing, &offset)
	WriteUint32(&buf, uid, &offset)
	WriteUint32(&buf, itemid, &offset)
	WriteUint8(&buf, 1, &offset)
	return buf[:offset]
}
