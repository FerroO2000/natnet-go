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

	count, offset, err := parseInt32(data)
	if err != nil {
		return nil, 0, err
	}
	dc.FrameCount = count

	dc.FrameValues = make([]float32, 0, dc.FrameCount)
	for range dc.FrameCount {
		val, tmpOffset, err := parseFloat(data[offset:])
		if err != nil {
			return nil, 0, err
		}
		dc.FrameValues = append(dc.FrameValues, val)
		offset += tmpOffset
	}

	return dc, offset, nil
}
