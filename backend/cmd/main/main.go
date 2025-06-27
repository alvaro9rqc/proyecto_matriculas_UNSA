package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/enrollment/config"
	course "github.com/enrollment/gen/course"
	enrollment "github.com/enrollment/gen/enrollment"
	"github.com/enrollment/gen/oauth"
	"github.com/enrollment/gen/queue"
	user "github.com/enrollment/gen/user"
	controllers "github.com/enrollment/internal/controllers"
	"github.com/enrollment/internal/db"
	repositories "github.com/enrollment/internal/repositories"

	"goa.design/clue/debug"
	"goa.design/clue/log"
)

func main() {
	cfg, err := config.NewMainConfig()
	if err != nil {
		panic(err)
	}

	conn, err := db.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	var (
		oauthRepo        = repositories.NewOauthRepository(conn)
		adminRepo        = repositories.NewAdministrativeRepository(conn)
		studentMajorRepo = repositories.NewStudentMajorRepository(conn)
		speakerRepo      = repositories.NewSpeakerRepository(conn)
	)

	var (
		courseSvc     course.Service
		enrollmentSvc enrollment.Service
		oauthSvc      oauth.Service
		queueSvc      queue.Service
		userSvc       user.Service
	)
	{
		courseSvc = controllers.NewCourse()
		enrollmentSvc = controllers.NewEnrollment()
		oauthSvc = controllers.NewOauth(cfg, oauthRepo)
		queueSvc = controllers.NewQueue()
		userSvc = controllers.NewUser(adminRepo, studentMajorRepo, speakerRepo)
	}

	var (
		courseEndpoints     *course.Endpoints
		enrollmentEndpoints *enrollment.Endpoints
		oauthEndpoints      *oauth.Endpoints
		queueEndpoints      *queue.Endpoints
		userEndpoints       *user.Endpoints
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

		queueEndpoints = queue.NewEndpoints(queueSvc)
		queueEndpoints.Use(debug.LogPayloads())
		queueEndpoints.Use(log.Endpoint)

		userEndpoints = user.NewEndpoints(userSvc)
		userEndpoints.Use(debug.LogPayloads())
		userEndpoints.Use(log.Endpoint)
	}

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(cfg.Ctx)

	handleHTTPServer(
		ctx,
		cfg.HttpPort,
		courseEndpoints,
		enrollmentEndpoints,
		oauthEndpoints,
		queueEndpoints,
		&wg,
		errc,
		cfg.Dbg,
	)

	log.Printf(ctx, "exiting (%v)", <-errc)

	cancel()

	wg.Wait()
	log.Printf(ctx, "exited")
}
