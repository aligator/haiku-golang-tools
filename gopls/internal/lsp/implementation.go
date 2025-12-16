// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"

	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/protocol"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/source"
	"github.com/aligator/haiku-golang-tools/gopls/internal/telemetry"
	"github.com/aligator/haiku-golang-tools/internal/event"
	"github.com/aligator/haiku-golang-tools/internal/event/tag"
)

func (s *Server) implementation(ctx context.Context, params *protocol.ImplementationParams) (_ []protocol.Location, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("implementation")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.implementation", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, source.Go)
	defer release()
	if !ok {
		return nil, err
	}
	return source.Implementation(ctx, snapshot, fh, params.Position)
}
