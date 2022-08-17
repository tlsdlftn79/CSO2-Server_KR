package supply

import (
	"math/rand"
	"net"
	"time"

	. "github.com/KouKouChan/CSO2-Server/blademaster/core/inventory"
	. "github.com/KouKouChan/CSO2-Server/blademaster/core/message"
	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/configure"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnSupplyOpenBox(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InOpenBoxPacket
	if !p.PraseOpenBoxPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error OpenBox packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to request openbox but not in server !")
		return
	}
	//发送数据
	if !uPtr.DecreaseItem(pkt.BoxID) {
		OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_UI_RAMDOMBOX_ALERT_000)
		DebugInfo(2, "User", uPtr.UserName, "open box", BoxList[pkt.BoxID].BoxName, "failed")
		return
	}

	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeInventory_Create),
		BuildInventoryInfoSingle(uPtr, pkt.BoxID))
	SendPacket(rst, uPtr.CurrentConnection)

	itemid := GetBoxItem(pkt.BoxID)
	uPtr.AddItem(itemid)

	rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeSupply),
		BuildSupplyOpenBox(itemid))
	SendPacket(rst, uPtr.CurrentConnection)

	rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeInventory_Create),
		BuildInventoryInfoSingle(uPtr, itemid))
	SendPacket(rst, uPtr.CurrentConnection)

	DebugInfo(2, "User", uPtr.UserName, "got item", ItemList[itemid].Name, "id", itemid, "by openning box", BoxList[pkt.BoxID].BoxName)

}

func BuildSupplyOpenBox(itemid uint32) []byte {
	buf := make([]byte, 25)
	offset := 0
	WriteUint8(&buf, openbox, &offset)
	WriteUint8(&buf, 1, &offset)
	WriteUint8(&buf, 0, &offset)
	WriteUint32(&buf, itemid, &offset) //itemid
	WriteUint16(&buf, 0, &offset)      //item count
	WriteUint16(&buf, 0, &offset)      //day
	WriteUint32(&buf, 0, &offset)      //mpoint
	return buf[:offset]
}

func GetBoxItem(boxid uint32) uint32 {
	if v, ok := BoxList[boxid]; ok {
		rand.Seed(time.Now().UnixNano())
		radV := rand.Intn(v.TotalValue)
		for _, item := range v.Items {
			radV -= item.Value
			if radV < 0 {
				return item.ItemID
			}
		}
	}
	DebugInfo(2, "Error : can't open box", boxid)
	return boxid
}
