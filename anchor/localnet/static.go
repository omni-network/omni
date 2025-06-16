package localnet

import _ "embed"

//go:embed solver_inbox.so
var InboxSO []byte

//go:embed solver_inbox-keypair.json
var InboxKeyPairJSON []byte

//go:embed dummy.so
var DummySO []byte

//go:embed dummy-keypair.json
var DummyKeyPairJSON []byte
