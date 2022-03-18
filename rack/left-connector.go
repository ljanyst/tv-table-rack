// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type LeftConnector struct {
	Primitive Primitive
	Cfg       Config
}

func NewLeftConnector(cfg Config) *LeftConnector {
	return &LeftConnector{Cfg: cfg}
}

func (o *LeftConnector) Build() Primitive {
	pinYOffset := (o.Cfg.Width - o.Cfg.BaseWidth) / 2
	pinZOffset := o.Cfg.PinHeight - o.Cfg.BaseHeight/4
	pinHeight := (o.Cfg.BaseHeight + o.Cfg.PinHeight)
	o.Primitive =
		NewList(
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.Width, o.Cfg.BaseHeight / 2}),
			NewTranslation(
				Vec3{0, pinYOffset, pinZOffset},
				NewCylinder(pinHeight, o.Cfg.PinRadius).SetFn(48)),
			NewTranslation(
				Vec3{0, -pinYOffset, pinZOffset},
				NewCylinder(pinHeight, o.Cfg.PinRadius).SetFn(48)),
		)
	return o.Primitive
}
