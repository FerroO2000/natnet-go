package natnetgo

type CameraDesc struct {
	Name        string
	Position    Vector3
	Orientation Quaternion
}

const cameraDescLen = stringMinLen + vector3Len + quaterionLen

func decodeCameraDesc(data []byte) (*CameraDesc, int, error) {
	if len(data) < cameraDescLen {
		return nil, 0, ErrTooShort
	}

	c := new(CameraDesc)

	name, offset, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	c.Name = name

	pos, tmpOffset, err := decodeVector3(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	c.Position = pos
	offset += tmpOffset

	rot, tmpOffset, err := decodeQuaternion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	c.Orientation = rot
	offset += tmpOffset

	return c, offset, nil
}
