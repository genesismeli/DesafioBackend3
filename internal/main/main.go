package main

import (
	"fmt"
	"github.com/genesismeli/DesafioBackend3/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	filename := "tickets.csv"
	ticketes := tickets.ReadFile(filename)
	fmt.Printf("Se leyeron %d tickets.\n", len(ticketes))
	var country string
	fmt.Print("Ingresa el país de destino: ")
	fmt.Scan(&country)

	travelersCount := tickets.CountTravelers(ticketes, country)
	fmt.Printf("Viajeros al país %s: %d\n", country, travelersCount)

	madrugadaCount := tickets.CountByTimeRange(ticketes, 0, 6)
	mananaCount := tickets.CountByTimeRange(ticketes, 7, 12)
	tardeCount := tickets.CountByTimeRange(ticketes, 13, 19)
	nocheCount := tickets.CountByTimeRange(ticketes, 20, 23)

	fmt.Printf("Viajeros en madrugada: %d\n", madrugadaCount)
	fmt.Printf("Viajeros en la mañana: %d\n", mananaCount)
	fmt.Printf("Viajeros en la tarde: %d\n", tardeCount)
	fmt.Printf("Viajeros en la noche: %d\n", nocheCount)
}
