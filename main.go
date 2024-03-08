package main

import (
	"bank/api"
	db "bank/db/sqlc"
	"bank/gapi"
	"bank/pb"
	"bank/util"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("❌ cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("😵‍💫 Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	// runGinServer(config, store)
	runGrpcServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("🚨 cannot create http server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("💀 cannot start http server: ", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("🚨 cannot create gRPC server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listner, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("🚨 cannot create listner: ", err)
	}
	log.Printf("🚀 started gRPC server at: %s", listner.Addr().String())

	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("💀 cannot start start server: ", err)
	}

}
