package natnetgo

type Skeleton struct {
	ID          ID
	RigidBodies *List[*RigidBody]
}

const skeletonMinLen = idLen + listMinLen

func parseSkeleton(data []byte) (*Skeleton, int, error) {
	if len(data) < skeletonMinLen {
		return nil, 0, ErrTooShort
	}

	s := new(Skeleton)

	id, _ := parseID(data)
	s.ID = id

	rigidBodies, offset, err := parseList(parseRigidBody, data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	s.RigidBodies = rigidBodies

	return s, offset, nil
}
