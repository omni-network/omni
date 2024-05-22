// This file contains build time developer tools used in the omni repo.
// We use go.mod to control the versions of these tools and dependabot to keep them up to date.
// To install all the tools run: go generate tools.go
// For more info on this pattern, see https://www.jvt.me/posts/2022/06/15/go-tools-dependency-management/.

// Build tag is require to fool go.mod into adding main packages as dependencies.
// See reference: https://github.com/golang/go/issues/32504
//go:build tools

package scripts

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/cosmos/gogoproto/protoc-gen-gocosmos"
	_ "golang.org/x/tools/cmd/stringer"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

//go:generate echo Installing tools: protobuf
//go:generate go install github.com/bufbuild/buf/cmd/buf
//go:generate go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint
//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
//go:generate go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar
//go:generate go install github.com/cosmos/gogoproto/protoc-gen-gocosmos
//go:generate go install cosmossdk.io/orm/cmd/protoc-gen-go-cosmos-orm

//go:generate echo Installing tools: misc
//go:generate go install golang.org/x/tools/cmd/stringer
//go:generate go install go.uber.org/mock/mockgen@latest
