package submit

import "github.com/spf13/pflag"

type Config struct {
	Enabled     bool
	PrivKey     string
	EVMENR      string
	RPCSubmit   string
	RPCArchive  string
	BatchSize   uint64
	Concurrency int64
	Genesis     string
}

func BindFlags(flags *pflag.FlagSet, cfg *Config) {
	flags.BoolVar(&cfg.Enabled, "evmredenom-submit-enable", cfg.Enabled, "Enable EVM redenomination account batch submission")
	flags.StringVar(&cfg.PrivKey, "evmredenom-privkey", cfg.PrivKey, "Path to redenom contract owner private key to submit transaction with")
	flags.StringVar(&cfg.EVMENR, "evmredenom-evm-enr", cfg.EVMENR, "EVM ENR to snap-sync account from. Must be local paired geth instance.")
	flags.StringVar(&cfg.RPCSubmit, "evmredenom-rpc-submit", cfg.RPCSubmit, "EVM RPC address to submit txs to. Must not be local paired geth instance.")
	flags.StringVar(&cfg.RPCArchive, "evmredenom-rpc-archive", cfg.RPCArchive, "EVM RPC address to query preimages from. Must be archive.")
	flags.StringVar(&cfg.Genesis, "evmredenom-genesis", cfg.Genesis, "Path to evm genesis file; for preimage lookups")
	flags.Uint64Var(&cfg.BatchSize, "evmredenom-batch-size", cfg.BatchSize, "Size of account batch transactions to submit (in bytes)")
	flags.Int64Var(&cfg.Concurrency, "evmredenom-concurrency", cfg.Concurrency, "Number of concurrent account batch submissions (1 = sequential)")
}
