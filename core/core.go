package core

import (
	"context"

	"github.com/golang/glog"
)

type NodeID string

type DeAINode struct {
	Identity     NodeID
	VideoNetwork net.VideoNetwork
	StreamDB     *StreamDB
	Eth          eth.Client
	IsTranscoder bool
}

func (n *DeAINode) Start() {
	//Connect to bootstrap node, ask for more peers

	//Connect to peers

	//Kick off process to periodically monitor peer connection by pinging them
}

func (n *DeAINode) monitorEth() {

}


func (n *DeAINode) BroadcastToNetwork(ctx context.Context) error {
	b := n.VideoNetwork.NewBroadcaster(strm.GetDeAIID())

	//Prepare the broadcast.  May have to send the MasterPlaylist as part of the handshake.

	//Kick off a go routine to broadcast the stream
	go func() {
		for {
			seg, err := strm.ReadHLSSegment()
			if err != nil {
				glog.Errorf("Error reading hls stream while broadcasting to network: %v", err)
				return //TODO: Should better handle error here
			}

			//Encode seg into []byte, then send it via b.Broadcast
		}
	}()

	select {
	case <-ctx.Done():
		glog.Errorf("Done Broadcasting")
		return nil
	}
}

//Create a new data stream
func (n *DeAINode) SubscribeFromNetwork(ctx context.Context, strmID DeAIID) (*stream.VideoStream, error) {

}

//UnsubscribeFromNetwork unsubscribes to a data stream on the network.
func (n *DeAINode) UnsubscribeFromNetwork(strmID DeAIID) error {
	
}
