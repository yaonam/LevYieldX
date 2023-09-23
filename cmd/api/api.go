package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"levyieldx/cmd/protocols"
)

func Start() {
	r := chi.NewRouter() // Basic CORS setup
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allowed origins, can be more specific like: []string{"http://localhost:3000"}
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Apply CORS middleware to the router
	r.Use(corsMiddleware.Handler)
	r.Use(middleware.Logger)

	r.Get("/", test)

	// TODO: Change for prod deployment
	log.Println("Starting server on port 8080")
	http.ListenAndServe("127.0.0.1:8080", r)
}

func test(w http.ResponseWriter, r *http.Request) {
	const protocol = "aavev3"
	const chain = "arbitrum"
	const wallet = "0x18dC22D776aEFefD2538079409176086fcB6C741"
	const token = "USDC"
	// Test Supply()
	p, _ := protocols.GetYieldProtocol(protocol)
	err := p.Connect(chain)
	if err != nil {
		log.Printf("Failed to connect to the %v protocol: %v", protocol, err)
		return
	}
	log.Println("Connected to the protocol")

	markets, err := p.GetMarkets()
	if err != nil {
		log.Printf("Failed to get markets: %v", err)
		return
	}
	res, err := json.Marshal(markets)
	if err != nil {
		log.Panicf("failed to marshal strategies: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
