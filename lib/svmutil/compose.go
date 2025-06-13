package svmutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"

	_ "embed"
)

var Version = "stable"

// Start starts a genesis solana node and returns a client, a funded private key and a stop function or an error.
// The dir parameter is the location of the docker compose.
func Start(ctx context.Context, composeDir string) (*rpc.Client, string, solana.PrivateKey, func(), error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute) // Allow 1 minute for edge case of pulling images.
	defer cancel()

	if !composeDown(ctx, composeDir) {
		return nil, "", nil, nil, errors.New("failure to clean up previous solana instance")
	}

	// Ensure ports are available
	port, err := getAvailablePort()
	if err != nil {
		return nil, "", nil, nil, errors.Wrap(err, "get available port")
	}

	if err := writeComposeFile(composeDir, port, Version); err != nil {
		return nil, "", nil, nil, errors.Wrap(err, "write compose file")
	}

	log.Info(ctx, "Starting solana")

	out, err := execCmd(ctx, composeDir, "docker", "compose", "up", "-d", "--remove-orphans")
	if err != nil {
		return nil, "", nil, nil, errors.Wrap(err, "docker compose up: "+out)
	}

	endpoint := "http://localhost:" + port

	cl := rpc.New(endpoint)

	stop := func() { //nolint:contextcheck // Fresh context required for stopping.
		// Fresh stop context since above context might be canceled.
		stopCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		composeDown(stopCtx, composeDir)
	}

	// Wait up to 10 secs for RPC to be available
	const retry = 20
	for i := 0; i < retry; i++ {
		if i == retry-1 {
			stop()
			return nil, "", nil, nil, errors.New("wait for RPC timed out")
		}

		select {
		case <-ctx.Done():
			return nil, "", nil, nil, errors.Wrap(ctx.Err(), "timeout")
		case <-time.After(time.Second):
		}

		_, err := cl.GetHealth(ctx)
		if err == nil {
			break
		}
		err = WrapRPCError(err, "health")

		if i > retry/2 {
			log.Warn(ctx, "Solana: waiting for RPC to be available", err)
		}
	}

	privKey, err := solana.PrivateKeyFromSolanaKeygenFile(filepath.Join(composeDir, "id.json"))
	if err != nil {
		return nil, "", nil, nil, errors.Wrap(err, "get private key")
	}

	log.Info(ctx, "Solana: RPC is available", "addr", endpoint)

	return cl, endpoint, privKey, stop, nil
}

// Rebuild rebuilds the program's so file with the provided private key.
// This effectively changes the program ID.
func Rebuild(ctx context.Context, program Program, key solana.PrivateKey, anchorDir string) (Program, error) {
	// Store privkey to <anchorDir>/tmp/<program>-keypair.json
	keyFile := filepath.Join(anchorDir, "tmp", program.Name+"-keypair.json")
	if err := SavePrivateKey(key, keyFile); err != nil {
		return Program{}, err
	}
	defer maybeDelete(keyFile)

	// Run <anchorDir>/build.sh tmp
	_, err := execCmd(ctx, anchorDir, "bash", "build.sh", "tmp")
	if err != nil {
		return Program{}, err
	}

	// Load the program from <anchorDir>/tmp/<program>.so
	soFile := filepath.Join(anchorDir, "tmp", program.SOFile())
	program.SharedObject, err = os.ReadFile(soFile)
	if err != nil {
		return Program{}, errors.Wrap(err, "read shared object file", "path", soFile)
	}

	key64, err := cast.Array64(key[:])
	if err != nil {
		return Program{}, err
	}
	program.KeyPairJSON, err = json.Marshal(key64)
	if err != nil {
		return Program{}, errors.Wrap(err, "marshal private key to JSON")
	}

	return program, nil
}

