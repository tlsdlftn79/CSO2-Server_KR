package host

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	users = 1
	cache = 2
)

func OnHostDataPacket(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InHostDataPacket
	if !p.PraseInHostDataPacket(&pkt) {
		DebugInfo(2, "Error : Recived a illegal hostData packet from", client.RemoteAddr().String())
		return
	}
	switch pkt.DataType {
	case users:
		OnHostDataUsersPacket(p, client)
	case cache:
		OnHostDataCachePacket(p, client)
	default:
		DebugInfo(2, "Unknown hostData packet", pkt.DataType, "from", client.RemoteAddr().String(), p.Data)
	}
}
