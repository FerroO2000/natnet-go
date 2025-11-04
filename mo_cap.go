package natnetgo

type MoCap struct {
	FramePrefix  *FramePrefix
	MarkerSets   List[MarkerSet]
	OtherMarkers List[Marker]
	RigidBodies  List[RigidBody]
	Skeletons    List[Skeleton]
	FrameSuffix  *FrameSuffix
}

func parseMoCap(data []byte) (MoCap, int, error) {
	mc := MoCap{}

	offset := 0

	// Get frame prefix
	fp, tmpOffset, err := parseFramePrefix(data)
	if err != nil {
		return mc, 0, err
	}
	mc.FramePrefix = fp
	offset += tmpOffset

	// Get the list of marker sets
	markerSets, tmpOffset, err := parseList(parseMarkerSet, data[offset:])
	if err != nil {
		return mc, 0, err
	}
	mc.MarkerSets = markerSets
	offset += tmpOffset

	// Get the list of other markers
	otherMarkers, tmpOffset, err := parseList(parseMarker, data[offset:])
	if err != nil {
		return mc, 0, err
	}
	mc.OtherMarkers = otherMarkers
	offset += tmpOffset

	// Get the list of rigid bodies
	rigidBodies, tmpOffset, err := parseList(parseRigidBody, data[offset:])
	if err != nil {
		return mc, 0, err
	}
	mc.RigidBodies = rigidBodies
	offset += tmpOffset

	// Get the list of skeletons
	skeletons, tmpOffset, err := parseList(parseSkeleton, data[offset:])
	if err != nil {
		return mc, 0, err
	}
	mc.Skeletons = skeletons
	offset += tmpOffset

	// TODO! assets
	// TODO! labeled markers
	// TODO! force plates
	// TODO! devices

	// Get the frame suffix
	fs, tmpOffset, err := parseFrameSuffix(data[offset:])
	if err != nil {
		return mc, 0, err
	}
	mc.FrameSuffix = fs
	offset += tmpOffset

	return mc, offset, nil
}
