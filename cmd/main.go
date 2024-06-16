package main

import (
	"log"
	"net"

	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"

	pb "auth-service/api"
	"database/sql"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("postgres", "host=db port=5432 user=admin password=1234 dbname=authdb sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, handler.NewAuthHandler(authService))
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Auth Service is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
