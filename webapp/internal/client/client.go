package client

import "net/url"

func Service() *MercuriusService {
	u, err := url.Parse("http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	return NewMercuriusService(u, "webclient", nil)
}
