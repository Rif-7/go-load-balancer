package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

// Backend represents a backend server
type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// SetAlive updates the alive status of backend
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// IsAlive returns true when backend is alive
func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}

// LoadBalancer represents a load balancer
type LoadBalancer struct {
	backends []*Backend
	current  uint64
}

// NextBackend returns the next available backend to handle the request
func (lb *LoadBalancer) NextBackend() *Backend {
	// Simple round-robin
	next := atomic.AddUint64(&lb.current, uint64(1)) % uint64(len(lb.backends))

	// Find the next available backend
	for i := 0; i < len(lb.backends); i++ {
		idx := (int(next) + i) % len(lb.backends)
		if lb.backends[idx].IsAlive() {
			return lb.backends[idx]
		}
	}
	return nil
}
