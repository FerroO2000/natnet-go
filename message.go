package natnetgo

type MessageType uint16

const (
	MessageTypeServerInfo          MessageType = 1
	MessageTypeResponse            MessageType = 3
	MessageTypeModelDescData       MessageType = 5
	MessageTypeMoCapData           MessageType = 7
	MessageTypeMessageString       MessageType = 8
	MessageTypeDisconnect          MessageType = 9
	MessageTypeDisconnectByTimeout MessageType = 11
	MessageTypeEchoRequest         MessageType = 12
	MessageTypeEchoResponse        MessageType = 13
	MessageTypeDiscovery           MessageType = 14
	MessageTypeUnrecognizedRequest MessageType = 100
)

type Message struct {
	Type MessageType
	Size uint16

	ServerInfo    *ServerInfo
	Response      *Response
	ModelDesc     *ModelDesc
	MoCap         *MoCap
	MessageString string
}

const messageMinLen = 2 * uint16Len

func DecodeMessage(data []byte) (*Message, error) {
	if len(data) < messageMinLen {
		return nil, ErrTooShort
	}

	m := new(Message)

	typ, offset, err := decodeUint16(data)
	if err != nil {
		return nil, err
	}
	m.Type = MessageType(typ)

	size, tmpOffset, err := decodeUint16(data[offset:])
	if err != nil {
		return nil, err
	}
	m.Size = size
	offset += tmpOffset

	switch m.Type {
	case MessageTypeServerInfo:
		serverInfo, _, err := decodeServerInfo(data[offset:])
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

	case MessageTypeModelDescData:
		desc, _, err := decodeModelDesc(data[offset:])
		if err != nil {
			return nil, err
		}
		m.ModelDesc = desc

	case MessageTypeMoCapData:
		moCap, _, err := decodeMoCap(data[offset:])
		if err != nil {
			return nil, err
		}
		m.MoCap = moCap

	case MessageTypeMessageString:
		str, _, err := decodeString(data[offset:])
		if err != nil {
			return nil, err
		}
		m.MessageString = str
	}

	return m, nil
}
