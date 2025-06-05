package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	auth "github.com/enrollment/gen/auth"
	course "github.com/enrollment/gen/course"
	enrollment "github.com/enrollment/gen/enrollment"
	oauth "github.com/enrollment/gen/oauth"
	controllers "github.com/enrollment/src/controllers"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	format := log.FormatJSON
	if log.IsTerminal() {
		format = log.FormatTerminal
	}
	ctx := log.Context(context.Background(), log.WithFormat(format))
	if *dbgF {
		ctx = log.Context(ctx, log.WithDebug())
		log.Debugf(ctx, "debug logs enabled")
	}
	log.Print(ctx, log.KV{K: "http-port", V: *httpPortF})

	// Initialize the services.
	var (
		courseSvc     course.Service
		enrollmentSvc enrollment.Service
		oauthSvc      oauth.Service
		authSvc       auth.Service
	)
	{
		courseSvc = controllers.NewCourse()
		enrollmentSvc = controllers.NewEnrollment()
		oauthSvc = controllers.NewOauth()
		authSvc = controllers.NewAuth()
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		courseEndpoints     *course.Endpoints
		enrollmentEndpoints *enrollment.Endpoints
		oauthEndpoints      *oauth.Endpoints
		authEndpoints       *auth.Endpoints
	)
	{
		courseEndpoints = course.NewEndpoints(courseSvc)
		courseEndpoints.Use(debug.LogPayloads())
		courseEndpoints.Use(log.Endpoint)
		enrollmentEndpoints = enrollment.NewEndpoints(enrollmentSvc)
		enrollmentEndpoints.Use(debug.LogPayloads())
		enrollmentEndpoints.Use(log.Endpoint)
		oauthEndpoints = oauth.NewEndpoints(oauthSvc)
		oauthEndpoints.Use(debug.LogPayloads())
		oauthEndpoints.Use(log.Endpoint)
		authEndpoints = auth.NewEndpoints(authSvc)
		authEndpoints.Use(debug.LogPayloads())
		authEndpoints.Use(log.Endpoint)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:80"
			u, err := url.Parse(addr)
			if err != nil {
				log.Fatalf(ctx, err, "invalid URL %#v\n", addr)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					log.Fatalf(ctx, err, "invalid URL %#v\n", u.Host)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, courseEndpoints, enrollmentEndpoints, oauthEndpoints, authEndpoints, &wg, errc, *dbgF)
		}

	default:
		log.Fatal(ctx, fmt.Errorf("invalid host argument: %q (valid hosts: localhost)", *hostF))
	}

	// Wait for signal.
	log.Printf(ctx, "exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	log.Printf(ctx, "exited")
}
