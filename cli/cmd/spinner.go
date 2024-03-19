package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

//nolint:containedctx // This is a CLI command and the context is not used in the function.
type spinnerDevFund struct {
	spinner spinner.Model
	// ctx is stored to preserve logging settings across the CLI command execution.
	// This is necessary to ensure that the logging context, which includes configurations
	// such as log level and output formatting.
	ctx  context.Context
	cfg  devnetFundConfig
	done bool
	err  error
}

type fundingCompleteMsg struct{}
type fundingFailedMsg struct {
	err error
}

func spinnedDevnetFund(ctx context.Context, cfg devnetFundConfig) error {
	p := spinnedDevnetFundProgram(ctx, cfg)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}

	return nil
}

func spinnedDevnetFundProgram(ctx context.Context, cfg devnetFundConfig) *tea.Program {
	return tea.NewProgram(devnetFundSpinner(ctx, cfg))
}

func devnetFundSpinner(ctx context.Context, cfg devnetFundConfig) tea.Model {
	return &spinnerDevFund{
		spinner: spinner.New(spinner.WithSpinner(spinner.MiniDot)),
		ctx:     ctx,
		cfg:     cfg,
	}
}

func typedDevnetFund(ctx context.Context, cfg devnetFundConfig) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second)
		err := devnetFund(ctx, cfg)
		if err != nil {
			return fundingFailedMsg{err: err}
		}

		return fundingCompleteMsg{}
	}
}

// Init initializes the spinner and starts the funding process.
// It returns a command that will be executed immediately after the model is
// initialized. `Init` is a required method of the `tea.Model` interface.
func (m spinnerDevFund) Init() tea.Cmd {
	// Assuming you have access to the context here, pass it to fundAccount
	// If you don't, you'll need to adjust your architecture to ensure it's available
	return tea.Batch(m.spinner.Tick, typedDevnetFund(m.ctx, m.cfg))
}

// Update handles messages and updates the model.
// It's a required method of the `tea.Model` interface.
func (m *spinnerDevFund) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

	case fundingCompleteMsg:
		m.done = true

	case fundingFailedMsg:
		m.done = true
		m.err = msg.err

	default:
		if m.done {
			time.Sleep(1500 * time.Millisecond)
			return m, tea.Quit
		}
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd
	}

	return m, nil
}

// View renders the spinner and the current status of the funding process.
// It's a required method of the `tea.Model` interface.
func (m *spinnerDevFund) View() string {
	if m.err != nil {
		return fmt.Sprintf("Funding failed: %v", m.err)
	}
	if m.done {
		return "Funding complete! ✓"
	}

	return m.spinner.View() + " Funding account..."
}
