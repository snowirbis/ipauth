package ipauth

import (
	"net"
)

type AuthHandler struct {
	allowedNets []*net.IPNet
}

func IPAuth(cidrs []string) (*AuthHandler, error) {
	var nets []*net.IPNet

	for _, cidr := range cidrs {
		_, n, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}
		nets = append(nets, n)
	}

	return &AuthHandler{nets}, nil
}

func (h *AuthHandler) Allowed(remote string) (bool, error) {

	addr, err := net.ResolveTCPAddr("tcp", remote+":0")

	if err != nil {
		return false, err
	}

	for _, n := range h.allowedNets {
		if n.Contains(addr.IP) {
			return true, nil
		}
	}

	return false, nil
}
