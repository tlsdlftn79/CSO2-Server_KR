package report

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/core/message"
	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	abusive       = "辱骂诽谤"
	refreshScreen = "聊天刷屏"
	illegalName   = "非法昵称"
	illegalProg   = "非法程序"
	useBug        = "恶意使用游戏漏洞"
	others        = "其他"
)

func OnReportMsg(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InReportMsgPacket
	if !p.PraseReportMsgPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error reportMsg packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to send reportMsg but not in server !")
		return
	}
	if IsExistsIngameName(pkt.Name) {
		switch pkt.Type {
		case 0:
			SaveReport(string(pkt.Name), abusive, pkt.Msg)
		case 1:
			SaveReport(string(pkt.Name), refreshScreen, pkt.Msg)
		case 2:
			SaveReport(string(pkt.Name), illegalName, pkt.Msg)
		case 3:
			SaveReport(string(pkt.Name), illegalProg, pkt.Msg)
		case 4:
			SaveReport(string(pkt.Name), useBug, pkt.Msg)
		default:
			SaveReport(string(pkt.Name), others, pkt.Msg)
		}
	}
	//发送返回数据包
	OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, GAME_REPORT_USER_SUCCEED)
}
