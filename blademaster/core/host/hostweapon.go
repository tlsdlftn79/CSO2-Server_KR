package host

import (
	"log"
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnHostWeaponPoint(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InWeaponPointPacket
	if !p.PraseInWeaponPointPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error WeaponPoint packet !")
		return
	}
	if pkt.KillerID <= 0 || pkt.KillerWeaponID <= 0 {
		return
	}
	//找到对应用户
	uPtr := GetUserFromID(pkt.KillerID)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		log.Println("Error : Client from", client.RemoteAddr().String(), "sent WeaponPoint but killer not in server !")
		return
	}
	//修改用户数据
	uPtr.CountWeaponKill(pkt.KillerWeaponID)
}
