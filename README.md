# BitGenGo : Pseudorandom Bit Generator 

A simple pseudorandom bit generator written in Go, based on processor context switch

# Overview

Randomness is hard to come by. Especially on deterministic machines like computers. Generally we settle for pseudorandomness, which closely resembles randomness, but actually isn't. 
The main sources of pseudorandomness in a computer are memory and CPU. There are other methods of generating random numbers with a computer using human input, such as keystrokes, mouseclicks and mouse movements. But those methods require constant new input from a human user.
The context switch of the CPU is a good candidate for a pseudorandom generator, since it is hard to predict when the context switch will actually occur.

# How does it work?

BitGenGo is actually a small program which spawns off two goroutines(threads) and they take turns setting a shared variable with a value(zero/one).
Every time the user requests a bit, the current value of the shared variable is returned. Since a context switch takes a certain amount of time, the user must set a time threshold for retreiving a bit (default is 25 milliseconds). 
This means a user can request the next bit only after the threshold time has passed since the retrieval of the previous bit. This ensures the shared variable will switch values several times before each request.
It is recommended the user set his own threshold for security reasons (minimum of 20 milliseconds is advised). 
That way, if many bits are requested at once, they will be generated in unknown(to the attacker) and large enough intervals. 

# Usage

Get the code:
 
```go get github.com/maxamel/BitGenGo```

Import it:
 
```import "github.com/maxamel/BitGenGo"```

Then you can start generating pseudorandom bits:

```
rnd := rand.NewRandomizer(23)
rnd.Powerup()
b := rnd.GetBit()
rnd.Shutdown()
// b contains a pseudorandom bit
```

# Randomness Quality

There is no bullet-proof way to measure the quality of randomness being produced from a generator. However, there are tools out there to provide an insight on how random is a sequence produced by such a generator.
One of them is the Linux utility dieharder (```apt-get install dieharder```), which conducts many statistical tests and determines how random the numbers really are. BitGenGo has been tested against dieharder, and compared to results of other generators. 
For more information, read the Benchmark section. 

# Benchmark

