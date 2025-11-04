package natnetgo

type Marker = Vector3

const markerLen = vector3Len

func parseMarker(data []byte) (Marker, int, error) {
	m := Marker{}

	if len(data) < markerLen {
		return m, 0, ErrTooShort
	}

	m, offset := parseVector3(data)

	return m, offset, nil
}

type MarkerSet struct {
	Name        string
	MarkerCount int32
	Markers     []Marker
}

const markerSetMinLen = 2 + countLen

func parseMarkerSet(data []byte) (MarkerSet, int, error) {
	ms := MarkerSet{}

	if len(data) < markerSetMinLen {
		return ms, 0, ErrTooShort
	}

	offset := 0

	// Get the marker set name
	name, tmpOffset, err := parseName(data)
	if err != nil {
		return ms, 0, err
	}
	ms.Name = name
	offset += tmpOffset

	// Get the number of markers of the set
	count, tmpOffset := parseCount(data[offset:])
	ms.MarkerCount = count
	offset += tmpOffset

	// Get the position of the markers
	ms.Markers = make([]Marker, 0, ms.MarkerCount)
	for range ms.MarkerCount {
		pos, tmpOffset, err := parseMarker(data[offset:])
		if err != nil {
			return ms, 0, err
		}
		ms.Markers = append(ms.Markers, pos)
		offset += tmpOffset
	}

	return ms, offset, nil
}
