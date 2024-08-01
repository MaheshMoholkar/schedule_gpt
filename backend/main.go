package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MaheshMoholkar/schedule_gpt/internal/database"
	"github.com/MaheshMoholkar/schedule_gpt/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func initDB() *sql.DB {
	cfg := database.DefaultSQLServerConfig()
	db, err := database.Open(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

// Logging middleware
func requestLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Continue stack
	err := c.Next()

	// Get the response status and body
	status := c.Response().StatusCode()

	duration := time.Since(start)

	// Log request and response details
	log.Printf("[%s] %s %s %s %d %s", start.Format(time.RFC3339), c.Method(), c.Path(), c.IP(), status, duration)

	return err
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}
	address := fmt.Sprintf(":%s", port)

	listenAddr := flag.String("listenAddr", address, "The listen address of the api server")
	flag.Parse()

	db := initDB()
	defer db.Close()

	store := database.New(db)

	var (
		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")

		// Initialize handlers
		userHandler     = handlers.NewUserHandler(store)
		ollamaHandler   = handlers.NewOllamaHandler(store)
		scheduleHandler = handlers.NewSchduleHandler(store)
	)

	// Use logging middleware
	app.Use(requestLogger)

	// Customize the CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// User handlers
	apiv1.Get("/user", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandleCreateUser)

	// Ollama handlers
	apiv1.Post("/generate", ollamaHandler.HandleGenerateRequest)
	apiv1.Post("/schedule", ollamaHandler.HandleGenerateSchedule)

	apiv1.Post("/chat", ollamaHandler.HandleChatRequest)

	// Schedule handlers
	apiv1.Get("/subjects", scheduleHandler.HandleGetSubjects)
	apiv1.Post("/subjects", scheduleHandler.HandleStoreSubjects)

	app.Listen(*listenAddr)
}
