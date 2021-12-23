package core

import (
	"errors"
	"math/big"
	"math/rand"
	"net/url"
	"sync"
	"time"
)

var MaxSessions = 10

type NodeType int

const (
	DefaultNode NodeType = iota
	BroadcasterNode
)

var nodeTypeStrs = map[NodeType]string{
	DefaultNode:      "default",
	BroadcasterNode:  "broadcaster",
}

func (t NodeType) String() string {
	str, ok := nodeTypeStrs[t]
	if !ok {
		return "unknown"
	}
	return str
}

//DeAINode handles videos going in and coming out of the DeAI network.
type DeAINode struct {

	// Common fields
	Eth      eth.DeAIEthClient
	WorkDir  string
	NodeType NodeType
	Database *common.DB

	// Broadcaster public fields
	Sender pm.Sender

	// Thread safety for config fields
	mu sync.RWMutex
	// Transcoder private fields
	priceInfo    *big.Rat
	serviceURI   url.URL
	segmentMutex *sync.RWMutex
}

//NewDeAINode creates a new DeAI Node. Eth can be nil.
func NewDeAINode(e eth.DeAIEthClient, wd string, dbh *common.DB) (*DeAINode, error) {
	rand.Seed(time.Now().UnixNano())
	return &DeAINode{
		Eth:             e,
		WorkDir:         wd,
		Database:        dbh,
		AutoAdjustPrice: true,
		SegmentChans:    make(map[ManifestID]SegmentChan),
		segmentMutex:    &sync.RWMutex{},
	}, nil
}

func (n *DeAINode) GetServiceURI() *url.URL {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return &n.serviceURI
}

func (n *DeAINode) SetServiceURI(newUrl *url.URL) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.serviceURI = *newUrl
}

// SetBasePrice sets the base price for an orchestrator on the node
func (n *DeAINode) SetBasePrice(price *big.Rat) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.priceInfo = price
}

// GetBasePrice gets the base price for an orchestrator
func (n *DeAINode) GetBasePrice() *big.Rat {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.priceInfo
}
