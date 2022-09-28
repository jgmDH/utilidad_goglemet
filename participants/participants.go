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
		"Daniel Abila",
		"Natalia Belen Alvarez",
		"Juan Fernando Calle Herrera",
		"Tomas Agustin Cambiasso",
		"Gianfranco Capodici",
		"Melany Nicole Carrizo",
		"Luis Alejandro Carvajal Villa",
		"Camila Lorena Conte",
		"Romina Andrea Corsiglia",
		"Marcela Cuellar Galvira (Marcela Cuellar Galvira)",
		"Tomas Cuello Rodriguez",
		"Dario Di Gulio",
		"Maria Sol Diaz Real",
		"Luis Rodrigo Durand Panez",
		"Adriana Echeverri Romero",
		"Mario Teodoro Escudero",
		"Sabrina Garcia",
		"Roberto Gattas",
		"Abril Martina Gil",
		"Juan Martin Guglielmone",
		"Juan Pablo Hernandez Jimenez",
		"Yuhong Huang",
		"Santiago Agustin Laiton Cubides",
		"Marcio Larramendi Rossi",
		"Jacobo Rave Londono",
		"Freyman Yohani Lozano Quiroga",
		"Mariano Macri",
		"Gaston Martin",
		"Nadia Martin Montesi",
		"Natali Martinez",
		"Victor Hugo Martinez",
		"Daniel De Jesus Medina Ortega",
		"Fabio Andres Mejia Parra (Meyi)",
		"Melany Lucia Morel",
		"Nicolas Baltazar Navarro Tapia",
		"Franco Damian Niz",
		"Julian Ochoa Barco",
		"Julie Padilla",
		"Franco Rodolfo Pesenda",
		"Juan Andrés Pérez Trejos",
		"Juan Jose Ramirez Gomez",
		"Pedro Maria Rocha",
		"Ivan Arturo Rodriguez Pineda",
		"Leidy Johanna Roldan Vargas",
		"Mauricio Andres Rosso",
		"Juan Camilo Salgado Meza",
		"Sarai San Juan",
		"Jeisson Fernando Santiesteban Mendivelso",
		"Milagros Stephanie Sassenus",
		"Dionys Seidel",
		"Gonzalo Sibona",
		"Laureano Silva",
		"Lucas Damian Soria Gava",
		"Zoe Agustina Tira",
		"German Daniel Torres",
		"Juan Urbano",
		"Martin Ignacio Urteaga Naya",
		"Luisa Alejandra Marin Valderrama",
		"Rosario Zamudio",
		"Fernando Agustin Zunda",
		"Matias De La Serna",
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
