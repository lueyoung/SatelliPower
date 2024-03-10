package main

type Item struct {
	key   string
	len   int
	value string
}

func make_item(key string, l int) *Item {
	var ret Item
	ret.key = key
	ret.len = l
	return &ret
}

func make_item60() []*Item {
	var Item60 []*Item

	Item60 = append(Item60, make_item("apid_0060", 11)) // only lower 4 bits are random
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

type Frame60 struct {
	apid_0060         int `json:apid_0060,string`
	packetSeqId_0060  int `json:packetSeqId_0060,string`
	packetCount_0060  int `json:packetCount_0060,string`
	packetLength_0060 int `json:packetLength_0060,string`
	TMN001_1          int `json:TMN001_1,string`
	TMN002_1          int `json:TMN002_1,string`
	TMN003_1          int `json:TMN003_1,string`
	TMN004_1          int `json:TMN004_1,string`
	TMN005_1          int `json:TMN005_1,string`
	TMN006            int `json:TMN006,string`
	TMN007            int `json:TMN007,string`
	TMN008            int `json:TMN008,string`
	TMN009            int `json:TMN009,string`
	TMN010            int `json:TMN010,string`
	TMN011            int `json:TMN011,string`
	TMN012            int `json:TMN012,string`
	TMN013            int `json:TMN013,string`
	TMN014            int `json:TMN014,string`
	TMN015            int `json:TMN015,string`
	TMN016            int `json:TMN016,string`
	TMN017            int `json:TMN017,string`
	TMN018            int `json:TMN018,string`
	TMN019            int `json:TMN019,string`
	TMN020            int `json:TMN020,string`
	TMN021            int `json:TMN021,string`
	TMN022            int `json:TMN022,string`
	TMN023            int `json:TMN023,string`
	TMN024            int `json:TMN024,string`
	TMN025            int `json:TMN025,string`
	TMN026            int `json:TMN026,string`
	TMN027            int `json:TMN027,string`
	TMN028            int `json:TMN028,string`
	TMN029            int `json:TMN029,string`
	TMN030            int `json:TMN030,string`
	TMN031            int `json:TMN031,string`
}
