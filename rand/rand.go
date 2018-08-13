package rand

import (
"sync"
"time"
"errors"
"math"
"strings"
"fmt"
"strconv"
)

type Randomizer interface {
  GetBit() (int, error)
  GetInt(int, int) (int, error)
  Powerup() error
  Shutdown() error
}

type randomizer struct {
  mux sync.Mutex
  lastcall int64
  bit int
  stop chan struct{}
  intervalInMillis int64
  running bool
}

// Minimum interval is 20
func NewRandomizer(intervalInMillis int64) (Randomizer, error){
  if intervalInMillis < 20 {
    return nil, errors.New("Minimum of 20 milliseconds interval is required")
  }
  r := randomizer{}
  r.intervalInMillis = intervalInMillis
  return &r, nil
}

func (r *randomizer) GetBit() (int, error) {
  if !r.running {
    return -1, errors.New("Randomizer not running")
  }
  a := time.Now().UnixNano() / int64(time.Millisecond)
  gap := a - r.lastcall
  if (gap < r.intervalInMillis) {
    time.Sleep(time.Duration(r.intervalInMillis-gap) * time.Millisecond)
  }
  r.lastcall = time.Now().UnixNano() / int64(time.Millisecond)
  return r.bit, nil
}

func (r *randomizer) GetInt(lowerBound int, upperBound int) (int, error) {
  normalizedRange := upperBound - lowerBound
  numBits := math.Ceil(math.Log2(float64(normalizedRange)))
  a := make([]int, int(numBits))
  b := false
  ans := -1
  for !b {
  for i:=0; i<int(numBits); i++ {
    bit, err := r.GetBit()
    if err != nil {
      return -1, err
    }
    a = append(a, bit)
  }
  s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
  fmt.Println(s)
  integer, err := strconv.ParseInt(s, 2, int(numBits+1))
  if err != nil {
    return -1, errors.New("Error in bit conversion to integer")
  }
  ans = int(integer) + lowerBound
  fmt.Println(ans)
  if ans < upperBound {
    b = true
  }
  }
  return ans, nil
}

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

func (r *randomizer) Shutdown() error {
  if !r.running {
    return errors.New("Cannot shutdown a non-running Randomizer")
  }
  close(r.stop)
  r.running = false
  return nil
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
