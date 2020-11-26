package net

import "context"

type VideoNetwork interface {
	NewBroadcaster(strmID string) Broadcaster
	NewSubscriber(strmID string) Subscriber
	GetSubscriber(strmID string) Subscriber
}

type Broadcaster interface {
	Broadcast(seqNo uint64, data []byte) error
	Finish() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, f func(seqNo uint64, data []byte)) error
	Unsubscribe() error
}

type TranscodeProfile struct {
	Name      string
	Bitrate   uint
	Framerate uint
}

type TranscodeConfig struct {
	StrmID   string
	Profiles []TranscodeProfile
}

type Transcoder interface {
	Transcode(strmID string, config TranscodeConfig, gotPlaylist func(masterPlaylist []byte)) error
}
