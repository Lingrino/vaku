package vaku

import (
	"github.com/hashicorp/vault/api"
)

// logical is functions from api.Logical() used by Vaku.
type logical interface {
	Delete(path string) (*api.Secret, error)
	List(path string) (*api.Secret, error)
	Read(path string) (*api.Secret, error)
	Write(path string, data map[string]interface{}) (*api.Secret, error)
}

// Client has all Vaku functions and wraps Vault API clients.
type Client struct {
	// src is the default client and also used as dst when dst is nil.
	src *api.Client
	dst *api.Client

	// wrap api.Client.Logical() in an interface.
	srcL logical
	dstL logical

	// workers is the max number of concurrent operations against vault.
	workers int

	// absolutepath if the absolution path is desired instead of the relative path.
	absolutepath bool
}

// Option configures a Client.
type Option interface {
	apply(c *Client) error
}

// WithVaultClient sets the default Vault client to be used.
func WithVaultClient(c *api.Client) Option {
	return withVaultClient{c}
}

// WithVaultSrcClient is an alias for WithVaultClient.
func WithVaultSrcClient(c *api.Client) Option {
	return withVaultClient{c}
}

type withVaultClient struct {
	client *api.Client
}

func (o withVaultClient) apply(c *Client) error {
	c.src = o.client
	c.srcL = o.client.Logical()
	return nil
}

// WithVaultDstClient sets a separate Vault client to be used only on operations that have a source
// and destination (copy, move, etc...). If unset the default client will be used as the source and
// destination.
func WithVaultDstClient(c *api.Client) Option {
	return withDstVaultClient{c}
}

type withDstVaultClient struct {
	client *api.Client
}

func (o withDstVaultClient) apply(c *Client) error {
	c.dst = o.client
	c.dstL = o.client.Logical()
	return nil
}

// WithWorkers sets the maximum number of goroutines that access Vault at any given time. Does not
// cap the number of goroutines overall. Default value is 10. A stable and well-operated Vault
// server should be able to handle 100 or more without issue. Use with caution and tune specifically
// to your environment and storage backend.
func WithWorkers(n int) Option {
	return withWorkers(n)
}

type withWorkers uint

func (o withWorkers) apply(c *Client) error {
	c.workers = int(o)
	return nil
}

// WithAbsolutePath sets the output format for all returned paths. Default path output is a relative
// path, trimmed up to the path input. Pass WithAbsolutePath(true) to set path output to the entire
// path. Example: List(secret/foo) -> "bar" OR "secret/foo/bar"
func WithAbsolutePath(b bool) Option {
	return withAbsolutePath(b)
}

type withAbsolutePath bool

func (o withAbsolutePath) apply(c *Client) error {
	c.absolutepath = bool(o)
	return nil
}

// NewClient returns a new Vaku Client based on the Vault API config.
func NewClient(opts ...Option) (*Client, error) {
	// set defaults
	client := &Client{
		workers: 10,
	}

	// apply options
	for _, opt := range opts {
		err := opt.apply(client)
		if err != nil {
			return nil, err
		}
	}

	// set dst to src if dst is unspecified
	if client.dst == nil && client.src != nil {
		client.dst = client.src
		client.dstL = client.srcL
	}

	return client, nil
}
