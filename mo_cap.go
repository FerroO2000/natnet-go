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

func decodeMoCap(data []byte) (*MoCap, int, error) {
	mc := new(MoCap)
	offset := 0

	// Get frame prefix
	fp, tmpOffset, err := decodeFramePrefix(data)
	if err != nil {
		return nil, 0, err
	}
	mc.FramePrefix = fp
	offset += tmpOffset

	// Get the list of marker sets
	markerSets, tmpOffset, err := decodeSizedList(decodeMarkerSet, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.MarkerSets = markerSets
	offset += tmpOffset

	// Get the list of other markers
	otherMarkers, tmpOffset, err := decodeSizedList(decodeVector3, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.OtherMarkers = otherMarkers
	offset += tmpOffset

	// Get the list of rigid bodies
	rigidBodies, tmpOffset, err := decodeSizedList(decodeRigidBody, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.RigidBodies = rigidBodies
	offset += tmpOffset

	// Get the list of skeletons
	skeletons, tmpOffset, err := decodeSizedList(decodeSkeleton, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.Skeletons = skeletons
	offset += tmpOffset

	// Get the list of assets
	assets, tmpOffset, err := decodeSizedList(decodeAsset, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.Assets = assets
	offset += tmpOffset

	// Get the list of labeled markers
	labeledMarkers, tmpOffset, err := decodeSizedList(decodeMarker, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.LabeledMarkes = labeledMarkers
	offset += tmpOffset

	// Get the list of force plates
	forcePlates, tmpOffset, err := decodeSizedList(decodeForcePlate, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.ForcePlates = forcePlates
	offset += tmpOffset

	// Get the list of devices
	devices, tmpOffset, err := decodeSizedList(decodeDevice, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.Devices = devices
	offset += tmpOffset

	// Get the frame suffix
	fs, tmpOffset, err := decodeFrameSuffix(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	mc.FrameSuffix = fs
	offset += tmpOffset

	return mc, offset, nil
}
