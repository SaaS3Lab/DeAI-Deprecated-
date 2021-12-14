package core

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ericxtang/m3u8"
	"github.com/golang/glog"
)
package core

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

var ErrNotFound = errors.New("NotFound")

const HLSWaitTime = time.Second * 10

type StreamDB struct {
	streams    map[StreamID]stream.VideoStream_
	SelfNodeID string
}

func NewStreamDB(selfNodeID string) *StreamDB {
	return &StreamDB{
		streams:    make(map[StreamID]stream.VideoStream_),
		SelfNodeID: selfNodeID}
}

func (s *StreamDB) GetHLSStream(id StreamID) stream.HLSVideoStream {
	strm, ok := s.streams[id]
	if !ok {
		return nil
	}
	if strm.GetStreamFormat() != stream.HLS {
		return nil
	}
	return strm.(stream.HLSVideoStream)
}

func (s *StreamDB) AddNewHLSStream(strmID StreamID) (strm stream.HLSVideoStream, err error) {
	strm = stream.NewBasicHLSVideoStream(strmID.String(), stream.DefaultSegWaitTime)
	s.streams[strmID] = strm

	// glog.Infof("Adding new video stream with ID: %v", strmID)
	return strm, nil
}


func (s *StreamDB) AddStream(strmID StreamID, strm stream.VideoStream_) (err error) {
	s.streams[strmID] = strm
	return nil
}

func (s *StreamDB) DeleteStream(strmID StreamID) {
	strm, ok := s.streams[strmID]
	if !ok {
		return
	}

	if strm.GetStreamFormat() == stream.HLS {
		//Remove all the variant lookups too
		hlsStrm := strm.(stream.HLSVideoStream)
		mpl, err := hlsStrm.GetMasterPlaylist()
		if err != nil {
			glog.Errorf("Error getting master playlist: %v", err)
		}
		for _, v := range mpl.Variants {
			vName := strings.Split(v.URI, ".")[0]
			delete(s.streams, StreamID(vName))
		}
	}
	delete(s.streams, strmID)
}

func (s StreamDB) String() string {
	streams := ""
	for vid, s := range s.streams {
		streams = streams + fmt.Sprintf("\nVariantID:%v, %v", vid, s)
	}

	return fmt.Sprintf("\nStreams:%v\n\n", streams)
}