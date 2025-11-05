package natnetgo

import (
	"bytes"
	"encoding/binary"
	"math"
	"time"
)

type ID int32

const idLen = 4

func parseID(data []byte) (ID, int) {
	return ID(binary.LittleEndian.Uint32(data[0:4])), idLen
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

const vector3Len = 12

func parseVector3(data []byte) (Vector3, int, error) {
	v := Vector3{}

	if len(data) < vector3Len {
		return v, 0, ErrTooShort
	}

	v.X = math.Float32frombits(binary.LittleEndian.Uint32(data[0:4]))
	v.Y = math.Float32frombits(binary.LittleEndian.Uint32(data[4:8]))
	v.Z = math.Float32frombits(binary.LittleEndian.Uint32(data[8:12]))

	return v, vector3Len, nil
}

const quaterionLen = 16

type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

func parseQuaternion(data []byte) (Quaternion, int) {
	q := Quaternion{}

	q.X = math.Float32frombits(binary.LittleEndian.Uint32(data[0:4]))
	q.Y = math.Float32frombits(binary.LittleEndian.Uint32(data[4:8]))
	q.Z = math.Float32frombits(binary.LittleEndian.Uint32(data[8:12]))
	q.W = math.Float32frombits(binary.LittleEndian.Uint32(data[12:16]))

	return q, quaterionLen
}

const nameMinLen = 2

func parseName(data []byte) (string, int, error) {
	if len(data) < nameMinLen {
		return "", 0, ErrTooShort
	}

	n := bytes.IndexByte(data, 0)
	if n == -1 {
		return "", 0, ErrInvalidName
	}

	return string(data[:n]), n + 1, nil
}

const int32Len = 4

func parseInt32(data []byte) (int32, int) {
	return int32(binary.LittleEndian.Uint32(data)), int32Len
}

const uint32Len = 4

func parseUint32(data []byte) (uint32, int) {
	return binary.LittleEndian.Uint32(data), uint32Len
}

const timestampLen = 8

func parseTimestamp(data []byte) (time.Time, int) {
	return time.Unix(0, int64(binary.LittleEndian.Uint64(data))), timestampLen
}

type Params uint16

const paramsLen = 2

func parseParams(data []byte) (Params, int) {
	return Params(binary.LittleEndian.Uint16(data)), paramsLen
}

const doubleLen = 8

func parseDouble(data []byte) (float64, int) {
	return math.Float64frombits(binary.LittleEndian.Uint64(data)), doubleLen
}

const floatLen = 4

func parseFloat(data []byte) (float32, int) {
	return math.Float32frombits(binary.LittleEndian.Uint32(data)), floatLen
}
