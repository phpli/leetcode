package packet

import (
	"bytes"
	"fmt"
)

const (
	CommandConn   = iota + 0x01 // 0x01
	CommandSubmit               // 0x02
)

const (
	CommandConnAck   = iota + 0x81 // 0x81
	CommandSubmitAck               //0x82
)

type Packet interface {
	Decode([]byte) error     //[]byte -> struct
	Encode() ([]byte, error) //struct->[]byte
}

type Submit struct {
	ID      string
	Payload []byte
}

func (s *Submit) Decode(bytes []byte) error {
	s.ID = string(bytes[:8])
	s.Payload = bytes[8:]
	return nil
}

func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), s.Payload}, nil), nil
}

type SubmitAck struct {
	ID     string
	Result uint8
}

func (s *SubmitAck) Decode(i []byte) error {
	s.ID = string(i[0:8])
	s.Result = uint8(i[8])
	return nil
}

func (s *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), []byte{s.Result}}, nil), nil
}

func Decode(packet []byte) (Packet, error) {
	commandID := packet[0]
	pktBody := packet[1:]
	switch commandID {
	case CommandConn:
		return nil, nil
	case CommandConnAck:
		return nil, nil
	case CommandSubmit:
		s := Submit{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	case CommandSubmitAck:
		s := SubmitAck{}
		err := s.Decode(pktBody)
		if err != nil {
			return nil, err
		}
		return &s, nil
	default:
		return nil, fmt.Errorf("unknown commandID [%d]", commandID)
	}
}

func Encode(p Packet) ([]byte, error) {
	var commandID uint8
	var pktBody []byte
	var err error
	switch t := p.(type) {
	case *Submit:
		//fmt.Println(t)
		commandID = CommandSubmit
		pktBody, err = p.Encode()
		//fmt.Println(pktBody)
		if err != nil {
			return nil, err
		}
	case *SubmitAck:
		fmt.Println(t)
		commandID = CommandSubmitAck
		pktBody, err = p.Encode()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown type [%s]", t)
	}
	return bytes.Join([][]byte{[]byte{commandID}, pktBody}, nil), nil
}
