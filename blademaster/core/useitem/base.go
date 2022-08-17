package useitem

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	useitem  = 1
	usepoint = 2
)

func OnUseItem(p *PacketData, client net.Conn) {
	var pkt InPointLottoPacket
	if p.PrasePointLottoPacket(&pkt) {
		switch pkt.Type {
		case usepoint:
			OnPointLottoUse(p, client)
		default:
			DebugInfo(2, "Unknown useitem packet", pkt.Type, "from", client.RemoteAddr().String(), p.Data)
		}
	} else {
		DebugInfo(2, "Error : Recived a illegal useitem packet from", client.RemoteAddr().String())
	}
}
