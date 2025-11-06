package natnetgo

type Skeleton struct {
	ID          ID
	RigidBodies []*RigidBody
}

const skeletonMinLen = idLen + listMinLen

func parseSkeleton(data []byte) (*Skeleton, int, error) {
	if len(data) < skeletonMinLen {
		return nil, 0, ErrTooShort
	}

	s := new(Skeleton)

	id, offset, err := parseID(data)
	if err != nil {
		return nil, 0, err
	}
	s.ID = id

	rigidBodies, offset, err := parseList(parseRigidBody, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	s.RigidBodies = rigidBodies

	return s, offset, nil
}

type SkeletonDesc struct {
	Name        string
	ID          ID
	RigidBodies []*RigidBodyDesc
}

const skeletonDescMinLen = nameMinLen + idLen + int32Len

func parseSkeletonDesc(data []byte) (*SkeletonDesc, int, error) {
	if len(data) < skeletonDescMinLen {
		return nil, 0, ErrTooShort
	}

	sd := new(SkeletonDesc)

	// Get the name
	name, offset, err := parseName(data)
	if err != nil {
		return nil, 0, err
	}
	sd.Name = name

	// Get the id
	id, tmpOffset, err := parseID(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	sd.ID = id
	offset += tmpOffset

	// Get the rigid bodies
	rigidBodies, offset, err := parseList(parseRigidBodyDesc, data[offset:])
	if err != nil {
		return nil, 0, err
	}
	sd.RigidBodies = rigidBodies

	return sd, offset, nil
}
