package host

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnHostBuyItem(p *PacketData, client net.Conn) {
	//检查数据包
	var pkt InHostBuyItemPacket
	if !p.PraseInHostBuyItemPacket(&pkt) {
		DebugInfo(2, "Error : Cannot prase a BuyItem packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromID(pkt.UserID)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : A host send BuyItem but not in server!")
		return
	}
	for _, v := range pkt.Items {
		DebugInfo(2, "User", uPtr.UserName, "bought", v)
	}
}
