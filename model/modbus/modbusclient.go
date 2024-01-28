package modbus

import (
	"encoding/binary"
	"errors"
	"net"
	"time"
)

type ModbusClient interface {
	Run() error
	Stop() error
}

type ModbusData struct {
	Request int
	Data    *[]byte
}

type modbusClient struct {
	Conf       Config
	conn       net.Conn
	dataOut    chan ModbusData
	readBuffer []byte
}

func New(conf *Config, dataOut chan ModbusData) ModbusClient {
	if conf == nil {
		conf = testConf()
	}
	return modbusClient{
		Conf:    *conf,
		dataOut: dataOut,
	}
}

// Stop implements ModbusTcpClient.
func (modbusClient) Stop() error {
	return errors.New("unimplemented")
}

/*
func NewTcpClient(ch chan ModbusData) {

	buf := make([]byte, 300)

	// Подключаемся к сокету
	conn, err := net.Dial("tcp", "172.30.240.1:502")
	if err != nil {
		return
	}

	var i uint16 = 0
	var req = mbTcpReq{mbTcpHeader{i, 0, 6}, mbReq{1, 3, 0, 10}}
	var ans mbAnsv
	for {

		binary.Write(conn, binary.BigEndian, &req)

		conn.Read(buf)
		ans.parseTcp(buf)

		ch <- ModbusData{0, &ans.data}
		i++

		time.Sleep(time.Millisecond * 500)
	}

}
*/

// Run implements ModbusTcpClient.
func (m modbusClient) Run() error {

	// connect
	switch m.Conf.Network {
	case "tcp":
		conn, err := net.Dial("tcp", m.Conf.IpPort)
		if err != nil {
			return err
		}
		m.conn = conn
	case "udp":
		return errors.New("Run udp no implemented")
	default:
		return errors.New("Error ModbusClient Network " + m.Conf.Network)
	}

	// run modbus master
	switch m.Conf.Proto {
	case "rtu":
		return errors.New("Run rtu not implemented")
	case "ascii":
		return errors.New("Run rtu not implemented")
	case "tcp":
		go m.workTcp()
	default:
		return errors.New("Error ModbusClient Proto " + m.Conf.Proto)

	}
	return nil
}

func (m *modbusClient) workTcp() {
	var id_trans uint16 = 0
	var ans mbAnsv
	m.readBuffer = make([]byte, 260)

	var requests []mbTcpReq
	for _, rc := range m.Conf.Requests {
		h := mbTcpHeader{0, 0, 6}
		r := mbReq{rc.Id, rc.FuncCode, rc.Start, rc.Length}
		requests = append(requests, mbTcpReq{h, r})
	}

	for {
		for i, r := range requests {
			r.header.transaction = id_trans

			binary.Write(m.conn, binary.BigEndian, &r)
			m.conn.Read(m.readBuffer)
			ans.parseTcp(m.readBuffer)
			m.dataOut <- ModbusData{i, &ans.data}

			time.Sleep(time.Millisecond * 500)

			id_trans += 1
		}
	}
}
