package modbus

import (
	"encoding/binary"
)

type mbTcpHeader struct {
	transaction uint16
	nulls       uint16
	length      uint16
}

type mbReq struct {
	id       uint8
	funcCode uint8
	start    uint16
	length   uint16
}

type mbAnsv struct {
	id         uint8
	funcCode   uint8
	dataLength uint8
	data       []byte
}

type mbTcpReq struct {
	header mbTcpHeader
	req    mbReq
}

type mbRtuReq struct {
	req mbReq
	crc uint16
}

func (ans *mbAnsv) parseTcp(b []byte) {

	ans.id = b[6]
	ans.funcCode = b[7]
	ans.dataLength = b[8]
	ans.data = b[9 : ans.dataLength+9]
}
func (ans *mbAnsv) parseRtu(b []byte) {
	ans.id = b[0]
	ans.funcCode = b[1]
	ans.dataLength = b[2]
	ans.data = b[3 : ans.dataLength+3]
}

func checkRtu(b []byte) bool {
	panic("calcCrcRtu not implemented")
}

func chekTcp(id_trasaction uint16, b []byte) bool {

	if binary.BigEndian.Uint16(b[0:2]) != id_trasaction {
		return false
	}
	if binary.BigEndian.Uint16(b[2:4]) != 0 {
		return false
	}
	return true
}
