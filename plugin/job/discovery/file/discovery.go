package file

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/netdata/go.d.plugin/pkg/logger"
	"github.com/netdata/go.d.plugin/plugin/job/confgroup"
)

type Config struct {
	Registry confgroup.Registry
	Read     []string
	Watch    []string
}

func validateConfig(cfg Config) error {
	if len(cfg.Registry) == 0 {
		return errors.New("empty config registry")
	}
	if len(cfg.Read)+len(cfg.Watch) == 0 {
		return errors.New("discoverers not set")
	}
	return nil
}

type (
	discoverer interface {
		Run(ctx context.Context, in chan<- []*confgroup.Group)
	}
	Discovery struct {
		discoverers []discoverer
		*logger.Logger
	}
)

func NewDiscovery(cfg Config) (*Discovery, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("file discovery config validation: %v", err)
	}

	d := Discovery{
		Logger: logger.New("discovery", "file manager"),
	}
	if err := d.registerDiscoverers(cfg); err != nil {
		return nil, fmt.Errorf("file discovery initialization: %v", err)
	}
	return &d, nil
}

func (d Discovery) String() string {
	return fmt.Sprintf("file discovery: %v", d.discoverers)
}

func (d *Discovery) registerDiscoverers(cfg Config) error {
	if len(cfg.Read) != 0 {
		d.discoverers = append(d.discoverers, NewReader(cfg.Registry, cfg.Read))
	}
	if len(cfg.Watch) != 0 {
		d.discoverers = append(d.discoverers, NewWatcher(cfg.Registry, cfg.Watch))
	}
	if len(d.discoverers) == 0 {
		return errors.New("zero registered discoverers")
	}
	return nil
}

func (d *Discovery) Run(ctx context.Context, in chan<- []*confgroup.Group) {
	d.Info("instance is started")
	defer func() { d.Info("instance is stopped") }()

	var wg sync.WaitGroup

	for _, dd := range d.discoverers {
		wg.Add(1)
		go func(dd discoverer) {
			defer wg.Done()
			d.runDiscoverer(ctx, dd, in)
		}(dd)
	}

	wg.Wait()
	<-ctx.Done()
}

func (d *Discovery) runDiscoverer(ctx context.Context, dd discoverer, in chan<- []*confgroup.Group) {
	updates := make(chan []*confgroup.Group)
	go dd.Run(ctx, updates)
	for {
		select {
		case <-ctx.Done():
			return
		case groups, ok := <-updates:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case in <- groups:
			}
		}
	}
}