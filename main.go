package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Data string `json:"data"`
}

var (
	quotes []Quote
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func init() {
	data, err := os.ReadFile("quotes.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &quotes)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// Handler
func hello(c *fiber.Ctx) error {
	rand.Seed(time.Now().UnixNano())
	data := quotes[rand.Intn(len(quotes))].Data
	return c.Status(http.StatusOK).JSON(data)
}
