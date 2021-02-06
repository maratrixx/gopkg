package netx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalIp(t *testing.T) {
	ip := InternalIp()
	assert.NotEqual(t, "", ip)

	if ip == "127.0.0.1" {
		t.Errorf("loopback not excluded")
	}
}
