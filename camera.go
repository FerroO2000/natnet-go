package natnetgo

type CameraDesc struct {
	Name        string
	Position    Vector3
	Orientation Quaternion
}

const cameraDescLen = nameMinLen + vector3Len + quaterionLen

func parseCameraDesc(data []byte) (*CameraDesc, int, error) {
	if len(data) < cameraDescLen {
		return nil, 0, ErrTooShort
	}

	c := new(CameraDesc)

	name, offset, err := parseName(data)
	if err != nil {
		return nil, 0, err
	}
	c.Name = name

	pos, tmpOffset, err := parseVector3(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	c.Position = pos
	offset += tmpOffset

	rot, tmpOffset, err := parseQuaternion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	c.Orientation = rot
	offset += tmpOffset

	return c, offset, nil
}
