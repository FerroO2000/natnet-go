package natnetgo

import "errors"

type MessageType uint16

const (
	MessageTypeConnect MessageType = iota // Command
	MessageTypeServerInfo
	MessageTypeRequest // Command
	MessageTypeResponse
	MessageTypeRequestModelDef // Command
	MessageTypeModelDef
	MessageTypeRequestFrameOfData // Command
	MessageTypeFrameOfData
	MessageTypeMessageString
	MessageTypeDisconnect // Command?
	MessageTypeKeepAlive  // Command
	MessageTypeDisconnectByTimeout
	MessageTypeEchoRequest
	MessageTypeEchoResponse
	MessageTypeDiscovery
	MessageTypeUnrecognizedRequest MessageType = 100
)

type Message struct {
	Type        MessageType
	Size        uint16
	ServerInfo  *ServerInfo
	Response    *Response
	Description *Desc
	MoCap       *MoCap
	String      string
}

const messageMinLen = 2 * uint16Len

func ParseMessage(data []byte) (*Message, error) {
	if len(data) < messageMinLen {
		return nil, ErrTooShort
	}

	m := new(Message)

	typ, offset, err := parseUint16(data)
	if err != nil {
		return nil, err
	}
	m.Type = MessageType(typ)

	size, tmpOffset, err := parseUint16(data[offset:])
	if err != nil {
		return nil, err
	}
	m.Size = size
	offset += tmpOffset

	switch m.Type {
	case MessageTypeServerInfo:
		serverInfo, _, err := parseServerInfo(data[offset:])
		if err != nil {
			return nil, err
		}
		m.ServerInfo = serverInfo

	case MessageTypeResponse:
		resp, _, err := decodeResponse(data[offset:])
		if err != nil {
			return nil, err
		}
		m.Response = resp

	case MessageTypeModelDef:
		desc, _, err := parseDesc(data[offset:])
		if err != nil {
			return nil, err
		}
		m.Description = desc

	case MessageTypeFrameOfData:
		moCap, _, err := parseMoCap(data[offset:])
		if err != nil {
			return nil, err
		}
		m.MoCap = moCap

	case MessageTypeMessageString:
		str, _, err := parseName(data[offset:])
		if err != nil {
			return nil, err
		}
		m.String = str

	case MessageTypeUnrecognizedRequest:
		return nil, errors.New("unrecognized request")
	}

	return m, nil
}
