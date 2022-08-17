package typestruct

import (
	"unsafe"

	. "github.com/KouKouChan/CSO2-Server/kerlong"
)

//每个房间的设置数据
type RoomSettings struct {
	LenOfName          uint8
	RoomName           []byte
	Unk00              uint8
	Unk01              uint8
	Unk02              uint32
	Unk03              uint32
	LenOfPassWd        uint8
	PassWd             []byte
	Unk10              uint16
	ForceCamera        uint8
	GameModeID         uint8
	MapID              uint8
	Unk13              uint8
	MaxPlayers         uint8
	WinLimit           uint8
	KillLimit          uint16
	Unk17              uint8
	Unk18              uint8
	WeaponRestrictions uint8
	Status             uint8
	Unk21              uint8
	MapCycleType       uint8
	Unk23              uint8
	Unk24              uint8
	Unk25              uint8
	LenOfMultiMaps     uint8
	MultiMaps          []byte
	TeamBalanceType    uint8
	Unk29              uint8
	Unk30              uint8
	Unk31              uint8
	Unk32              uint8
	Unk33              uint8
	AreBotsEnabled     uint8
	BotDifficulty      uint8
	NumCtBots          uint8
	NumTrBots          uint8
	Unk35              uint8
	Unk36              uint8
	Unk37              uint8
	Unk38              uint8
	Unk39              uint8
	StartMoney         uint16
	ChangeTeams        uint8
	Unk43              uint8
	HltvEnabled        uint8
	Unk45              uint8
	RespawnTime        uint8
	NextMapEnabled     uint8
	Difficulty         uint8
	IsIngame           uint8
}

