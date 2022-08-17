package mail

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnMailList(p *PacketData, client net.Conn) {
	//检索数据报
	var pkt InMailListPacket
	if !p.PraseMailListPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a illegal maillist packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to request maillist but not in server !")
		return
	}
	//发送数据
	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeMail), BuildMailList())
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Sent a null mail list to User", uPtr.UserName)

}

func BuildMailList() []byte {
	return []byte{0x00, 0x01}
}
