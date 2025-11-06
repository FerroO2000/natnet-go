package natnetgo

import "errors"

type Desc struct {
	Datasets []*DescDataset
}

const descMinLen = int32Len

func parseDesc(data []byte) (*Desc, int, error) {
	if len(data) < descMinLen {
		return nil, 0, ErrTooShort
	}

	d := new(Desc)

	// Get the list of datasets
	datasets, offset, err := parseList(parseDescDataset, data)
	if err != nil {
		return nil, 0, err
	}
	d.Datasets = datasets

	return d, offset, nil
}

type DescDatasetType = uint8

const (
	DescDatasetTypeMarkerSet DescDatasetType = iota
	DescDatasetTypeRigidBody
	DescDatasetTypeSkeleton
	DescDatasetTypeForcePlate
	DescDatasetTypeDevice
	DescDatasetTypeCamera
	DescDatasetTypeAsset
)

type DescDataset struct {
	Type       DescDatasetType
	MarkerSet  *MarkerSetDesc
	RigidBody  *RigidBodyDesc
	Skeleton   *SkeletonDesc
	ForcePlate *ForcePlateDesc
	Device     *DeviceDesc
	Camera     *CameraDesc
	Asset      *AssetDesc
}

const descDatasetMinLen = 2 * int32Len

func parseDescDataset(data []byte) (*DescDataset, int, error) {
	if len(data) < descDatasetMinLen {
		return nil, 0, ErrTooShort
	}

	dd := new(DescDataset)

	// Get the type
	typ, _, err := parseInt32(data)
	if err != nil {
		return nil, 0, err
	}
	dd.Type = DescDatasetType(typ)

	// Check the size of the dataset
	sizeByte, _, err := parseInt32(data[int32Len:])
	if err != nil {
		return nil, 0, err
	}

	startOffset := descDatasetMinLen
	endOffset := startOffset + int(sizeByte)

	if len(data) < endOffset {
		return nil, 0, ErrTooShort
	}

	switch dd.Type {
	case DescDatasetTypeMarkerSet:
		markerSet, _, err := parseMarkerSetDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.MarkerSet = markerSet

	case DescDatasetTypeRigidBody:
		rigidBody, _, err := parseRigidBodyDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.RigidBody = rigidBody

	case DescDatasetTypeSkeleton:
		skeleton, _, err := parseSkeletonDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.Skeleton = skeleton

	case DescDatasetTypeForcePlate:
		forcePlate, _, err := parseForcePlateDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.ForcePlate = forcePlate

	case DescDatasetTypeDevice:
		device, _, err := parseDeviceDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.Device = device

	case DescDatasetTypeCamera:
		camera, _, err := parseCameraDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.Camera = camera

	case DescDatasetTypeAsset:
		asset, _, err := parseAssetDesc(data[startOffset:endOffset])
		if err != nil {
			return nil, 0, err
		}
		dd.Asset = asset

	default:
		return nil, 0, errors.New("invalid description type")
	}

	return dd, endOffset, nil
}
