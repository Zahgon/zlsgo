//go:build go1.18
// +build go1.18

package zpool

import (
	"errors"

	"github.com/sohaha/zlsgo/zsync"
	"github.com/sohaha/zlsgo/zutil"
)

// BalancerStrategy represents the load balancing strategy
type BalancerStrategy int

const (
	// StrategyRandom represents random selection strategy
	StrategyRandom BalancerStrategy = iota
	// StrategyLeastConn represents least connections selection strategy
	StrategyLeastConn
	// StrategyRoundRobin represents round-robin selection strategy
	StrategyRoundRobin
	// StrategyWeighted represents weighted selection strategy
	StrategyWeighted
)

var (
	ErrKeyRequired      = errors.New("key is required")
	ErrNodeExists       = errors.New("node already exists")
	ErrNodeNotFound     = errors.New("node not found")
	ErrNoAvailableNodes = errors.New("no available nodes")
	ErrEmptyCallback    = errors.New("callback function cannot be empty")
	ErrNoNodesAdded     = errors.New("please add nodes first")
)

type Balancer[T any] struct {
	nodes       map[string]*balancerNode[T]
	mu          *zsync.RBMutex
	nodeKeys    []string
	lastNodeIdx uint64
}

type balancerNode[T any] struct {
	node     T
	total    *zutil.Int64
	failedAt *zutil.Int64
	cooldown *zutil.Int64
	max      int64
	weight   *zutil.Uint64
}

type BalancerNodeOptions struct {
	// Maximum concurrent connections per node, fairness not guaranteed
	MaxConns int64
	// Node weight, default is 1
	Weight uint64
	// Node cooldown period after failure, default is 1000ms
	Cooldown int64
}

// BalancerNodeInfo contains complete information about a balancer node
type BalancerNodeInfo[T any] struct {
	Node      T      // Node data
	Weight    uint64 // Current weight
	MaxConns  int64  // Maximum connections
	Cooldown  int64  // Cooldown period in milliseconds
	Available bool   // Whether the node is available
	Active    int64  // Current active connections
}

// NewBalancer creates a new load balancer
func NewBalancer[T any]() *Balancer[T] { _ = "STUB: not implemented"; return nil }

// Get returns the node with the given key
func (b *Balancer[T]) Get(key string) (node T, available bool, exists bool) {
	_ = "STUB: not implemented"
	return *new(T), false, false
}

// GetWeight returns the weight of the node with the given key
func (b *Balancer[T]) GetWeight(key string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetWeight updates the weight of the node with the given key
func (b *Balancer[T]) SetWeight(key string, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Add adds a new node to the load balancer
func (b *Balancer[T]) Add(key string, node T, opt ...func(opts *BalancerNodeOptions)) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove removes a node from the load balancer
func (b *Balancer[T]) Remove(key string) { _ = "STUB: not implemented"; return }

// Mark marks a node as available or not
func (b *Balancer[T]) Mark(key string, available bool) { _ = "STUB: not implemented"; return }

// getAvailableNodes returns all available nodes
func (b *Balancer[T]) getAvailableNodes(keys ...string) []*balancerNode[T] {
	_ = "STUB: not implemented"
	return nil
}

// selectNode selects a node based on the given strategy
func (b *Balancer[T]) selectNode(nodes []*balancerNode[T], strategy BalancerStrategy) (*balancerNode[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allNodesHaveSameWeight checks if all nodes have the same weight
func allNodesHaveSameWeight[T any](nodes []*balancerNode[T]) bool {
	_ = "STUB: not implemented"
	return false
}

var rngSeed uint32

func fastRand(n int) int { _ = "STUB: not implemented"; return 0 }

// Run runs the given function on the selected node
func (b *Balancer[T]) Run(fn func(node T) (normal bool, err error), strategy ...BalancerStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Balancer[T]) RunByKeys(keys []string, fn func(node T) (normal bool, err error), strategy ...BalancerStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

// rc.BackOffDelay = true

// WalkNodes walks all nodes
func (b *Balancer[T]) WalkNodes(fn func(node T, available bool) (normal bool)) {
	_ = "STUB: not implemented"
	return
}

// Keys returns all node keys
func (b *Balancer[T]) Keys() []string { _ = "STUB: not implemented"; return nil }

// Len returns the number of nodes
func (b *Balancer[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// GetNodeInfo returns complete information about the node with the given key
func (b *Balancer[T]) GetNodeInfo(key string) (BalancerNodeInfo[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}
