package views

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type ParkingSpace struct {
	Container     *fyne.Container
	Background    *canvas.Rectangle
	CarImage      *canvas.Image
	NumberLabel   *canvas.Text
	StatusText    *canvas.Text
	OccupiedImage string
}

var (
	backgroundColor = &color.NRGBA{R: 40, G: 44, B: 52, A: 255}
	availableImage  = "src/assets/si.png"
	occupiedImages  = []string{"src/assets/cars/car-black.png", "src/assets/cars/car-blue.png", "src/assets/cars/car-green.png", "src/assets/cars/car-red.png", "src/assets/cars/car-purple.png"}
	borderColor     = &color.NRGBA{R: 255, G: 250, B: 0, A: 255}
	textColor       = &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
)

func NewParkingSpace(number int) *ParkingSpace {
	space := &ParkingSpace{}

	const (
		parkingSpaceWidth  = 60
		parkingSpaceHeight = 100
	)

	space.Background = canvas.NewRectangle(backgroundColor)
	space.Background.SetMinSize(fyne.NewSize(parkingSpaceWidth, parkingSpaceHeight))
	space.Background.Resize(fyne.NewSize(parkingSpaceWidth, parkingSpaceHeight))
	space.Background.StrokeWidth = 0.5
	space.Background.StrokeColor = borderColor

	space.CarImage = canvas.NewImageFromFile(availableImage)
	space.CarImage.SetMinSize(fyne.NewSize(90, 150))
	space.CarImage.FillMode = canvas.ImageFillContain
	space.NumberLabel = canvas.NewText(fmt.Sprintf("%d", number), textColor)
	space.NumberLabel.TextSize = 16
	space.NumberLabel.TextStyle = fyne.TextStyle{Bold: true}
	space.NumberLabel.Alignment = fyne.TextAlignCenter

	space.StatusText = canvas.NewText("LIBRE", textColor)
	space.StatusText.TextSize = 12
	space.StatusText.Alignment = fyne.TextAlignCenter

	space.Container = container.NewStack(
		space.Background,
		container.NewPadded(
			container.NewVBox(
				container.NewCenter(space.NumberLabel),
				container.NewCenter(space.CarImage),
				container.NewCenter(space.StatusText),
			),
		),
	)

	return space
}

func (p *ParkingSpace) UpdateStatus(occupied bool, carID int) {
	if occupied {
		if p.OccupiedImage == "" {
			p.OccupiedImage = occupiedImages[rand.Intn(len(occupiedImages))]
		}
		p.CarImage.File = p.OccupiedImage
		p.StatusText.Text = fmt.Sprintf("Carro #%d", carID)
	} else {
		p.CarImage.File = availableImage
		p.StatusText.Text = "LIBRE"
		p.OccupiedImage = ""
	}
	p.CarImage.Refresh()
	p.StatusText.Refresh()
}
