package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
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
	r.Get("/strats", getStrats)
	r.Post("/transactions", getTransactions)

	// TODO: Change for prod deployment
	log.Println("Starting server on port 8080")
	http.ListenAndServe("127.0.0.1:8080", r)
}

func test(w http.ResponseWriter, r *http.Request) {
	ps := []string{"aavev3", "compoundv3"}
	chains := []string{"arbitrum", "ethereum"}
	baseTokens := []string{"ETH", "USDC", "USDT"}
	maxLevels := 2
	initialAmountUSD := new(big.Int).SetInt64(100000000000)
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

func getStrats(w http.ResponseWriter, r *http.Request) {
	ps := strings.Split(r.URL.Query().Get("protocols"), ",")
	chains := strings.Split(r.URL.Query().Get("chains"), ",")
	baseTokens := strings.Split(r.URL.Query().Get("baseTokens"), ",")
	maxLevels, _ := strconv.Atoi(r.URL.Query().Get("maxLevels"))
	initialAmountUSD, _ := new(big.Int).SetString(r.URL.Query().Get("initialAmountUSD"), 10)
	safetyFactor, _ := new(big.Int).SetString(r.URL.Query().Get("safetyFactor"), 10)

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

func getTransactions(w http.ResponseWriter, r *http.Request) {
	var request schema.TransactionsRequest
	if err := render.Bind(r, &request); err != nil {
		log.Panicf("failed to bind input: %v", err)
	}

	txs := make([][]*types.Transaction, len(request.Strategies))
	yps := make(map[string]protocols.YieldProtocol)
	bp, err := protocols.GetBridgeProtocol("synapse")
	if err != nil {
		log.Panicf("Failed to get protocol: %v", err)
	}
	for i, strat := range request.Strategies {
		for _, step := range strat.Steps {
			if step.Market.Protocol == "synapse" {
				// Bridge protocol
				newTxs, err := bp.GetTransactions(request.Wallet, step)
				if err != nil {
					log.Panicf("failed to get transactions: %v", err)
				}
				txs[i] = append(txs[i], newTxs...)
			} else {
				// Yield protocol
				p, ok := yps[step.Market.Protocol]
				if !ok {
					var err error
					p, err = protocols.GetYieldProtocol(step.Market.Protocol)
					if err != nil {
						log.Panicf("Failed to get protocol: %v", err)
					}
					yps[step.Market.Protocol] = p
				}
				p.Connect(step.Market.Chain)
				newTxs, err := p.GetTransactions(request.Wallet, step)
				if err != nil {
					log.Panicf("failed to get transactions: %v", err)
				}
				txs[i] = append(txs[i], newTxs...)
			}
		}
	}

	res, err := json.Marshal(txs)
	if err != nil {
		log.Panicf("failed to marshal strategies: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
