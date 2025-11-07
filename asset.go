package natnetgo

type Asset struct {
	ID          ID
	RigidBodies []*RigidBody
	Markers     []*Marker
}

const assetMinLen = 2 * (idLen + listMinLen)

func decodeAsset(data []byte) (*Asset, int, error) {
	if len(data) < assetMinLen {
		return nil, 0, ErrTooShort
	}

	ass := new(Asset)

	// Get the id
	id, offset, err := decodeID(data)
	if err != nil {
		return nil, 0, err
	}
	ass.ID = id

	// Get the rigid bodies
	rigidBodies, offset, err := decodeList(decodeRigidBody, data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	ass.RigidBodies = rigidBodies

	// Get the markers
	markers, offset, err := decodeList(decodeMarker, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ass.Markers = markers

	return ass, offset, nil
}

type AssetDesc struct {
	Name        string
	Type        int32
	ID          ID
	RigidBodies []*RigidBodyDesc
	Markers     []*MarkerDesc
}

const assetDescMinLen = stringMinLen + int32Len + idLen + 2*listMinLen

func decodeAssetDesc(data []byte) (*AssetDesc, int, error) {
	if len(data) < assetDescMinLen {
		return nil, 0, ErrTooShort
	}

	ad := new(AssetDesc)

	// Get the name
	name, offset, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	ad.Name = name

	// Get the type
	typ, tmpOffset, err := decodeInt32(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ad.Type = typ
	offset += tmpOffset

	// Get the id
	id, tmpOffset, err := decodeID(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ad.ID = id
	offset += tmpOffset

	// Get the rigid bodies
	rigidBodies, offset, err := decodeList(decodeRigidBodyDesc, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ad.RigidBodies = rigidBodies

	// Get the markers
	markers, offset, err := decodeList(decodeMarkerDesc, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	ad.Markers = markers

	return ad, offset, nil
}
