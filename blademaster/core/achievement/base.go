package achievement

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	campaign = 3
	boss     = 4
)

func OnAchievement(p *PacketData, client net.Conn) {
	var pkt InAchievementPacket
	if p.PraseInAchievementPacket(&pkt) {
		switch pkt.Type {
		case campaign:
			OnAchievementCampaign(p, client)
		default:
			DebugInfo(2, "Unknown achievement packet", pkt.Type, "from", client.RemoteAddr().String())
		}
	} else {
		DebugInfo(2, "Error : Recived a illegal achievement packet from", client.RemoteAddr().String())
	}
}
