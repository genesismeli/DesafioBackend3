package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/genesismeli/DesafioBackend3/internal/tickets"
)

const (
	filename = "./tickets.csv"
)

func main() {
	filename := "tickets.csv"
	
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()


	storage := tickets.Storage{
		Tickets: tickets.ReadFile(filename),
	}

	fmt.Printf("Se leyeron %d tickets.\n", len(storage.Tickets))
	
	//Crear Canales para comunicarnos con las GoRoutines
	TotalTickets := make(chan string)
	defer close(TotalTickets)

	TicketsPorHorario :=make(chan int)
	defer close(TicketsPorHorario)

	PorcentajeDestino := make(chan string)
	defer close(PorcentajeDestino)
	

	canalErr := make(chan error)
	defer close(canalErr)

	//Requerimiento 1: Contar total de tickets
	
	var country string
	fmt.Print("Ingresa el país de destino: ")
	_, err := fmt.Scan(&country)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	go func (chan string, chan error) {
		travelersCount, err := storage.CountTravelers(country, storage.Tickets)
		if err != nil {
			canalErr <- err
			return
		}
		mensaje := fmt.Sprintf("Viajeros al país %s: %d\n", country, travelersCount)
		TotalTickets <- mensaje
	}(TotalTickets, canalErr)

	time.Sleep(time.Millisecond * 100)
	
	

	//Requerimiento 2: Contar total de viajantes por rango horario
	var wg sync.WaitGroup
	wg.Add(4) //Para poder hacer que todas las goroutines finalicen antes de cerrar el canal
	
	go func(tickets []tickets.Ticket, start, end int, resultChan chan<- int) {
		count := storage.CountByTimeRange(tickets, start, end)
		resultChan <- count
	}(storage.Tickets, 0, 6, TicketsPorHorario)

	go func(tickets []tickets.Ticket, start, end int, resultChan chan<- int) {
		count := storage.CountByTimeRange(tickets, start, end)
		resultChan <- count
	}(storage.Tickets, 7, 12, TicketsPorHorario)

	go func(tickets []tickets.Ticket, start, end int, resultChan chan<- int) {
		count := storage.CountByTimeRange(tickets, start, end)
		resultChan <- count
	}(storage.Tickets, 13, 19, TicketsPorHorario)

	go func(tickets []tickets.Ticket, start, end int, resultChan chan<- int) {
		count := storage.CountByTimeRange(tickets, start, end)
		resultChan <- count
	}(storage.Tickets, 20, 23, TicketsPorHorario)

	go func() {
		wg.Wait()
		close(TotalTickets)
	}()
	
	// Recuperar los resultados de la goroutine y sumarlos
	madrugadaCount := <-TicketsPorHorario
	mananaCount := <-TicketsPorHorario
	tardeCount := <-TicketsPorHorario
	nocheCount := <-TicketsPorHorario

	fmt.Printf("Viajeros en madrugada: %d\n", madrugadaCount)
	fmt.Printf("Viajeros en la mañana: %d\n", mananaCount)
	fmt.Printf("Viajeros en la tarde: %d\n", tardeCount)
	fmt.Printf("Viajeros en la noche: %d\n", nocheCount)


	//Requerimiento 3: Calcular porcentaje de personas que viajan a un pais determinado de un día

	var inputPaisPorcentaje string

	fmt.Print("Ingrese pais elegido para calcular el porcentaje: ")
	_, errPorcentaje := fmt.Scan(&inputPaisPorcentaje)
	
	if errPorcentaje != nil {
		log.Fatal(errPorcentaje)
		os.Exit(1)
	}

	go func(chan string, chan error) {
		totalTickets := 0
		for i := 0; i < len(storage.Tickets); i++ {
			totalTickets++
		}

		porcentaje, err := storage.AverageDestination(inputPaisPorcentaje, storage.Tickets)
		if err != nil {
			canalErr <- err
			return
		}
		mensaje := fmt.Sprintf("El porcentaje de personas que viajan al destino %s es %.2f.", inputPaisPorcentaje, porcentaje)

		PorcentajeDestino <- mensaje
	}(PorcentajeDestino, canalErr)

	time.Sleep(time.Millisecond * 100)


	//Impresion de Canales
	select {
	//Primer requerimiento
	case totalTicket := <-TotalTickets:
		fmt.Println(totalTicket)
	//Aca va el 2
	//Segundo requerimiento
	case porcentajeDestino := <-PorcentajeDestino:
		fmt.Println(porcentajeDestino)

	case err := <-canalErr:
		fmt.Println(err)
		os.Exit(1)
	}
}
