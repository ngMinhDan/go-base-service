/*
Package server: Provides method to create server to serve HTTP request
Package Functionality: This package allows you to create , start, stop server
Author: MinhDan <nguyenmd.works@gmail.com>
*/
package server

import (
	"base/pkg/config"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Server Struct
type Server struct {
	srv *http.Server
}

// Server Configuration Struct
type serverConfig struct {
	IP   string
	Port string
}

// Server Configuration Variable
var ServerCfg serverConfig

func init() {
	ServerCfg.IP = config.Config.GetString("SERVER_IP")
	ServerCfg.Port = config.Config.GetString("SERVER_PORT")
}

// NewServer Function to Create a New Server Handler
func NewServer(handler http.Handler) *Server {
	// Initialize New Server
	return &Server{
		srv: &http.Server{
			Addr:    net.JoinHostPort(ServerCfg.IP, ServerCfg.Port),
			Handler: handler,
		},
	}
}

// Start Method for Server
func (s *Server) Start() {
	// Initialize Context Handler Without Timeout
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start server
	fmt.Println(config.Config.GetString("SERVER_NAME") + " started at pid " + strconv.Itoa(os.Getpid()) +
		" listening on " + net.JoinHostPort(ServerCfg.IP, ServerCfg.Port) + "...")

	// Server handle all incoming request
	// ListenAndServe will be block the execution of program until the server stopped
	// Use Goroutine for gracefully stop
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Error:", err)
		}
	}()

}

// Stop Method for Server
func (s *Server) Stop() {
	// Initialize Timeout
	timeout := 5 * time.Second
	// start server
	fmt.Println(config.Config.GetString("SERVER_NAME") + " stoped at pid " + strconv.Itoa(os.Getpid()) +
		" listening on " + net.JoinHostPort(ServerCfg.IP, ServerCfg.Port))

	// Initialize Context Handler With Timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Hanlde Any Error While Stopping Server
	if err := s.srv.Shutdown(ctx); err != nil {
		if err = s.srv.Close(); err != nil {
			log.Fatalln("{\"label\":\"server-http\",\"level\":\"error\",\"msg\":\"" + err.Error() +
				"\",\"service\":\"" + config.Config.GetString("SERVER_NAME") + "\",\"time\":" +
				fmt.Sprint(time.Now().Format(time.RFC3339Nano)) + "\"}")
			return
		}
	}
}
