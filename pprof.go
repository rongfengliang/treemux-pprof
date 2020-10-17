package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/vmihailenco/treemux"
)

const (
	defaultPrefix string = ""
)

// RouterpprofRegister for treemux
func RouterpprofRegister(rg *treemux.TreeMux) {
	prefixRouter := rg.NewGroup(defaultPrefix)
	{
		prefixRouter.GET("/", pprofHandler(pprof.Index))
		prefixRouter.GET("/cmdline", pprofHandler(pprof.Cmdline))
		prefixRouter.GET("/profile", pprofHandler(pprof.Profile))
		prefixRouter.POST("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/symbol", pprofHandler(pprof.Symbol))
		prefixRouter.GET("/trace", pprofHandler(pprof.Trace))
		prefixRouter.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		prefixRouter.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		prefixRouter.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		prefixRouter.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		prefixRouter.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		prefixRouter.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

func pprofHandler(h http.HandlerFunc) treemux.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(w http.ResponseWriter, req treemux.Request) error {
		handler.ServeHTTP(w, req.Request)
		return nil
	}
}
