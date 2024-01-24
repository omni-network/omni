package main

import (
	"os"
	"strings"

	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/runner/docker"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
	cmtdocker "github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"

	"github.com/spf13/cobra"
)

func main() {
	libcmd.Main(NewCLI().root)
}

// CLI is the Cobra-based command-line interface.
type CLI struct {
	root     *cobra.Command
	testnet  *e2e.Testnet
	preserve bool
	infp     infra.Provider
}

// NewCLI sets up the CLI.
func NewCLI() *CLI {
	cli := &CLI{}
	cli.root = &cobra.Command{
		Use:           "runner",
		Short:         "End-to-end test runner",
		SilenceUsage:  true,
		SilenceErrors: true, // we'll output them ourselves in Run()
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return errors.Wrap(err, "getting file")
			}
			m, err := e2e.LoadManifest(file)
			if err != nil {
				return errors.Wrap(err, "loading manifest")
			}

			ifd, err := e2e.NewDockerInfrastructureData(m)
			if err != nil {
				return errors.Wrap(err, "creating docker infrastructure data")
			}

			testnet, err := e2e.LoadTestnet(file, ifd)
			if err != nil {
				return errors.Wrap(err, "loading testnet")
			}

			cli.testnet = adaptTestnet(testnet)
			cli.infp = docker.NewProvider(testnet, ifd, defaultServices())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if err := Cleanup(ctx, cli.testnet); err != nil {
				return err
			}
			if err := Setup(ctx, cli.testnet, cli.infp); err != nil {
				return err
			}

			if err := Start(ctx, cli.testnet, cli.infp); err != nil {
				return err
			}

			if err := DeployContracts(ctx); err != nil {
				return err
			}

			if err := Wait(ctx, cli.testnet, 5); err != nil { // allow some txs to go through
				return err
			}

			if cli.testnet.HasPerturbations() {
				return errors.New("perturbations not supported yet")
			}

			if cli.testnet.Evidence > 0 {
				return errors.New("evidence injection not supported yet")
			}

			if err := Wait(ctx, cli.testnet, 5); err != nil { // wait for network to settle before tests
				return err
			}

			if !cli.preserve {
				if err := Cleanup(ctx, cli.testnet); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cli.root.PersistentFlags().StringP("file", "f", "", "Testnet TOML manifest")
	_ = cli.root.MarkPersistentFlagRequired("file")

	cli.root.Flags().BoolVarP(&cli.preserve, "preserve", "p", false,
		"Preserves the running of the test net after tests are completed")

	cli.root.AddCommand(&cobra.Command{
		Use:   "setup",
		Short: "Generates the testnet directory and configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Setup(cmd.Context(), cli.testnet, cli.infp)
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starts the testnet, waiting for nodes to become available",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			_, err := os.Stat(cli.testnet.Dir)
			if os.IsNotExist(err) {
				err = Setup(ctx, cli.testnet, cli.infp)
			}
			if err != nil {
				return errors.Wrap(err, "setup")
			}

			return Start(ctx, cli.testnet, cli.infp)
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "wait",
		Short: "Waits for a few blocks to be produced and all nodes to catch up",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Wait(cmd.Context(), cli.testnet, 5)
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "stop",
		Short: "Stops the testnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info(cmd.Context(), "Stopping testnet")

			return cli.infp.StopTestnet(cmd.Context())
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "cleanup",
		Short: "Removes the testnet directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Cleanup(cmd.Context(), cli.testnet)
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "logs",
		Short: "Shows the testnet logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmtdocker.ExecComposeVerbose(cmd.Context(), cli.testnet.Dir, "logs")
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "tail",
		Short: "Tails the testnet logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmtdocker.ExecComposeVerbose(cmd.Context(), cli.testnet.Dir, "logs", "--follow")
		},
	})

	return cli
}

func adaptTestnet(testnet *e2e.Testnet) *e2e.Testnet {
	// Move test dir: path/test/e2e/manifests/single -> path/test/e2e/runs/single
	testnet.Dir = strings.Replace(testnet.Dir, "manifests", "runs", 1)
	testnet.VoteExtensionsEnableHeight = 1
	testnet.UpgradeVersion = "omniops/halo:latest"
	for i := range testnet.Nodes {
		testnet.Nodes[i] = adaptNode(testnet.Nodes[i])
	}

	return testnet
}

func adaptNode(node *e2e.Node) *e2e.Node {
	node.Version = "omniops/halo:latest"
	node.PrivvalKey = k1.GenPrivKey()

	return node
}
