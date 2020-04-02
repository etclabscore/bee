// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package addressbook

import (
	"github.com/ethersphere/bee/pkg/swarm"

	ma "github.com/multiformats/go-multiaddr"
)

type GetPutter interface {
	Getter
	Putter
}

type GetPutRemover interface {
	GetPutter
	Remover
}

type Getter interface {
	Get(overlay swarm.Address) (addr ma.Multiaddr, exists bool)
}

type Putter interface {
	Put(overlay swarm.Address, addr ma.Multiaddr) (exists bool)
}

type Remover interface {
	Remove(overlay swarm.Address) (exists bool)
}
