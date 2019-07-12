package cmd

import (
    "context"
    "fmt"
    "github.com/hitman99/hashing/controllers"
    "github.com/hitman99/hashing/gen/hashing"
    hashingsvr "github.com/hitman99/hashing/gen/http/hashing/server"
    "github.com/spf13/cobra"
    goahttp "goa.design/goa/v3/http"
    httpmdlwr "goa.design/goa/v3/http/middleware"
    "goa.design/goa/v3/middleware"
    "log"
    "net/http"
    "net/url"
    "os"
    "os/signal"
    "sync"
    "time"
)

var httpServerCmd = &cobra.Command{
    Use:   "server",
    Short: "serves an application by starting an HTTP server",
    Run: func(cmd *cobra.Command, args []string) {
        server(debug)
    },
}

func server(debug bool) {
    logger := log.New(os.Stderr, "[hashingapi] ", log.Ltime)
    hashingSvc := controllers.NewHashing(logger)
    hashingEndpoints := hashing.NewEndpoints(hashingSvc)
    errchan := make(chan error)

    // Catch signals
    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        errchan <- fmt.Errorf("%s", <-c)
    }()

    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())

    // Start the servers and send errors (if any) to the error channel.
    u, err := url.Parse("http://0.0.0.0:8080")
    if err != nil {
        fmt.Fprintf(os.Stderr, "invalid URL")
        os.Exit(1)
    }
    handleHTTPServer(ctx, u, hashingEndpoints, &wg, errchan, logger, debug)

    // Wait for signal.
    logger.Printf("exiting (%v)", <-errchan)
    // Send cancellation signal to the goroutines.
    cancel()
    wg.Wait()
    logger.Println("exited")
}

func handleHTTPServer(ctx context.Context, u *url.URL, hashingEndpoints *hashing.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {
    adapter := middleware.NewLogger(logger)
    // Provide the transport specific request decoder and response encoder.
    // The goa http package has built-in support for JSON, XML and gob.
    // Other encodings can be used by providing the corresponding functions,
    // see goa.design/encoding.
    var (
        dec = goahttp.RequestDecoder
        enc = goahttp.ResponseEncoder
        mux = goahttp.NewMuxer()
        eh = errorHandler(logger)
        hashingServer = hashingsvr.New(hashingEndpoints, mux, dec, enc, eh)
    )
    // Configure the mux.
    hashingsvr.Mount(mux, hashingServer)

    // Wrap the multiplexer with additional middlewares. Middlewares mounted
    // here apply to all the service endpoints.
    var handler http.Handler = mux
    {
        if debug {
            handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
        }
        handler = httpmdlwr.Log(adapter)(handler)
        handler = httpmdlwr.RequestID()(handler)
    }

    // Start HTTP server using default configuration, change the code to
    // configure the server as required by your service.
    srv := &http.Server{Addr: u.Host, Handler: handler}
    for _, m := range hashingServer.Mounts {
        logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
    }

    (*wg).Add(1)
    go func() {
        defer (*wg).Done()

        // Start HTTP server in a separate goroutine.
        go func() {
            logger.Printf("HTTP server listening on %q", u.Host)
            errc <- srv.ListenAndServe()
        }()

        <-ctx.Done()
        logger.Printf("shutting down HTTP server at %q", u.Host)

        // Shutdown gracefully with a 30s timeout.
        ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
        defer cancel()

        srv.Shutdown(ctx)
    }()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
    return func(ctx context.Context, w http.ResponseWriter, err error) {
        id := ctx.Value(middleware.RequestIDKey).(string)
        w.Write([]byte("[" + id + "] encoding: " + err.Error()))
        logger.Printf("[%s] ERROR: %s", id, err.Error())
    }
}