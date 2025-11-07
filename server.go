package natnetgo

type ServerInfo struct {
	AppName       string
	AppVersion    Version
	NatNetVersion Version
}

const serverInfoMinLen = 256 + 2*versionLen

func decodeServerInfo(data []byte) (*ServerInfo, int, error) {
	if len(data) < serverInfoMinLen {
		return nil, 0, ErrTooShort
	}

	si := new(ServerInfo)

	// Get the application name.
	// It is always 256 bytes long
	appName, _, err := decodeString(data)
	if err != nil {
		return nil, 0, err
	}
	si.AppName = appName
	offset := 256

	// Get the application version
	appVersion, tmpOffset, err := decodeVersion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	si.AppVersion = appVersion
	offset += tmpOffset

	// Get the natnet version
	natNetVersion, tmpOffset, err := decodeVersion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	si.NatNetVersion = natNetVersion
	offset += tmpOffset

	return si, offset, nil
}
