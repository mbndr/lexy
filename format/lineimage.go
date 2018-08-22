package format

// NOT COMPLETE / WORKING

/*
// TODO use other color lib 

import (
	"io"
	"image"
	"image/color"
	"image/draw"
	//"log"
	"image/png"

	"gopkg.in/go-playground/colors.v1"

	"github.com/mbndr/lexy"
)

var (
	// black
	defaultForeground = &color.RGBA{0, 0, 0, 255}
	// white
	defaultBackground = &color.RGBA{255, 255, 255, 255}
)

// basically the same as lexy.Style, only with converted colors
// TODO do better
// TODO on hljs theme download -> to color.RGBA??
// if color not parsable -> foreground
// if foreground not parsable -> defaultForeground
// warning if sth is not parsable
type RGBAStyle struct {
	Foreground *color.RGBA
	Background *color.RGBA

	TokenColors map[lexy.TokenType]*color.RGBA
}

func styleToRGBAStyle(s lexy.Style) RGBAStyle {
	rgbas := RGBAStyle{}
	rgbas.TokenColors = make(map[lexy.TokenType]*color.RGBA)

	rgbas.Foreground = stringColorToRGBA(s.Foreground)
	rgbas.Background = stringColorToRGBA(s.Background)

	if rgbas.Foreground == nil {
		rgbas.Foreground = defaultForeground
	}
	if rgbas.Background == nil {
		rgbas.Background = defaultBackground
	}

	rgbas.TokenColors[lexy.TokenKeyword] = stringColorToRGBA(s.TokenColors[lexy.TokenKeyword])
	rgbas.TokenColors[lexy.TokenLiteral] = stringColorToRGBA(s.TokenColors[lexy.TokenLiteral])
	rgbas.TokenColors[lexy.TokenBuiltin] = stringColorToRGBA(s.TokenColors[lexy.TokenBuiltin])
	rgbas.TokenColors[lexy.TokenOperator] = stringColorToRGBA(s.TokenColors[lexy.TokenOperator])
	rgbas.TokenColors[lexy.TokenComment] = stringColorToRGBA(s.TokenColors[lexy.TokenComment])
	rgbas.TokenColors[lexy.TokenString] = stringColorToRGBA(s.TokenColors[lexy.TokenString])
	rgbas.TokenColors[lexy.TokenNumber] = stringColorToRGBA(s.TokenColors[lexy.TokenNumber])

	return rgbas
}

func stringColorToRGBA(s string) *color.RGBA {
	c, err := colors.Parse(s)
	if err != nil {
		return nil
		//return color.RGBA{255, 153, 0, 255} // TODO other indicator (currently orange)
	}

	rgba := c.ToRGBA()
	return &color.RGBA{rgba.R, rgba.G, rgba.R, 255}
}



type LineImageFormatter struct {
	Style lexy.Style
	w io.Writer
	rgbaStyle RGBAStyle
}

func NewLineImage(s lexy.Style, w io.Writer) LineImageFormatter {
	return LineImageFormatter{Style: s, w: w, rgbaStyle: styleToRGBAStyle(s)}
}


func (f *LineImageFormatter) Format(tokens []lexy.Token) error {

	// TODO fill complete image with bg

	img := image.NewRGBA(image.Rect(0, 0, 100, 500))

	draw.Draw(img, img.Bounds(), &image.Uniform{f.rgbaStyle.Background}, image.ZP, draw.Src)

	//log.Fatalf("%#v", f.rgbaStyle.TokenColors[lexy.TokenInvalid])

	curLine := 1
	curCol := 1

	for _, t := range tokens {

		if t.Typ == lexy.TokenWS {
			// for each char in ws
			for _, ch := range t.Val {
				if ch == '\n' {
					curLine++
					curCol = 1
				}
				if ch == '\t' {
					curCol += 4 // TODO configurable
				}
				if ch == ' ' {
					curCol++
				}
			}
			continue
		}

		// no whitespace token
		for i := 0; i < len(t.Val); i++ {

			c := f.rgbaStyle.TokenColors[t.Typ]
			if c == nil {
				c = f.rgbaStyle.Foreground
			}

			img.Set(curCol, curLine * 2, c)
			curCol++
		}





		// draw line always with one pixel margin

	}

	

	png.Encode(f.w, img)

	return nil
}
*/