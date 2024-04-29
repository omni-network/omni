package coingecko

type options struct {
	Host string
}

func WithHost(host string) func(*options) {
	return func(o *options) {
		o.Host = host
	}
}

func defaultOpts() options {
	return options{
		Host: prodHost,
	}
}
