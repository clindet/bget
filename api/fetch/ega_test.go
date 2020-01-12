package fetch

import (
	cnet "github.com/openbiox/ligo/net"
	"testing"
)

func TestEgafetch(t *testing.T) {
	opt := &cnet.Params{
		Timeout: 35,
		Proxy:   "",
	}
	Egafetch("", "", "", opt)
}
