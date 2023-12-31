package client

import "github.com/khulnasoft-lab/shipyard-cli/pkg/requests"

type Client struct {
	Requester requests.Requester
	Org       string
}

func New(r requests.Requester, org string) Client {
	return Client{Requester: r, Org: org}
}
