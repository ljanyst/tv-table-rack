// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	//	"github.com/ljanyst/ghostscad/lib/utils"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// Base of the board drawer
type DrawerBase struct {
	Primitive Primitive
	Cfg       Config

	// Coordinate of the corner anchors: front-top, front-bottom,
	// back-top, back-bottom
	Corners []Vec2

	BaseScrew *Anchor
}

func NewDrawerBase(cfg Config) *DrawerBase {
	return &DrawerBase{
		Cfg: cfg,
	}
}

func (o *DrawerBase) Build() Primitive {
	// Offset of the front/back blocks
	offsetY := (o.Cfg.Depth - o.Cfg.BaseWidth) / 2

	// Offset oft the left/right blocks
	offsetX := (o.Cfg.Height/2 - o.Cfg.BaseHingeWidth)

	// Lenght of the screw attachment
	attLength := o.Cfg.BaseHingeWidth + o.Cfg.BaseWidth

	cx := o.Cfg.Height/2 - o.Cfg.BaseWidth
	cy := o.Cfg.Depth/2 - o.Cfg.BaseWidth
	o.Corners = []Vec2{
		{-cx, -cy},
		{cx, -cy},
		{-cx, cy},
		{cx, cy},
	}

	o.BaseScrew = NewAnchor()

	o.Primitive = NewList(
		// Back
		NewTranslation(
			Vec3{0, offsetY, 0},

			// Back block
			NewCube(Vec3{o.Cfg.Height, o.Cfg.BaseWidth, o.Cfg.DrawerWidth})),

		// Front
		NewTranslation(
			Vec3{0, -offsetY, 0},

			// Front block
			NewCube(Vec3{o.Cfg.Height, o.Cfg.BaseWidth, o.Cfg.DrawerWidth})),

		// Right
		NewTranslation(
			Vec3{offsetX, 0, 0},
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.Depth, o.Cfg.DrawerWidth})),

		// Left
		NewTranslation(
			Vec3{-offsetX, 0, 0},
			NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.Depth, o.Cfg.DrawerWidth})),

		// Screw attachment
		NewTranslation(
			Vec3{
				-offsetX - o.Cfg.BaseWidth/2 + 1.5*o.Cfg.BaseHingeHeight,
				offsetY + o.Cfg.BaseHingeWidth/2,
				(o.Cfg.BaseHingeWidth + o.Cfg.DrawerWidth) / 2},
			NewRotation(
				Vec3{0, -90, 0},
				NewDifference(
					NewCube(Vec3{
						o.Cfg.BaseHingeWidth,
						attLength,
						o.Cfg.BaseHingeHeight,
					}),
					NewTranslation(
						Vec3{0, (attLength - o.Cfg.BaseHingeWidth) / 2, 0},
						NewCylinder(3*o.Cfg.BaseHingeHeight, 1.5).SetFn(48),
						NewTranslation(
							Vec3{0, 0, o.Cfg.BaseHingeHeight / 2},
							o.BaseScrew),
					)))),
	)
	return o.Primitive
}
