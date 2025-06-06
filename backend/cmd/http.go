package main

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	auth "github.com/enrollment/gen/auth"
	course "github.com/enrollment/gen/course"
	enrollment "github.com/enrollment/gen/enrollment"
	authsvr "github.com/enrollment/gen/http/auth/server"
	coursesvr "github.com/enrollment/gen/http/course/server"
	enrollmentsvr "github.com/enrollment/gen/http/enrollment/server"
	oauthsvr "github.com/enrollment/gen/http/oauth/server"
	oauth "github.com/enrollment/gen/oauth"
	"goa.design/clue/debug"
	"goa.design/clue/log"
	goahttp "goa.design/goa/v3/http"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, courseEndpoints *course.Endpoints, enrollmentEndpoints *enrollment.Endpoints, oauthEndpoints *oauth.Endpoints, authEndpoints *auth.Endpoints, wg *sync.WaitGroup, errc chan error, dbg bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and mount debug and profiler
	// endpoints in debug mode.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
		if dbg {
			// Mount pprof handlers for memory profiling under /debug/pprof.
			debug.MountPprofHandlers(debug.Adapt(mux))
			// Mount /debug endpoint to enable or disable debug logs at runtime.
			debug.MountDebugLogEnabler(debug.Adapt(mux))
		}
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		courseServer     *coursesvr.Server
		enrollmentServer *enrollmentsvr.Server
		oauthServer      *oauthsvr.Server
		authServer       *authsvr.Server
	)
	{
		eh := errorHandler(ctx)
		courseServer = coursesvr.New(courseEndpoints, mux, dec, enc, eh, nil)
		enrollmentServer = enrollmentsvr.New(enrollmentEndpoints, mux, dec, enc, eh, nil)
		oauthServer = oauthsvr.New(oauthEndpoints, mux, dec, enc, eh, nil)
		authServer = authsvr.New(authEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	coursesvr.Mount(mux, courseServer)
	enrollmentsvr.Mount(mux, enrollmentServer)
	oauthsvr.Mount(mux, oauthServer)
	authsvr.Mount(mux, authServer)

	var handler http.Handler = mux
	if dbg {
		// Log query and response bodies if debug logs are enabled.
		handler = debug.HTTP()(handler)
	}
	handler = log.HTTP(ctx)(handler)

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range courseServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range enrollmentServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range oauthServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range authServer.Mounts {
		log.Printf(ctx, "HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			log.Printf(ctx, "HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		log.Printf(ctx, "shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf(ctx, "failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logCtx context.Context) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf(logCtx, "ERROR: %s", err.Error())
	}
}
