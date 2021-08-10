package proxy

import (
	"crypto/tls"
	"ggate/internal/config"
	"ggate/internal/context"
	"ggate/internal/route"
	"ggate/internal/service"

	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type Proxy struct {
	routes         []*route.Route
	services       *map[string]service.Service
	client         *http.Client
	defaultFilters []route.Filter
	certs          map[string]tls.Certificate
	pool           sync.Pool
	cfg            *config.ProxyConfig
}

func New(cfg config.ProxyConfig) *Proxy {
	routes, err := route.GetRoutesBy(cfg.RoutesFile)
	if err != nil {
		logrus.Panic(err)
	}
	services, err := service.GetAllServicesBy(cfg.ServicesFile)
	if err != nil {
		logrus.Panic(err)
	}
	filters, err := route.GetDefaultFilters()
	if err != nil {
		logrus.Panic(err)
	}
	proxy := &Proxy{
		routes:         routes,
		services:       &services,
		defaultFilters: filters,
		cfg:            &cfg,
	}

	proxy.client = proxy.initHttpClient()
	proxy.pool.New = func() interface{} {
		return &context.Context{}
	}

	return proxy
}

func (p *Proxy) Run() {

	server := http.Server{
		Addr:    p.cfg.Addr,
		Handler: p,
	}

	logrus.Info("proxy Running at : http://" + server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		logrus.Panic(err)
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := p.pool.Get().(*context.Context)
	ctx.Request = r
	ctx.ResponseWriter = w
	oneMatch := false
	for _, route := range p.routes {
		match := true
		for _, pre := range route.Predicates {
			match = match && pre.Apply(ctx)
		}
		if !match {
			continue
		}
		oneMatch = true
		ctx.Handle()
	}
	if !oneMatch {
		p.handleNotMatch()
	}
	ctx.Reset()
	p.pool.Put(ctx)

	//all predicate.Apply(ctx) == true
	//build filter
	//ctx.do -> do filter -> redirect
}
func (p *Proxy) initHttpClient() *http.Client {
	maxIdleConns := p.cfg.MaxIdleConns
	MaxIdleConnsPerHost := p.cfg.MaxIdleConnsPerHost
	tansport := http.DefaultTransport
	pTransPort, _ := tansport.(*http.Transport)

	pTransPort.MaxIdleConns = maxIdleConns
	pTransPort.MaxIdleConnsPerHost = MaxIdleConnsPerHost
	return &http.Client{
		Transport: pTransPort,
	}
}

func (p *Proxy) handleNotMatch() {

}
