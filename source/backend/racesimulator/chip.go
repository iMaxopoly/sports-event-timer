package racesimulator

import (
	"errors"
	"log"
	"strings"

	"github.com/rs/xid"
)

var (
	errChipSetIdentifierInvalid = errors.New("chip identifier set is invalid")
)

// IChip is the interface that wraps the underlying Identifier and SetIdentifier methods.
// Indentifier returns a valid string and SetIdentifier sets it.
// This helps form consistency with derivative structures.
type IChip interface {
	Identifier() string
	SetIdentifier(string) error
}

// chip is an implementation struct of IChip interface and serves as
// the chip to be used in both cases; the athlete as well as the timepoint marks
// on the racetrack.
type Chip struct{ identifier string }

func (c *Chip) Identifier() string { return c.identifier }

func (c *Chip) SetIdentifier(id string) error {
	if strings.TrimSpace(id) == "" {
		return errChipSetIdentifierInvalid
	}

	c.identifier = id
	return nil
}

// NewChip returns a new interface of IChip with valid and unique string.
// The unique string generated with the help of xid, a 3rd party library.
func NewChip() IChip {
	var c Chip
	err := c.SetIdentifier(xid.New().String())
	if err != nil {
		log.Fatal(err)
	}
	return &c
}
