package natnetgo

type commandType uint16

const (
	commandTypeConnect              commandType = 0
	commandTypeRequest              commandType = 2
	commandTypeRequestModelDescData commandType = 4
	commandTypeRequestMoCapData     commandType = 6
	commandTypeKeepAlive            commandType = 10
)

func encodeCommand(typ commandType, cmd string) ([]byte, int) {
	switch typ {
	case commandTypeConnect:
		cmd = "Ping"
	case commandTypeRequestModelDescData, commandTypeRequestMoCapData, commandTypeKeepAlive:
		cmd = ""
	}

	cmdBytes := []byte(cmd)
	cmdBytesLen := len(cmdBytes)
	dataLen := 2*uint16Len + cmdBytesLen + 1
	data := make([]byte, dataLen)

	encodeUint16(data, uint16(typ))
	encodeUint16(data[2:4], uint16(cmdBytesLen+1))

	offset := 4

	copy(data[offset:offset+cmdBytesLen], cmdBytes)
	data[offset+cmdBytesLen] = 0

	return data, dataLen
}

// EncodeConnect encodes a connect command.
// It is used to connect to the server.
func EncodeConnect() ([]byte, int) {
	return encodeCommand(commandTypeConnect, "")
}

// EncodeRequest encodes a request command.
// It is used to send custom commands to the server.
func EncodeRequest(cmd string) ([]byte, int) {
	return encodeCommand(commandTypeRequest, cmd)
}

// EncodeRequestModelDescData encodes a request model description data command.
// This command is used to request the model description
// from the server.
func EncodeRequestModelDescData() ([]byte, int) {
	return encodeCommand(commandTypeRequestModelDescData, "")
}

// EncodeRequestMoCapData encodes a request of motion capture data command.
// This command is used to request the motion capture data
// from the server.
func EncodeRequestMoCapData() ([]byte, int) {
	return encodeCommand(commandTypeRequestMoCapData, "")
}

// EncodeKeepAlive encodes a keep alive command.
func EncodeKeepAlive() ([]byte, int) {
	return encodeCommand(commandTypeKeepAlive, "")
}
