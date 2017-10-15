package main

import (
	"fmt"

	"github.com/ssut/cryptoticker"
)

type poolWorkerConfig struct {
	parser     *cryptoticker.Parser
	subscriber *cryptoticker.Subscriber
}

type worker struct {
	parser     *cryptoticker.Parser
	subscriber *cryptoticker.Subscriber
	running    bool
}

type pool struct {
	workers []*worker
	cap     int
	size    int
}

func newPool(size int) *pool {
	p := &pool{
		workers: make([]*worker, size),
		size:    0,
		cap:     size,
	}

	return p
}

func (p *pool) addWorker(config *poolWorkerConfig) error {
	if config.parser == nil || config.subscriber == nil {
		return fmt.Errorf("either of cryptoticker needs to be set")
	} else if p.size == p.size {
		return fmt.Errorf("the number of worker limit exceeded")
	}

	p.size += 1
	p.workers[p.size] = &worker{
		parser:     config.parser,
		subscriber: config.subscriber,
	}

	return nil
}
