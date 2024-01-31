package main

import (
	"context"
	"os"
	"strings"

	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/runner/docker"
	"github.com/omni-network/omni/test/e2e/runner/netman"

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
	root    *cobra.Command
	testnet *e2e.Testnet
	infp    infra.Provider
	mngr    netman.Manager

	preserve       bool
	skipTests      bool
	deployKeyFile  string
	relayerKeyFile string
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
			m, err := LoadManifest(file)
			if err != nil {
				return errors.Wrap(err, "loading manifest")
			}

			mngr, err := netman.NewManager(m.Network, cli.deployKeyFile, cli.relayerKeyFile)
			if err != nil {
				return errors.Wrap(err, "get network")
			}

			ifd, err := e2e.NewDockerInfrastructureData(m.Manifest)
			if err != nil {
				return errors.Wrap(err, "creating docker infrastructure data")
			}

			testnet, err := e2e.LoadTestnet(file, ifd)
			if err != nil {
				return errors.Wrap(err, "loading testnet")
			}

			cli.testnet = adaptTestnet(testnet)
			cli.mngr = mngr
			cli.infp = docker.NewProvider(testnet, ifd, mngr.AdditionalService())

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if err := Cleanup(ctx, cli.testnet); err != nil {
				return err
			}

			// Deploy public portals first so their addresses are available for setup.
			if err := cli.mngr.DeployPublicPortals(ctx); err != nil {
				return err
			}

			if err := Setup(ctx, cli.testnet, cli.infp, cli.mngr); err != nil {
				return err
			}

			if err := Start(ctx, cli.testnet, cli.infp); err != nil {
				return err
			}

			if err := cli.mngr.DeployPrivatePortals(ctx); err != nil {
				return err
			}

			sendCtx, sendCancel := context.WithCancel(ctx)
			defer sendCancel()
			if err := StartSendingXMsgs(sendCtx, cli.mngr.Portals()); err != nil {
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

			sendCancel() // Stop sending messages

			if err := Wait(ctx, cli.testnet, 5); err != nil { // wait for network to settle before tests
				return err
			}

			if cli.skipTests {
				log.Info(ctx, "Skipping tests; --skip-tests=true")
			} else {
				if err := Test(ctx, cli.testnet, cli.infp.GetInfrastructureData(), cli.mngr.HostNetwork()); err != nil {
					return err
				}
			}

			if err := LogMetrics(ctx, cli.testnet, cli.mngr); err != nil {
				return err
			}

			if cli.preserve {
				log.Warn(ctx, "Docker containers not stopped, --preserve=true", nil)
			} else if err := Cleanup(ctx, cli.testnet); err != nil {
				return err
			}

			return nil
		},
	}

	cli.root.PersistentFlags().StringP("file", "f", "", "Testnet TOML manifest")
	_ = cli.root.MarkPersistentFlagRequired("file")

	cli.root.Flags().BoolVarP(&cli.skipTests, "skip-tests", "s", false,
		"Skips running tests, useful to just bootstrap a devnet (if used with -p)")

	cli.root.Flags().BoolVarP(&cli.preserve, "preserve", "p", false,
		"Preserves the running of the test net after tests are completed")

	cli.root.Flags().StringVar(&cli.deployKeyFile, "deploy-key", "",
		"Hex private key used to deploy public chain contracts (needs to be funded)")

	cli.root.Flags().StringVar(&cli.relayerKeyFile, "relayer-key", "",
		"Relayer's hex private key used for submissions on public chains (needs to be funded)")

	cli.root.AddCommand(&cobra.Command{
		Use:   "setup",
		Short: "Generates the testnet directory and configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Setup(cmd.Context(), cli.testnet, cli.infp, cli.mngr)
		},
	})

	cli.root.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starts the testnet, waiting for nodes to become available",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			_, err := os.Stat(cli.testnet.Dir)
			if os.IsNotExist(err) {
				err = Setup(ctx, cli.testnet, cli.infp, cli.mngr)
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
		Short: "Shows the container logs (except anvil)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			args := append([]string{"logs"}, cli.mngr.AdditionalService()...)

			return cmtdocker.ExecComposeVerbose(cmd.Context(), cli.testnet.Dir, args...)
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
	testnet.UpgradeVersion = "omniops/halo:main"
	for i := range testnet.Nodes {
		testnet.Nodes[i] = adaptNode(testnet.Nodes[i])
	}

	return testnet
}

func adaptNode(node *e2e.Node) *e2e.Node {
	node.Version = "omniops/halo:main"
	node.PrivvalKey = k1.GenPrivKey()

	return node
}
