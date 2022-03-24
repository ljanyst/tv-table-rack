// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	//	"github.com/ljanyst/ghostscad/lib/shapes"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// Disk drawer
type DrawerDisk struct {
	Primitive Primitive
	Cfg       Config
	Dims      Vec3

	BaseScrew *Anchor
}

func NewDrawerDisk(cfg Config, dims Vec3) *DrawerDisk {
	return &DrawerDisk{
		Cfg:  cfg,
		Dims: dims,
	}
}

func (o *DrawerDisk) Build() Primitive {
	base := NewDrawerBase(o.Cfg)
	base.Build()
	o.BaseScrew = base.BaseScrew

	rail := func(down bool) Primitive {
		rotation := Vec3{}
		if down {
			rotation = Vec3{0, 180, 0}
		}
		return NewRotation(
			rotation,
			NewDifference(
				NewCube(Vec3{o.Dims[2], 30, o.Dims[2]}),
				NewTranslation(
					Vec3{o.Dims[2] / 2, 0, 0},
					NewRotation(Vec3{90, 0, 0},
						NewCylinder(35, o.Dims[2]/2).SetFn(48)))))
	}

	offsetX := o.Dims[0] / 2
	offsetY := (o.Cfg.Depth - 30) / 2
	barWidth := (o.Cfg.Depth - o.Dims[1]) / 2
	o.Primitive = NewList(
		base.Primitive,
		NewTranslation(
			Vec3{0, 0, (o.Cfg.DrawerWidth + o.Dims[2]) / 2},
			NewTranslation(
				Vec3{offsetX, offsetY, 0},
				rail(true)),
			NewTranslation(
				Vec3{offsetX, -offsetY, 0},
				rail(true)),
			NewTranslation(
				Vec3{-offsetX, offsetY, 0},
				rail(false)),
			NewTranslation(
				Vec3{-offsetX, -offsetY, 0},
				rail(false)),
			NewTranslation(
				Vec3{0, (-o.Cfg.Depth + barWidth) / 2, 0},
				NewCube(Vec3{o.Dims[0], barWidth, o.Dims[2]})),
		),
	)
	return o.Primitive
}
