package natnetgo

type Version struct {
	Major    uint8
	Minor    uint8
	Build    uint8
	Revision uint8
}

const versionLen = 4

func parseVersion(data []byte) (Version, int, error) {
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

type ServerInfo struct {
	AppName       string
	AppVersion    Version
	NatNetVersion Version
}

const serverInfoMinLen = 256 + 2*versionLen

func parseServerInfo(data []byte) (*ServerInfo, int, error) {
	if len(data) < serverInfoMinLen {
		return nil, 0, ErrTooShort
	}

	si := new(ServerInfo)

	// Get the application name.
	// It is always 256 bytes long
	appName, _, err := parseName(data)
	if err != nil {
		return nil, 0, err
	}
	si.AppName = appName
	offset := 256

	// Get the application version
	appVersion, tmpOffset, err := parseVersion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	si.AppVersion = appVersion
	offset += tmpOffset

	// Get the natnet version
	natNetVersion, tmpOffset, err := parseVersion(data[offset:])
	if err != nil {
		return nil, 0, err
	}
	si.NatNetVersion = natNetVersion
	offset += tmpOffset

	return si, offset, nil
}
