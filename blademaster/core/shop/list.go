package shop

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/configure"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

var (
	ShopReply []byte
)

func OnShopList(p *PacketData, client net.Conn) {
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to request shop list but not in server !")
		return
	}
	//发送数据
	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeShop), BuildShopList())
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Send shop list to User", uPtr.UserName)

}
func BuildShopList() []byte {
	if Conf.EnableShop == 0 {
		return []byte{0, 0, 0}
	}
	return ShopReply
}

func InitShopReply() {
	buf := make([]byte, 3)
	offset, optIdx := 0, 0
	WriteUint8(&buf, outshoplist, &offset)
	WriteUint16(&buf, uint16(len(ShopItemList)), &offset)
	for _, item := range ShopItemList {
		tmp := make([]byte, 128)
		offset = 0

		WriteUint32(&tmp, item.ItemID, &offset)
		WriteUint8(&tmp, item.Currency, &offset)
		WriteUint8(&tmp, 1, &offset)            //numopt
		WriteUint32(&tmp, item.ItemID, &offset) //optidx
		WriteUint16(&tmp, 0, &offset)           //quantity
		WriteUint64(&tmp, 0, &offset)           //continue~day
		WriteUint8(&tmp, 0, &offset)
		WriteUint16(&tmp, 1, &offset)
		WriteUint32(&tmp, item.Price, &offset)
		WriteUint32(&tmp, item.Price, &offset)
		WriteUint8(&tmp, 0, &offset) //discount
		WriteUint32(&tmp, 0, &offset)
		WriteUint32(&tmp, 0, &offset)
		WriteUint8(&tmp, 0, &offset) //flags
		WriteUint8(&tmp, 0, &offset)
		WriteUint8(&tmp, 1, &offset)
		WriteUint8(&tmp, 0, &offset)
		WriteUint32(&tmp, 0, &offset)
		WriteUint8(&tmp, 0, &offset)
		WriteUint8(&tmp, 0, &offset)

		optIdx++
		buf = BytesCombine(buf, tmp[:offset])
	}
	ShopReply = buf
}
