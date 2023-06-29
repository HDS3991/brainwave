package router

import "brainwave/internal/router/base"

type Gr struct {
	base.Router
}

var Group = new(Gr)
