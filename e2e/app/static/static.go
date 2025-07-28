package static

import _ "embed"

//go:embed devnet-anvil-state.json
var elAnvilState []byte

func GetDevnetAnvilState() []byte {
	return elAnvilState
}

//go:embed devnet-deployments.json
var elDeployments []byte

func GetDevnetDeployments() []byte {
	return elDeployments
}
