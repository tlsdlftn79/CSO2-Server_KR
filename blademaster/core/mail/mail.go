package mail

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	maillist = 0
)

func OnMail(p *PacketData, client net.Conn) {
	var pkt InMailPacket
	if p.PraseMailPacket(&pkt) {
		switch pkt.InMailType {
		case maillist:
			OnMailList(p, client)
		default:
			DebugInfo(2, "Unknown mail packet", pkt.InMailType, "from", client.RemoteAddr().String())
		}
	} else {
		DebugInfo(2, "Error : Recived a illegal mail packet from", client.RemoteAddr().String(), p.Data)
	}
}
