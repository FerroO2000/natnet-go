package natnetgo

type Device struct {
	ID       ID
	Channels *List[*Channel]
}

const deviceMinLen = idLen + listMinLen

func parseDevice(data []byte) (*Device, int, error) {
	if len(data) < deviceMinLen {
		return nil, 0, ErrTooShort
	}

	d := new(Device)

	id, _ := parseID(data)
	d.ID = id

	channels, offset, err := parseList(parseChannel, data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	d.Channels = channels

	return d, offset, nil
}
