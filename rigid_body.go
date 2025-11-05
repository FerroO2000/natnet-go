package natnetgo

type RigidBody struct {
	ID        ID
	Position  Vector3
	Rotation  Quaternion
	MeanError float32
	Params    Params
}

const rigidBodyLen = idLen + vector3Len + quaterionLen + floatLen + paramsLen

func parseRigidBody(data []byte) (*RigidBody, int, error) {
	if len(data) < rigidBodyLen {
		return nil, 0, ErrTooShort
	}

	rb := new(RigidBody)

	// Get the id
	id, _ := parseID(data)
	rb.ID = id

	// Get the position
	pos, _, _ := parseVector3(data[idLen:])
	rb.Position = pos

	// Get the rotation quaternion
	rot, _ := parseQuaternion(data[idLen+vector3Len:])
	rb.Rotation = rot

	// Get the mean error
	meanErr, _ := parseFloat(data[idLen+vector3Len+quaterionLen:])
	rb.MeanError = meanErr

	// Get the params
	params, _ := parseParams(data[idLen+vector3Len+quaterionLen+floatLen:])
	rb.Params = params

	return rb, rigidBodyLen, nil
}

type RigidBodyParams = Params

const (
	RigidBodyParamsTracking RigidBodyParams = 1 << iota
)

func (rb *RigidBody) ParamsIs(target RigidBodyParams) bool {
	return rb.Params&target != 0
}
