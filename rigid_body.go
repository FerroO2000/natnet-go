package natnetgo

type RigidBody struct {
	ID        ID
	Position  Vector3
	Rotation  Quaternion
	MeanError float32
	Params    Params
}

const rigidBodyLen = idLen + vector3Len + quaterionLen + floatLen + paramsLen

func decodeRigidBody(data []byte) (*RigidBody, int, error) {
	if len(data) < rigidBodyLen {
		return nil, 0, ErrTooShort
	}

	rb := new(RigidBody)

	// Get the id
	id, offset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	rb.ID = id

	// Get the position
	pos, tmpOffset, err := decodeVector3(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rb.Position = pos
	offset += tmpOffset

	// Get the rotation quaternion
	rot, tmpOffset, err := decodeQuaternion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rb.Rotation = rot
	offset += tmpOffset

	// Get the mean error
	meanErr, tmpOffset, err := decodeFloat(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rb.MeanError = meanErr
	offset += tmpOffset

	// Get the params
	params, tmpOffset, err := decodeParams(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rb.Params = params
	offset += tmpOffset

	return rb, offset, nil
}

type RigidBodyParams = Params

const (
	RigidBodyParamsTracking RigidBodyParams = 1 << iota
)

func (rb *RigidBody) ParamsIs(target RigidBodyParams) bool {
	return rb.Params&target != 0
}

type RigidBodyDesc struct {
	Name               string
	ID                 ID
	ParentID           ID
	OffsetPosition     Vector3
	OffsetRotation     Quaternion
	MarkerCount        int32
	MarkerPositions    []Vector3
	MarkerActiveLabels []int32
	MarkerNames        []string
}

const rigidBodyDescMinLen = stringMinLen + idLen + idLen + vector3Len + quaterionLen + int32Len

func decodeRigidBodyDesc(data []byte) (*RigidBodyDesc, int, error) {
	if len(data) < rigidBodyDescMinLen {
		return nil, 0, ErrTooShort
	}

	rbd := new(RigidBodyDesc)

	// Get the name
	name, offset, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	rbd.Name = name

	// Get the ids
	id, tmpOffset, err := decodeID(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rbd.ID = id
	offset += tmpOffset

	parentID, tmpOffset, err := decodeID(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rbd.ParentID = parentID
	offset += tmpOffset

	// Get the position/rotation offset
	pos, tmpOffset, _ := decodeVector3(data[offset:])
	rbd.OffsetPosition = pos
	offset += tmpOffset

	rot, tmpOffset, err := decodeQuaternion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rbd.OffsetRotation = rot
	offset += tmpOffset

	// Get the marker count
	markerCount, tmpOffset, err := decodeInt32(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	rbd.MarkerCount = markerCount
	offset += tmpOffset

	// Check the minimum length based on the marker count
	minMarkerSectionLen := int(markerCount) * (vector3Len + int32Len + stringMinLen)
	if len(data) < rigidBodyDescMinLen+minMarkerSectionLen {
		return nil, 0, ErrTooShort
	}

	// Get the marker positions
	rbd.MarkerPositions = make([]Vector3, 0, markerCount)
	for range markerCount {
		pos, tmpOffset, err := decodeVector3(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		rbd.MarkerPositions = append(rbd.MarkerPositions, pos)
		offset += tmpOffset
	}

	// Get the marker active labels
	rbd.MarkerActiveLabels = make([]int32, 0, markerCount)
	for range markerCount {
		label, tmpOffset, err := decodeInt32(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		rbd.MarkerActiveLabels = append(rbd.MarkerActiveLabels, label)
		offset += tmpOffset
	}

	// Get the marker names
	rbd.MarkerNames = make([]string, 0, markerCount)
	for range markerCount {
		name, tmpOffset, err := decodeString(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		rbd.MarkerNames = append(rbd.MarkerNames, name)
		offset += tmpOffset
	}

	return rbd, offset, nil
}
