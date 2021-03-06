package dns

import (
	"context"

	"github.com/Dreamacro/clash/component/dialer"

	D "github.com/miekg/dns"
)

type client struct {
	*D.Client
	Address string
}

func (c *client) Exchange(m *D.Msg) (msg *D.Msg, err error) {
	return c.ExchangeContext(context.Background(), m)
}

func (c *client) ExchangeContext(ctx context.Context, m *D.Msg) (msg *D.Msg, err error) {
	c.Client.Dialer = dialer.Dialer()

	// Please note that miekg/dns ExchangeContext doesn't respond to context cancel.
	msg, _, err = c.Client.ExchangeContext(ctx, m, c.Address)
	return
}
