package restgopher

import (
	"fmt"
	"net/http"

	"github.com/agoradevjsj/RESTGopher/utils/timeexec"
)

// Structures

// Application structure
type Server struct {
	DomainServer string
	PortServer   string
	middleware   []func(http.HandlerFunc) http.HandlerFunc
	message      string
}

// App configurte
type ConfigApp struct {
	DomainApp string
	PortApp   string
	Message   string
}

// Variables

var port_expose string
var domain_expose string
var time_exec any

// Functionality

// Initialization variables privates
func init() {
	port_expose = ":8000"
	domain_expose = "127.0.0.1"
}

func CreateConfigApp(portApp string, domainApp string, messageApp string) ConfigApp {
	return ConfigApp{
		PortApp:   portApp,
		Message:   messageApp,
		DomainApp: domainApp,
	}
}

// Create a new app
func NewServerGopher(config ...ConfigApp) *Server {
	timeexec.StartTimer()
	application := &Server{}
	if len(config) > 0 {
		if config[0].PortApp != "" {
			application.PortServer = config[0].PortApp
		} else {
			application.PortServer = port_expose
		}
		if config[0].DomainApp != "" {
			application.DomainServer = config[0].DomainApp
		} else {
			application.DomainServer = domain_expose
		}
		application.message = config[0].Message
	}
	return application
}

func (s *Server) Use(mws ...func(http.HandlerFunc) http.HandlerFunc) {
	s.middleware = append(s.middleware, mws...)
}

func (s *Server) GetMessage() string {
	return s.message
}

// Start listining the server
func (s *Server) StartServerGopher() error {
	timeexec.StartTimer()
	server := http.NewServeMux()

	// Add middlewares at server
	var handler http.HandlerFunc
	for i := len(s.middleware) - 1; i >= 0; i-- {
		handler = s.middleware[i](handler)
	}

	if len(s.middleware) > 0 {
		server.HandleFunc("/", handler)
	}

	timeexec.StopTime()
	time_exec = timeexec.GetTime().Seconds()
	fmt.Printf("------------  RESTGother  ------------\n\n          🚀 Server running\n\n  👍 Host: %v%v\n  ⏲️  Time: %vs\n\n----------------  END  ---------------\n\n", s.DomainServer, s.PortServer, time_exec)
	fmt.Printf("%v\n\n", s.message)
	return http.ListenAndServe(s.PortServer, server)
}
