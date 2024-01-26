package main

import (
	_ "embed"
	"image/color"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

const (
	WIDTH  = 320
	HEIGHT = 240
)

var rainbow []color.RGBA
var pressed uint8
var quit bool

func Badge() {
	quit = false
	display.FillScreen(colors[BLACK])

	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	for {
		myNameIsRainbow("Denis")
		if quit {
			break
		}
		blinkyRainbow("@zwindler", "Denis Germain")
		if quit {
			break
		}
		comeToMyTalk()
	}
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 10

	// black corners detail
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

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

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "HELLO")
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 34, "HELLO", colors[WHITE])

	w32, _ = tinyfont.LineWidth(&freesans.Oblique12pt7b, "my name is")
	tinyfont.WriteLine(&display, &freesans.Oblique12pt7b, (WIDTH-int16(w32))/2, 54, "my name is", colors[WHITE])

	// gophers
	tinyfont.WriteLineColors(&display, &gophers.Regular58pt, WIDTH-84, 208, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
}

func myNameIsRainbow(name string) {
	myNameIs(name)

	w32, size := getFontWidthSize(name)
	for i := 0; i < 100; i++ {
		if size == 24 {
			tinyfont.WriteLineColors(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else if size == 18 {
			tinyfont.WriteLineColors(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else if size == 12 {
			tinyfont.WriteLineColors(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		} else {
			tinyfont.WriteLineColors(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 140, name, rainbow[i:])
		}
		i += 2
		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func blinkyRainbow(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, sizetop := getFontWidthSize(topline)
	w32bottom, sizebottom := getFontWidthSize(bottomline)
	for i := int16(0); i < 12; i++ {
		// show black text
		if sizetop == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32top))/2, 90, topline, getRainbowRGB(uint8(i*12)))
		} else if sizetop == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32top))/2, 90, topline, getRainbowRGB(uint8(i*12)))
		} else if sizetop == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32top))/2, 90, topline, getRainbowRGB(uint8(i*12)))
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32top))/2, 90, topline, getRainbowRGB(uint8(i*12)))
		}
		if sizebottom == 24 {
			tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (WIDTH-int16(w32bottom))/2, 180, bottomline, getRainbowRGB(uint8(i*12)))
		} else if sizebottom == 18 {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32bottom))/2, 180, bottomline, getRainbowRGB(uint8(i*12)))
		} else if sizebottom == 12 {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 180, bottomline, getRainbowRGB(uint8(i*12)))
		} else {
			tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32bottom))/2, 180, bottomline, getRainbowRGB(uint8(i*12)))
		}

		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func comeToMyTalk() {
	display.FillScreen(colors[WHITE])

	for i := 0; i < 10; i++ {
		w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, "Come to my talk")
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 35, "Come to my talk", colors[RED])

		w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "Putting an end to")
		tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 90, "Putting an end to", colors[BLACK])

		w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "Makefiles in go proj.")
		tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 130, "Makefiles in go proj.", colors[BLACK])

		w32, _ = tinyfont.LineWidth(&freesans.Oblique18pt7b, "with GoReleaser")
		tinyfont.WriteLine(&display, &freesans.Oblique18pt7b, (WIDTH-int16(w32))/2, 170, "with GoReleaser", colors[BLACK])

		w32, _ = tinyfont.LineWidth(&freesans.Bold18pt7b, "ud2218a Sat. 15:00")
		tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (WIDTH-int16(w32))/2, 225, "ud2218a Sat. 15:00", colors[RED])

		if !btnB.Get() {
			quit = true
			break
		}
	}
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
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
