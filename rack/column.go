// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// Column connecting the bases
type Column struct {
	Primitive Primitive
	Cfg       Config

	BaseBottom *Anchor
	BaseTop    *Anchor
}

func NewColumn(cfg Config) *Column {
	return &Column{Cfg: cfg}
}

func (o *Column) Build() Primitive {
	o.BaseBottom = NewAnchor()
	o.BaseTop = NewAnchor()

	o.Primitive =
		NewDifference(

			// Volumn block
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.BaseWidth, o.Cfg.Height}),

			// Top pin hole
			NewTranslation(
				Vec3{0, 0, o.Cfg.Height / 2},
				NewCylinder(2.2*o.Cfg.PinHeight, o.Cfg.PinRadius+0.1),
				o.BaseTop),

			// Bottom pin hole
			NewTranslation(
				Vec3{0, 0, -o.Cfg.Height / 2},
				NewCylinder(2.2*o.Cfg.PinHeight, o.Cfg.PinRadius+0.1),
				o.BaseBottom),
		)
	return o.Primitive
}
