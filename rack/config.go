// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

type Config struct {
	// Distance from the front to the back of the rack
	Depth float64

	// Distance from the bottom to the top
	Height float64

	// The width of a base support block
	BaseWidth float64

	// The height of a base support block
	BaseHeight float64

	// Radious of a connector pin
	PinRadius float64

	// Height of a connector pin
	PinHeight float64

	// Distance between a column and a base hinge
	DrawerHingeWidth float64

	// Width of the base of the drawer
	DrawerWidth float64

	// Width of the base block that holds the drawer against the column
	BaseHingeWidth float64

	// Height of the block that holds the drawer agains the columtn
	BaseHingeHeight float64
}

func NewDefaultConfig() Config {
	return Config{
		Depth:            100,
		Height:           120,
		BaseWidth:        10,
		BaseHeight:       5,
		PinRadius:        2.5,
		PinHeight:        5,
		DrawerHingeWidth: 2.5,
		DrawerWidth:      2.5,
		BaseHingeWidth:   5,
		BaseHingeHeight:  2.5,
	}
}

// Total size required to accomodate a hinge
func (c Config) HingeWidth() float64 {
	return c.BaseHingeWidth + c.DrawerHingeWidth
}
