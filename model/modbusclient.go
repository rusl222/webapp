package model

import (
	"encoding/binary"
	"net"
	"time"
)

type mbTcpHeader struct {
	num      uint16
	mbcode   uint32
	id       uint8
	funcCode uint8
}

type mbReq struct {
	header mbTcpHeader
	start  uint16
	length uint16
}

type mbAnsv struct {
	header     mbTcpHeader
	dataLength uint8
	data       []byte
}

func (ans *mbAnsv) parse(b []byte) {
	ans.header.num = binary.BigEndian.Uint16(b[0:2])
	ans.header.mbcode = binary.BigEndian.Uint32(b[2:6])
	ans.header.id = b[6]
	ans.header.funcCode = b[7]

	ans.dataLength = b[8]
	ans.data = b[9 : ans.dataLength+9]
}

type ConfigModbusRequest struct {
	Id       uint8
	FuncCode uint8
	Start    uint16
	Length   uint16
}

type ConfigTcp struct {
	Network  string
	IpPort   string
	Requests map[int]ConfigModbusRequest
}

func NewTcpClient(ch chan *[]byte) {

	buf := make([]byte, 300)

	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "172.30.240.1:502")

	var i uint16 = 0
	var req mbReq
	req.header = mbTcpHeader{0, 6, 1, 3}
	var ans mbAnsv
	for {

		req.header.num = i
		req.start = 0
		req.length = 10

		binary.Write(conn, binary.BigEndian, &req)

		conn.Read(buf)
		ans.parse(buf)

		ch <- &ans.data
		i++

		time.Sleep(time.Millisecond * 500)
	}

}
