package submit

import "github.com/spf13/pflag"

type Config struct {
	PrivKey   string
	EVMENR    string
	EVMAddr   string
	BatchSize uint64
}

func (c Config) Enabled() bool {
	return c.PrivKey != "" && c.BatchSize != 0 && c.EVMAddr != "" && c.EVMENR != ""
}

func BindFlags(flags *pflag.FlagSet, cfg *Config) {
	flags.StringVar(&cfg.PrivKey, "evmredenom-privkey", cfg.PrivKey, "Path to private key to submit transaction with")
	flags.StringVar(&cfg.EVMENR, "evmredenom-evm-enr", cfg.EVMENR, "EVM ENR to snap-sync account from")
	flags.StringVar(&cfg.EVMAddr, "evmredenom-evm-addr", cfg.EVMENR, "EVM RPC address to query and submit txs to")
	flags.Uint64Var(&cfg.BatchSize, "evmredenom-batch-size", cfg.BatchSize, "Size of account batch transactions to submit")
}
