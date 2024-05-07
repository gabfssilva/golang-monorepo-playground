package clients

import (
	"github.com/gabfssilva/gomonorepo/libs/async"
	"github.com/gabfssilva/gomonorepo/libs/result"
	"github.com/go-resty/resty/v2"
)

func GetMyIp() result.Result[string] {
	client := resty.New()

	responseResult := result.Try(func() (*resty.Response, error) {
		return client.R().
			EnableTrace().
			Get("https://api.ipify.org?format=json")
	})

	return result.Map(responseResult, func(r *resty.Response) string {
		return r.String()
	})
}

func GetMyIps() result.Result[string] {
	client := resty.New()

	// TODO: how to make Result & Zip async to work together?
	ipify, myip, err := async.ZipEffectAsync(func() (*resty.Response, error) {
		return client.R().
			EnableTrace().
			Get("https://api.ipify.org?format=json")
	}, func() (*resty.Response, error) {
		return client.R().
			EnableTrace().
			Get("https://api.my-ip.io/v2/ip.json")
	})

	if err != nil {
		return result.Err[string](err)
	}

	return result.Ok(ipify.String() + myip.String())
}
