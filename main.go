// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/rode/new-collector-template/proto/v1alpha1"
	"github.com/rode/new-collector-template/server"
	pb "github.com/rode/rode/proto/v1alpha1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/rode/new-collector-template/config"
)

func main() {
	conf, err := config.Build(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	logger, err := createLogger(conf.Debug)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GrpcPort))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
	}
	if conf.RodeConfig.Insecure {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := grpc.DialContext(ctx, conf.RodeConfig.Host, dialOptions...)
	if err != nil {
		logger.Info(fmt.Sprintf("conf %v", conf.RodeConfig))
		logger.Fatal("failed to establish grpc connection to Rode", zap.Error(err))
	}
	defer conn.Close()

	rodeClient := pb.NewRodeClient(conn)
	grpcServer := grpc.NewServer()

	if conf.Debug {
		reflection.Register(grpcServer)
	}

	collectorServer := server.NewNewCollectorTemplateServer(logger, rodeClient)
	v1alpha1.RegisterNewCollectorTemplateServer(grpcServer, collectorServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("failed to serve", zap.Error(err))
		}
	}()

	httpServer, err := createGrpcGateway(context.Background(), lis.Addr().String(), fmt.Sprintf(":%d", conf.HttpPort))
	if err != nil {
		logger.Fatal("failed to start gateway", zap.Error(err))
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Fatal("failed to start serve on http port", zap.Error(err))
		}
	}()

	logger.Info("listening", zap.String("host", lis.Addr().String()))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	terminationSignal := <-sig

	logger.Info("shutting down...", zap.String("termination signal", terminationSignal.String()))

	grpcServer.GracefulStop()
	httpServer.Shutdown(context.Background())
}

func createGrpcGateway(ctx context.Context, grpcAddress, httpPort string) (*http.Server, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddress,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err := v1alpha1.RegisterNewCollectorTemplateHandler(ctx, gwmux, conn); err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    httpPort,
		Handler: gwmux,
	}, nil
}

func createLogger(debug bool) (*zap.Logger, error) {
	if debug {
		return zap.NewDevelopment()
	}

	return zap.NewProduction()
}
