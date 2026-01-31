package main

import (
	"crm/internal/activity"
	"crm/internal/customer"
	"crm/internal/deal"
	"crm/internal/lead"
	"crm/internal/note"
	"crm/internal/pipeline_stage"
	"crm/internal/reminder"
	"crm/internal/tenant"
	"crm/internal/user"
	"crm/pkg/config"
	"crm/pkg/middlewares"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %d - %s %s %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.StatusCode,
			param.Method,
			param.Path,			param.Latency,
		)
	}))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.CorsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.Use(middlewares.ErrorHandler())

	// === Database ===
	if err := config.Connect(cfg); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// === Migrate Database ===
	db := config.GetDB()
	tables := []interface{}{
		&tenant.Tenant{},
		&user.User{},
		&lead.Lead{},
		&customer.Customer{},
		&pipeline_stage.PipelineStage{},
		&deal.Deal{},
		&note.Note{},
		&activity.Activity{},
		&reminder.Reminder{},
	}
	if err := db.AutoMigrate(tables...); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("✅ Database migration successful.")

	// === Home Route ===
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "Welcome to CRM SAAS API",
			"version":   "1.0.0",
			"timestamp": time.Now(),
		})
	})

	// === Setup Routes ===
	// Tenant
	tenantRepo := tenant.NewRepository(db)
	tenantService := tenant.NewService(tenantRepo)
	tenantController := tenant.NewController(tenantService)
	tenant.SetupRoute(r, tenantController, cfg)

	// User
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, tenantRepo, cfg, db)
	userController := user.NewController(userService, cfg)
	user.SetupRoute(r, userController, cfg)
	// Lead
	leadRepo := lead.NewRepository(db)
	leadService := lead.NewService(leadRepo)
	leadController := lead.NewController(leadService)
	lead.SetupRoute(r, leadController, cfg)

	// Customer
	customerRepo := customer.NewRepository(db)
	customerService := customer.NewService(customerRepo)
	customerController := customer.NewController(customerService)
	customer.SetupRoute(r, customerController, cfg)

	// Pipeline Stage
	pipelineStageRepo := pipeline_stage.NewRepository(db)
	pipelineStageService := pipeline_stage.NewService(pipelineStageRepo)
	pipelineStageController := pipeline_stage.NewController(pipelineStageService)
	pipeline_stage.SetupRoute(r, pipelineStageController, cfg)

	// Deal
	dealRepo := deal.NewRepository(db)
	dealService := deal.NewService(dealRepo)
	dealController := deal.NewController(dealService)
	deal.SetupRoute(r, dealController, cfg)

	// Note
	noteRepo := note.NewRepository(db)
	noteService := note.NewService(noteRepo)
	noteController := note.NewController(noteService)
	note.SetupRoute(r, noteController, cfg)

	// Activity
	activityRepo := activity.NewRepository(db)
	activityService := activity.NewService(activityRepo)
	activityController := activity.NewController(activityService)
	activity.SetupRoutes(r, activityController, cfg)

	// Reminder
	reminderRepo := reminder.NewRepository(db)
	reminderService := reminder.NewService(reminderRepo)
	reminderController := reminder.NewController(reminderService)
	reminder.SetupRoutes(r, reminderController, cfg)

	// 404 Not Found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})

	// clean.CleanupUnusedUploads(db, "./uploads")

	// === Start Server ===
	log.Printf("Server running on port %s", cfg.Port)
	log.Printf("Local: http://localhost:%s", cfg.Port)
	log.Printf("Environment: %s", cfg.NodeEnv)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
