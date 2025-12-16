// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"

	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/protocol"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/source"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/template"
	"github.com/aligator/haiku-golang-tools/gopls/internal/telemetry"
	"github.com/aligator/haiku-golang-tools/internal/event"
	"github.com/aligator/haiku-golang-tools/internal/event/tag"
)

func (s *Server) references(ctx context.Context, params *protocol.ReferenceParams) (_ []protocol.Location, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("references")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.references", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, source.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	if snapshot.FileKind(fh) == source.Tmpl {
		return template.References(ctx, snapshot, fh, params)
	}
	return source.References(ctx, snapshot, fh, params.Position, params.Context.IncludeDeclaration)
}
