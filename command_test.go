package natnetgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeConnect(t *testing.T) {
	assert := assert.New(t)

	expected := []byte{
		0, 0, // command type
		5, 0, // command length
		80, 105, 110, 103, // "Ping"
		0,
	}

	data, dataLen := EncodeConnect()
	assert.Equal(expected, data)
	assert.Equal(len(expected), dataLen)
}

func Test_EncodeRequest(t *testing.T) {
	assert := assert.New(t)

	expected := []byte{
		2, 0, // command type
		5, 0, // command length
		116, 101, 115, 116, // "test"
		0,
	}

	data, dataLen := EncodeRequest("test")
	assert.Equal(expected, data)
	assert.Equal(len(expected), dataLen)
}

func Test_EncodeRequestModelDescData(t *testing.T) {
	assert := assert.New(t)

	expected := []byte{
		4, 0, // command type
		1, 0, // command length
		0,
	}

	data, dataLen := EncodeRequestModelDescData()
	assert.Equal(expected, data)
	assert.Equal(len(expected), dataLen)
}

func Test_EncodeRequestMoCapData(t *testing.T) {
	assert := assert.New(t)

	expected := []byte{
		6, 0, // command type
		1, 0, // command length
		0,
	}

	data, dataLen := EncodeRequestMoCapData()
	assert.Equal(expected, data)
	assert.Equal(len(expected), dataLen)
}

func Test_EncodeKeepAlive(t *testing.T) {
	assert := assert.New(t)

	expected := []byte{
		10, 0, // command type
		1, 0, // command length
		0,
	}

	data, dataLen := EncodeKeepAlive()
	assert.Equal(expected, data)
	assert.Equal(len(expected), dataLen)
}