// Deploy deploys a program to the rpc network using the provided keys.
// It returns the confirmed transaction result or an error.
func Deploy(ctx context.Context, rpcAddr string, program Program, deployer, upgrader solana.PrivateKey) (*rpc.GetTransactionResult, error) {
	cl := rpc.New(rpcAddr)
	_, err := cl.GetAccountInfoWithOpts(ctx, program.MustPublicKey(), &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	if err == nil {
		return nil, errors.Wrap(err, "program already deployed", "account", program.MustPublicKey())
	} else if !errors.Is(err, rpc.ErrNotFound) {
		return nil, errors.Wrap(err, "get account info", "account", program.MustPublicKey())
	}

	return deploy(ctx, rpcAddr, program, deployer, upgrader)
}

// Redeploy redeployes/upgrades the program to the rpc network using the provided keys.
// It returns the confirmed transaction result or an error.
func Redeploy(ctx context.Context, rpcAddr string, program Program, upgrader solana.PrivateKey) (*rpc.GetTransactionResult, error) {
	cl := rpc.New(rpcAddr)
	info, err := cl.GetAccountInfoWithOpts(ctx, program.MustPublicKey(), &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get account info", "account", program.MustPublicKey())
	} else if !info.Value.Executable {
		return nil, errors.New("program not executable", "account", program.MustPublicKey())
	}

	return deploy(ctx, rpcAddr, program, upgrader, upgrader)
}

func deploy(ctx context.Context, rpcAddr string, program Program, deployer, upgrader solana.PrivateKey) (*rpc.GetTransactionResult, error) {
	hostDir, err := filepath.Abs("deploy")
	if err != nil {
		return nil, errors.Wrap(err, "absolute path")
	}
	if err := os.MkdirAll(hostDir, 0o755); err != nil {
		return nil, errors.Wrap(err, "mkdir temp")
	}
	defer maybeDelete(hostDir)

	m := mount{
		HostPath:      hostDir,
		ContainerPath: "/deploy",
	}

	// Store SO file
	if err := os.WriteFile(m.OnHost(program.SOFile()), program.SharedObject, 0644); err != nil {
		return nil, errors.Wrap(err, "write shared object file")
	}
	// Store program keypair file
	if err := SavePrivateKey(program.MustPrivateKey(), m.OnHost(program.KeyPairFile())); err != nil {
		return nil, errors.Wrap(err, "write keypair file")
	}
	// Store deployer keypair file
	deployerKeyFile := "deployer.json"
	if err := SavePrivateKey(deployer, m.OnHost(deployerKeyFile)); err != nil {
		return nil, errors.Wrap(err, "write deployer keypair file")
	}
	// Store upgrader keypair file
	upgraderKeyFile := "upgrader.json"
	if err := SavePrivateKey(upgrader, m.OnHost(upgraderKeyFile)); err != nil {
		return nil, errors.Wrap(err, "write upgrader keypair file")
	}

	// Since we run the program in a docker container, we need to replace localhost with host IP
	dockerAddr := rpcAddr
	if strings.HasPrefix(dockerAddr, "http://localhost") {
		hostIP, err := getHostIP(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get host IP")
		}
		dockerAddr = strings.ReplaceAll(dockerAddr, "localhost", hostIP)
		log.Debug(ctx, "Replacing localhost with host IP", "host_ip", hostIP, "docker_addr", dockerAddr, "rpc_addr", rpcAddr)
	}

	// TODO(corver): Switch to program-v4 once 'stable' supports --program-keypair flag

	args := []string{
		"run",
		"--rm",          // Remove the container after execution
		"-v", m.Mount(), // Mount the host directory to the container
		"--entrypoint", "solana", // Run CLI (instead of test-validator)
		"anzaxyz/agave:" + Version,                // Container
		"program",                                 // Program command
		"deploy", m.InContainer(program.SOFile()), // Loader subcommand
		// Args...
		"--program-id", m.InContainer(program.KeyPairFile()),
		"--keypair", m.InContainer(deployerKeyFile),
		"--upgrade-authority", m.InContainer(upgraderKeyFile),
		"--verbose",
		"--commitment", string(rpc.CommitmentConfirmed),
		"--url", dockerAddr,
		"--output", "json",
		"--use-rpc",
	}

	var stdOut, stdErr bytes.Buffer
	cmd := exec.CommandContext(ctx, "docker", args...)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "docker run deploy", "stderr", stdErr.String(), "stdout", stdOut.String(), "args", args)
	}

	var resp struct {
		ProgramID string `json:"programId"`
		Signature string `json:"signature"`
	}
	if err := json.Unmarshal(stdOut.Bytes(), &resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal deploy response", "stdout", stdOut.String())
	}

	if resp.ProgramID != program.MustPublicKey().String() {
		return nil, errors.New("unexpected program ID", "expected", program.MustPublicKey(), "got", resp.ProgramID)
	}

	txSig, err := solana.SignatureFromBase58(resp.Signature)
	if err != nil {
		return nil, errors.Wrap(err, "parse signature")
	}

	cl := rpc.New(rpcAddr)

	tx, err := AwaitConfirmedTransaction(ctx, cl, txSig)
	if err != nil {
		return nil, errors.Wrap(err, "await confirmed transaction")
	}

	info, err := cl.GetAccountInfoWithOpts(ctx, program.MustPublicKey(), &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get account info")
	} else if !info.Value.Executable {
		return nil, errors.New("program not executable", "account", program.MustPublicKey())
	}

	return tx, nil
}

// getHostIP retrieves the host's IP address by running a shell command.
// Ref: https://stackoverflow.com/questions/13322485/how-to-get-the-primary-ip-address-of-the-local-machine-on-linux-and-os-x
func getHostIP(ctx context.Context) (string, error) {
	out, err := exec.CommandContext(ctx, "bash", "-c", `ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.' | grep -v '172.' | head -n1`).CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, "get host IP", "output", string(out))
	}

	return strings.TrimSpace(string(out)), nil
}

