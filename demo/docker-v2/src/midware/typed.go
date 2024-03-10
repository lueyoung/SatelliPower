package main

type Item struct {
	name string
	len  int
}

func make_item(name string, l int) *Item {
	var ret Item
	ret.name = name
	ret.len = l
	return &ret
}

func make_item60() []*Item {
	var Item60 []*Item

	Item60 = append(Item60, make_item("apid_0060", 16))
	Item60 = append(Item60, make_item("packetSeqId_0060", 2))
	Item60 = append(Item60, make_item("packetCount_0060", 14))
	Item60 = append(Item60, make_item("packetLength_0060", 16))
	Item60 = append(Item60, make_item("TMN001_1", 16))
	Item60 = append(Item60, make_item("TMN002_1", 16))
	Item60 = append(Item60, make_item("TMN003_1", 16))
	Item60 = append(Item60, make_item("TMN004_1", 16))
	Item60 = append(Item60, make_item("TMN005_1", 16))
	Item60 = append(Item60, make_item("TMN006", 16))
	Item60 = append(Item60, make_item("TMN007", 16))
	Item60 = append(Item60, make_item("TMN008", 16))
	Item60 = append(Item60, make_item("TMN009", 1))
	Item60 = append(Item60, make_item("TMN010", 1))
	Item60 = append(Item60, make_item("TMN011", 1))
	Item60 = append(Item60, make_item("TMN012", 1))
	Item60 = append(Item60, make_item("TMN013", 4))
	Item60 = append(Item60, make_item("TMN014", 8))
	Item60 = append(Item60, make_item("TMN015", 8))
	Item60 = append(Item60, make_item("TMN016", 8))
	Item60 = append(Item60, make_item("TMN017", 8))
	Item60 = append(Item60, make_item("TMN018", 8))
	Item60 = append(Item60, make_item("TMN019", 8))
	Item60 = append(Item60, make_item("TMN020", 1))
	Item60 = append(Item60, make_item("TMN021", 1))
	Item60 = append(Item60, make_item("TMN022", 1))
	Item60 = append(Item60, make_item("TMN023", 1))
	Item60 = append(Item60, make_item("TMN024", 1))
	Item60 = append(Item60, make_item("TMN025", 1))
	Item60 = append(Item60, make_item("TMN026", 2))
	Item60 = append(Item60, make_item("TMN027", 8))
	Item60 = append(Item60, make_item("TMN028", 8))
	Item60 = append(Item60, make_item("TMN029", 8))
	Item60 = append(Item60, make_item("TMN030", 8))
	Item60 = append(Item60, make_item("TMN031", 8))

	return Item60
}
