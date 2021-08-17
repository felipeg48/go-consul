package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	CONSUL_SERVER = "localhost:8500"
	SERVICE = "another"
	TAG = "primary"
	ADDRESS = "127.0.0.1"
	PORT = 8000
)

var (
	C *client
	Error error
)

func TestMain(m *testing.M) {
	C, Error = NewConsulClient(CONSUL_SERVER);
	code := m.Run()
	os.Exit(code)
}


func TestClient_Register(t *testing.T) {
	assert.NoError(t, Error)
	assert.NotNil(t, C)

	error := C.Register(SERVICE,SERVICE,ADDRESS,PORT, []string { TAG })
	assert.NoError(t, error)
}

func TestClient_Service(t *testing.T) {
	assert.NoError(t, Error)
	assert.NotNil(t, C)

	s,_, error := C.Service(SERVICE,TAG)
	assert.NoError(t, error)
	assert.NotNil(t, s)

	assert.Equal(t,SERVICE,s[0].Service.Service)
	assert.Equal(t, ADDRESS,s[0].Service.Address)
	assert.Equal(t, PORT,s[0].Service.Port)
}

func TestClient_DeRegister(t *testing.T) {
	assert.NoError(t, Error)
	assert.NotNil(t, C)

	error := C.DeRegister(SERVICE)
	assert.NoError(t, error)
}
