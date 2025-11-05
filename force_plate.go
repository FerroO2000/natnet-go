package natnetgo

type ForcePlate struct {
	ID       ID
	Channels *List[*Channel]
}

const forcePlateMinLen = idLen + listMinLen

func parseForcePlate(data []byte) (*ForcePlate, int, error) {
	if len(data) < forcePlateMinLen {
		return nil, 0, ErrTooShort
	}

	fp := new(ForcePlate)

	id, _ := parseID(data)
	fp.ID = id

	channels, offset, err := parseList(parseChannel, data[idLen:])
	if err != nil {
		return nil, 0, err
	}
	fp.Channels = channels

	return fp, offset, nil
}
