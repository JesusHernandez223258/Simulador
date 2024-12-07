package config

import "image/color"

var (
    BackgroundColor = color.NRGBA{R: 40, G: 44, B: 52, A: 255}
    TextColor       = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
    
    // Rutas de imágenes
    AvailableImage = "src/assets/si.png"
    OccupiedImages = []string{
        "src/assets/cars/car-black.png",
        "src/assets/cars/car-blue.png",
        "src/assets/cars/car-green.png",
        "src/assets/cars/car-purple.png",
        "src/assets/cars/car-red.png",
    }
)
