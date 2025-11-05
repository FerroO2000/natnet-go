package natnetgo

import (
	"time"
)

type FramePrefix struct {
	FrameNumber int32
}

const framePrefixLen = int32Len

func parseFramePrefix(data []byte) (*FramePrefix, int, error) {
	if len(data) < framePrefixLen {
		return nil, 0, ErrTooShort
	}

	fp := new(FramePrefix)

	// Get the frame number
	frameNum, _ := parseInt32(data)
	fp.FrameNumber = frameNum

	return fp, framePrefixLen, nil
}

type FrameSuffix struct {
	Timecode                          uint32
	TimecodeSubframe                  uint32
	SoftwareTimestamp                 float64
	CameraMidExposureTimestamp        time.Time
	CameraDataReceivedTimestamp       time.Time
	TransmitTimestamp                 time.Time
	PrecisionTimestampSeconds         uint32
	PrecisionTimestampFractionSeconds uint32
	Params                            Params
}

const frameSuffixLen = 4*uint32Len + doubleLen + 3*timestampLen + paramsLen

func parseFrameSuffix(data []byte) (*FrameSuffix, int, error) {
	if len(data) < frameSuffixLen {
		return nil, 0, ErrTooShort
	}

	fs := new(FrameSuffix)

	offset := 0

	// Get the timecode
	timecode, tmpOffset := parseUint32(data[offset:])
	fs.Timecode = timecode
	offset += tmpOffset

	// Get the timecode subframe
	timecodeSubframe, tmpOffset := parseUint32(data[offset:])
	fs.TimecodeSubframe = timecodeSubframe
	offset += tmpOffset

	// Get the software timestamp
	swTime, tmpOffset := parseDouble(data[offset:])
	fs.SoftwareTimestamp = swTime
	offset += tmpOffset

	// Get the camera mid exposure timestamp
	cameraMidExpTime, tmpOffset := parseTimestamp(data[offset:])
	fs.CameraMidExposureTimestamp = cameraMidExpTime
	offset += tmpOffset

	// Get the camera data received timestamp
	cameraDataReceivedTime, tmpOffset := parseTimestamp(data[offset:])
	fs.CameraDataReceivedTimestamp = cameraDataReceivedTime
	offset += tmpOffset

	// Get the transmit timestamp
	transmitTime, tmpOffset := parseTimestamp(data[offset:])
	fs.TransmitTimestamp = transmitTime
	offset += tmpOffset

	// Get the precision timestamp seconds
	precisionTimestampSeconds, tmpOffset := parseUint32(data[offset:])
	fs.PrecisionTimestampSeconds = precisionTimestampSeconds
	offset += tmpOffset

	// Get the precision timestamp fraction seconds
	precisionTimestampFractionSeconds, tmpOffset := parseUint32(data[offset:])
	fs.PrecisionTimestampFractionSeconds = precisionTimestampFractionSeconds
	offset += tmpOffset

	// Get the params
	params, tmpOffset := parseParams(data[offset:])
	fs.Params = params
	offset += tmpOffset

	return fs, frameSuffixLen, nil
}

type FrameSuffixParams = Params

const (
	FrameSuffixParamsRecording FrameSuffixParams = 1 << iota
	FrameSuffixParamsTrackedModelsChanged
)

func (fs *FrameSuffix) ParamsIs(target FrameSuffixParams) bool {
	return fs.Params&target != 0
}
