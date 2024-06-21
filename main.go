package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/leonardodf95/api-go/internal/events/infra/http"
	"github.com/leonardodf95/api-go/internal/events/infra/repository"
	"github.com/leonardodf95/api-go/internal/events/usecase"
)

// @title Events API
// @version 1.0
// @description This is a server for managing events. Imersão Full Cycle
// @host localhost:8080
// @BasePath /
func main() {
	jsonData, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	// Repositório
	eventRepo, err := repository.NewEventRepository(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	reserveSpotUseCase := usecase.NewReserveSpotUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)

	// Handlers HTTP
	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		getEventUseCase,
		reserveSpotUseCase,
		listSpotsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /events/{eventId}/reserve", eventsHandler.ReserveSpots)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Canal para escutar sinais do sistema operacional
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Recebido sinal de interrupção, iniciando o graceful shutdown
		log.Println("Recebido sinal de interrupção, iniciando o graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Erro no graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Iniciando o servidor HTTP
	log.Println("Servidor HTTP rodando na porta 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Servidor HTTP finalizado")
}
