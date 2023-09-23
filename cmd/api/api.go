package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"levyieldx/cmd/arbitrage"
	"levyieldx/cmd/protocols"
	"levyieldx/cmd/protocols/schema"
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
	ps := []string{"aavev3", "compoundv3"}
	chains := []string{"arbitrum", "ethereum"}
	baseTokens := []string{"ETH", "USDC", "USDT", "DAI"}
	maxLevels := 3
	initialAmountUSD := new(big.Int).SetInt64(1000)
	safetyFactor := new(big.Int).SetInt64(9000)

	var pcs []*schema.ProtocolChain
	psMap := make(map[string]*protocols.YieldProtocol)
	for _, protocol := range ps {
		p, err := protocols.GetYieldProtocol(protocol)
		if err != nil {
			log.Panicf("Failed to get protocol: %v", err)
		}
		psMap[protocol] = &p
		for _, chain := range chains {
			p.Connect(chain)
			pms, err := p.GetMarkets()
			if err != nil {
				log.Panicf("failed to get markets: %v", err)
			}
			pcs = append(pcs, pms...)
		}
	}

	log.Println("Calculating all strats...")
	startTime := time.Now()
	strats := arbitrage.GetAllStrats(pcs, baseTokens, maxLevels)
	log.Printf("Time elapsed: %v", time.Since(startTime))
	log.Println("Generating all steps...")
	strategies, err := arbitrage.CalcBridgedStrats(psMap, strats, initialAmountUSD, safetyFactor)
	if err != nil {
		log.Panicf("failed to calc strategies: %v", err)
	}
	arbitrage.SortStrategies(strategies)

	res, err := json.Marshal(strategies)
	if err != nil {
		log.Panicf("failed to marshal strategies: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
