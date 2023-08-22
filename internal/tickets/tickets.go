package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          string
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      string
}

type Storage struct {
	Tickets []Ticket
}

func ReadFile(filename string) []Ticket {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var resultado []Ticket
	for i := 0; i < len(data); i++ {
		if len(data[i]) > 0 {
			fields := strings.Split(data[i], ",")
			if len(fields) >= 6 {
				ticket := Ticket{
					ID:          fields[0],
					Nombre:      fields[1],
					Email:       fields[2],
					PaisDestino: fields[3],
					HoraVuelo:   fields[4],
					Precio:      fields[5],
				}
				resultado = append(resultado, ticket)
			} else {
				fmt.Printf("Error en la línea %d: no hay suficientes campos\n", i+1)
			}
		}
	}
	return resultado
}
func CountTravelers(tickets []Ticket, country string) int {
	count := 0
	for _, ticket := range tickets {
		if ticket.PaisDestino == country {
			count++
		}
	}

	return count
}

func CountByTimeRange(tickets []Ticket, startHour, endHour int) int {
	count := 0

	for _, ticket := range tickets {
		hourStr := strings.Split(ticket.HoraVuelo, ":")[0]
		hour, err := strconv.Atoi(hourStr)
		if err != nil {
			continue
		}

		if hour >= startHour && hour <= endHour {
			count++
		}
	}

	return count
}
