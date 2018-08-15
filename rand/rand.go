package rand

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Randomizer is the interface of the randomizer object which creates pseudorandom bits
type Randomizer interface {
	GetBit() (int, error)
	GetInt(int, int) (int, error)
	Powerup() error
	Shutdown() error
}

// randomizer is a the basic structure which creates random bits
type randomizer struct {
	mux              sync.Mutex
	lastcall         int64
	bit              int
	stop             chan struct{}
	intervalInMillis int64
	running          bool
}

// NewRandomizer creates a new Randomizer object
// Minimum interval is 20
func NewRandomizer(intervalInMillis int64) (Randomizer, error) {
	if intervalInMillis < 20 {
		return nil, errors.New("Minimum of 20 milliseconds interval is required")
	}
	r := randomizer{}
	r.intervalInMillis = intervalInMillis
	return &r, nil
}

// GetBit returns a new pseudorandom bit
func (r *randomizer) GetBit() (int, error) {
	if !r.running {
		return -1, errors.New("Randomizer not running")
	}
	a := time.Now().UnixNano() / int64(time.Millisecond)
	gap := a - r.lastcall
	if gap < r.intervalInMillis {
		time.Sleep(time.Duration(r.intervalInMillis-gap) * time.Millisecond)
	}
	r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
	return r.bit, nil
}

// GetInt returns a new pseudorandom integer
func (r *randomizer) GetInt(lowerBound int, upperBound int) (int, error) {
	normalizedRange := upperBound - lowerBound
	numBits := math.Ceil(math.Log2(float64(normalizedRange)))
	b := false
	ans := -1
	for !b {
		a := make([]int, int(numBits))
		for i := 0; i < int(numBits); i++ {
			bit, err := r.GetBit()
			if err != nil {
				return -1, err
			}
			a = append(a, bit)
		}
		s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
		integer, err := strconv.ParseInt(s, 2, int(numBits+1))
		if err != nil {
			return -1, errors.New("Error in bit conversion to integer")
		}
		ans = int(integer) + lowerBound
		if ans < upperBound {
			b = true
		}
	}
	return ans, nil
}

// Start the Randomizer
func (r *randomizer) Powerup() error {
	if r.running {
		return errors.New("Cannot power up a running Randomizer")
	}
	r.stop = make(chan struct{})
	go r.randomize(0)
	go r.randomize(1)
	r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
	r.running = true
	return nil
}

// Stop the Randomizer
func (r *randomizer) Shutdown() error {
	if !r.running {
		return errors.New("Cannot shutdown a non-running Randomizer")
	}
	close(r.stop)
	r.running = false
	return nil
}

// randomize is an infinite loop which tries to set a bit to its parameter
func (r *randomizer) randomize(bitwise int) {
	for {
		select {
		default:
			if r.bit != bitwise {
				r.mux.Lock()
				r.bit = bitwise
				r.mux.Unlock()
			}
		case <-r.stop:
			return
		}
	}
}