func (dest *Room) ToUpdateSetting(src *InUpSettingReq) {
	flags := src.Flags
	lowFlag := *(*uint32)(unsafe.Pointer(&flags))
	//右移32比特位
	flags = flags >> 32
	highFlag := *(*uint32)(unsafe.Pointer(&flags))
	if lowFlag&0x1 != 0 {
		dest.Setting.LenOfName = src.LenOfRoomName
		dest.Setting.RoomName = src.RoomName
	}
	if lowFlag&0x2 != 0 {
		dest.Setting.Unk00 = src.Unk00
	}
	if lowFlag&0x4 != 0 {
		dest.Setting.Unk01 = src.Unk01
		dest.Setting.Unk02 = src.Unk02
		dest.Setting.Unk03 = src.Unk03
	}
	if lowFlag&0x8 != 0 {
		dest.Setting.LenOfPassWd = src.LenOfpasswd
		dest.Setting.PassWd = src.Passwd
		if dest.Setting.LenOfPassWd > 0 {
			dest.PasswordProtected = 1
		} else {
			dest.PasswordProtected = 0
		}
	}
	if lowFlag&0x10 != 0 {
		dest.Setting.Unk10 = src.Unk10
	}
	if lowFlag&0x20 != 0 {
		dest.Setting.ForceCamera = src.ForceCamera
	}
	if lowFlag&0x40 != 0 {
		dest.Setting.GameModeID = src.GameModeID
	}
	if lowFlag&0x80 != 0 {
		dest.Setting.MapID = src.MapID
		dest.Setting.Unk13 = src.Unk13
	}
	if lowFlag&0x100 != 0 {
		dest.Setting.MaxPlayers = src.MaxPlayers
	}
	if lowFlag&0x200 != 0 {
		dest.Setting.WinLimit = src.WinLimit
	}
	if lowFlag&0x400 != 0 {
		dest.Setting.KillLimit = src.KillLimit
	}
	if lowFlag&0x800 != 0 {
		dest.Setting.Unk17 = src.Unk17
	}
	if lowFlag&0x1000 != 0 {
		dest.Setting.Unk18 = src.Unk18
	}
	if lowFlag&0x2000 != 0 {
		dest.Setting.WeaponRestrictions = src.WeaponRestrictions
	}
	if lowFlag&0x4000 != 0 {
		dest.Setting.Status = src.Status
	}
	if lowFlag&0x8000 != 0 {
		dest.Setting.Unk21 = src.Unk21
		dest.Setting.MapCycleType = src.MapCycleType
		dest.Setting.Unk23 = src.Unk23
		dest.Setting.Unk24 = src.Unk24
	}
	if lowFlag&0x10000 != 0 {
		dest.Setting.Unk25 = src.Unk25
	}
	if lowFlag&0x20000 != 0 {
		dest.Setting.LenOfMultiMaps = src.NumOfMultiMaps
		dest.Setting.MultiMaps = make([]byte, src.NumOfMultiMaps)
		for i := 0; i < int(dest.Setting.LenOfMultiMaps); i++ {
			dest.Setting.MultiMaps[i] = src.MultiMaps[i]
		}
	}
	if lowFlag&0x40000 != 0 {
		dest.Setting.TeamBalanceType = src.TeamBalanceType
	}
	if lowFlag&0x80000 != 0 {
		dest.Setting.Unk29 = src.Unk29
	}
	if lowFlag&0x100000 != 0 {
		dest.Setting.Unk30 = src.Unk30
	}
	if lowFlag&0x200000 != 0 {
		dest.Setting.Unk31 = src.Unk31
	}
	if lowFlag&0x400000 != 0 {
		dest.Setting.Unk32 = src.Unk32
	}
	if lowFlag&0x800000 != 0 {
		dest.Setting.Unk33 = src.Unk33
	}
	if lowFlag&0x1000000 != 0 {
		dest.Setting.AreBotsEnabled = src.BotEnabled
		dest.Setting.BotDifficulty = src.BotDifficulty
		dest.Setting.NumCtBots = src.NumCtBots
		dest.Setting.NumTrBots = src.NumTrBots
	}

	if lowFlag&0x2000000 != 0 {
		dest.Setting.Unk35 = src.Unk35
	}

	if lowFlag&0x4000000 != 0 {
		dest.Setting.Unk36 = src.Unk36
	}

	if lowFlag&0x8000000 != 0 {
		dest.Setting.Unk37 = src.Unk37
	}

	if lowFlag&0x10000000 != 0 {
		dest.Setting.Unk38 = src.Unk38
	}

	if lowFlag&0x20000000 != 0 {
		dest.Setting.Unk39 = src.Unk39
	}

	if lowFlag&0x40000000 != 0 {
		dest.Setting.IsIngame = src.IsIngame
	}

	if lowFlag&0x80000000 != 0 {
		dest.Setting.StartMoney = src.StartMoney
	}

	if highFlag&0x1 != 0 {
		dest.Setting.ChangeTeams = src.ChangeTeams
	}

	if highFlag&0x2 != 0 {
		dest.Setting.Unk43 = src.Unk43
	}

	if highFlag&0x4 != 0 {
		dest.Setting.HltvEnabled = src.HltvEnabled
	}

	if highFlag&0x8 != 0 {
		dest.Setting.Unk45 = src.Unk45
	}

	if highFlag&0x10 != 0 {
		dest.Setting.RespawnTime = src.RespawnTime
	}
}

