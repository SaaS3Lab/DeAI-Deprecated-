package core

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

)

var ErrSourceID = errors.New("ErrSourceID")

const (
	DefaultSourceIDLength = 4
)

type DataParameters struct {
	SourceID       SourceID
	OS               drivers.OSSession
	RecordOS         drivers.OSSession
	Capabilities     *Capabilities
}

type SourceID string

func SplitSourceIDString(str string) SourceID {
	parts := strings.SplitN(str, "/", 2)
	return MakeSourceIDFromString(parts[0], parts[1])
}

func (id SourceID) String() string {
	return fmt.Sprintf("%v/%v", id.SourceID, id.Rendition)
}

func RandomSourceID() SourceID {
	return SourceID(common.RandomIDGenerator(DefaultSourceIDLength))
}