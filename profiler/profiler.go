package profiler

import (
	"errors"
	"net/http/pprof"
	"strings"

	"github.com/nano-interactive/go-utils"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var (
	pprofIndex        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
	pprofCmdline      = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	pprofProfile      = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	pprofSymbol       = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	pprofTrace        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	pprofAllocs       = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("allocs").ServeHTTP)
	pprofBlock        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("block").ServeHTTP)
	pprofGoroutine    = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("goroutine").ServeHTTP)
	pprofHeap         = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("heap").ServeHTTP)
	pprofMutex        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("mutex").ServeHTTP)
	pprofThreadcreate = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("threadcreate").ServeHTTP)
)

var ErrNotProfiler = errors.New("not a profiler route")

func Handler(c *fasthttp.RequestCtx) error {
	path := utils.UnsafeString(c.Path())

	if len(path) < 12 || !strings.HasPrefix(path, "/debug/pprof") {
		return ErrNotProfiler
	}

	switch path {
	case "/debug/pprof/":
		pprofIndex(c)
	case "/debug/pprof/cmdline":
		pprofCmdline(c)
	case "/debug/pprof/profile":
		pprofProfile(c)
	case "/debug/pprof/symbol":
		pprofSymbol(c)
	case "/debug/pprof/trace":
		pprofTrace(c)
	case "/debug/pprof/allocs":
		pprofAllocs(c)
	case "/debug/pprof/block":
		pprofBlock(c)
	case "/debug/pprof/goroutine":
		pprofGoroutine(c)
	case "/debug/pprof/heap":
		pprofHeap(c)
	case "/debug/pprof/mutex":
		pprofMutex(c)
	case "/debug/pprof/threadcreate":
		pprofThreadcreate(c)
	default:
		// pprof index only works with trailing slash
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
		} else {
			path = "/debug/pprof/"
		}

		c.Redirect(path, fasthttp.StatusFound)
	}

	return nil
}
