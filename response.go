package natnetgo

import "strings"

type Response struct {
	Code          int32
	Message       string
	StreamVersion string
}

const responseMinLen = int32Len

const bitstreamKeyworkd = "Bitstream"

func decodeResponse(data []byte) (*Response, int, error) {
	if len(data) < responseMinLen {
		return nil, 0, ErrTooShort
	}

	r := new(Response)

	if len(data) == responseMinLen {
		code, offset, _ := decodeInt32(data)
		r.Code = code

		return r, offset, nil
	}

	message, offset, err := decodeString(data[responseMinLen:])
	if err != nil {
		return nil, 0, err
	}
	r.Message = message

	// Decode the Bitstream version if present
	splMessages := strings.Split(message, ",")
	if len(splMessages) < 2 || splMessages[0] != bitstreamKeyworkd {
		return r, offset, nil
	}
	bsVersion := splMessages[1]
	r.StreamVersion = bsVersion

	return r, offset, nil
}
