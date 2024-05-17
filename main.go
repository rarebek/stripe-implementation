package main

import (
	"log"
	"net/http"
	"restaurant-payment/config"
	"restaurant-payment/handlers"
	"restaurant-payment/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	stripeService := services.NewStripeService(cfg)
	paymentHandler := handlers.NewPaymentHandler(stripeService)

	router := mux.NewRouter()
	router.HandleFunc("/pay", paymentHandler.HandlePayment).Methods("POST")

	http.Handle("/", router)
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
