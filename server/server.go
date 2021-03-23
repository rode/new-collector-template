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

package server

import (
	"context"
	"fmt"

	"github.com/rode/new-collector-template/proto/v1alpha1"
	pb "github.com/rode/rode/proto/v1alpha1"
	"go.uber.org/zap"
)

type NewCollectorTemplateServer struct {
	logger *zap.Logger
	rode   pb.RodeClient
}

func NewNewCollectorTemplateServer(logger *zap.Logger, rode pb.RodeClient) *NewCollectorTemplateServer {
	return &NewCollectorTemplateServer{
		logger,
		rode,
	}
}

func (s *NewCollectorTemplateServer) Hello(_ context.Context, request *v1alpha1.HelloRequest) (*v1alpha1.HelloResponse, error) {
	return &v1alpha1.HelloResponse{
		Hello: fmt.Sprintf("Hello, %s", request.Name),
	}, nil
}
