package rand

import (
"sync"
"time"
)

type Randomizer interface {
  GetBit() int
  Powerup()
  Shutdown()
}

type randomizer struct {
  mux sync.Mutex
  lastcall int64
  bit int
  stop chan struct{}
  intervalInMillis int64
}

// Minimum interval is 20
func NewRandomizer(intervalInMillis int64) Randomizer{
  r := randomizer{}
  r.intervalInMillis = 20
  if intervalInMillis > 19 {
    r.intervalInMillis = intervalInMillis
  }
  return &r
}

func (r *randomizer) GetBit() (int) {
  a := time.Now().UnixNano() / int64(time.Millisecond)
  gap := a - r.lastcall
  if (gap < r.intervalInMillis) {
    time.Sleep(time.Duration(r.intervalInMillis-gap) * time.Millisecond)
  }
  r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
  return r.bit
}

func (r *randomizer) Powerup() {
  r.stop = make(chan struct{})
  go r.randomize(0)
  go r.randomize(1)
  r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
}

func (r *randomizer) Shutdown() {
  close(r.stop)
}

func (r *randomizer) randomize(bitwise int) {
  for {
    select {
      default:
        if r.bit != bitwise {
          r.mux.Lock()
          r.bit = bitwise
          r.mux.Unlock()
        }
      case <- r.stop:
        return
    }
  }
}
