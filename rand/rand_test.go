package rand

import (
	"testing"
	"time"
)

func TestBasicBit(t *testing.T) {
	rnd, err := NewRandomizer(50)
	if err != nil {
		t.Errorf("Failed to create a new Randomizer")
	}
	err = rnd.Powerup()
	if err != nil {
		t.Errorf("Failed to powerup the Randomizer")
	}
	_, err = rnd.GetBit()
	if err != nil {
		t.Errorf("Failed to get a bit from the Randomizer")
	}
	err = rnd.Shutdown()
	if err != nil {
		t.Errorf("Failed to shutdown the Randomizer")
	}
}

func TestBasicBits(t *testing.T) {
	rnd, err := NewRandomizer(34)
	if err != nil {
		t.Errorf("Failed to create a new Randomizer")
	}
	err = rnd.Powerup()
	if err != nil {
		t.Errorf("Failed to powerup the Randomizer")
	}
	bits, err := rnd.GetBits(30)
	if err != nil || len(bits) != 30 {
		t.Errorf("Failed to get bits from the Randomizer")
	}
	err = rnd.Shutdown()
	if err != nil {
		t.Errorf("Failed to shutdown the Randomizer")
	}
}

func TestBasicInt(t *testing.T) {
	rnd, err := NewRandomizer(44)
	if err != nil {
		t.Errorf("Failed to create a new Randomizer")
	}
	err = rnd.Powerup()
	if err != nil {
		t.Errorf("Failed to powerup the Randomizer")
	}
	bit, err := rnd.GetInt(20, 100)
	if err != nil {
		t.Errorf("Failed to get an int from the Randomizer")
	}
	if bit > 100 || bit < 20 {
		t.Errorf("Failed to get an int in the given range")
	}
	err = rnd.Shutdown()
	if err != nil {
		t.Errorf("Failed to shutdown the Randomizer")
	}
}

func TestBasicInts(t *testing.T) {
        rnd, err := NewRandomizer(27)
        if err != nil {
                t.Errorf("Failed to create a new Randomizer")
        }
        err = rnd.Powerup()
        if err != nil {
                t.Errorf("Failed to powerup the Randomizer")
        }
        ints, err := rnd.GetInts(20, 20, 100)
        if err != nil || len(ints) != 20{
                t.Errorf("Failed to get ints from the Randomizer")
        }
	for _, num := range ints {
        	if num > 100 || num < 20 {
                	t.Errorf("Failed to maintain ints in the given range")
        	}
	}
        err = rnd.Shutdown()
        if err != nil {
                t.Errorf("Failed to shutdown the Randomizer")
        }
}


func TestEnforcedMinimumFrequency(t *testing.T) {
	rnd, _ := NewRandomizer(50)
	rnd.Powerup()
	rnd.GetBit()
	time1 := time.Now().UnixNano() / int64(time.Millisecond)
	rnd.GetBit()
	time2 := time.Now().UnixNano() / int64(time.Millisecond)
	rnd.Shutdown()
	if time2 - time1 < 50 {
		t.Errorf("Minimum frequency time not enforced")
	}
}

func Test1000MeanVariance(t *testing.T) {
	rnd, _ := NewRandomizer(20)
        rnd.Powerup()
        sum := 0
        for i := 1; i < 1000; i++ {
                b, _ := rnd.GetBit()
                sum = sum + b
        }
        rnd.Shutdown()
        var mean float64 = float64(sum) / 1000
        if mean < 0.45 || mean > 0.55 {
                t.Errorf("Error in mean of the bits")
        }
}

func TestDoublePowerup(t *testing.T) {
	rnd, _ := NewRandomizer(30)
	rnd.Powerup()
	err := rnd.Powerup()
	if err == nil {
		t.Fatalf("Double Powerup should crash!")
	}
}

func TestDoubleShutdown(t *testing.T) {
	rnd, _ := NewRandomizer(34)
	rnd.Powerup()
	rnd.Shutdown()
	err := rnd.Shutdown()
	if err == nil {
		t.Fatalf("Double Shutdown should crash!")
	}
}

func TestGetBitOfShutdownRandomizer(t *testing.T) {
	rnd, _ := NewRandomizer(40)
	_, err := rnd.GetBit()
	if err == nil {
		t.Fatalf("Getting bit from shutdown randomizer should crash!")
	}
}

func TestMinimumInterval(t *testing.T) {
	_, err := NewRandomizer(19)
	if err == nil {
		t.Fatalf("Minimum interval is 20 milliseconds")
	}
}