type mount struct {
	HostPath      string
	ContainerPath string
}

func (m mount) Mount() string {
	return fmt.Sprintf("%s:%s", m.HostPath, m.ContainerPath)
}

func (m mount) OnHost(path string) string {
	return filepath.Join(m.HostPath, path)
}

func (m mount) InContainer(path string) string {
	return filepath.Join(m.ContainerPath, path)
}

func SendSimple(ctx context.Context, cl *rpc.Client, privkey solana.PrivateKey, instrs ...solana.Instruction) (solana.Signature, error) {
	return Send(ctx, cl, WithPrivateKeys(privkey), WithInstructions(instrs...))
}

type sendOpts struct {
	Instructions []solana.Instruction
	PrivateKeys  []solana.PrivateKey
}

func WithInstructions(instrs ...solana.Instruction) func(*sendOpts) {
	return func(opts *sendOpts) {
		opts.Instructions = append(opts.Instructions, instrs...)
	}
}

func WithPrivateKeys(privkeys ...solana.PrivateKey) func(*sendOpts) {
	return func(opts *sendOpts) {
		opts.PrivateKeys = append(opts.PrivateKeys, privkeys...)
	}
}

func Send(ctx context.Context, cl *rpc.Client, opts ...func(*sendOpts)) (solana.Signature, error) {
	var o sendOpts
	for _, opt := range opts {
		opt(&o)
	}
	if len(o.PrivateKeys) == 0 {
		return solana.Signature{}, errors.New("no private keys provided")
	}

	recent, err := cl.GetLatestBlockhash(ctx, rpc.CommitmentConfirmed)
	if err != nil {
		return solana.Signature{}, errors.Wrap(err, "get latest blockhash")
	}

	tx, err := solana.NewTransaction(
		o.Instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(o.PrivateKeys[0].PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, errors.Wrap(err, "new tx")
	}

	sigs, err := tx.Sign(func(pub solana.PublicKey) *solana.PrivateKey {
		for _, privkey := range o.PrivateKeys {
			if privkey.PublicKey().Equals(pub) {
				return &privkey
			}
		}

		return nil
	})
	if err != nil {
		return solana.Signature{}, errors.Wrap(err, "sign tx")
	} else if len(sigs) != len(o.PrivateKeys) {
		return solana.Signature{}, errors.New("unexpected number of signatures", "count", len(sigs))
	}

	txSig, err := cl.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{
		SkipPreflight: true, // Preflights cause "Program not deployed" errors, so skip for now.
	})
	if err != nil {
		return solana.Signature{}, errors.Wrap(err, "send tx")
	}

	return txSig, nil
}

// composeDown runs docker-compose down in the provided directory.
func composeDown(ctx context.Context, dir string) bool {
	if _, err := os.Stat(dir + "/compose.yaml"); os.IsNotExist(err) {
		return true
	}

	out, err := execCmd(ctx, dir, "docker", "compose", "down")
	if err != nil {
		log.Error(ctx, "Error: docker compose down", err, "out", out)
		return false
	}

	log.Debug(ctx, "Solana: docker compose down: ok")

	return true
}

func execCmd(ctx context.Context, dir string, cmd string, args ...string) (string, error) {
	c := exec.CommandContext(ctx, cmd, args...)
	c.Dir = dir

	out, err := c.CombinedOutput()
	if err != nil {
		return string(out), errors.Wrap(err, "exec", "out", string(out))
	}

	return string(out), nil
}

//go:embed compose.yaml.tmpl
var composeTpl []byte

func writeComposeFile(dir string, port, version string) error {
	tpl, err := template.New("").Parse(string(composeTpl))
	if err != nil {
		return errors.Wrap(err, "parse compose template")
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, struct {
		Port    string
		Version string
	}{
		Port:    port,
		Version: version,
	})
	if err != nil {
		return errors.Wrap(err, "execute compose template")
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "mkdir", "path", dir)
	}

	err = os.WriteFile(dir+"/compose.yaml", buf.Bytes(), 0644)
	if err != nil {
		return errors.Wrap(err, "write compose file")
	}

	return nil
}

func getAvailablePort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", errors.Wrap(err, "resolve addr")
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", errors.Wrap(err, "listen")
	}
	defer l.Close()

	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return "", errors.Wrap(err, "split host port")
	}

	return port, nil
}

func WrapRPCError(err error, endpoint string, attrs ...any) error {
	if err == nil {
		return nil
	}

	attrs = append(attrs, "endpoint", endpoint)

	rpcErr := new(jsonrpc.RPCError)
	if errors.As(err, &rpcErr) {
		attrs = append(attrs, "code", rpcErr.Code, "message", rpcErr.Message, "data", fmt.Sprintf("%+v", rpcErr.Data))
		return errors.New("solana rpc error", attrs...)
	}

	return errors.Wrap(err, "solana rpc error", attrs...)
}
