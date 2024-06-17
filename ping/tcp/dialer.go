package tcp

import (
	"context"
	"golang.org/x/net/proxy"
	"net"
)

// dialer 是一个封装了SOCKS5代理的net.Dialer
type dialer struct {
	d        *net.Dialer
	proxy    proxy.Dialer
	hasProxy bool
}

// DialContext 使用封装的SOCKS5代理进行TCP连接
func (d *dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if d.hasProxy && d.proxy != nil {
		return d.proxy.Dial(network, address)
	}

	return d.d.DialContext(ctx, network, address)
}
