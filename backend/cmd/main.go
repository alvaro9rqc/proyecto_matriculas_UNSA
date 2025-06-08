package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	course "github.com/enrollment/gen/course"
	enrollment "github.com/enrollment/gen/enrollment"
	controllers "github.com/enrollment/src/controllers"
	"github.com/joho/godotenv"
	"goa.design/clue/debug"
	"goa.design/clue/log"
)

func getEnv(key string, fallback string) *string {
	if value, exists := os.LookupEnv(key); exists {
		return &value
	}
	return &fallback
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = getEnv("HOST", "localhost")
		domainF   = getEnv("DOMAIN", "")
		httpPortF = getEnv("HTTP_PORT", "8080")
	)

	secureF, _ := strconv.ParseBool(*getEnv("SECURE", "false"))
	dbgF, _ := strconv.ParseBool(*getEnv("DEBUG", "false"))

	// Setup logger. Replace logger with your own log package of choice.
	format := log.FormatJSON
	if log.IsTerminal() {
		format = log.FormatTerminal
	}
	ctx := log.Context(context.Background(), log.WithFormat(format))
	if dbgF {
		ctx = log.Context(ctx, log.WithDebug())
		log.Debugf(ctx, "debug logs enabled")
	}
	log.Print(ctx, log.KV{K: "http-port", V: *httpPortF})

	// Initialize the services.
	var (
		courseSvc     course.Service
		enrollmentSvc enrollment.Service
	)
	{
		courseSvc = controllers.NewCourse()
		enrollmentSvc = controllers.NewEnrollment()
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		courseEndpoints     *course.Endpoints
		enrollmentEndpoints *enrollment.Endpoints
	)
	{
		courseEndpoints = course.NewEndpoints(courseSvc)
		courseEndpoints.Use(debug.LogPayloads())
		courseEndpoints.Use(log.Endpoint)
		enrollmentEndpoints = enrollment.NewEndpoints(enrollmentSvc)
		enrollmentEndpoints.Use(debug.LogPayloads())
		enrollmentEndpoints.Use(log.Endpoint)
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
			if secureF {
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
			handleHTTPServer(ctx, u, courseEndpoints, enrollmentEndpoints, &wg, errc, dbgF)
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
