package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/victor99z/blockchain-app/handlers"
)

func main() {
	godotenv.Load()
	app := fiber.New(fiber.Config{
		StreamRequestBody: true,
	})
	// app := fiber.New()
	// app := fiber.New(fiber.Config{
	// 	BodyLimit: 1000 * 1024 * 1024,
	// })
	api := app.Group("/api")

	parseEnvs, _ := json.MarshalIndent(os.Environ(), "", "  ")

	fmt.Println(string(parseEnvs))

	v1 := api.Group("/v1")
	api_version := os.Getenv("VERSION_API")
	if api_version == "onchain" {
		handler := handlers.NewOnchainHandler()
		v1.Get("/record/:id", handler.GetRecords)
		v1.Get("/record/:id/counter", handler.GetRecordCounterByAddr)
		v1.Get("/record/:id/:index", handler.GetRecord)
		v1.Post("/record", handler.SaveRecord)
	} else if api_version == "sidechain" {
		handler := handlers.NewSidechainHandler()
		v1.Get("/record/:id", handler.GetRecords)
		v1.Get("/record/:id/counter", handler.GetRecordCounterByAddr)
		v1.Get("/record/:id/:index", handler.GetRecord)
		v1.Post("/record", handler.SaveRecord)
	}
	log.Fatal(app.Listen(":5003"))
}
