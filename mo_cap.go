package natnetgo

type MoCap struct {
	FramePrefix   *FramePrefix
	MarkerSets    []*MarkerSet
	OtherMarkers  []Vector3
	RigidBodies   []*RigidBody
	Skeletons     []*Skeleton
	Assets        []*Asset
	LabeledMarkes []*Marker
	ForcePlates   []*ForcePlate
	Devices       []*Device
	FrameSuffix   *FrameSuffix
}

func parseMoCap(data []byte) (*MoCap, int, error) {
	mc := new(MoCap)
	offset := 0

	// Get frame prefix
	fp, tmpOffset, err := parseFramePrefix(data)
	if err != nil {
		return nil, 0, err
	}
	mc.FramePrefix = fp
	offset += tmpOffset

	// Get the list of marker sets
	markerSets, tmpOffset, err := parseSizedList(parseMarkerSet, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.MarkerSets = markerSets
	offset += tmpOffset

	// Get the list of other markers
	otherMarkers, tmpOffset, err := parseSizedList(parseVector3, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.OtherMarkers = otherMarkers
	offset += tmpOffset

	// Get the list of rigid bodies
	rigidBodies, tmpOffset, err := parseSizedList(parseRigidBody, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.RigidBodies = rigidBodies
	offset += tmpOffset

	// Get the list of skeletons
	skeletons, tmpOffset, err := parseSizedList(parseSkeleton, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.Skeletons = skeletons
	offset += tmpOffset

	// Get the list of assets
	assets, tmpOffset, err := parseSizedList(parseAsset, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.Assets = assets
	offset += tmpOffset

	// Get the list of labeled markers
	labeledMarkers, tmpOffset, err := parseSizedList(parseMarker, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.LabeledMarkes = labeledMarkers
	offset += tmpOffset

	// Get the list of force plates
	forcePlates, tmpOffset, err := parseSizedList(parseForcePlate, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.ForcePlates = forcePlates
	offset += tmpOffset

	// Get the list of devices
	devices, tmpOffset, err := parseSizedList(parseDevice, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.Devices = devices
	offset += tmpOffset

	// Get the frame suffix
	fs, tmpOffset, err := parseFrameSuffix(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.FrameSuffix = fs
	offset += tmpOffset

	return mc, offset, nil
}
