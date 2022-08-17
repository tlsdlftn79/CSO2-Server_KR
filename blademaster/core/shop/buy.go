package shop

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/core/inventory"
	. "github.com/KouKouChan/CSO2-Server/blademaster/core/message"
	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/configure"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnShopBuyItem(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InShopBuyItemPacket
	if !p.PraseShopBuyItemPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error buyitem packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to request buyitem but not in server !")
		return
	}
	//找到购买的物品
	for _, item := range ShopItemList {
		if item.ItemID == pkt.ItemID {
			//发送数据
			switch item.Currency {
			case 0: //credit
				if !uPtr.UseCredits(item.Price) {
					OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_BUY_FAIL_CN_0X11_NO_CASH)
					DebugInfo(2, "User", uPtr.UserName, "try to buy item", pkt.ItemID, "but not enough cash")
					return
				}
			case 1: //point
				if !uPtr.UsePoints(uint64(item.Price)) {
					OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_BUY_FAIL_NO_POINT)
					DebugInfo(2, "User", uPtr.UserName, "try to buy item", pkt.ItemID, "but not enough points")
					return
				}
			case 2: //mpoint
				if !uPtr.UseMPoints(item.Price) {
					OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_BUY_FAIL_NO_MILEAGE)
					DebugInfo(2, "User", uPtr.UserName, "try to buy item", pkt.ItemID, "but not enough mpoints")
					return
				}
			default:
				OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_BUY_FAIL_CN_0X12_DB_ERROR)
				DebugInfo(2, "User", uPtr.UserName, "try to buy item", pkt.ItemID, "but unkown currency", item.Currency)
				return
			}
			uPtr.AddItem(pkt.ItemID)
			OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_BUY_SUCCEED)
			rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeInventory_Create),
				BuildInventoryInfoSingle(uPtr, pkt.ItemID))
			SendPacket(rst, uPtr.CurrentConnection)
			//UserInfo部分
			rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeUserInfo), BuildUserInfo(0XFFFFFFFF, NewUserInfo(uPtr), uPtr.Userid, true))
			SendPacket(rst, uPtr.CurrentConnection)
			DebugInfo(2, "User", uPtr.UserName, "bought item", pkt.ItemID)
			return
		}
	}
	//未找到物品
	OnSendMessage(uPtr.CurrentSequence, uPtr.CurrentConnection, MessageDialogBox, CSO2_BUY_FAIL_CN_0X12_DB_ERROR)
	DebugInfo(2, "User", uPtr.UserName, "try to buy item", pkt.ItemID, "but failed")
}
