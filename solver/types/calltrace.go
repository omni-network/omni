package types

// CallTrace is the result of debug_traceCall. Types kept to strings for simplicity.
// TODO(kevin): move to lib/ethclient.
type CallTrace struct {
	From    string      `json:"from"`
	To      string      `json:"to"`
	Data    string      `json:"data"`
	Value   string      `json:"value"`
	Type    string      `json:"type"`
	Gas     string      `json:"gas"`
	GasUsed string      `json:"gasUsed"`
	Input   string      `json:"input"`
	Output  string      `json:"output"`
	Error   string      `json:"error"`
	Calls   []CallTrace `json:"calls,omitempty"`
}

// Map converts CallTrace to a map[string]any.
func (t CallTrace) Map() map[string]any {
	out := map[string]any{
		"from":    t.From,
		"to":      t.To,
		"data":    t.Data,
		"value":   t.Value,
		"type":    t.Type,
		"gas":     t.Gas,
		"gasUsed": t.GasUsed,
		"input":   t.Input,
	}

	if t.Output != "" {
		out["output"] = t.Output
	}

	if t.Error != "" {
		out["error"] = t.Error
	}

	if len(t.Calls) > 0 {
		calls := make([]map[string]any, len(t.Calls))
		for i, call := range t.Calls {
			calls[i] = call.Map()
		}

		out["calls"] = calls
	}

	return out
}
