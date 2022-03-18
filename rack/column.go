// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type Column struct {
	Primitive Primitive
	Cfg       Config
}

func NewColumn(cfg Config) *Column {
	return &Column{Cfg: cfg}
}

func (o *Column) Build() Primitive {
	o.Primitive =
		NewDifference(
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.BaseWidth, o.Cfg.Height}),
			NewTranslation(
				Vec3{0, 0, o.Cfg.Height / 2},
				NewCylinder(2.2*o.Cfg.PinHeight, o.Cfg.PinRadius+0.1).SetFn(48)),
			NewTranslation(
				Vec3{0, 0, -o.Cfg.Height / 2},
				NewCylinder(2.2*o.Cfg.PinHeight, o.Cfg.PinRadius+0.1).SetFn(48)),
		)
	return o.Primitive
}
