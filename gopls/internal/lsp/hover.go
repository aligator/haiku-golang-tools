// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"

	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/mod"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/protocol"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/source"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/template"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/work"
	"github.com/aligator/haiku-golang-tools/gopls/internal/telemetry"
	"github.com/aligator/haiku-golang-tools/internal/event"
	"github.com/aligator/haiku-golang-tools/internal/event/tag"
)

func (s *Server) hover(ctx context.Context, params *protocol.HoverParams) (_ *protocol.Hover, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("hover")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.hover", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, source.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	switch snapshot.FileKind(fh) {
	case source.Mod:
		return mod.Hover(ctx, snapshot, fh, params.Position)
	case source.Go:
		return source.Hover(ctx, snapshot, fh, params.Position)
	case source.Tmpl:
		return template.Hover(ctx, snapshot, fh, params.Position)
	case source.Work:
		return work.Hover(ctx, snapshot, fh, params.Position)
	}
	return nil, nil
}
