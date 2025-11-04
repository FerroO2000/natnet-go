package natnetgo

import "encoding/binary"

type Skeleton struct {
	ID             ID
	RigidBodyCount int32
	RigidBodies    []RigidBody
}

const skeletonDataLen = idLen + 4

func parseSkeleton(data []byte) (Skeleton, int, error) {
	s := Skeleton{}

	if len(data) < skeletonDataLen {
		return s, 0, ErrTooShort
	}

	id, _ := parseID(data)
	s.ID = id

	s.RigidBodyCount = int32(binary.LittleEndian.Uint32(data[idLen:]))
	s.RigidBodies = make([]RigidBody, 0, s.RigidBodyCount)

	offset := idLen + 4
	for range s.RigidBodyCount {
		rb, tmpOffset, err := parseRigidBody(data[offset:])
		if err != nil {
			return s, 0, err
		}

		s.RigidBodies = append(s.RigidBodies, rb)
		offset += tmpOffset
	}

	return s, offset, nil
}
