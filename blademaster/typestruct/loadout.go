package typestruct

type UserLoadout struct {
	Items []uint32
}

func CreateDefaultLoadout() []UserLoadout {
	return []UserLoadout{
		{[]uint32{14, 2, 27, 4, 23, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]uint32{3, 15, 27, 4, 23, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]uint32{101, 15, 80, 4, 23, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
}

func CreateFullLoadout() []UserLoadout {
	return []UserLoadout{
		{[]uint32{5336, 5356, 5330, 4, 23, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]uint32{5285, 5294, 5231, 4, 23, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]uint32{5206, 5356, 5365, 4, 23, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
}
