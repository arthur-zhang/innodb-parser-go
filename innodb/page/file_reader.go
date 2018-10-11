package page

import (
	"encoding/binary"
)

type FileReader struct {
	data        []byte
	cursorIndex uint32
}

func (self *FileReader) seek(n uint32) {
	self.cursorIndex = n
}

func NewFileReader(data []byte) *FileReader {
	return &FileReader{data, 0}
}

func (self *FileReader) readUint8() uint8 {
	val := self.data[self.cursorIndex]
	self.cursorIndex += 1
	return val
}

// u2
func (self *FileReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data[self.cursorIndex:])
	self.cursorIndex += 2
	return val
}

// u4
func (self *FileReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data[self.cursorIndex:])
	self.cursorIndex += 4
	return val
}

func (self *FileReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data[self.cursorIndex:])
	self.cursorIndex += 8
	return val
}