//创建房间设置数据包
func BuildRoomSetting(room *Room, flags uint64) []byte {
	buf := make([]byte, 128+room.Setting.LenOfName+ //实际计算是最大63字节+长度
		room.Setting.LenOfPassWd+
		room.Setting.LenOfMultiMaps)
	offset := 0
	WriteUint8(&buf, OUTUpdateSettings, &offset)
	WriteUint64(&buf, flags, &offset)
	lowFlag := *(*uint32)(unsafe.Pointer(&flags))
	flags = flags >> 32
	highFlag := *(*uint32)(unsafe.Pointer(&flags))
	if lowFlag&0x1 != 0 {
		WriteString(&buf, room.Setting.RoomName, &offset)
	}
	if lowFlag&0x2 != 0 {
		WriteUint8(&buf, room.Setting.Unk00, &offset)
	}
	if lowFlag&0x4 != 0 {
		WriteUint8(&buf, room.Setting.Unk01, &offset)
		WriteUint32(&buf, room.Setting.Unk02, &offset)
		WriteUint32(&buf, room.Setting.Unk03, &offset)
	}
	if lowFlag&0x8 != 0 {
		WriteString(&buf, room.Setting.PassWd, &offset)
	}
	if lowFlag&0x10 != 0 {
		WriteUint16(&buf, room.Setting.Unk10, &offset)
	}
	if lowFlag&0x20 != 0 {
		WriteUint8(&buf, room.Setting.ForceCamera, &offset)
	}
	if lowFlag&0x40 != 0 {
		WriteUint8(&buf, room.Setting.GameModeID, &offset)
	}
	if lowFlag&0x80 != 0 {
		WriteUint8(&buf, room.Setting.MapID, &offset)
		WriteUint8(&buf, room.Setting.Unk13, &offset)
	}
	if lowFlag&0x100 != 0 {
		WriteUint8(&buf, room.Setting.MaxPlayers, &offset)
	}
	if lowFlag&0x200 != 0 {
		WriteUint8(&buf, room.Setting.WinLimit, &offset)
	}
	if lowFlag&0x400 != 0 {
		WriteUint16(&buf, room.Setting.KillLimit, &offset)
	}
	if lowFlag&0x800 != 0 {
		WriteUint8(&buf, room.Setting.Unk17, &offset)
	}
	if lowFlag&0x1000 != 0 {
		WriteUint8(&buf, room.Setting.Unk18, &offset)
	}
	if lowFlag&0x2000 != 0 {
		WriteUint8(&buf, room.Setting.WeaponRestrictions, &offset)
	}
	if lowFlag&0x4000 != 0 {
		WriteUint8(&buf, room.Setting.Status, &offset)
	}
	if lowFlag&0x8000 != 0 {
		WriteUint8(&buf, room.Setting.Unk21, &offset)
		WriteUint8(&buf, room.Setting.MapCycleType, &offset)
		WriteUint8(&buf, room.Setting.Unk23, &offset)
		WriteUint8(&buf, room.Setting.Unk24, &offset)
	}
	if lowFlag&0x10000 != 0 {
		WriteUint8(&buf, room.Setting.Unk21, &offset)
	}
	if lowFlag&0x20000 != 0 {
		WriteUint8(&buf, room.Setting.LenOfMultiMaps, &offset)
		for _, v := range room.Setting.MultiMaps {
			WriteUint8(&buf, v, &offset)
		}
	}
	if lowFlag&0x40000 != 0 {
		WriteUint8(&buf, room.Setting.TeamBalanceType, &offset)
	}
	if lowFlag&0x80000 != 0 {
		WriteUint8(&buf, room.Setting.Unk29, &offset)
	}
	if lowFlag&0x100000 != 0 {
		WriteUint8(&buf, room.Setting.Unk30, &offset)
	}
	if lowFlag&0x200000 != 0 {
		WriteUint8(&buf, room.Setting.Unk31, &offset)
	}
	if lowFlag&0x400000 != 0 {
		WriteUint8(&buf, room.Setting.Unk32, &offset)
	}
	if lowFlag&0x800000 != 0 {
		WriteUint8(&buf, room.Setting.Unk33, &offset)
	}
	if lowFlag&0x1000000 != 0 {
		WriteUint8(&buf, room.Setting.AreBotsEnabled, &offset)
		if room.Setting.AreBotsEnabled != 0 {
			WriteUint8(&buf, room.Setting.BotDifficulty, &offset)
			WriteUint8(&buf, room.Setting.NumCtBots, &offset)
			WriteUint8(&buf, room.Setting.NumTrBots, &offset)
		}
	}

	if lowFlag&0x2000000 != 0 {
		WriteUint8(&buf, room.Setting.Unk35, &offset)
	}

	if lowFlag&0x4000000 != 0 {
		WriteUint8(&buf, room.Setting.Unk36, &offset)
	}

	if lowFlag&0x8000000 != 0 {
		WriteUint8(&buf, room.Setting.Unk37, &offset)
	}

	if lowFlag&0x10000000 != 0 {
		WriteUint8(&buf, room.Setting.Unk38, &offset)
	}

	if lowFlag&0x20000000 != 0 {
		WriteUint8(&buf, room.Setting.Unk39, &offset)
	}

	if lowFlag&0x40000000 != 0 {
		WriteUint8(&buf, room.Setting.IsIngame, &offset)
	}

	if lowFlag&0x80000000 != 0 {
		WriteUint16(&buf, room.Setting.StartMoney, &offset)
	}

	if highFlag&0x1 != 0 {
		WriteUint8(&buf, room.Setting.ChangeTeams, &offset)
	}

	if highFlag&0x2 != 0 {
		WriteUint8(&buf, room.Setting.Unk43, &offset)
	}

	if highFlag&0x4 != 0 {
		WriteUint8(&buf, room.Setting.HltvEnabled, &offset)
	}

	if highFlag&0x8 != 0 {
		WriteUint8(&buf, room.Setting.Unk45, &offset)
	}

	if highFlag&0x10 != 0 {
		WriteUint8(&buf, room.Setting.RespawnTime, &offset)
	}
	return buf[:offset]
}

