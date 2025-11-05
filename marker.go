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

func parseMarker(data []byte) (*Marker, int, error) {
	if len(data) < markerLen {
		return nil, 0, ErrTooShort
	}

	m := new(Marker)

	// Get the ids.
	// The lower half is the marker id, the upper half is the model id
	tmpID, _ := parseID(data)
	m.MarkerID = int16(tmpID & 0x0000ffff)
	m.ModelID = int16(tmpID >> 16)

	// Get the position
	pos, _, _ := parseVector3(data[idLen:])
	m.Position = pos

	// Get the size
	size, _ := parseFloat(data[idLen+vector3Len:])
	m.Size = size

	// Get the params
	params, _ := parseParams(data[idLen+vector3Len+floatLen:])
	m.Params = params

	// Get the residual
	residual, _ := parseFloat(data[idLen+vector3Len+floatLen+paramsLen:])
	m.Residual = residual

	return m, markerLen, nil
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
	Markers *List[Vector3]
}

const markerSetMinLen = nameMinLen + listMinLen

func parseMarkerSet(data []byte) (*MarkerSet, int, error) {
	if len(data) < markerSetMinLen {
		return nil, 0, ErrTooShort
	}

	ms := new(MarkerSet)

	offset := 0

	// Get the marker set name
	name, tmpOffset, err := parseName(data)
	if err != nil {
		return nil, 0, err
	}
	ms.Name = name
	offset += tmpOffset

	// Get the list of markers of the set
	markers, tmpOffset, err := parseList(parseVector3, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ms.Markers = markers
	offset += tmpOffset

	return ms, offset, nil
}
