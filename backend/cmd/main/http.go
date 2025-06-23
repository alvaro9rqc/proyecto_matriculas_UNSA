package main

import (
	"context"
	"net/http"
	"sync"
	"time"

	coursesvr "github.com/enrollment/gen/http/course/server"
	enrollmentsvr "github.com/enrollment/gen/http/enrollment/server"
	oauthsvr "github.com/enrollment/gen/http/oauth/server"

	"github.com/enrollment/gen/course"
	"github.com/enrollment/gen/enrollment"
	"github.com/enrollment/gen/oauth"
	"github.com/enrollment/gen/queue"

	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

func handleHTTPServer(
	ctx context.Context,
	port string,
	courseEndpoints *course.Endpoints,
	enrollmentEndpoints *enrollment.Endpoints,
	oauthEndpoints *oauth.Endpoints,
	queueEndpoints *queue.Endpoints,
	wg *sync.WaitGroup,
	errc chan error,
	dbg bool) {

	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
		if dbg {
			debug.MountPprofHandlers(debug.Adapt(mux))
			debug.MountDebugLogEnabler(debug.Adapt(mux))
		}
	}

	var (
		courseServer     *coursesvr.Server
		enrollmentServer *enrollmentsvr.Server
		oauthServer      *oauthsvr.Server
		// queueServer      *queuesvr.Server
	)
	{
		eh := errorHandler(ctx)
		courseServer = coursesvr.New(courseEndpoints, mux, dec, enc, eh, nil)
		enrollmentServer = enrollmentsvr.New(enrollmentEndpoints, mux, dec, enc, eh, nil)
		oauthServer = oauthsvr.New(oauthEndpoints, mux, dec, enc, eh, nil)
		// queueServer = queuesvr.New(queueEndpoints, mux, dec, enc, eh, nil)
	}

	coursesvr.Mount(mux, courseServer)
	enrollmentsvr.Mount(mux, enrollmentServer)
	oauthsvr.Mount(mux, oauthServer)

	var handler http.Handler = mux
	if dbg {
		handler = debug.HTTP()(handler)
	}
	handler = log.HTTP(ctx)(handler)

	srv := &http.Server{Addr: ":" + port, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range courseServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range enrollmentServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range oauthServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	// for _, m := range queueServer.Mounts {
	// 	log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	// }

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		go func() {
			log.Printf(ctx, "HTTP server listening on port %q", port)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down HTTP server at port %q", port)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf(ctx, "failed to shutdown: %v", err)
		}
	}()
}

func errorHandler(logCtx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf(logCtx, "ERROR: %s", err.Error())
	}
}
