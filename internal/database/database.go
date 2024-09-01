package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	lg "gorm.io/gorm/logger"
)

// Service represents a service that interacts with a database.
type Service interface {
	Getpdb() *gorm.DB
}

type service struct {
	Pdb *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)
var GORM_DB *gorm.DB
var SQL_DB *sql.DB
var DB_MIGRATOR gorm.Migrator

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	newLogger := lg.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		lg.Config{
			LogLevel:                  lg.Error, // Log level
			IgnoreRecordNotFoundError: true,     // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)
	// Open the database connection
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		GORM_DB = db
		SQL_DB, _ = db.DB()
		DB_MIGRATOR = db.Migrator()
	}
	dbInstance = &service{
		Pdb: db,
	}
	return dbInstance
}
func (s *service) Getpdb() *gorm.DB {
	return s.Pdb
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
// func (s *service) Health() map[string]string {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	stats := make(map[string]string)

// 	// Ping the database
// 	err := s.db.PingContext(ctx)
// 	if err != nil {
// 		stats["status"] = "down"
// 		stats["error"] = fmt.Sprintf("db down: %v", err)
// 		log.Fatalf(fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
// 		return stats
// 	}

// 	// Database is up, add more statistics
// 	stats["status"] = "up"
// 	stats["message"] = "It's healthy"

// 	// Get database stats (like open connections, in use, idle, etc.)
// 	dbStats := s.db.Stats()
// 	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
// 	stats["in_use"] = strconv.Itoa(dbStats.InUse)
// 	stats["idle"] = strconv.Itoa(dbStats.Idle)
// 	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
// 	stats["wait_duration"] = dbStats.WaitDuration.String()
// 	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
// 	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

// 	// Evaluate stats to provide a health message
// 	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
// 		stats["message"] = "The database is experiencing heavy load."
// 	}

// 	if dbStats.WaitCount > 1000 {
// 		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
// 	}

// 	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
// 		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
// 	}

// 	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
// 		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
// 	}

// 	return stats
// }

// // Close closes the database connection.
// // It logs a message indicating the disconnection from the specific database.
// // If the connection is successfully closed, it returns nil.
// // If an error occurs while closing the connection, it returns the error.
// func (s *service) Close() error {
// 	log.Printf("Disconnected from database: %s", database)
// 	return s.db.Close()
// }
