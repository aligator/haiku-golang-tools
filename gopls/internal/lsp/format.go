// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"

	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/mod"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/protocol"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/source"
	"github.com/aligator/haiku-golang-tools/gopls/internal/lsp/work"
	"github.com/aligator/haiku-golang-tools/internal/event"
	"github.com/aligator/haiku-golang-tools/internal/event/tag"
)

func (s *Server) formatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	ctx, done := event.Start(ctx, "lsp.Server.formatting", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, source.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	switch snapshot.FileKind(fh) {
	case source.Mod:
		return mod.Format(ctx, snapshot, fh)
	case source.Go:
		return source.Format(ctx, snapshot, fh)
	case source.Work:
		return work.Format(ctx, snapshot, fh)
	}
	return nil, nil
}
