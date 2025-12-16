// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gopls (pronounced “go please”) is an LSP server for Go.
// The Language Server Protocol allows any text editor
// to be extended with IDE-like features;
// see https://langserver.org/ for details.
//
// See https://github.com/golang/tools/blob/master/gopls/README.md
// for the most up-to-date documentation.
package main // import "github.com/aligator/haiku-golang-tools/gopls"

//go:generate go run doc/generate.go

import (
	"context"
	"os"

	"github.com/aligator/haiku-golang-tools/gopls/internal/hooks"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/cmd"
	"github.com/aligator/haiku-golang-tools/gopls/internal/telemetry"
	"github.com/aligator/haiku-golang-tools/internal/tool"
)

func main() {
	telemetry.Start()
	ctx := context.Background()
	tool.Main(ctx, cmd.New("gopls", "", nil, hooks.Options), os.Args[1:])
}
