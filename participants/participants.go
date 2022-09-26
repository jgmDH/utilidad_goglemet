package participants

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"golang.org/x/exp/slices"
)

var (
	bootcampers = []string{
		"Abila, Daniel",
		"Alvarez, Natalia",
		"Calle, Juan Fernando",
		"Cambiasso, Tomas",
		"Capodici, Gianfranco",
		"Carrizo, Melany",
		"Carvajal, Luis Alejandro",
		"Conte, Camila Lorena",
		"Corsiglia, Romina",
		"Cuellar, Marcela",
		"Cuello Rodriguez, Tomas",
		"Di Gulio, Dario",
		"Diaz Real, Maria Sol",
		"Durand, Luis",
		"Echeverri, Adriana",
		"Escudero, Mario",
		"García, Sabrina",
		"Gattás, Roberto",
		"Gil, Abril",
		"Guglielmone, Juan Martin",
		"Hernandez, Juan Pablo",
		"Huang, Yuhong",
		"Laiton Cubides, Santiago Agustin",
		"Larramendi, Marcio",
		"Londoño, Jacobo",
		"Lozano Quiroga, Freyman Yohani",
		"Macri, Mariano",
		"Martin, Gaston",
		"Martin Montesi, Nadia",
		"Martinez, Natali",
		"Martínez Camacho, Víctor Hugo",
		"Medina, Daniel",
		"Mejía Parra, Fabio Andrés",
		"Morel, Melany Lucía",
		"Navarro Tapia, Nicolás",
		"Niz, Franco",
		"Ochoa Barco, Julian",
		"Padilla, Julie",
		"Pesenda, Franco Rodolfo",
		"Pérez Trejos, Juan Andrés",
		"Ramírez Gómez, Juan José",
		"Rocha, Pedro",
		"Rodriguez Pineda, Ivan",
		"Roldán, Leidy Johanna",
		"Rosso, Mauricio",
		"Salgado Meza, Juan Camilo",
		"San Juan Escalona, Sarai",
		"Santiesteban Mendivelso, Jeisson Fernando",
		"Sassenus, Milagros",
		"Seidel, Dionys",
		"Sibona, Gonzalo",
		"Silva, Laureano",
		"Soria Gava, Lucas Damián",
		"Tira, Zoé Agustina",
		"Torres, Germán",
		"Urbano Montaña, Juan Carlos",
		"Urteaga Naya, Martín",
		"Valderrama, Luisa",
		"Zamudio, Rosario",
		"Zunda, Fernando Agustin",
		"De la Serna, Matías",
	}
)

func BootcampersAusentes(pts []string) {
	var ausentes []string
	for _, p := range bootcampers {
		if !slices.Contains(pts, p) {
			ausentes = append(ausentes, p)
		}
	}

	now := time.Now()
	fileName := "ausentes_" + now.Format("2006-01-02 ") + now.Format(time.Kitchen) + ".csv"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatal("Cannot create file - err ", err)
		}
	}
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	for _, a := range ausentes {
		record := []string{a}
		if err := csvwriter.Write(record); err != nil {
			panic(err)
		}
	}

	csvwriter.Flush()
}
