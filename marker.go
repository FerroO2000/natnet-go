package natnetgo

type Marker struct {
	MarkerID int16
	ModelID  int16
	Position Vector3
	Size     float32
	Params   Params
	Residual float32
}

const markerLen = idLen + vector3Len + 2*floatLen + paramsLen

func decodeMarker(data []byte) (*Marker, int, error) {
	if len(data) < markerLen {
		return nil, 0, ErrTooShort
	}

	m := new(Marker)

	// Get the ids.
	// The lower half is the marker id, the upper half is the model id
	tmpID, offset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	m.MarkerID = int16(tmpID & 0x0000ffff)
	m.ModelID = int16(tmpID >> 16)

	// Get the position
	pos, tmpOffset, err := decodeVector3(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	m.Position = pos
	offset += tmpOffset

	// Get the size
	size, tmpOffset, err := decodeFloat(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	m.Size = size
	offset += tmpOffset

	// Get the params
	params, tmpOffset, err := decodeParams(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	m.Params = params
	offset += tmpOffset

	// Get the residual
	residual, tmpOffset, err := decodeFloat(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	m.Residual = residual
	offset += tmpOffset

	return m, offset, nil
}

type MarkerParams = Params

const (
	MarkerParamsOccluded MarkerParams = 1 << iota
	MarkerParamsPointCloudSolved
	MarkerParamsModelFilled
	MarkerParamsHasModel
	MarkerParamsUnlabeled
	MarkerParamsActive
	MarkerParamsEstablished
	MarkerParamsMeasurement
)

func (m *Marker) ParamsIs(target MarkerParams) bool {
	return m.Params&target != 0
}

type MarkerSet struct {
	Name    string
	Markers []Vector3
}

const markerSetMinLen = stringMinLen + listMinLen

func decodeMarkerSet(data []byte) (*MarkerSet, int, error) {
	if len(data) < markerSetMinLen {
		return nil, 0, ErrTooShort
	}

	ms := new(MarkerSet)

	offset := 0

	// Get the marker set name
	name, tmpOffset, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	ms.Name = name
	offset += tmpOffset

	// Get the list of markers of the set
	markers, tmpOffset, err := decodeList(decodeVector3, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ms.Markers = markers
	offset += tmpOffset

	return ms, offset, nil
}

type MarkerDesc struct {
	Name            string
	ID              ID
	InitialPosition Vector3
	Size            float32
	Params          Params
}

const markerDescMinLen = stringMinLen + idLen + vector3Len + paramsLen

type MarkerDescParams = Params

const (
	MarkerDescParamsActive MarkerDescParams = 1 << iota
)

func decodeMarkerDesc(data []byte) (*MarkerDesc, int, error) {
	if len(data) < markerDescMinLen {
		return nil, 0, ErrTooShort
	}

	md := new(MarkerDesc)

	// Get the name
	name, offset, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	md.Name = name

	// Get the id
	id, tmpOffset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	md.ID = id
	offset += tmpOffset

	// Get the initial position
	pos, tmpOffset, err := decodeVector3(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	md.InitialPosition = pos
	offset += tmpOffset

	// Get the size
	size, tmpOffset, err := decodeFloat(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	md.Size = size
	offset += tmpOffset

	// Get the params
	params, tmpOffset, err := decodeParams(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	md.Params = params
	offset += tmpOffset

	return md, 0, nil
}

func (md *MarkerDesc) ParamsIs(target MarkerDescParams) bool {
	return md.Params&target != 0
}

type MarkerSetDesc struct {
	Name        string
	MarkerNames []string
}

const markerSetDescMinLen = stringMinLen + listMinLen

func decodeMarkerSetDesc(data []byte) (*MarkerSetDesc, int, error) {
	if len(data) < markerSetDescMinLen {
		return nil, 0, ErrTooShort
	}

	msd := new(MarkerSetDesc)

	name, offset, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	msd.Name = name

	markerNames, tmpOffset, err := decodeList(decodeString, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	msd.MarkerNames = markerNames
	offset += tmpOffset

	return msd, offset, nil
}
