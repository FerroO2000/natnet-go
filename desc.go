package natnetgo

import "errors"

type ModelDesc struct {
	Descriptions []*Desc
}

const modelDescMinLen = int32Len

func decodeModelDesc(data []byte) (*ModelDesc, int, error) {
	if len(data) < modelDescMinLen {
		return nil, 0, ErrTooShort
	}

	md := new(ModelDesc)

	// Get the list of model descriptions
	descriptions, offset, err := decodeList(decodeDesc, data)
	if err != nil {
		return nil, 0, err
	}
	md.Descriptions = descriptions

	return md, offset, nil
}

type DescType = uint8

const (
	DescTypeMarkerSet DescType = iota
	DescTypeRigidBody
	DescTypeSkeleton
	DescTypeForcePlate
	DescTypeDevice
	DescTypeCamera
	DescTypeAsset
)

type Desc struct {
	Type       DescType
	MarkerSet  *MarkerSetDesc
	RigidBody  *RigidBodyDesc
	Skeleton   *SkeletonDesc
	ForcePlate *ForcePlateDesc
	Device     *DeviceDesc
	Camera     *CameraDesc
	Asset      *AssetDesc
}

const descMinLen = 2 * int32Len

func decodeDesc(data []byte) (*Desc, int, error) {
	if len(data) < descMinLen {
		return nil, 0, ErrTooShort
	}

	d := new(Desc)

	// Get the type
	typ, _, err := decodeInt32(data)
	if err != nil {
		return nil, 0, err
	}
	d.Type = DescType(typ)

	// Check the size of the dataset
	sizeByte, _, err := decodeInt32(data[int32Len:])
	if err != nil {
		return nil, 0, err
	}

	startOffset := descMinLen
	endOffset := startOffset + int(sizeByte)

	if len(data) < endOffset {
		return nil, 0, ErrTooShort
	}

	switch d.Type {
	case DescTypeMarkerSet:
		markerSet, _, err := decodeMarkerSetDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.MarkerSet = markerSet

	case DescTypeRigidBody:
		rigidBody, _, err := decodeRigidBodyDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.RigidBody = rigidBody

	case DescTypeSkeleton:
		skeleton, _, err := decodeSkeletonDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.Skeleton = skeleton

	case DescTypeForcePlate:
		forcePlate, _, err := decodeForcePlateDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.ForcePlate = forcePlate

	case DescTypeDevice:
		device, _, err := decodeDeviceDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.Device = device

	case DescTypeCamera:
		camera, _, err := decodeCameraDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.Camera = camera

	case DescTypeAsset:
		asset, _, err := decodeAssetDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		d.Asset = asset

	default:
		return nil, 0, errors.New("invalid description type")
	}

	return d, endOffset, nil
}
