package natnetgo

import (
	"bytes"
	"encoding/binary"
	"math"
	"time"
)

const stringMinLen = 2

func decodeString(data []byte) (string, int, error) {
	if len(data) < stringMinLen {
		return "", 0, ErrTooShort
	}

	n := bytes.IndexByte(data, 0)
	if n == -1 {
		return "", 0, ErrInvalidName
	}

	return string(data[:n]), n + 1, nil
}

const uint16Len = 2

func decodeUint16(data []byte) (uint16, int, error) {
	if len(data) < uint16Len {
		return 0, 0, ErrTooShort
	}

	return binary.LittleEndian.Uint16(data), uint16Len, nil
}

func encodeUint16(data []byte, val uint16) {
	binary.LittleEndian.PutUint16(data, val)
}

const int32Len = 4

func decodeInt32(data []byte) (int32, int, error) {
	if len(data) < int32Len {
		return 0, 0, ErrTooShort
	}

	return int32(binary.LittleEndian.Uint32(data)), int32Len, nil
}

const uint32Len = 4

func decodeUint32(data []byte) (uint32, int) {
	return binary.LittleEndian.Uint32(data), uint32Len
}

const doubleLen = 8

func decodeDouble(data []byte) (float64, int, error) {
	if len(data) < doubleLen {
		return 0, 0, ErrTooShort
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(data)), doubleLen, nil
}

const floatLen = 4

func decodeFloat(data []byte) (float32, int, error) {
	if len(data) < floatLen {
		return 0, 0, ErrTooShort
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(data)), floatLen, nil
}

const timestampLen = 8

func decodeTimestamp(data []byte) (time.Time, int, error) {
	if len(data) < timestampLen {
		return time.Time{}, 0, ErrTooShort
	}

	return time.Unix(0, int64(binary.LittleEndian.Uint64(data))), timestampLen, nil
}

type ID int32

const idLen = 4

func decodeID(data []byte) (ID, int, error) {
	if len(data) < idLen {
		return 0, 0, ErrTooShort
	}

	return ID(binary.LittleEndian.Uint32(data[0:4])), idLen, nil
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

const vector3Len = 12

func decodeVector3(data []byte) (Vector3, int, error) {
	v := Vector3{}

	if len(data) < vector3Len {
		return v, 0, ErrTooShort
	}

	v.X = math.Float32frombits(binary.LittleEndian.Uint32(data[0:4]))
	v.Y = math.Float32frombits(binary.LittleEndian.Uint32(data[4:8]))
	v.Z = math.Float32frombits(binary.LittleEndian.Uint32(data[8:12]))

	return v, vector3Len, nil
}

type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

const quaterionLen = 16

func decodeQuaternion(data []byte) (Quaternion, int, error) {
	q := Quaternion{}

	if len(data) < quaterionLen {
		return q, 0, ErrTooShort
	}

	q.X = math.Float32frombits(binary.LittleEndian.Uint32(data[0:4]))
	q.Y = math.Float32frombits(binary.LittleEndian.Uint32(data[4:8]))
	q.Z = math.Float32frombits(binary.LittleEndian.Uint32(data[8:12]))
	q.W = math.Float32frombits(binary.LittleEndian.Uint32(data[12:16]))

	return q, quaterionLen, nil
}

type Params uint16

const paramsLen = 2

func decodeParams(data []byte) (Params, int, error) {
	if len(data) < paramsLen {
		return 0, 0, ErrTooShort
	}

	return Params(binary.LittleEndian.Uint16(data)), paramsLen, nil
}

type Version struct {
	Major    uint8
	Minor    uint8
	Build    uint8
	Revision uint8
}

const versionLen = 4

func decodeVersion(data []byte) (Version, int, error) {
	v := Version{}

	if len(data) < versionLen {
		return v, 0, ErrTooShort
	}

	v.Major = data[0]
	v.Minor = data[1]
	v.Build = data[2]
	v.Revision = data[3]

	return v, versionLen, nil
}

type decodeItemFn[T any] func(data []byte) (T, int, error)

const listMinLen = int32Len

const sizedListMinLen = 2 * listMinLen

func decodeSizedList[T any](decodeItem decodeItemFn[T], data []byte) ([]T, int, error) {
	if len(data) < sizedListMinLen {
		return nil, 0, ErrTooShort
	}

	count, _, err := decodeInt32(data)
	if err != nil {
		return nil, 0, err
	}

	items := make([]T, 0, count)
	offset := sizedListMinLen
	for range count {
		item, tmpOffset, err := decodeItem(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		items = append(items, item)
		offset += tmpOffset
	}

	return items, offset, nil
}

func decodeList[T any](parseItem decodeItemFn[T], data []byte) ([]T, int, error) {
	if len(data) < listMinLen {
		return nil, 0, ErrTooShort
	}

	count, offset, err := decodeInt32(data)
	if err != nil {
		return nil, 0, err
	}

	items := make([]T, 0, count)
	for range count {
		item, tmpOffset, err := parseItem(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		items = append(items, item)
		offset += tmpOffset
	}

	return items, offset, nil
}

type CalibrationMatrix [12][12]float32

const calibrationMatrixLen = 12 * 12 * floatLen

func decodeCalibrationMatrix(data []byte) (CalibrationMatrix, int, error) {
	cm := CalibrationMatrix{}

	if len(data) < calibrationMatrixLen {
		return cm, 0, ErrTooShort
	}

	offset := 0
	for i := range 12 {
		for j := range 12 {
			val, tmpOffset, err := decodeFloat(data[offset:])
			if err != nil {
				return cm, 0, err
			}
			cm[i][j] = val
			offset += tmpOffset
		}
	}

	return cm, calibrationMatrixLen, nil
}

type Corners [4][3]float32

const cornersLen = 4 * 3 * floatLen

func decodeCorners(data []byte) (Corners, int, error) {
	c := Corners{}

	if len(data) < cornersLen {
		return c, 0, ErrTooShort
	}

	offset := 0
	for i := range 4 {
		for j := range 3 {
			val, tmpOffset, err := decodeFloat(data[offset:])
			if err != nil {
				return c, 0, err
			}
			c[i][j] = val
			offset += tmpOffset
		}
	}

	return c, cornersLen, nil
}
