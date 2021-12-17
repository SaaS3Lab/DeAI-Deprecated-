package server

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/ericxtang/m3u8"
	"github.com/golang/glog"
	"github.com/nareix/joy4/av"
)

var ErrNotFound = errors.New("NotFound")
var ErrAlreadyExists = errors.New("requestAlreadyExists")
var HLSWaitTime = time.Second * 10
var HLSBufferCap = uint(43200) //12 hrs assuming 1s segment
var HLSBufferWindow = uint(5)
var SegOptions = segmenter.SegmenterOptions{SegLength: 8 * time.Second}
var HLSUnsubWorkerFreq = time.Second * 5

type Server struct {
	LPMS         *lpms.LPMS
	HttpPort     string

	hlsSubTimer           map[core.requestID]time.Time
	hlsWorkerRunning      bool
}


func (s *Server) StartLPMS(ctx context.Context) error {
	s.hlsSubTimer = make(map[core.requestID]time.Time)
	go s.startHlsUnsubscribeWorker(time.Second*10, HLSUnsubWorkerFreq)


	http.HandleFunc("/transcode", func(w http.ResponseWriter, r *http.Request) {
		//Tell the LivepeerNode to create a job
	})

	http.HandleFunc("/createrequest", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/localrequests", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/peersCount", func(w http.ResponseWriter, r *http.Request) {
	})

	http.HandleFunc("/requesterStatus", func(w http.ResponseWriter, r *http.Request) {
	})

	err := s.LPMS.Start(ctx)
	if err != nil {
		glog.Errorf("Cannot start LPMS: %v", err)
		return err
	}

	return nil
}



func parserequestID(reqPath string) core.requestID {
	var strmID string
	regex, _ := regexp.Compile("\\/request\\/([[:alpha:]]|\\d)*")
	match := regex.FindString(reqPath)
	if match != "" {
		strmID = strings.Replace(match, "/request/", "", -1)
	}
	return core.requestID(strmID)
}
