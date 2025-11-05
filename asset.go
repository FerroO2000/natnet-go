package natnetgo

type Asset struct {
	ID          ID
	RigidBodies *List[*RigidBody]
	Markers     *List[*Marker]
}

const assetMinLen = 2 * (idLen + listMinLen)

func parseAsset(data []byte) (*Asset, int, error) {
	if len(data) < assetMinLen {
		return nil, 0, ErrTooShort
	}

	ass := new(Asset)

	// Get the id
	id, _ := parseID(data)
	ass.ID = id

	// Get the rigid bodies
	rigidBodies, offset, err := parseList(parseRigidBody, data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	ass.RigidBodies = rigidBodies

	// Get the markers
	markers, offset, err := parseList(parseMarker, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ass.Markers = markers

	return ass, offset, nil
}
