// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type RightConnector struct {
	Primitive Primitive
	Cfg       Config
}

func NewRightConnector(cfg Config) *RightConnector {
	return &RightConnector{Cfg: cfg}
}

func (o *RightConnector) Build() Primitive {
	pinHoleOffset := (o.Cfg.Width - o.Cfg.BaseWidth) / 2
	o.Primitive =
		NewDifference(
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.Width, o.Cfg.BaseHeight / 2}),
			NewTranslation(
				Vec3{0, pinHoleOffset, 0},
				NewCylinder(3*o.Cfg.BaseHeight, o.Cfg.PinRadius+0.1).SetFn(48)),
			NewTranslation(
				Vec3{0, -pinHoleOffset, 0},
				NewCylinder(3*o.Cfg.BaseHeight, o.Cfg.PinRadius+0.1).SetFn(48)),
		)
	return o.Primitive
}
