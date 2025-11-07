package natnetgo

type Device struct {
	ID       ID
	Channels []*Channel
}

const deviceMinLen = idLen + listMinLen

func decodeDevice(data []byte) (*Device, int, error) {
	if len(data) < deviceMinLen {
		return nil, 0, ErrTooShort
	}

	d := new(Device)

	id, offset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	d.ID = id

	channels, tmpOffset, err := decodeList(decodeChannel, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	d.Channels = channels
	offset += tmpOffset

	return d, offset, nil
}

type DeviceDesc struct {
	ID              ID
	Name            string
	SerialNumber    string
	DeviceType      int32
	ChannelDataType int32
	ChannelNames    []string
}

const deviceDescMinLen = idLen + 2*stringMinLen

func decodeDeviceDesc(data []byte) (*DeviceDesc, int, error) {
	if len(data) < deviceDescMinLen {
		return nil, 0, ErrTooShort
	}

	dd := new(DeviceDesc)

	// Get the id
	id, offset, _ := decodeID(data)
	dd.ID = id

	// Get the name and serial number
	name, tmpOffset, _ := decodeString(data[offset:])
	dd.Name = name
	offset += tmpOffset

	serialNumber, tmpOffset, _ := decodeString(data[offset:])
	dd.SerialNumber = serialNumber
	offset += tmpOffset

	// Get the name of the channels
	channelNames, tmpOffset, err := decodeList(decodeString, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	dd.ChannelNames = channelNames
	offset += tmpOffset

	return dd, offset, nil
}
