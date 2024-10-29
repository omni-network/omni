package coingecko

type options struct {
	Host   string
	APIKey string
}

func WithHost(host string) func(*options) {
	return func(o *options) {
		o.Host = host
	}
}

// WithAPIKey sets the API key for the client and updates the host to the pro API host if the host is the default.
func WithAPIKey(apikey string) func(*options) {
	return func(o *options) {
		if apikey == "" {
			return
		}

		o.APIKey = apikey

		if o.Host == defaultProdHost {
			o.Host = proProdHost
		}
	}
}

func defaultOpts() options {
	return options{
		Host: defaultProdHost,
	}
}
