package page

var PAGE_TYPE_MAP = map[uint16]string{
	0:     "ALLOCATED",      //# Freshly allocated page
	2:     "UNDO_LOG",       //# Undo log page
	3:     "INODE",          //# Index node
	4:     "IBUF_FREE_LIST", //# Insert buffer Free list
	5:     "IBUF_BITMAP",    //# Insert buffer bitmap
	6:     "SYS",            // # System page
	7:     "TRX_SYS",        //# Transaction system data
	8:     "FSP_HDR",        //# File space header
	9:     "XDES",           //# Extent descriptor page
	10:    "BLOB",           //# Uncompressed BLOB page
	11:    "ZBLOB",          //# First compressed BLOB page
	12:    "ZBLOB2",         //# Subsequent compressed BLOB page
	17855: "INDEX",          //# B-tree node
}

type fil_header struct {
	Checksum  uint32
	Offset    uint32
	Prev      uint32
	Next      uint32
	Lsn       uint64
	Page_type string
	Flush_lsn uint64
	Space_id  uint32
}

type page_header struct {
	N_dir_slots uint16
	Heap_top    uint16
	N_heap      uint16
	Free        uint16
	Garbage     uint16
	Last_insert uint16
	Direction   string
	N_direction uint16
	N_recs      uint16
	Max_trx_id  uint64
	Level       uint16
	Index_id    uint64
	Format      string
}

type Page struct {
	FilHeader  fil_header
	PageHeader page_header
}



func (self *Page) Read(reader *FileReader) {
	self.readFilHeader(reader)
	if self.FilHeader.Page_type == "INDEX" {
		self.readPageHeader(reader)
	}
}

var FIL_HEADER_START = 0
var FIL_HEADER_SIZE = 38
var PAGE_HEADER_START = uint32(FIL_HEADER_START + FIL_HEADER_SIZE)

var PAGE_DIRECTION = map[uint16]string{
	1: "left",
	2: "right",
	3: "same_rec",
	4: "same_page",
	5: "no_direction",
}

func (self *Page) readPageHeader(reader *FileReader) {
	reader.seek(PAGE_HEADER_START)
	self.PageHeader = page_header{
		N_dir_slots: reader.readUint16(),
		Heap_top:    reader.readUint16(),
		N_heap:      reader.readUint16(),
		Free:        reader.readUint16(),
		Garbage:     reader.readUint16(),
		Last_insert: reader.readUint16(),
		Direction:   PAGE_DIRECTION[reader.readUint16()],
		N_direction: reader.readUint16(),
		N_recs:      reader.readUint16(),
		Max_trx_id:  reader.readUint64(),
		Level:       reader.readUint16(),
		Index_id:    reader.readUint64(),
		//Format:(N_heap & 1<<15) == 0 ? :redundant : :compact
		Format: "compact",

	}
}
func (self *Page) readFilHeader(reader *FileReader) {
	reader.seek(0)
	self.FilHeader = fil_header{
		Checksum:  reader.readUint32(),
		Offset:    reader.readUint32(),
		Prev:      reader.readUint32(),
		Next:      reader.readUint32(),
		Lsn:       reader.readUint64(),
		Page_type: PAGE_TYPE_MAP[reader.readUint16()],
		Flush_lsn: reader.readUint64(),
		Space_id:  reader.readUint32(),
	}
}
