package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/enrollment/config"
	// course "github.com/enrollment/gen/course"
	// enrollment "github.com/enrollment/gen/enrollment"
	"github.com/enrollment/gen/institution"
	"github.com/enrollment/gen/oauth"
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

	// Initialize repositories
	var (
		oauthRepo       = repositories.NewOauthRepository(conn)
		institutionRepo = repositories.NewInstitutionRepository(conn)
		processRepo     = repositories.NewProcessRepository(conn)
		studentRepo     = repositories.NewStudentRepository(conn)
		courseRepo      = repositories.NewCourseRepository(conn)
	)

	//Initialize services
	var (
		institutionSvc institution.Service
		oauthSvc       oauth.Service
		// queueSvc      queue.Service
		userSvc user.Service
	)
	{
		institutionSvc = controllers.NewInstitution(oauthRepo, institutionRepo, processRepo, studentRepo, courseRepo)
		oauthSvc = controllers.NewOauth(cfg, oauthRepo)
		// queueSvc = controllers.NewQueue()
		// userSvc = controllers.NewUser(cfg, adminRepo, studentMajorRepo, speakerRepo)
	}

	var (
		institutionEndpoints *institution.Endpoints
		oauthEndpoints       *oauth.Endpoints
		// queueEndpoints      *queue.Endpoints
		userEndpoints *user.Endpoints
	)
	{
		institutionEndpoints = institution.NewEndpoints(institutionSvc)
		institutionEndpoints.Use(debug.LogPayloads())
		institutionEndpoints.Use(log.Endpoint)

		oauthEndpoints = oauth.NewEndpoints(oauthSvc)
		oauthEndpoints.Use(debug.LogPayloads())
		oauthEndpoints.Use(log.Endpoint)

		// queueEndpoints = queue.NewEndpoints(queueSvc)
		// queueEndpoints.Use(debug.LogPayloads())
		// queueEndpoints.Use(log.Endpoint)

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
		institutionEndpoints,
		oauthEndpoints,
		// queueEndpoints,
		&wg,
		errc,
		cfg.Dbg,
	)

	log.Printf(ctx, "exiting (%v)", <-errc)

	cancel()

	wg.Wait()
	log.Printf(ctx, "exited")
}
