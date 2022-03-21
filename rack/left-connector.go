// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// The connector on the left side of the base
type LeftConnector struct {
	Primitive Primitive
	Cfg       Config
	Type      BaseType

	BaseAttachment *Anchor
}

func NewLeftConnector(cfg Config, typ BaseType) *LeftConnector {
	return &LeftConnector{
		Cfg:  cfg,
		Type: typ,
	}
}

func (o *LeftConnector) Build() Primitive {
	pinHoleOffset := (o.Cfg.Depth - o.Cfg.BaseWidth) / 2
	attachmentZOffset := -o.Cfg.BaseHeight / 4

	if o.Type == Top {
		attachmentZOffset *= -1
	}

	o.BaseAttachment = NewAnchor()
	o.Primitive =
		NewDifference(
			// Base block
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.Depth, o.Cfg.BaseHeight / 2}),

			// Back pin hole
			NewTranslation(
				Vec3{0, pinHoleOffset, 0},
				NewCylinder(3*o.Cfg.BaseHeight, o.Cfg.PinRadius+0.1).SetFn(48)),

			// Front pin hole
			NewTranslation(
				Vec3{0, -pinHoleOffset, 0},
				NewCylinder(3*o.Cfg.BaseHeight, o.Cfg.PinRadius+0.1).SetFn(48),

				// The anchor for attaching the connector to a base
				NewTranslation(
					Vec3{o.Cfg.BaseWidth / 2, -o.Cfg.BaseWidth / 2, attachmentZOffset},
					o.BaseAttachment)))
	return o.Primitive
}
