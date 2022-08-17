package room

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnRoomRequest(p *PacketData, client net.Conn) {
	var pkt InRoomPaket
	if p.PraseRoomPacket(&pkt) {
		switch pkt.InRoomType {
		case NewRoomRequest:
			OnNewRoom(p, client)
		case JoinRoomRequest:
			OnJoinRoom(p, client)
		case LeaveRoomRequest:
			OnLeaveRoom(client, false)
		case ToggleReadyRequest:
			OnToggleReady(p, client)
		case GameStartRequest:
			OnGameStart(p, client)
		case UpdateSettings:
			OnUpdateRoom(p, client)
		case OnCloseResultWindow:
			OnCloseResultRequest(p, client)
		case SetUserTeamRequest:
			OnChangeTeam(p, client)
		case GameStartCountdownRequest:
			OnGameStartCountdown(p, client)
		case Feedback:
			OnRoomFeedback(p, client)
		default:
			DebugInfo(2, "Unknown room packet", pkt.InRoomType, "from", client.RemoteAddr().String())
		}
	} else {
		DebugInfo(2, "Error : Recived a illegal room packet from", client.RemoteAddr().String())
	}
}
