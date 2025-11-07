package natnetgo

type ForcePlate struct {
	ID       ID
	Channels []*Channel
}

const forcePlateMinLen = idLen + listMinLen

func decodeForcePlate(data []byte) (*ForcePlate, int, error) {
	if len(data) < forcePlateMinLen {
		return nil, 0, ErrTooShort
	}

	fp := new(ForcePlate)

	id, offset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	fp.ID = id

	channels, offset, err := decodeList(decodeChannel, data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	fp.Channels = channels

	return fp, offset, nil
}

type ForcePlateDesc struct {
	ID                ID
	SerialNumber      string
	Width             float32
	Length            float32
	Origin            Vector3
	CalibrationMatrix CalibrationMatrix
	Corners           Corners
	PlateType         int32
	ChannelType       int32
	ChannelNames      []string
}

const forcePlateDescMinLen = idLen + stringMinLen + 2*int32Len + (2+12*12+4*3)*floatLen + vector3Len + listMinLen

func decodeForcePlateDesc(data []byte) (*ForcePlateDesc, int, error) {
	if len(data) < forcePlateDescMinLen {
		return nil, 0, ErrTooShort
	}

	fpd := new(ForcePlateDesc)

	// Get the id
	id, offset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	fpd.ID = id

	// Get the serial number
	sn, tmpOffset, err := decodeString(data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	fpd.SerialNumber = sn
	offset += tmpOffset

	// Get the width and length
	width, tmpOffset, err := decodeFloat(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.Width = width
	offset += tmpOffset

	length, tmpOffset, err := decodeFloat(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.Length = length
	offset += tmpOffset

	// Get the origin
	origin, tmpOffset, err := decodeVector3(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.Origin = origin
	offset += tmpOffset

	// Get the calibration matrix
	calibrationMat, tmpOffset, err := decodeCalibrationMatrix(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.CalibrationMatrix = calibrationMat
	offset += tmpOffset

	// Get the corners
	corners, tmpOffset, err := decodeCorners(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.Corners = corners
	offset += tmpOffset

	// Get the plate type
	plateType, tmpOffset, err := decodeInt32(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.PlateType = plateType
	offset += tmpOffset

	// Get the channel type
	channelType, tmpOffset, err := decodeInt32(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.ChannelType = channelType
	offset += tmpOffset

	// Get the channel names
	channelNames, tmpOffset, err := decodeList(decodeString, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	fpd.ChannelNames = channelNames
	offset += tmpOffset

	return fpd, offset, nil
}
