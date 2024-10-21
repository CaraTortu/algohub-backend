package main

import (
	"log"
	"net"
	"os"

	pb "algohub.dev/backend/proto"
	"algohub.dev/backend/servers"
	"algohub.dev/backend/structs"
	"algohub.dev/backend/utils/db"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	port = ":8080"
)

// Register all the servers
func registerServers(s *grpc.Server) {
	env := getEnvironment()

	db, err := gorm.Open(postgres.Open(env.DB_URL))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	pb.RegisterExampleServer(s, &servers.ExampleServer{DB: db, Env: &env})
}

// Used to get environment variables
func getEnvironment() structs.Env {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	jwt_secret := os.Getenv("JWT_SECRET")
	if jwt_secret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	return structs.Env{DB_URL: db_url, JWT_SECRET: []byte(jwt_secret)}
}

func startServer() {
	// Listen on port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer(grpc.MaxSendMsgSize(1024 * 1024 * 20))
	registerServers(s)

	// Serve the gRPC server
	log.Println("Server started on " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	rootCmd := &cobra.Command{Use: "eikoiq-backend"}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start the gRPC server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "erase_db",
		Short: "Erase the database",
		Run: func(cmd *cobra.Command, args []string) {
			env := getEnvironment()
			db.EraseDB(&env)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "seed_db",
		Short: "Seed the database",
		Run: func(cmd *cobra.Command, args []string) {
			//env := getEnvironment()
			//db.SeedDB(&env)
		},
	})

	rootCmd.Execute()
}
