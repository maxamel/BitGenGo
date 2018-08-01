package rand

import (
"sync"
"time"
//"fmt"
)

type Rand struct {
  mux sync.Mutex
  lastcall int64
  bit int
  stop chan struct{}
  interval int
}

func (r *Rand) GetBit() (int) {
  a := time.Now().UnixNano() / int64(time.Millisecond)
  gap := a - r.lastcall
  if (gap < 100) {
    time.Sleep(time.Duration(100-gap) * time.Millisecond)
  }
  r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
  return r.bit
}

func (r *Rand) Powerup() {
  r.stop = make(chan struct{})
  go r.randomize(0)
  go r.randomize(1)
  r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
}

func (r *Rand) Shutdown() {
  close(r.stop)
}

func (r *Rand) randomize(bitwise int) {
  for {
    select {
      default:
        //fmt.Println("BITWISE %d\n", bitwise)
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
