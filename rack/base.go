// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	"github.com/ljanyst/ghostscad/lib/utils"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

type BaseType int

const (
	Bottom BaseType = iota
	Top
)

// Top or Bottom support of the rack
type Base struct {
	Primitive Primitive
	Cfg       Config
	Type      BaseType
	Width     float64

	LeftConnector  *Anchor
	FrontColumn    *Anchor
	BackColumn     *Anchor
	RightConnector *Anchor
	DrawerScrew    *Anchor
}

func NewBase(cfg Config, width float64, typ BaseType) *Base {
	return &Base{
		Cfg:   cfg,
		Type:  typ,
		Width: width,
	}
}

func (o *Base) Build() Primitive {
	// Width of the base is the width required for the drawer and the hinge
	width := o.Width + o.Cfg.HingeWidth()

	// Offset of the front/back blocks
	offsetY := (o.Cfg.Depth - o.Cfg.BaseWidth) / 2

	// Z offset of the right connector's connecting anchor
	rightConnectorZOffset := -o.Cfg.BaseHeight / 2

	// Z offest of the hinge
	hingeZOffset := o.Cfg.BaseHeight / 2

	// There is not scre attachment in the bottom base
	var screwAttachment Primitive
	screwAttachment = NewNothing()

	// We need to make some adjustments for the top base
	if o.Type == Top {
		// Hinges are at the bottom if we're creating the top base
		rightConnectorZOffset *= -1
		hingeZOffset *= -1
		hingeZOffset -= o.Cfg.BaseHingeHeight

		// The small extension of the drawer hinge to allow for screwing the drawer
		// to the base
		o.DrawerScrew = NewAnchor()
		screwAttachment = NewTranslation(
			Vec3{
				o.Cfg.BaseHingeWidth / 2,
				o.Cfg.BaseWidth + o.Cfg.BaseHingeWidth/2,
				o.Cfg.BaseHingeHeight / 2},
			NewDifference(
				NewCube(Vec3{
					o.Cfg.BaseHingeWidth,
					o.Cfg.BaseHingeWidth,
					o.Cfg.BaseHingeHeight,
				}),
				NewCylinder(3*o.Cfg.BaseHingeHeight, 1.5).SetFn(48),
				NewTranslation(
					Vec3{0, 0, -o.Cfg.BaseHingeHeight / 2},
					o.DrawerScrew),
			))
	}

	// Connectors connecting this base to other bases
	lc := NewLeftConnector(o.Cfg, o.Type)
	rc := NewRightConnector(o.Cfg, o.Type)
	lc.Build()
	rc.Build()

	o.LeftConnector = rc.LeftConnector
	o.RightConnector = lc.RightConnector
	o.FrontColumn = rc.FrontColumn
	o.BackColumn = rc.BackColumn

	o.Primitive =
		NewList(
			// Back
			NewTranslation(
				Vec3{0, offsetY, 0},

				// Back block
				NewCube(Vec3{width, o.Cfg.BaseWidth, o.Cfg.BaseHeight}),

				// Back hinge
				NewTranslation(
					Vec3{
						-width/2 + o.Cfg.DrawerHingeWidth + 0.25,
						-o.Cfg.BaseWidth / 2,
						hingeZOffset,
					},

					// Back hinge block
					NewCube(Vec3{
						o.Cfg.BaseHingeWidth,
						o.Cfg.BaseWidth,
						o.Cfg.BaseHingeHeight,
					}).SetCenter(false),

					// Back screw attachment if it's a Top base
					screwAttachment)),

			// Front
			NewTranslation(
				Vec3{0, -offsetY, 0},

				// Front block
				NewCube(Vec3{width, o.Cfg.BaseWidth, o.Cfg.BaseHeight}),

				// Right connector
				NewTranslation(
					Vec3{width / 2, -o.Cfg.BaseWidth / 2, rightConnectorZOffset},
					utils.AlignHere(rc.BaseAttachment).Add(rc.Primitive)),

				// Left connector
				NewTranslation(
					Vec3{-width / 2, -o.Cfg.BaseWidth / 2, 0},
					utils.AlignHere(lc.BaseAttachment).Add(lc.Primitive)),

				// Front hinge
				NewTranslation(
					Vec3{
						-width/2 + o.Cfg.DrawerHingeWidth + 0.25,
						-o.Cfg.BaseWidth / 2,
						hingeZOffset,
					},

					// Front hinge block
					NewCube(Vec3{
						o.Cfg.BaseHingeWidth,
						o.Cfg.BaseWidth,
						o.Cfg.BaseHingeHeight,
					}).SetCenter(false))),
		)
	return o.Primitive
}
