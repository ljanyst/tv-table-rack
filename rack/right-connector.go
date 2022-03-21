// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// The connector on the right side of the base
type RightConnector struct {
	Primitive Primitive
	Cfg       Config
	Type      BaseType

	BaseAttachment *Anchor
	LeftConnector  *Anchor
	FrontColumn    *Anchor
	BackColumn     *Anchor
}

func NewRightConnector(cfg Config, typ BaseType) *RightConnector {
	return &RightConnector{
		Cfg:  cfg,
		Type: typ,
	}
}

func (o *RightConnector) Build() Primitive {
	pinYOffset := (o.Cfg.Depth - o.Cfg.BaseWidth) / 2
	pinZOffset := o.Cfg.PinHeight - o.Cfg.BaseHeight/4
	pinHeight := (o.Cfg.BaseHeight + o.Cfg.PinHeight)
	attachmentZOffset := -o.Cfg.BaseHeight / 4
	leftConnectorZBase := -o.Cfg.BaseHeight / 4
	leftConnectorZOffset := o.Cfg.BaseHeight / 2
	columnZOffset := o.Cfg.BaseHeight

	// We do stuff at the other side if it's a top base
	if o.Type == Top {
		pinZOffset *= -1
		attachmentZOffset *= -1
		leftConnectorZBase *= -1
		leftConnectorZOffset = -o.Cfg.BaseHeight
		columnZOffset *= -1
	}

	o.BaseAttachment = NewAnchor()
	o.LeftConnector = NewAnchor()
	o.FrontColumn = NewAnchor()
	o.BackColumn = NewAnchor()

	o.Primitive =
		NewList(
			// Base block
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.Depth, o.Cfg.BaseHeight / 2}),

			// Back pin
			NewTranslation(
				Vec3{0, pinYOffset, pinZOffset},
				NewCylinder(pinHeight, o.Cfg.PinRadius).SetFn(48)),

			// Front pin
			NewTranslation(
				Vec3{0, -pinYOffset, pinZOffset},
				NewCylinder(pinHeight, o.Cfg.PinRadius).SetFn(48),
			),

			// Left anchor
			NewTranslation(
				Vec3{0, -pinYOffset, leftConnectorZBase},

				// Left connector anchor
				NewTranslation(
					Vec3{0, 0, leftConnectorZOffset},
					o.LeftConnector),

				// Front left column
				NewTranslation(
					Vec3{0, 0, columnZOffset},
					o.FrontColumn),
			),

			// Left connector anchor
			NewTranslation(
				Vec3{0, pinYOffset, leftConnectorZBase + columnZOffset},
				o.BackColumn),

			// The anchor for attaching the connector to a base
			NewTranslation(
				Vec3{-o.Cfg.BaseWidth / 2, -pinYOffset - o.Cfg.BaseWidth/2, attachmentZOffset},
				o.BaseAttachment))

	return o.Primitive
}
