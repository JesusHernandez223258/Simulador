package main

import (
	"Parking-Simulator/src/models"
	view "Parking-Simulator/src/views"
	"log"

	"fyne.io/fyne/v2/app"
)

const (
	capacidadEstacionamiento = 20
	totalCarros              = 100
)

func main() {
	application := app.New()

	// Modificamos para manejar el error retornado por NewParking
	parkingLot, err := models.NewParking(capacidadEstacionamiento)
	if err != nil {
		log.Fatalf("Error al inicializar el estacionamiento: %v", err)
	}

	window := view.CreateWindow(application, parkingLot, totalCarros)
	if window == nil {
		log.Fatalf("Error al crear la ventana de la aplicación")
	}

	window.ShowAndRun()
}
