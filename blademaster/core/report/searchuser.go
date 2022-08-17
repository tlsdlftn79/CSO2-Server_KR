package report

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	notfound = 0
	found    = 1
)

func OnReportSearchUser(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InReportSearchUserPacket
	if !p.PraseReportSearchUserPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error SearchUser packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to searchReportUser but not in server !")
		return
	}
	//发送返回数据包
	var rst []byte
	if IsExistsIngameName(pkt.Name) {
		rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeReport), BuildSearchResult(found))
		SendPacket(rst, uPtr.CurrentConnection)
		DebugInfo(2, "Send report-SearchResult-Found of destUser", string(pkt.Name), " to User", uPtr.UserName)
		return
	}
	rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeReport), BuildSearchResult(notfound))
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Send report-SearchResult-NotFound of destUser", string(pkt.Name), " to User", uPtr.UserName)

}

func BuildSearchResult(isfound uint8) []byte {
	buf := make([]byte, 3)
	offset := 0
	WriteUint8(&buf, reportUser, &offset)
	WriteUint8(&buf, searchUser, &offset)
	WriteUint8(&buf, isfound, &offset)
	return buf
}
