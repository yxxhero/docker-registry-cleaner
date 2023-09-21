package util

import (
	"crypto/tls"
	"net/http"

	"github.com/heroku/docker-registry-client/registry"
)

func WarpAuthTransport(registryURL, uname, passwd string, insecure bool) http.RoundTripper {
	baseTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	return registry.WrapTransport(baseTransport, registryURL, uname, passwd)
}
