package fetch

import (
	cnet "github.com/openbiox/butils/net"
	"testing"
)

func TestEgafetch(t *testing.T) {
	opt := &cnet.BnetParams{
		Timeout: 35,
		Proxy:   "",
	}
	Egafetch("", "", "", opt)
}
