package rpc

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"golang.org/x/net/context"
	"net/http"
)

type Server struct {
	srv *http.Server
}

//":8080"
func NewServer(portnumber string, handler http.Handler) *Server {
	return &Server{srv: &http.Server{Addr: portnumber, Handler: handler}}
}

func (s *Server) Start() {
	//log.Info("Starting HTTP Server at 8080")
	//srv := 	&http.Server{Addr: ":8080", Handler:&msg.JsonRpcHandler{}}
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Infof("Httpserver: ListenAndServe() error: %s", err)

		}
	}()

	log.Info("HTTP Server started")
}

func (s *Server) Stop() {
	log.Info("Stopping Http Server")
	if err := s.srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	log.Info("Server stopped")

}
