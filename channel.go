package natnetgo

type Channel struct {
	FrameCount  int32
	FrameValues []float32
}

const channelMinLen = int32Len

func parseChannel(data []byte) (*Channel, int, error) {
	if len(data) < channelMinLen {
		return nil, 0, ErrTooShort
	}

	dc := new(Channel)

	count, _ := parseInt32(data)
	dc.FrameCount = count

	offset := channelMinLen
	dc.FrameValues = make([]float32, 0, dc.FrameCount)
	for range dc.FrameCount {
		val, tmpOffset := parseFloat(data[offset:])
		dc.FrameValues = append(dc.FrameValues, val)
		offset += tmpOffset
	}

	return dc, offset, nil
}
