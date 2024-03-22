package static

import _ "embed"

//go:embed el-anvil-state.json
var elAnvilState []byte

func GetDevnetElAnvilState() []byte {
	return elAnvilState
}

//go:embed el-deployments.json
var elDeployments []byte

func GetDevnetElDeployments() []byte {
	return elDeployments
}
