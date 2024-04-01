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
	DiskDims  Vec3
	UsbDims   Vec3

	BaseScrew *Anchor
}

func NewDrawerDisk(cfg Config, diskDims Vec3, usbDims Vec3) *DrawerDisk {
	return &DrawerDisk{
		Cfg:      cfg,
		DiskDims: diskDims,
		UsbDims:  usbDims,
	}
}

func (o *DrawerDisk) Build() Primitive {
	base := NewDrawerBase(o.Cfg)
	base.Build()
	o.BaseScrew = base.BaseScrew

	diskRail := func(down bool) Primitive {
		rotation := Vec3{}
		if down {
			rotation = Vec3{0, 180, 0}
		}
		return NewRotation(
			rotation,
			NewDifference(
				NewCube(Vec3{o.DiskDims[2], 30, o.DiskDims[2]}),
				NewTranslation(
					Vec3{o.DiskDims[2] / 2, 0, 0},
					NewRotation(Vec3{90, 0, 0},
						NewCylinder(35, o.DiskDims[2]/2)))))
	}

	diskOffsetX := o.DiskDims[0] / 2
	diskOffsetY := (o.Cfg.Depth - 30) / 2
	diskBarWidth := (o.Cfg.Depth - o.DiskDims[1]) / 2
	diskDrOffsetX := (o.Cfg.Height-o.DiskDims[0]-o.DiskDims[2])/2 - o.Cfg.BaseHingeHeight

	diskRails := NewList(
		NewTranslation(
			Vec3{diskOffsetX, diskOffsetY, 0},
			diskRail(true)),
		NewTranslation(
			Vec3{diskOffsetX, -diskOffsetY, 0},
			diskRail(true)),
		NewTranslation(
			Vec3{-diskOffsetX, diskOffsetY, 0},
			diskRail(false)),
		NewTranslation(
			Vec3{-diskOffsetX, -diskOffsetY, 0},
			diskRail(false)),
		NewTranslation(
			Vec3{0, (-o.Cfg.Depth + diskBarWidth) / 2, 0},
			NewCube(Vec3{o.DiskDims[0], diskBarWidth, o.DiskDims[2]})),
	)

	usbRail := func(down, front bool) Primitive {
		translation := Vec3{o.Cfg.BaseWidth / 2, 0, o.UsbDims[2] / 2}
		if down {
			translation = Vec3{-o.Cfg.BaseWidth / 2, 0, o.UsbDims[2] / 2}

		}
		rail := NewList(
			NewCube(Vec3{o.Cfg.BaseHingeHeight, o.Cfg.BaseWidth, o.UsbDims[2] + o.Cfg.BaseHingeHeight}),
			NewTranslation(
				translation,
				NewCube(Vec3{o.Cfg.BaseWidth, o.Cfg.BaseWidth, o.Cfg.BaseHingeHeight})),
		)
		if front {
			limitTrX := (o.Cfg.BaseWidth/2 + o.Cfg.BaseHingeHeight) / 2
			if down {
				limitTrX = -limitTrX
			}
			rail.Add(
				NewTranslation(
					Vec3{limitTrX, -(o.Cfg.BaseWidth - o.Cfg.BaseHingeHeight) / 2, 0},
					NewCube(Vec3{o.Cfg.BaseWidth / 2, o.Cfg.BaseHingeHeight, o.UsbDims[2]})),
			)
		}
		return rail
	}

	usbOffsetX := (o.UsbDims[0] + o.Cfg.BaseHingeHeight) / 2
	usbOffsetY := (o.Cfg.Depth - o.Cfg.BaseWidth) / 2
	usbDrOffsetX := (diskDrOffsetX-o.DiskDims[0])/2 - o.DiskDims[2]
	usbRails := NewList(
		NewTranslation(
			Vec3{usbOffsetX, usbOffsetY, 0},
			usbRail(true, false)),
		NewTranslation(
			Vec3{usbOffsetX, -usbOffsetY, 0},
			usbRail(true, true)),
		NewTranslation(
			Vec3{-usbOffsetX, usbOffsetY, 0},
			usbRail(false, false)),
		NewTranslation(
			Vec3{-usbOffsetX, -usbOffsetY, 0},
			usbRail(false, true)),
	)

	o.Primitive = NewList(
		base.Primitive,
		NewTranslation(
			Vec3{diskDrOffsetX, 0, (o.Cfg.DrawerWidth + o.DiskDims[2]) / 2},
			diskRails,
		),
		NewTranslation(
			Vec3{usbDrOffsetX, 0, (o.Cfg.DrawerWidth + o.UsbDims[2]) / 2},
			usbRails,
		),
	)
	return o.Primitive
}
