// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

type Config struct {
	Width            float64
	Height           float64
	BaseWidth        float64
	BaseHeight       float64
	PinRadius        float64
	PinHeight        float64
	DrawerHingeWidth float64
	DrawerWidth      float64
	BaseHingeWidth   float64
}

func NewDefaultConfig() Config {
	return Config{
		Width:            100,
		Height:           120,
		BaseWidth:        10,
		BaseHeight:       5,
		PinRadius:        2.5,
		PinHeight:        5,
		DrawerHingeWidth: 2.5,
		DrawerWidth:      5,
		BaseHingeWidth:   5,
	}
}

func (c Config) HingeWidth() float64 {
	return c.BaseHingeWidth + c.DrawerHingeWidth
}
