package useitem

import (
	"math/rand"
	"net"
	"time"

	. "github.com/KouKouChan/CSO2-Server/blademaster/core/inventory"
	. "github.com/KouKouChan/CSO2-Server/blademaster/core/message"
	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

const (
	lotto_base       = 1000 //银币
	lotto_max        = 9000
	lotto_event_base = 100 //铜币
	lotto_event_max  = 12900
	lotto_gold_base  = 10000 //金币
	lotto_gold_max   = 10000
)

func OnPointLottoUse(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InPointLottoUsePacket
	if !p.PrasePointLottoUsePacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error pointlottouse packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to request use pointlotto but not in server !")
		return
	}
	//发送数据
	lottoID := uPtr.GetItemIDBySeq(pkt.ItemSeq)
	switch lottoID {
	case 2008: //银币 id 2008
	case 2013: //铜币
	case 2014: //金币
	default:
		DebugInfo(2, "User", uPtr.UserName, "try using pointlotto but itemid is", lottoID)
		return
	}
	if !uPtr.DecreaseItem(lottoID) {
		OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_UI_RAMDOMBOX_ALERT_000)
		DebugInfo(2, "User", uPtr.UserName, "use pointlotto", lottoID, "failed")
		return
	}

	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeInventory_Create),
		BuildInventoryInfoSingle(uPtr, lottoID))
	SendPacket(rst, uPtr.CurrentConnection)

	point := UsePointLotto(lottoID)
	uPtr.GetPoints(point)

	rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeUseItem),
		buildUsePoint(uint32(point)))
	SendPacket(rst, uPtr.CurrentConnection)

	//UserInfo部分
	rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeUserInfo), BuildUserInfo(0XFFFFFFFF, NewUserInfo(uPtr), uPtr.Userid, true))
	SendPacket(rst, uPtr.CurrentConnection)

	DebugInfo(2, "User", uPtr.UserName, "got point", point, "by using pointlotto", lottoID)

}

func buildUsePoint(point uint32) []byte {
	buf := make([]byte, 25)
	offset := 0
	WriteUint8(&buf, usepoint, &offset)
	WriteUint8(&buf, 5, &offset)
	WriteUint8(&buf, 1, &offset)      //unk00
	WriteUint32(&buf, 0, &offset)     //unk01
	WriteUint32(&buf, point, &offset) //mpoint
	return buf[:offset]
}

func UsePointLotto(itemid uint32) uint64 {
	rand.Seed(time.Now().UnixNano())
	switch itemid {
	case 2008: //银币 id 2008
		return uint64(lotto_base + rand.Intn(lotto_max))
	case 2013: //铜币
		return uint64(lotto_event_base + rand.Intn(lotto_event_max))
	case 2014: //金币
		return uint64(lotto_gold_base + rand.Intn(lotto_gold_max))
	default:
		return 0
	}
}
