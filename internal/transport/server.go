﻿// MIT License

// Copyright (c) 2024 ISSuh

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package transport

import (
	"net"

	"github.com/ISSuh/skvs/internal/transport/handler"
	"github.com/ISSuh/skvs/internal/transport/rpc"
)

type Server struct {
	rpcServer rpc.ServerAdaptor
}

func NewServer(handler handler.Storage) *Server {
	return &Server{
		rpcServer: rpc.NewServerAdaptor(handler),
	}
}

func (s *Server) Run(address string) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	if err := s.rpcServer.Run(listen); err != nil {
		return err
	}
	return nil
}