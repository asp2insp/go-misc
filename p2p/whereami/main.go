package main

import (
	"fmt"

	"flag"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/asp2insp/go-misc/p2p/netutils"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {
	args := flag.Args()
	if len(args) != 1 {
		fmt.Print("usage: whereami portnum\n")
		os.Exit(1)
	}

	portString := args[0]
	port, err := strconv.Atoi(portString)
	if err != nil {
		fmt.Print("usage: whereami portnum\n")
		os.Exit(1)
	}
	source, err := QrForPort(port)
	if err != nil {
		fmt.Printf("Failed to generate QR code for '%s': %v\n", netutils.GetLocalAddress(port), err)
		os.Exit(1)
	}

	theme := flags.CreateTheme(driver)
	img := theme.CreateImage()

	mx := source.Bounds().Max
	window := theme.CreateWindow(mx.X, mx.Y, "WhereAmI")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(img)

	// Copy the image to a RGBA format before handing to a gxui.Texture
	rgba := image.NewRGBA(source.Bounds())
	draw.Draw(rgba, source.Bounds(), source, image.ZP, draw.Src)
	texture := driver.CreateTexture(rgba, 1)
	img.SetTexture(texture)

	window.OnClose(driver.Terminate)
}

func main() {
	flag.Parse()
	gl.StartDriver(appMain)
}

func QrForPort(port int) (barcode.Barcode, error) {
	code, err := qr.Encode(netutils.GetLocalAddress(port), qr.H, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(code, 400, 400)
}
