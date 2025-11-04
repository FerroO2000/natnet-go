package natnetgo

import (
	"encoding/binary"
	"math"
)

type RigidBody struct {
	ID        ID
	Position  Vector3
	Rotation  Quaternion
	MeanError float32
	Tracking  bool

	// TODO! marker list
}

const rigidBodyLen = idLen + vector3Len + quaterionLen + 6

func parseRigidBody(data []byte) (RigidBody, int, error) {
	rb := RigidBody{}

	if len(data) < rigidBodyLen {
		return rb, 0, ErrTooShort
	}

	id, _ := parseID(data)
	rb.ID = id

	pos, _ := parseVector3(data[idLen:])
	rb.Position = pos

	rot, _ := parseQuaternion(data[idLen+vector3Len:])
	rb.Rotation = rot

	tmpMeanErr := binary.LittleEndian.Uint32(data[idLen+vector3Len+quaterionLen:])
	rb.MeanError = math.Float32frombits(tmpMeanErr)

	tmpTraking := binary.LittleEndian.Uint16(data[idLen+vector3Len+quaterionLen+4:])
	rb.Tracking = (tmpTraking & 0x01) != 0

	return rb, rigidBodyLen, nil
}
