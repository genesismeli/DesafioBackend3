package tickets

import (
	"os"
	"strings"

	"github.com/genesismeli/DesafioBackend3/internal/tickets"
)



type Ticket struct {

	id string
	nombre string
	email string
	paisDestino string 
	horaVuelo string
	precio string
}

type Storage struct {
	Ticket []Ticket
}

func readFile(filename string) []tickets.Ticket{

		file, err := os.ReadFile(filename)

		if err != nil {
			panic(err)
		}

		data := strings.Split(string(file), "\n")

		var resultado []tickets.Ticket 
		for i := 0; i< len(data); i++ {
			if len(data[i])>0 {
				file := strings.Split(string(data[i]), ",")
				tickets := tickets.Ticket{
						id: file[0],
						nombre: file[1],
						email: file[2],
						paisDestino: file[3], 
						horaVuelo: file[4],
						precio: file[5],
				}
				resultado = append(resultado, tickets)
			}
		}
		return resultado
}



