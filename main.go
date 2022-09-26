package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/juanmachuca95/server-gogle-meet/participants"
)

var (
	fileName string = "chat.txt"
)

type MessageDTO struct {
	SenderName         string   `json:"sender_name"`
	FormattedTimestamp string   `json:"formatted_timestamp"`
	Messages           []string `json:"messages"`
}

type AsistenciaDTO struct {
	Participantes []string `json:"participants"`
}

// Format message from google meet script javascript
// {"sender-name":"Tú","formatted-timestamp":"19:38","messages":["lkjljl"]}
func main() {
	http.HandleFunc("/asistencia", asistencia)
	http.HandleFunc("/data", data)

	// Server on port 8000 initialized
	log.Println("Server initialized on port 8000 🚀")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err.Error())
	}
}

// Asistencia
func asistencia(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := new(AsistenciaDTO)
	err = json.Unmarshal(data, p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	participants.BootcampersAusentes(p.Participantes)
	w.WriteHeader(http.StatusOK)
}

// Captura de mensajes de google meet
func data(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := new(MessageDTO)
	err = json.Unmarshal(data, m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Cada vez que recibimos información guardamos el resultado en chat.txt
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		file, err = os.Create(fileName)
		if err != nil {
			log.Fatal("Cannot create file - err ", err)
		}
	}

	// ignoramos el error, ya que no deveriamos retornar info al script js
	_, _ = file.WriteString("User: " + m.SenderName + " - " + m.FormattedTimestamp + "\n")
	var messages string
	for _, value := range m.Messages {
		_, err := file.WriteString(value + "\n")
		if err != nil {
			log.Fatal("Cannot create row text on file ", err)
		}

		messages += value + "\n"
	}

	err = beeep.Notify(m.SenderName, messages, "go.png")
	if err != nil {
		panic(err)
	}

	// That's also possible with the package de github.com/bitfield/script
	// script.Exec("notify-send " + message).Stdout()
	_, _ = file.WriteString("\n")
}
