package models

import (
	"context"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"sync"
	"time"
)

var logger = log.New(log.Writer(), "", 0)

type Parking struct {
	capacity       int           // Capacidad total del estacionamiento
	mutex          sync.RWMutex  // Mutex para acceso concurrente
	Queue          chan *Car     // Canal para esperar vehículos
	availableSpots chan struct{} // Canal que indica espacios disponibles
	entryExitMutex chan struct{} // Mutex para controlar la entrada/salida
	occupiedSpaces []bool        // Array para saber qué espacios están ocupados
	carIDs         []int         // Array para almacenar IDs de vehículos
	nextSpotIndex  int           // Índice para el próximo espacio disponible
	EntryColor     color.Color
	WaitColor      color.Color
	wg             sync.WaitGroup // WaitGroup para sincronizar goroutines
}

// Cambiamos la función para retornar un error en caso de fallo
func NewParking(capacity int) (*Parking, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("la capacidad debe ser mayor a cero")
	}

	parking := &Parking{
		capacity:       capacity,
		Queue:          make(chan *Car, capacity),     // Inicializa el canal para la cola
		availableSpots: make(chan struct{}, capacity), // Canal para espacios disponibles
		entryExitMutex: make(chan struct{}, 1),        // Mutex para entrada/salida
		occupiedSpaces: make([]bool, capacity),        // Array de espacios ocupados
		carIDs:         make([]int, capacity),         // Array para IDs de vehículos
		nextSpotIndex:  0,                             // Comenzamos con el primer espacio
	}
	// Llenamos el canal de espacios disponibles al iniciar
	for i := 0; i < capacity; i++ {
		parking.availableSpots <- struct{}{}
	}

	return parking, nil
}

func (p *Parking) Capacity() int {
	return p.capacity
}

// Devuelve los espacios ocupados y los IDs de los vehículos
func (p *Parking) OccupiedSpaces() ([]bool, []int) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	occupiedSpaces := make([]bool, p.capacity)
	carIDs := make([]int, p.capacity)

	for i := 0; i < p.capacity; i++ {
		occupiedSpaces[i] = p.occupiedSpaces[i]
		if p.occupiedSpaces[i] {
			carIDs[i] = p.carIDs[i]
		}
	}
	return occupiedSpaces, carIDs
}

// Busca el próximo espacio libre
func (p *Parking) findNextSpot() int {
	for i := range p.occupiedSpaces {
		if !p.occupiedSpaces[i] { // Si encontramos un espacio libre
			return i // Retornamos el índice del espacio libre
		}
	}
	return -1 // Si no hay espacio disponible, devolvemos -1
}


// Simula la llegada de vehículos al estacionamiento
func generateCars(ctx context.Context, carChannel chan<- *Car, arrivalRate float64) {
	defer close(carChannel)
	carID := 1
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(rand.ExpFloat64()/arrivalRate) * time.Second):
			carChannel <- &Car{ID: carID}
			carID++
		}
	}
}

func processCars(parking *Parking, carChannel <-chan *Car) {
	for car := range carChannel {
		parking.Enter(car)
	}
}

func Simulate(parking *Parking, arrivalRate float64, ctx context.Context) {
	carChannel := make(chan *Car)
	go generateCars(ctx, carChannel, arrivalRate)
	go processCars(parking, carChannel)
	parking.wg.Wait()
}