func GetFlags(room Room) uint64 {
	lowFlag := 0
	highFlag := 0

	if room.Setting.RoomName != nil {
		lowFlag |= 0x1
	}
	if room.Setting.Unk00 != 0 {
		lowFlag |= 0x2
	}
	if room.Setting.Unk01 != 0 &&
		room.Setting.Unk02 != 0 &&
		room.Setting.Unk03 != 0 {
		lowFlag |= 0x4
	}
	if room.Setting.PassWd != nil {
		lowFlag |= 0x8
	}
	if room.Setting.Unk10 != 0 {
		lowFlag |= 0x10
	}
	if room.Setting.ForceCamera != 0 {
		lowFlag |= 0x20
	}
	if room.Setting.GameModeID != 0 {
		lowFlag |= 0x40
	}
	if room.Setting.MapID != 0 && room.Setting.Unk13 != 0 {
		lowFlag |= 0x80
	}
	if room.Setting.MaxPlayers != 0 {
		lowFlag |= 0x100
	}
	if room.Setting.WinLimit != 0 {
		lowFlag |= 0x200
	}
	if room.Setting.KillLimit != 0 {
		lowFlag |= 0x400
	}
	if room.Setting.Unk17 != 0 {
		lowFlag |= 0x800
	}
	if room.Setting.Unk18 != 0 {
		lowFlag |= 0x1000
	}
	if room.Setting.WeaponRestrictions != 0 {
		lowFlag |= 0x2000
	}
	if room.Setting.Status != 0 {
		lowFlag |= 0x4000
	}
	if room.Setting.Unk21 != 0 &&
		room.Setting.MapCycleType != 0 &&
		room.Setting.Unk23 != 0 &&
		room.Setting.Unk24 != 0 {
		lowFlag |= 0x8000
	}
	if room.Setting.Unk25 != 0 {
		lowFlag |= 0x10000
	}
	if room.Setting.MultiMaps != nil {
		lowFlag |= 0x20000
	}
	if room.Setting.TeamBalanceType != 0 {
		lowFlag |= 0x40000
	}
	if room.Setting.Unk29 != 0 {
		lowFlag |= 0x80000
	}
	if room.Setting.Unk30 != 0 {
		lowFlag |= 0x100000
	}
	if room.Setting.Unk31 != 0 {
		lowFlag |= 0x200000
	}
	if room.Setting.Unk32 != 0 {
		lowFlag |= 0x400000
	}
	if room.Setting.Unk33 != 0 {
		lowFlag |= 0x800000
	}
	if room.Setting.AreBotsEnabled != 0 {
		lowFlag |= 0x1000000
	}

	if room.Setting.Unk35 != 0 {
		lowFlag |= 0x2000000
	}

	if room.Setting.Unk36 != 0 {
		lowFlag |= 0x4000000
	}

	if room.Setting.Unk37 != 0 {
		lowFlag |= 0x8000000
	}

	if room.Setting.Unk38 != 0 {
		lowFlag |= 0x10000000
	}

	if room.Setting.Unk39 != 0 {
		lowFlag |= 0x20000000
	}

	if room.Setting.IsIngame != 0 {
		lowFlag |= 0x40000000
	}

	if room.Setting.StartMoney != 0 {
		lowFlag |= 0x80000000
	}

	if room.Setting.ChangeTeams != 0 {
		highFlag |= 0x1
	}

	if room.Setting.Unk43 != 0 {
		highFlag |= 0x2
	}

	if room.Setting.HltvEnabled != 0 {
		highFlag |= 0x4
	}

	if room.Setting.Unk45 != 0 {
		highFlag |= 0x8
	}

	if room.Setting.RespawnTime != 0 {
		highFlag |= 0x10
	}

	flags := uint64(highFlag)
	flags = flags << 32
	return flags + uint64(lowFlag)
}