The program was run against a dieharder suite and compared to the results of running dieharder 
against random bits generated from [random.org](https://www.random.org), which is supposed to be truely random (as the source of randomness is atmospheric noise).
This experiment was repeated mutltiple times, with similar results. Below you can find sample runs of dieharder with default parameters.
The output of running ```dieharder -a -f output.txt``` where output is a file containing a stream of bits by BitGenGo, and the machine is a VM - Intel(R) Core(TM) i5 CPU, 2.67GHz:
```
0,0,1,0,0,0,1,0,1,1,0,1,0,0,1,0,1,0,0,0,1,1,1,0,1,1,0,0,1,0,1,0,1,1,1,1,1,1,0,0,0,0,1,0,1,1,1,0,0,1,
0,0,0,1,0,0,1,0,0,0,1,1,1,0,0,0,1,1,0,0,0,0,0,0,1,1,1,0,1,0,0,1,0,1,0,0,1,1,0,0,0,0,1,0,1,1,1,1,1,1,
1,1,0,0,1,1,0,0,1,0,1,0,1,0,1,1,1,1,0,1,1,1,1,1,1,1,1,1,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,1,1,0,0,
1,0,1,1,0,0,1,1,1,0,1,0,0,1,1,1,1,1,1,0,0,1,1,1,0,0,0,0,0,1,1,1,1,1,1,1,0,1,0,1,1,1,0,1,0,1,0,0,1,0,
0,0,0,1,0,0,0,1,0,1,1,0,0,1,1,1,1,0,1,0,0,1,0,1,0,0,1,1,1,1,0,0,0,0,0,0,1,1,0,1,1,1,1,0,1,0,0,1,0,0,
0,1,0,0,0,0,1,1,0,0,0,1,1,0,1,0,0,0,1,1,0,1,0,0,0,1,1,1,0,1,1,1,1,0,1,1,0,0,1,1,1,0,1,1,1,1,0,0,1,0,
0,0,1,1,0,0,0,0,1,1,0,1,0,1,1,0,0,1,1,0,1,0,1,0,1,0,1,1,1,0,0,0,1,1,0,1,0,0,0,0,1,1,0,1,0,0,0,1,1,0,
0,0,1,0,0,0,1,0,1,1,1,0,0,1,1,0,0,0,1,0,1,1,1,0,0,0,0,0,0,1,0,0,0,0,1,1,1,0,0,1,0,1,1,0,1,0,1,0,0,1,
1,1,0,1,0,1,1,0,0,1,1,0,0,0,0,0,1,0,0,1,0,1,0,1,1,0,1,0,1,1,1,1,1,1,1,1,1,0,1,1,1,1,1,1,1,1,0,0,1,1,
1,0,0,0,1,1,1,1,1,1,0,0,1,1,1,0,1,1,0,1,1,1,1,0,0,1,1,1,1,1,1,1,1,1,1,1,0,1,0,0,1,1,0,0,0,1,1,0,0,1,
0,0,1,0,1,0,0,1,1,0,1,1,0,1,0,0,0,0,0,1,0,0,0,1,0,0,0,1,1,1,1,1,1,1,0,0,0,1,0,1,1,0,0,1,1,1,1,0,0,0,
1,1,1,1,0,0,0,1,1,0,1,1,1,0,1,1,0,1,0,0,1,0,0,0,1,0,1,1,1,0,0,0,0,0,1,1,1,1,0,0,1,0,1,1,1,1,1,1,0,0,
1,0,0,0,0,1,0,0,0,0,1,1,0,0,1,1,0,0,1,0,0,0,0,0,1,0,0,0,1,0,0,0,0,1,0,1,0,1,1,0,1,0,1,1,0,0,1,1,0,0,
1,1,0,1,1,0,0,1,1,0,1,1,1,0,0,1,1,1,1,1,0,1,0,1,0,0,0,0,1,0,0,0,1,1,0,0,0,0,0,1,0,0,0,1,1,0,1,0,0,0,
0,1,1,1,1,0,0,0,0,1,1,1,1,1,0,0,1,0,0,0,0,1,0,1,0,0,0,1,1,0,1,0,0,1,0,0,0,1,0,0,0,0,0,0,1,0,1,0,0,1,
0,1,1,0,1,0,1,0,1,0,1,0,0,0,0,0,1,1,1,0,1,0,0,1,1,0,1,0,1,1,1,0,1,0,1,0,1,0,0,0,0,0,0,1,0,1,0,1,0,0,
0,0,0,1,0,0,1,0,0,0,1,1,1,0,0,0,1,0,0,0,0,0,1,0,1,1,1,0,0,1,1,1,1,1,1,0,0,0,0,0,1,0,1,1,1,0,0,1,0,1,
0,0,0,1,1,0,0,0,1,0,0,0,0,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,0,0,0,1,1,1,1,1,0,0,1,0,0,1,1,0,0,0,0,0,1,
0,0,1,1,0,0,0,0,0,0,0,0,1,1,1,1,0,0,0,0,1,0,1,1,0,1,0,0,1,0,0,0,0,1,0,0,1,0,1,0,0,0,0,0,0,1,1,0,1,0,
0,1,1,1,0,0,1,0,0,0,1,1,0,0,0,0,1,0,1,0,0,1,1,1,0,0,0,1,1,0,1,1,0,1,0,1,0,0,1,1,1,0,1,1,0,0,1,0,1,0

#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
        mt19937|                        randomgo|  3.97e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.14645319|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.58258306|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.26509050|  PASSED  
    diehard_rank_6x8|   0|    100000|     100|0.88581102|  PASSED  
   diehard_bitstream|   0|   2097152|     100|0.84882352|  PASSED  
        diehard_opso|   0|   2097152|     100|0.16036919|  PASSED  
        diehard_oqso|   0|   2097152|     100|0.35491955|  PASSED  
         diehard_dna|   0|   2097152|     100|0.06265474|  PASSED  
diehard_count_1s_str|   0|    256000|     100|0.02711762|  PASSED  
diehard_count_1s_byt|   0|    256000|     100|0.40173057|  PASSED  
 diehard_parking_lot|   0|     12000|     100|0.43824831|  PASSED  
    diehard_2dsphere|   2|      8000|     100|0.78779991|  PASSED  
    diehard_3dsphere|   3|      4000|     100|0.03947804|  PASSED  
     diehard_squeeze|   0|    100000|     100|0.48464408|  PASSED  
        diehard_sums|   0|       100|     100|0.88161131|  PASSED  
        diehard_runs|   0|    100000|     100|0.62935448|  PASSED  
        diehard_runs|   0|    100000|     100|0.15721506|  PASSED  
       diehard_craps|   0|    200000|     100|0.21382755|  PASSED  
       diehard_craps|   0|    200000|     100|0.89466284|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.44160385|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.94716786|  PASSED  
         sts_monobit|   1|    100000|     100|0.51016024|  PASSED  
            sts_runs|   2|    100000|     100|0.94200664|  PASSED  
          sts_serial|   1|    100000|     100|0.31836211|  PASSED  
          sts_serial|   2|    100000|     100|0.48766889|  PASSED  
          sts_serial|   3|    100000|     100|0.85996072|  PASSED  
          sts_serial|   3|    100000|     100|0.17974708|  PASSED  
          sts_serial|   4|    100000|     100|0.97979360|  PASSED  
          sts_serial|   4|    100000|     100|0.74230018|  PASSED  
          sts_serial|   5|    100000|     100|0.08834865|  PASSED  
          sts_serial|   5|    100000|     100|0.12955736|  PASSED  
          sts_serial|   6|    100000|     100|0.09682240|  PASSED  
          sts_serial|   6|    100000|     100|0.76362183|  PASSED  
          sts_serial|   7|    100000|     100|0.18370740|  PASSED  
          sts_serial|   7|    100000|     100|0.69436163|  PASSED  
          sts_serial|   8|    100000|     100|0.39994848|  PASSED  
          sts_serial|   8|    100000|     100|0.35910306|  PASSED  
          sts_serial|   9|    100000|     100|0.92711207|  PASSED  
          sts_serial|   9|    100000|     100|0.44987712|  PASSED  
          sts_serial|  10|    100000|     100|0.12048954|  PASSED  
          sts_serial|  10|    100000|     100|0.04354006|  PASSED  
          sts_serial|  11|    100000|     100|0.00583977|  PASSED  
          sts_serial|  11|    100000|     100|0.91567852|  PASSED  
          sts_serial|  12|    100000|     100|0.73583989|  PASSED  
          sts_serial|  12|    100000|     100|0.34358994|  PASSED  
          sts_serial|  13|    100000|     100|0.01875896|  PASSED  
          sts_serial|  13|    100000|     100|0.00180133|   WEAK   
          sts_serial|  14|    100000|     100|0.58391481|  PASSED  
          sts_serial|  14|    100000|     100|0.71582112|  PASSED  
          sts_serial|  15|    100000|     100|0.33275385|  PASSED  
          sts_serial|  15|    100000|     100|0.62331385|  PASSED  
          sts_serial|  16|    100000|     100|0.42103242|  PASSED  
          sts_serial|  16|    100000|     100|0.79005509|  PASSED  
         rgb_bitdist|   1|    100000|     100|0.80515041|  PASSED  
         rgb_bitdist|   2|    100000|     100|0.41057991|  PASSED  
         rgb_bitdist|   3|    100000|     100|0.05407465|  PASSED  
         rgb_bitdist|   4|    100000|     100|0.12503986|  PASSED  
         rgb_bitdist|   5|    100000|     100|0.71355278|  PASSED  
         rgb_bitdist|   6|    100000|     100|0.45489383|  PASSED  
         rgb_bitdist|   7|    100000|     100|0.89858801|  PASSED  
         rgb_bitdist|   8|    100000|     100|0.24359458|  PASSED  
         rgb_bitdist|   9|    100000|     100|0.65207382|  PASSED  
         rgb_bitdist|  10|    100000|     100|0.90008269|  PASSED  
         rgb_bitdist|  11|    100000|     100|0.76669106|  PASSED  
         rgb_bitdist|  12|    100000|     100|0.92969446|  PASSED  
rgb_minimum_distance|   2|     10000|    1000|0.34930473|  PASSED  
rgb_minimum_distance|   3|     10000|    1000|0.75643402|  PASSED  
rgb_minimum_distance|   4|     10000|    1000|0.02216321|  PASSED  
rgb_minimum_distance|   5|     10000|    1000|0.84796766|  PASSED  
    rgb_permutations|   2|    100000|     100|0.46102437|  PASSED  
    rgb_permutations|   3|    100000|     100|0.91557603|  PASSED  
    rgb_permutations|   4|    100000|     100|0.41965353|  PASSED  
    rgb_permutations|   5|    100000|     100|0.18085561|  PASSED  
      rgb_lagged_sum|   0|   1000000|     100|0.98983058|  PASSED  
      rgb_lagged_sum|   1|   1000000|     100|0.98148563|  PASSED  
      rgb_lagged_sum|   2|   1000000|     100|0.94315273|  PASSED  
      rgb_lagged_sum|   3|   1000000|     100|0.96660101|  PASSED  
      rgb_lagged_sum|   4|   1000000|     100|0.01520343|  PASSED  
      rgb_lagged_sum|   5|   1000000|     100|0.84046122|  PASSED  
      rgb_lagged_sum|   6|   1000000|     100|0.39765759|  PASSED  
      rgb_lagged_sum|   7|   1000000|     100|0.78608922|  PASSED  
      rgb_lagged_sum|   8|   1000000|     100|0.51781081|  PASSED  
      rgb_lagged_sum|   9|   1000000|     100|0.87443790|  PASSED  
      rgb_lagged_sum|  10|   1000000|     100|0.69314392|  PASSED  
      rgb_lagged_sum|  11|   1000000|     100|0.86979691|  PASSED  
      rgb_lagged_sum|  12|   1000000|     100|0.98641683|  PASSED  
      rgb_lagged_sum|  13|   1000000|     100|0.73669787|  PASSED  
      rgb_lagged_sum|  14|   1000000|     100|0.36742420|  PASSED  
      rgb_lagged_sum|  15|   1000000|     100|0.04380533|  PASSED  
      rgb_lagged_sum|  16|   1000000|     100|0.78765231|  PASSED  
      rgb_lagged_sum|  17|   1000000|     100|0.64937734|  PASSED  
      rgb_lagged_sum|  18|   1000000|     100|0.16596301|  PASSED  
      rgb_lagged_sum|  19|   1000000|     100|0.90832752|  PASSED  
      rgb_lagged_sum|  20|   1000000|     100|0.89498862|  PASSED  
      rgb_lagged_sum|  21|   1000000|     100|0.77707384|  PASSED  
      rgb_lagged_sum|  22|   1000000|     100|0.93530502|  PASSED  
      rgb_lagged_sum|  23|   1000000|     100|0.53756476|  PASSED  
      rgb_lagged_sum|  24|   1000000|     100|0.99995424|   WEAK   
      rgb_lagged_sum|  25|   1000000|     100|0.35188642|  PASSED  
      rgb_lagged_sum|  26|   1000000|     100|0.42811134|  PASSED  
      rgb_lagged_sum|  27|   1000000|     100|0.05295153|  PASSED  
      rgb_lagged_sum|  28|   1000000|     100|0.41155895|  PASSED  
      rgb_lagged_sum|  29|   1000000|     100|0.05748314|  PASSED  
      rgb_lagged_sum|  30|   1000000|     100|0.97335320|  PASSED  
      rgb_lagged_sum|  31|   1000000|     100|0.54750045|  PASSED  
      rgb_lagged_sum|  32|   1000000|     100|0.01535850|  PASSED  
     rgb_kstest_test|   0|     10000|    1000|0.53113556|  PASSED  
     dab_bytedistrib|   0|  51200000|       1|0.69810655|  PASSED  
             dab_dct| 256|     50000|       1|0.77418164|  PASSED  
Preparing to run test 207.  ntuple = 0
        dab_filltree|  32|  15000000|       1|0.74650143|  PASSED  
        dab_filltree|  32|  15000000|       1|0.32808436|  PASSED  
Preparing to run test 208.  ntuple = 0
       dab_filltree2|   0|   5000000|       1|0.68034268|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.96396563|  PASSED  
Preparing to run test 209.  ntuple = 0
        dab_monobit2|  12|  65000000|       1|0.08748562|  PASSED 
```

Running dieharder against a random.org generated sequence of bits:
```
0,0,0,1,0,1,0,0,0,0,0,0,0,0,1,0,1,0,0,0,1,0,0,1,1,1,0,1,1,0,0,0,1,0,1,1,0,0,0,0,1,1,0,0,0,1,0,0,0,1,
1,0,1,0,0,1,0,1,0,0,0,0,1,1,0,1,0,1,1,1,0,0,1,1,1,0,1,0,1,0,0,0,0,1,1,1,0,1,1,0,0,1,0,1,1,0,1,1,1,1,
1,0,1,1,0,0,0,1,1,0,1,0,1,1,1,0,0,1,1,1,0,1,1,1,0,0,0,0,0,1,1,0,1,0,1,0,1,1,0,0,1,0,0,0,0,1,0,0,0,0,
1,1,0,0,1,0,1,0,0,0,0,0,0,0,1,0,1,0,0,1,1,0,1,1,1,0,0,0,0,0,1,0,0,0,1,1,1,0,0,1,1,1,0,1,0,0,0,1,1,1,
1,0,1,1,0,0,0,0,1,0,1,1,1,0,0,0,1,1,0,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,0,0,0,0,1,1,0,1,0,1,0,1,0,1,1,0,
0,0,1,0,1,0,1,0,1,1,0,0,0,0,0,0,0,0,0,0,1,0,1,0,0,1,1,1,0,0,0,1,0,0,1,1,0,1,0,0,1,1,1,1,0,1,1,0,0,0,
0,1,1,0,0,1,1,1,0,1,1,1,1,0,1,0,1,0,1,0,0,0,1,1,1,1,1,0,1,1,1,1,1,1,1,1,0,1,0,0,1,1,1,0,0,1,1,1,0,1,
1,0,0,1,1,1,1,1,1,1,1,1,1,0,1,0,0,0,0,1,1,0,1,1,0,1,0,1,1,1,1,0,0,1,0,0,0,0,1,0,1,1,0,1,1,0,0,0,1,0,
0,0,0,0,1,0,1,1,1,1,0,0,1,1,0,0,0,0,0,1,1,0,0,0,1,1,1,1,0,0,0,0,1,1,1,0,1,1,1,1,1,0,1,0,1,1,1,0,1,1,
0,1,0,1,0,0,1,0,1,0,0,0,1,1,0,1,0,0,1,0,0,0,0,0,0,1,1,1,1,0,0,1,1,0,1,1,0,0,0,0,0,0,0,0,1,0,0,0,0,1,
0,0,1,0,0,1,0,0,0,0,1,0,1,0,1,1,1,1,0,0,1,1,1,1,1,1,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,1,1,1,1,0,1,1,1,0,
0,0,1,1,1,1,0,1,1,1,1,1,1,1,0,1,0,1,1,1,0,0,1,1,1,1,1,1,0,0,1,1,1,0,1,0,0,1,1,1,0,0,0,0,0,0,0,0,0,0,
0,0,1,1,1,0,1,1,0,1,0,1,1,1,1,1,1,0,1,0,0,0,1,0,0,1,0,0,0,0,1,1,0,0,0,1,0,0,1,0,1,1,0,0,0,0,0,1,1,1,
1,0,0,1,1,1,1,1,0,1,1,0,0,1,0,1,1,0,1,0,1,0,0,1,1,1,1,1,1,1,1,1,1,1,0,0,1,1,0,0,0,1,0,1,0,1,1,0,0,0,
0,0,1,1,0,1,0,1,0,1,0,1,0,0,1,1,0,0,1,1,1,1,0,1,1,1,1,1,1,1,0,0,0,0,1,1,0,1,0,0,0,0,0,0,0,1,1,1,0,1,
0,1,1,0,1,1,0,1,0,1,1,0,1,0,1,0,1,1,0,1,1,1,1,0,1,0,1,0,0,1,1,0,0,0,1,1,0,1,1,0,0,0,0,0,0,1,1,0,0,0,
0,0,0,0,0,1,0,1,1,1,0,0,0,1,0,1,1,1,0,1,1,1,0,1,0,1,0,0,0,1,1,1,1,0,1,0,0,0,1,1,1,1,1,1,0,1,0,1,0,1,
1,1,1,0,1,1,0,1,0,1,1,1,0,1,1,0,0,1,0,1,1,0,1,1,0,1,0,1,0,0,1,0,1,0,1,0,0,0,1,0,1,0,1,0,1,0,0,0,1,0,
0,1,1,0,0,0,0,1,1,1,1,0,0,1,1,1,1,0,1,1,0,0,0,0,1,0,0,1,0,1,1,0,1,1,1,1,0,1,0,1,1,1,0,1,1,1,1,0,1,0,
0,1,1,0,0,0,0,0,1,1,0,0,1,1,1,1,0,1,1,1,1,1,1,0,0,0,1,1,1,0,0,0,1,1,0,0,0,0,0,1,0,0,1,1,1,1,1,1,0,0

#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
        mt19937|                      randomnums|  1.69e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.76021038|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.87789200|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.93188811|  PASSED  
    diehard_rank_6x8|   0|    100000|     100|0.97648642|  PASSED  
   diehard_bitstream|   0|   2097152|     100|0.91484347|  PASSED  
        diehard_opso|   0|   2097152|     100|0.91264493|  PASSED  
        diehard_oqso|   0|   2097152|     100|0.40305522|  PASSED  
         diehard_dna|   0|   2097152|     100|0.75510708|  PASSED  
diehard_count_1s_str|   0|    256000|     100|0.34667112|  PASSED  
diehard_count_1s_byt|   0|    256000|     100|0.80613635|  PASSED  
 diehard_parking_lot|   0|     12000|     100|0.16555261|  PASSED  
    diehard_2dsphere|   2|      8000|     100|0.37350508|  PASSED  
    diehard_3dsphere|   3|      4000|     100|0.19300114|  PASSED  
     diehard_squeeze|   0|    100000|     100|0.50776984|  PASSED  
        diehard_sums|   0|       100|     100|0.72990953|  PASSED  
        diehard_runs|   0|    100000|     100|0.12719027|  PASSED  
        diehard_runs|   0|    100000|     100|0.13112442|  PASSED  
       diehard_craps|   0|    200000|     100|0.62592003|  PASSED  
       diehard_craps|   0|    200000|     100|0.85143334|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.35118735|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.52471592|  PASSED  
         sts_monobit|   1|    100000|     100|0.90128514|  PASSED  
            sts_runs|   2|    100000|     100|0.69157545|  PASSED  
          sts_serial|   1|    100000|     100|0.07029455|  PASSED  
          sts_serial|   2|    100000|     100|0.58345893|  PASSED  
          sts_serial|   3|    100000|     100|0.01658495|  PASSED  
          sts_serial|   3|    100000|     100|0.38747318|  PASSED  
          sts_serial|   4|    100000|     100|0.37470336|  PASSED  
          sts_serial|   4|    100000|     100|0.85738119|  PASSED  
          sts_serial|   5|    100000|     100|0.20124750|  PASSED  
          sts_serial|   5|    100000|     100|0.53397791|  PASSED  
          sts_serial|   6|    100000|     100|0.99855222|   WEAK   
          sts_serial|   6|    100000|     100|0.56569057|  PASSED  
          sts_serial|   7|    100000|     100|0.85635292|  PASSED  
          sts_serial|   7|    100000|     100|0.52623501|  PASSED  
          sts_serial|   8|    100000|     100|0.93016566|  PASSED  
          sts_serial|   8|    100000|     100|0.57008488|  PASSED  
          sts_serial|   9|    100000|     100|0.99953071|   WEAK   
          sts_serial|   9|    100000|     100|0.77072022|  PASSED  
          sts_serial|  10|    100000|     100|0.89889022|  PASSED  
          sts_serial|  10|    100000|     100|0.20196711|  PASSED  
          sts_serial|  11|    100000|     100|0.35271495|  PASSED  
          sts_serial|  11|    100000|     100|0.32620586|  PASSED  
          sts_serial|  12|    100000|     100|0.45588983|  PASSED  
          sts_serial|  12|    100000|     100|0.83232332|  PASSED  
          sts_serial|  13|    100000|     100|0.47234473|  PASSED  
          sts_serial|  13|    100000|     100|0.62397219|  PASSED  
          sts_serial|  14|    100000|     100|0.47925238|  PASSED  
          sts_serial|  14|    100000|     100|0.17785689|  PASSED  
          sts_serial|  15|    100000|     100|0.84423576|  PASSED  
          sts_serial|  15|    100000|     100|0.78887692|  PASSED  
          sts_serial|  16|    100000|     100|0.23324668|  PASSED  
          sts_serial|  16|    100000|     100|0.13210633|  PASSED  
         rgb_bitdist|   1|    100000|     100|0.90902857|  PASSED  
         rgb_bitdist|   2|    100000|     100|0.48849832|  PASSED  
         rgb_bitdist|   3|    100000|     100|0.50252503|  PASSED  
         rgb_bitdist|   4|    100000|     100|0.64082451|  PASSED  
         rgb_bitdist|   5|    100000|     100|0.99894802|   WEAK   
         rgb_bitdist|   6|    100000|     100|0.59094661|  PASSED  
         rgb_bitdist|   7|    100000|     100|0.90737236|  PASSED  
         rgb_bitdist|   8|    100000|     100|0.60371197|  PASSED  
         rgb_bitdist|   9|    100000|     100|0.43716740|  PASSED  
         rgb_bitdist|  10|    100000|     100|0.67061561|  PASSED  
         rgb_bitdist|  11|    100000|     100|0.12145088|  PASSED  
         rgb_bitdist|  12|    100000|     100|0.79585037|  PASSED  
rgb_minimum_distance|   2|     10000|    1000|0.85779718|  PASSED  
rgb_minimum_distance|   3|     10000|    1000|0.31981110|  PASSED  
rgb_minimum_distance|   4|     10000|    1000|0.94245513|  PASSED  
rgb_minimum_distance|   5|     10000|    1000|0.20227614|  PASSED  
    rgb_permutations|   2|    100000|     100|0.19543365|  PASSED  
    rgb_permutations|   3|    100000|     100|0.47565995|  PASSED  
    rgb_permutations|   4|    100000|     100|0.97073560|  PASSED  
    rgb_permutations|   5|    100000|     100|0.82556367|  PASSED  
      rgb_lagged_sum|   0|   1000000|     100|0.10798801|  PASSED  
      rgb_lagged_sum|   1|   1000000|     100|0.76569848|  PASSED  
      rgb_lagged_sum|   2|   1000000|     100|0.87782267|  PASSED  
      rgb_lagged_sum|   3|   1000000|     100|0.92803995|  PASSED  
      rgb_lagged_sum|   4|   1000000|     100|0.68954243|  PASSED  
      rgb_lagged_sum|   5|   1000000|     100|0.68294065|  PASSED  
      rgb_lagged_sum|   6|   1000000|     100|0.70084783|  PASSED  
      rgb_lagged_sum|   7|   1000000|     100|0.32059376|  PASSED  
      rgb_lagged_sum|   8|   1000000|     100|0.83511698|  PASSED  
      rgb_lagged_sum|   9|   1000000|     100|0.76500218|  PASSED  
      rgb_lagged_sum|  10|   1000000|     100|0.42151679|  PASSED  
      rgb_lagged_sum|  11|   1000000|     100|0.99980510|   WEAK   
      rgb_lagged_sum|  12|   1000000|     100|0.75108028|  PASSED  
      rgb_lagged_sum|  13|   1000000|     100|0.25371872|  PASSED  
      rgb_lagged_sum|  14|   1000000|     100|0.94110845|  PASSED  
      rgb_lagged_sum|  15|   1000000|     100|0.42336887|  PASSED  
      rgb_lagged_sum|  16|   1000000|     100|0.95470641|  PASSED  
      rgb_lagged_sum|  17|   1000000|     100|0.95464046|  PASSED  
      rgb_lagged_sum|  18|   1000000|     100|0.27113912|  PASSED  
      rgb_lagged_sum|  19|   1000000|     100|0.77094179|  PASSED  
      rgb_lagged_sum|  20|   1000000|     100|0.30122851|  PASSED  
      rgb_lagged_sum|  21|   1000000|     100|0.86826027|  PASSED  
      rgb_lagged_sum|  22|   1000000|     100|0.95232510|  PASSED  
      rgb_lagged_sum|  23|   1000000|     100|0.97846721|  PASSED  
      rgb_lagged_sum|  24|   1000000|     100|0.96943131|  PASSED  
      rgb_lagged_sum|  25|   1000000|     100|0.43883895|  PASSED  
      rgb_lagged_sum|  26|   1000000|     100|0.31614349|  PASSED  
      rgb_lagged_sum|  27|   1000000|     100|0.38367613|  PASSED  
      rgb_lagged_sum|  28|   1000000|     100|0.38993935|  PASSED  
      rgb_lagged_sum|  29|   1000000|     100|0.11205049|  PASSED  
      rgb_lagged_sum|  30|   1000000|     100|0.12733273|  PASSED  
      rgb_lagged_sum|  31|   1000000|     100|0.62217458|  PASSED  
      rgb_lagged_sum|  32|   1000000|     100|0.22435488|  PASSED  
     rgb_kstest_test|   0|     10000|    1000|0.09361493|  PASSED  
     dab_bytedistrib|   0|  51200000|       1|0.71341985|  PASSED  
             dab_dct| 256|     50000|       1|0.08468585|  PASSED  
Preparing to run test 207.  ntuple = 0
        dab_filltree|  32|  15000000|       1|0.63628160|  PASSED  
        dab_filltree|  32|  15000000|       1|0.01162037|  PASSED  
Preparing to run test 208.  ntuple = 0
       dab_filltree2|   0|   5000000|       1|0.63080808|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.34832488|  PASSED  
Preparing to run test 209.  ntuple = 0
        dab_monobit2|  12|  65000000|       1|0.88664311|  PASSED  
```

As can be seen, both runs have only a small amount of weak tests. As the dieharder tests are statistical tests, it is quite normal for a good pseudorandom generator to produce a few weak results, and even fail a couple.
For more information regarding the significance of the results refer to the dieharder documentation (```man dieharder```). 
