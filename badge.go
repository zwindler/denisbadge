package main

import (
	_ "embed"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

var (
	ledColors []color.RGBA
	rainbow   []color.RGBA
	pressed   uint8
	quit      bool
)

func Badge() {
	quit = false
	display.FillScreen(colors[BLACK])

	// Initialize the "2-led strip"
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds = ws2812.New(neo)
	ledColors = make([]color.RGBA, 2)
	ledColors[0] = color.RGBA{0, 0, 100, 255}
	ledColors[1] = color.RGBA{100, 0, 0, 255}
	leds.WriteColors(ledColors)

	for {
		myNameIs("Denis")
		time.Sleep(time.Second * 5)
		blinkyDenis("@zwindler", "Denis Germain")
		time.Sleep(time.Second * 5)
		comeToMyTalk()
		time.Sleep(time.Second * 5)
	}
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 10
	/*
		// black corners detail
		display.FillRectangle(0, 0, r, r, colors[BLACK])
		display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
		display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
		display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])
	*/
	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[RED])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[RED])

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, colors[RED])
	display.FillRectangle(0, r, WIDTH, 54, colors[RED])

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[RED])
	display.FillRectangle(0, HEIGHT-3*r-1, WIDTH, 2*r, colors[RED])

	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "HELLO")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 34, "HELLO", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique12pt7b, "my name is")
	tinyfont.WriteLine(&display, &freesans.Oblique12pt7b, (WIDTH-int16(w32))/2, 54, "my name is", colors[WHITE])

	// gophers
	tinyfont.WriteLine(&display, &gophers.Regular58pt, WIDTH-84, 208, "BE", colors[RED])

	w32, _ = getFontWidthSize(name)
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, colors[RED])
}

func blinkyDenis(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := getFontWidthSize(topline)
	w32bottom, _ := getFontWidthSize(bottomline)
	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 90, topline, colors[RED])

	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 180, bottomline, colors[RED])
}

func comeToMyTalk() {
	display.FillScreen(colors[WHITE])

	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "Come to my talk")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 35, "Come to my talk", colors[RED])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "Putting an end to")
	tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 90, "Putting an end to", colors[BLUE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "Makefiles in go proj.")
	tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 130, "Makefiles in go proj.", colors[BLUE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "with GoReleaser")
	tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 170, "with GoReleaser", colors[BLUE])

	w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, "ud2218a Sat. 15:00")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 225, "ud2218a Sat. 15:00", colors[RED])

}

func getFontWidthSize(text string) (w32 uint32, size byte) {
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, text)
	size = 24
	if w32 < 300 {
		size = 24
	} else {
		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, text)
		if w32 < 300 {
			size = 18
		} else {
			w32, _ = tinyfont.LineWidth(&freesans.Bold12pt7b, text)
			if w32 < 300 {
				size = 12
			} else {
				w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, text)
				size = 9
			}
		}
	}
	return
}
