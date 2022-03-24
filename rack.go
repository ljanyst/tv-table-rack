package main

import (
	"github.com/ljanyst/ghostscad/lib/utils"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
	. "github.com/ljanyst/tv-table-rack/rack"
)

func main() {
	cfg := NewDefaultConfig()

	basesWidth := []float64{
		30, // Power board
		30, // Power board
		30, // Power board
		35, // Router board
		35, // Router board
		35, // KVM
		25, // Pi
		10, // Disk T5
		10, // Disk T7
	}
	basesBottom := []*Base{}
	basesTop := []*Base{}
	rack := NewList()

	// Deal with the bases
	for i, w := range basesWidth {
		// Create the bases
		basesBottom = append(basesBottom, NewBase(cfg, w, Bottom))
		basesTop = append(basesTop, NewBase(cfg, w, Top))
		basesBottom[i].Build()
		basesTop[i].Build()

		// Compute the transform of the base
		var tr *Transform
		if i == 0 {
			tr = NewTranslation(Vec3{})
		} else {
			tr = utils.Align(
				basesBottom[i-1].RightConnector,
				basesBottom[i].LeftConnector,
			)
		}

		// Attach the bottom base
		baseB := basesBottom[i]
		rack.Add(tr.Add(baseB.Primitive))

		// Attach the columns
		col := NewColumn(cfg)
		col.Build()
		tr = utils.Align(baseB.FrontColumn, col.BaseBottom)
		rack.Add(tr.Add(col.Primitive))
		col = NewColumn(cfg)
		col.Build()
		tr = utils.Align(baseB.BackColumn, col.BaseBottom)
		rack.Add(tr.Add(col.Primitive))

		// Attach the top base
		baseT := basesTop[i]
		tr = utils.Align(col.BaseTop, baseT.BackColumn)
		rack.Add(tr.Add(baseT.Primitive))
	}

	// Leftmost blind
	rBlindBottom := NewRightConnector(cfg, Bottom)
	rBlindTop := NewRightConnector(cfg, Top)
	rBlindBottom.Build()
	rBlindTop.Build()
	tr := utils.Align(
		basesBottom[len(basesBottom)-1].RightConnector,
		rBlindBottom.LeftConnector,
	)
	rack.Add(tr.Add(rBlindBottom.Primitive))
	tr = utils.Align(
		basesTop[len(basesTop)-1].RightConnector,
		rBlindTop.LeftConnector,
	)
	rack.Add(tr.Add(rBlindTop.Primitive))

	// Leftmost column
	col := NewColumn(cfg)
	col.Build()
	tr = utils.Align(rBlindBottom.FrontColumn, col.BaseBottom)
	rack.Add(tr.Add(col.Primitive))
	col = NewColumn(cfg)
	col.Build()
	tr = utils.Align(rBlindBottom.BackColumn, col.BaseBottom)
	rack.Add(tr.Add(col.Primitive))

	// Rightmost blind
	lBlindBottom := NewLeftConnector(cfg, Bottom)
	lBlindTop := NewLeftConnector(cfg, Top)
	lBlindBottom.Build()
	lBlindTop.Build()
	tr = utils.Align(basesBottom[0].LeftConnector, lBlindBottom.RightConnector)
	rack.Add(tr.Add(lBlindBottom.Primitive))
	tr = utils.Align(basesTop[0].LeftConnector, lBlindTop.RightConnector)
	rack.Add(tr.Add(lBlindTop.Primitive))

	// Power drawers
	var pd *DrawerPower
	for i := 0; i < 3; i++ {
		pd = NewDrawerPower(cfg)
		pd.Build()
		tr = utils.Align(basesTop[i].DrawerScrew, pd.BaseScrew)
		rack.Add(tr.Add(pd.Primitive))
	}

	// Router drawers
	var rd *DrawerRouter
	for i := 3; i < 5; i++ {
		rd = NewDrawerRouter(cfg)
		rd.Build()
		tr = utils.Align(basesTop[i].DrawerScrew, rd.BaseScrew)
		rack.Add(tr.Add(rd.Primitive))
	}

	// Raspberry Pis
	rpd := []*DrawerRPi{
		NewDrawerRPi(cfg, true),
		NewDrawerRPi(cfg, false),
	}

	for i := 5; i < 7; i++ {
		rpd[i-5].Build()
		tr = utils.Align(basesTop[i].DrawerScrew, rpd[i-5].BaseScrew)
		rack.Add(tr.Add(rpd[i-5].Primitive))
	}

	// Disks
	dd := []*DrawerDisk{
		NewDrawerDisk(cfg, Vec3{57, 85, 8}),
		NewDrawerDisk(cfg, Vec3{57, 75, 11}),
	}

	for i := 7; i < 9; i++ {
		dd[i-7].Build()
		tr = utils.Align(basesTop[i].DrawerScrew, dd[i-7].BaseScrew)
		rack.Add(tr.Add(dd[i-7].Primitive))
	}

	sys.RenderMultiple(map[string]Primitive{
		"right-blind":        rBlindBottom.Primitive,
		"left-blind":         lBlindBottom.Primitive,
		"column":             col.Primitive,
		"base-power-bottom":  basesBottom[0].Primitive,
		"base-power-top":     basesTop[0].Primitive,
		"base-router-bottom": basesBottom[3].Primitive,
		"base-router-top":    basesTop[3].Primitive,
		"power-drawer":       pd.Primitive,
		"base-rpi-bottom":    basesBottom[6].Primitive,
		"base-rpi-top":       basesTop[6].Primitive,
		"base-kvm-bottom":    basesBottom[5].Primitive,
		"base-kvm-top":       basesTop[5].Primitive,
		"base-disk-bottom":   basesBottom[7].Primitive,
		"base-disk-top":      basesTop[7].Primitive,
		"router-drawer":      rd.Primitive,
		"rpi-drawer":         rpd[1].Primitive,
		"kvm-drawer":         rpd[0].Primitive,
		"disk-drawer-t5":     dd[1].Primitive,
		"disk-drawer-t7":     dd[0].Primitive,
		"rack":               rack,
	}, "rack")
}
