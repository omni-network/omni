package submit

import "github.com/spf13/pflag"

type Config struct {
	PrivKey   string
	GethENR   string
	BatchSize uint64
}

func BindFlags(flags *pflag.FlagSet, cfg *Config) {
	flags.StringVar(&cfg.PrivKey, "evmredenom-privkey", cfg.PrivKey, "Path to private key to submit transaction with")
	flags.StringVar(&cfg.GethENR, "evmredenom-gethenr", cfg.GethENR, "Geth ENR to snap-sync account from")
	flags.StringVar(&cfg.GethENR, "evmredenom-batchsize", cfg.GethENR, "Size of account batch transactions to submit")
}
