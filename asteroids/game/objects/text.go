package objects

import (
	"asteroids/game/core"
	"asteroids/game/math"
	"image/color"
	"strings"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	ebiten_text "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type AlignVertical uint8

const (
	AlignVerticalTop AlignVertical = iota
	AlignVerticalCenter
	AlignVerticalBottom
)

type AlignHorizontal uint8

const (
	AlignHorizontalLeft AlignHorizontal = iota
	AlignHorizontalCenter
	AlignHorizontalRight
)

type GrowVertical uint8

const (
	GrowVerticalDown GrowVertical = iota
	GrowVerticalUp
	GrowVerticalNone
)

type GrowHorizontal uint8

const (
	GrowHorizontalRight GrowHorizontal = iota
	GrowHorizontalLeft
	GrowHorizontalNone
)

func NewText(
	text string,
	position core.Vector[float64],
	fontFace font.Face,
	color color.Color,
	lineSpacing float64,
) (Text, error) {
	m := fontFace.Metrics()
	capHeight := math.Abs(float64(m.CapHeight.Floor()))
	lineHeight := float64(m.Height.Floor())

	if lineSpacing != 1 {
		h := float64(m.Height.Round()) * lineSpacing
		fontFace = ebiten_text.FaceWithLineHeight(fontFace, math.Round(h))
	}

	//position.Y += float64(fontFace.Metrics().Ascent.Floor())
	return Text{
		X:          position.X,
		Y:          position.Y,
		Text:       text,
		Color:      color,
		Visible:    true,
		fontFace:   fontFace,
		capHeight:  capHeight,
		lineHeight: lineHeight,
	}, nil
}

type Text struct {
	X float64
	Y float64

	Width  float64
	Height float64

	Text string

	Color           color.Color
	BackgroundColor color.RGBA

	AlignVertical   AlignVertical
	AlignHorizontal AlignHorizontal
	GrowVertical    GrowVertical
	GrowHorizontal  GrowHorizontal

	Visible bool

	fontFace   font.Face
	capHeight  float64
	lineHeight float64
}

func (text *Text) estimateHeight(numLines int) float64 {
	// Начинаем с высоты, которая нам потребуется для первой строки
	estimatedHeight := text.capHeight
	if numLines >= 2 {
		// Добавляем высоту для всех остальных строк
		estimatedHeight += (float64(numLines) - 1) * text.lineHeight
	}
	return estimatedHeight
}

func (text *Text) Draw(screen *ebiten.Image) {
	if !text.Visible || text.Text == "" {
		return
	}

	posX := text.X
	posY := text.Y + text.capHeight

	var (
		containerX0 float64
		containerY0 float64
		containerX1 float64
		containerY1 float64
	)

	bounds, _ := font.BoundString(text.fontFace, text.Text)
	boundsWidth := float64(bounds.Max.X)
	boundsHeight := float64(bounds.Max.Y)

	if text.Width == 0 && text.Height == 0 {
		// Автоматическое проставление рабочей области
		containerX0 = posX
		containerY0 = posY
		containerX1 = posX + boundsWidth
		containerY1 = posY + boundsHeight
	} else {
		containerX0 = posX
		containerY0 = posY
		containerX1 = posX + text.Width
		containerY1 = posY + text.Height
		if delta := boundsWidth - text.Width; delta > 0 {
			switch text.GrowHorizontal {
			case GrowHorizontalRight:
				containerX1 += delta
			case GrowHorizontalLeft:
				containerX0 -= delta
			case GrowHorizontalNone:
				// Ничего не делаем
			}
		}
		if delta := boundsHeight - text.Height; delta > 0 {
			switch text.GrowVertical {
			case GrowVerticalDown:
				containerY1 += delta
			case GrowVerticalUp:
				containerY0 -= delta
				posY -= delta
			case GrowVerticalNone:
				// Ничего не делаем
			}
		}
	}

	var (
		containerWidth  float64 = containerX1 - containerX0
		containerHeight float64 = containerY1 - containerY0
	)

	if text.BackgroundColor.A != 0 {
		// Пытаюсь уместить вызов DrawRect по ширине...
		x0 := containerX0
		y0 := containerY0 - text.capHeight
		w := containerWidth
		h := containerHeight
		vector.DrawFilledRect(
			screen,
			float32(x0),
			float32(y0),
			float32(w),
			float32(h),
			text.BackgroundColor,
			false,
		)
	}

	numLines := strings.Count(text.Text, "\n") + 1
	switch text.AlignVertical {
	case AlignVerticalTop:
		// Ничего не делаем
	case AlignVerticalCenter:
		posY += (containerHeight - text.estimateHeight(numLines)) / 2
	case AlignVerticalBottom:
		posY += containerHeight - text.estimateHeight(numLines)
	}

	options := &ebiten.DrawImageOptions{}
	options.ColorScale.ScaleWithColor(text.Color)

	if text.AlignHorizontal == AlignHorizontalLeft {
		options.GeoM.Translate(posX, posY)
		ebiten_text.DrawWithOptions(screen, text.Text, text.fontFace, options)
		return
	}

	// Нужно обрабатывать текст построчно, выравнивая каждую
	// строку отдельно
	textRemaining := text.Text
	offsetY := 0.0
	for {
		nextLine := strings.IndexByte(textRemaining, '\n')
		lineText := textRemaining

		if nextLine != -1 {
			lineText = textRemaining[:nextLine]
			textRemaining = textRemaining[nextLine+len("\n"):]
		}

		lineBounds, _ := font.BoundString(text.fontFace, lineText)
		lineBoundsWidth := float64(lineBounds.Max.X)
		offsetX := 0.0

		switch text.AlignHorizontal {
		case AlignHorizontalCenter:
			offsetX = (containerWidth - lineBoundsWidth) / 2
		case AlignHorizontalRight:
			offsetX = containerWidth - lineBoundsWidth
		}

		options.GeoM.Reset()
		options.GeoM.Translate(posX+offsetX, posY+offsetY)
		ebiten_text.DrawWithOptions(screen, lineText, text.fontFace, options)

		if nextLine == -1 {
			break
		}

		offsetY += text.lineHeight
	}
}

func (text *Text) Update() error {
	return nil
}
