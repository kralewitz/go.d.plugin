package springboot2

import (
	"github.com/netdata/go.d.plugin/modules"
	"github.com/netdata/go.d.plugin/pkg/prometheus"
	"github.com/netdata/go.d.plugin/pkg/utils"
	"github.com/netdata/go.d.plugin/pkg/web"
)

// New creates Springboot2 with default values
func New() *Springboot2 {
	return &Springboot2{}
}

// Springboot2 Spring boot 2 module
type Springboot2 struct {
	modules.Base

	web.HTTP `yaml:",inline"`

	prom prometheus.Prometheus
}

type metrics struct {
	ThreadsDaemon int64 `stm:"threads_daemon"`
	Threads       int64 `stm:"threads"`
}

// Cleanup makes cleanup
func (Springboot2) Cleanup() {}

// Init makes initialization
func (s *Springboot2) Init() bool {
	s.prom = prometheus.New(s.CreateHTTPClient(), s.RawRequest)
	return true
}

// Check makes check
func (s *Springboot2) Check() bool {
	rawMetrics, err := s.prom.Scrape()
	if err != nil {
		s.Error(err)
		return false
	}
	jvmMemory := rawMetrics.FindByName("jvm_memory_used_bytes")

	return len(jvmMemory) > 0
}

// Charts creates Charts
func (Springboot2) Charts() *Charts {
	return charts.Copy()
}

// GatherMetrics gathers metrics
func (s *Springboot2) GatherMetrics() map[string]int64 {
	rawMetrics, err := s.prom.Scrape()
	if err != nil {
		return nil
	}

	var m metrics
	m.ThreadsDaemon = int64(rawMetrics.FindByName("jvm_threads_daemon").Max())
	m.Threads = int64(rawMetrics.FindByName("jvm_threads_live").Max())

	return utils.ToMap(m)
}

func init() {
	creator := modules.Creator{
		Create: func() modules.Module { return New() },
	}

	modules.Register("springboot2", creator)
}