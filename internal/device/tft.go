package device

import (
	"image/color"
	"machine"

	"github.com/jypelle/xship/internal/img"
	"tinygo.org/x/drivers/st7735"
)

type tftDev struct {
	st7735.Device
	tft565Buffer  []uint8
	width, height int32
}

func NewTftDevice(width, height int32) *tftDev {
	return &tftDev{
		Device:       st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE),
		tft565Buffer: make([]uint8, height*width*2),
		width:        width,
		height:       height,
	}
}

func (d *tftDev) configure() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 16000000,
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
	})
	d.Device.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
		Model:    st7735.GREENTAB,
	})
	d.FillScreen(color.RGBA{})
}

func (d *tftDev) SetPixel(x, y int32, color color.RGBA) {
	if color.A > 0 && x >= 0 && x < d.width && y >= 0 && y < d.height {
		c565 := st7735.RGBATo565(color)
		d.tft565Buffer[(y*d.width+x)*2] = uint8(c565 >> 8)
		d.tft565Buffer[(y*d.width+x)*2+1] = uint8(c565)
	}
}

var blackC565 = st7735.RGBATo565(color.RGBA{R: 0, G: 0, B: 0, A: 255})

func (d *tftDev) DrawImageFromAsset(imageAsset img.ImageAsset, offset img.Position, frame int) {
	fullXOffset := imageAsset.Offset[frame%len(imageAsset.Offset)].X
	fullYOffset := imageAsset.Offset[frame%len(imageAsset.Offset)].Y
	var pixelOffset int32
	for y := int32(0); y < imageAsset.Height; y++ {
		for x := int32(0); x < imageAsset.Width; x++ {
			pixelOffset = ((fullYOffset+y)*imageAsset.Asset.Width + fullXOffset + x) * 2
			d.SetPixelFromC565Alpha(
				int32(offset.X)+x, int32(offset.Y)+y,
				uint16(imageAsset.Asset.Data[pixelOffset])|uint16(imageAsset.Asset.Data[pixelOffset+1])<<8,
			)
		}
	}
}

func (d *tftDev) SetPixelFromC565(x, y int32, c565 uint16) {
	if x >= 0 && x < d.width && y >= 0 && y < d.height {
		d.tft565Buffer[(y*d.width+x)*2] = uint8(c565 >> 8)
		d.tft565Buffer[(y*d.width+x)*2+1] = uint8(c565)
	}
}

func (d *tftDev) SetPixelFromC565Alpha(x, y int32, c565 uint16) {
	if c565 > 0 && x >= 0 && x < d.width && y >= 0 && y < d.height {
		d.tft565Buffer[(y*d.width+x)*2] = uint8(c565 >> 8)
		d.tft565Buffer[(y*d.width+x)*2+1] = uint8(c565)
	}
}

func (d *tftDev) Clear() {
	for y := int32(0); y < d.height; y++ {
		for x := int32(0); x < d.width; x++ {
			d.SetPixelFromC565(x, y, blackC565)
		}
	}
}

func (d *tftDev) Refresh() error {

	// Set window
	d.Tx([]uint8{st7735.CASET}, true)
	d.Tx([]uint8{uint8(0 >> 8), uint8(0), uint8((d.width - 1) >> 8), uint8(d.width - 1)}, false)
	d.Tx([]uint8{st7735.RASET}, true)
	d.Tx([]uint8{uint8(0 >> 8), uint8(0), uint8((d.height - 1) >> 8), uint8(d.height - 1)}, false)
	d.Command(st7735.RAMWR)

	// Send pixels data
	d.Tx(d.tft565Buffer, false)

	return nil
}

func (d *tftDev) Width() int32 {
	return d.width
}

func (d *tftDev) Height() int32 {
	return d.height
}
