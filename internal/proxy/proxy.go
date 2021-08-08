package proxy

import (
	"crypto/tls"
	"ggate/internal/config"
	"ggate/internal/route"
	"ggate/internal/service"
	"ggate/pkg/requests"
	"io"

	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type Proxy struct {
	routes         []route.Route
	services       *map[string]service.Service
	client         *http.Client
	defaultFilters []route.Filter
	certs          map[string]tls.Certificate
	pool           sync.Pool
	cfg            *config.ProxyConfig
}

func New(cfg config.ProxyConfig) *Proxy {
	routes, err := route.GetRoutesBy(cfg.RoutesReaderType, cfg.RoutesReaderArg)
	if err != nil {
		logrus.Panic(err)
	}
	services, err := service.GetAllServicesBy(cfg.ServicesReaderType, cfg.ServicesReaderArg)
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
		return &Context{}
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
	result := requests.Request("http://www.douban.com", "GET", p.client).
		Headers(r.Header).
		Send()

	resp := result.Resp
	for name, values := range resp.Header {
		w.Header()[name] = values
	}
	w.WriteHeader(resp.StatusCode)
	_, err := io.Copy(w, resp.Body)
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()

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
