package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"innodb_parse/innodb/page"
	"os"
)

var PAGE_SIZE = 16384

func main() {

	var file, _ = os.Open("/Users/Arthur/cvt_dev/open_source/mysql-5.5.59/build_out/data/my_test/t1.ibd")
	fs, _ := file.Stat()
	ibdFileSize := fs.Size()

	println(ibdFileSize)
	pageSize := int(ibdFileSize / int64(PAGE_SIZE))

	for i := 0; i < pageSize; i++ {
		offset := int64(i * PAGE_SIZE)
		file.Seek(offset, 0)
		reader := bufio.NewReader(file)
		bytes := make([]byte, PAGE_SIZE)
		reader.Read(bytes)

		fileReader := page.NewFileReader(bytes)
		page := page.Page{}
		page.Read(fileReader)

		//if page.FilHeader.Page_type == "INDEX" {
		//
		//}
		if page.FilHeader.Page_type == "INDEX" {
			data, _ := json.MarshalIndent(page, "", "\t")
			fmt.Println(string(data))

			//fmt.Printf("%+v\n", page)
		}


		//page := FileReader{reader: bufio.NewReader(innodb)}
		//readFilHeader := page.readFilHeader()
		//fmt.Printf("%+v\n", readFilHeader)
	}

}

//func parseFilHeader(reader io.Reader) fil_header {
//	readFilHeader := fil_header{}
//	readFilHeader.checksum = readUInt32(reader)
//	readFilHeader.offset = readUInt32(reader)
//	readFilHeader.prev = readUInt32(reader)
//	readFilHeader.next = readUInt32(reader)
//	readFilHeader.lsn = readUInt64(reader)
//	readFilHeader.page_type = PAGE_TYPE_MAP[readUInt16(reader)]
//	readFilHeader.flush_lsn = readUInt64(reader)
//	readFilHeader.space_id = readUInt32(reader)
//	return readFilHeader
//}

//func readUInt16(reader io.Reader) uint16 {
//	var result uint16
//	binary.Read(reader, binary.BigEndian, &result)
//	return result
//}
//func readUInt32(reader io.Reader) uint32 {
//	var result uint32
//	binary.Read(reader, binary.BigEndian, &result)
//	return result
//}
//func readUInt64(reader io.Reader) uint64 {
//	var result uint64
//	binary.Read(reader, binary.BigEndian, &result)
//	return result
//}
