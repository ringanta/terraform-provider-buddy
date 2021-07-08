package buddy

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

type Config struct {
	BuddyURL  string
	Token     string
	VerifySSL bool
}

type buddyAdapter struct {
	BuddyURL string
	Token    string
	*http.Client
}

func newBuddyClient(c *Config) *buddyAdapter {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	config := &tls.Config{
		InsecureSkipVerify: !c.VerifySSL,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	httpClient := &http.Client{Transport: tr}

	return &buddyAdapter{BuddyURL: c.BuddyURL, Token: c.Token, Client: httpClient}
}
