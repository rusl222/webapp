package modbus

import "errors"

type ModbusRequestConfig struct {
	Id       uint8
	FuncCode uint8
	Start    uint16
	Length   uint16
}

type Config struct {
	Network  string
	Proto    string
	IpPort   string
	Requests []ModbusRequestConfig
}

func testConf() *Config {
	var req []ModbusRequestConfig
	req = append(req, ModbusRequestConfig{
		Id:       1,
		FuncCode: 3,
		Start:    0,
		Length:   10,
	})

	return &Config{
		Network:  "tcp",
		Proto:    "tcp",
		IpPort:   "172.30.240.1:502",
		Requests: req,
	}
}

const (
	Id       = "id"
	FuncCode = "funcCode"
	Start    = "start"
	Quantity = "quantity"
)

func (req mbReq) setConfig(param string, value interface{}) error {

	switch param {
	case Id:
		if v, ok := value.(uint8); ok {
			req.id = v
			return nil
		}
	case FuncCode:
		if v, ok := value.(uint8); ok {
			req.funcCode = v
			return nil
		}
	case Start:
		if v, ok := value.(uint16); ok {
			req.start = v
			return nil
		}
	case Quantity:
		if v, ok := value.(uint16); ok {
			req.length = v
			return nil
		}
	}
	return errors.New("set config error")
}

// setConfig implements ModbusTcpClient.
func (modbusClient) setConfig(param string, value interface{}) error {
	return errors.New("unimplemented")
}
