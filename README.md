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
That way, if many bits are requested at once, they will be generated in unknown and large enough intervals. 

# Usage

Get the code: ```go get github.com/maxamel/BitGenGo```

Import it: ```github.com/maxamel/BitGenGo```

Then you can start generating pseudorandom bits:

```
rnd := rand,Rand{IntervalInMillis: 23}
rnd.Powerup()
b := rnd.GetBit()
rnd.Shutdown()
```

# Measurement

There is no bullet-proof way to measure the quality of randomness being produced from a generator. However, there are tools out there to provide an insight on how random is a sequence produced by such a generator.
One of them is the Linux utility dieharder (```apt-get install dieharder```), which conducts many statistical tests and determines how random the numbers really are. BitGenGo has been tested against dieharder, and compared to results of other generators. 
For more information, read the Benchmark section. 

# Benchmark

The program was run against a dieharder suite and compared to the results of running dieharder 
against random bits generated from [random.org](https://www.random.org), which is supposed to be truely random (as the source of randomness is atmospheric noise).
This experiment was repeated mutltiple times, with similar results. Below you can find sample runs of dieharder with default parameters.
The output of running ```dieharder -a -f output.txt``` where output is a file containing a stream of bits by BitGenGo, and the machine is a VM - Intel(R) Core(TM) i5 CPU, 2.67GHz:
```
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
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
        mt19937|                      randomnums|  4.72e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.31471433|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.67553700|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.74682937|  PASSED  
    diehard_rank_6x8|   0|    100000|     100|0.93361068|  PASSED  
   diehard_bitstream|   0|   2097152|     100|0.25015633|  PASSED  
        diehard_opso|   0|   2097152|     100|0.54238539|  PASSED  
        diehard_oqso|   0|   2097152|     100|0.52163831|  PASSED  
         diehard_dna|   0|   2097152|     100|0.03203937|  PASSED  
diehard_count_1s_str|   0|    256000|     100|0.68493273|  PASSED  
diehard_count_1s_byt|   0|    256000|     100|0.13279341|  PASSED  
 diehard_parking_lot|   0|     12000|     100|0.30484250|  PASSED  
    diehard_2dsphere|   2|      8000|     100|0.70509677|  PASSED  
    diehard_3dsphere|   3|      4000|     100|0.66369793|  PASSED  
     diehard_squeeze|   0|    100000|     100|0.35308888|  PASSED  
        diehard_sums|   0|       100|     100|0.09841737|  PASSED  
        diehard_runs|   0|    100000|     100|0.23256038|  PASSED  
        diehard_runs|   0|    100000|     100|0.76579859|  PASSED  
       diehard_craps|   0|    200000|     100|0.64610285|  PASSED  
       diehard_craps|   0|    200000|     100|0.33875344|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.99178020|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.49820928|  PASSED  
         sts_monobit|   1|    100000|     100|0.42133313|  PASSED  
            sts_runs|   2|    100000|     100|0.97798267|  PASSED  
          sts_serial|   1|    100000|     100|0.87788064|  PASSED  
          sts_serial|   2|    100000|     100|0.66925144|  PASSED  
          sts_serial|   3|    100000|     100|0.88862612|  PASSED  
          sts_serial|   3|    100000|     100|0.91285016|  PASSED  
          sts_serial|   4|    100000|     100|0.92850602|  PASSED  
          sts_serial|   4|    100000|     100|0.17711415|  PASSED  
          sts_serial|   5|    100000|     100|0.99775151|   WEAK   
          sts_serial|   5|    100000|     100|0.79837990|  PASSED  
          sts_serial|   6|    100000|     100|0.43001101|  PASSED  
          sts_serial|   6|    100000|     100|0.94002566|  PASSED  
          sts_serial|   7|    100000|     100|0.82884820|  PASSED  
          sts_serial|   7|    100000|     100|0.73929878|  PASSED  
          sts_serial|   8|    100000|     100|0.58696943|  PASSED  
          sts_serial|   8|    100000|     100|0.89739927|  PASSED  
          sts_serial|   9|    100000|     100|0.13859173|  PASSED  
          sts_serial|   9|    100000|     100|0.89194067|  PASSED  
          sts_serial|  10|    100000|     100|0.95615061|  PASSED  
          sts_serial|  10|    100000|     100|0.15396427|  PASSED  
          sts_serial|  11|    100000|     100|0.64425068|  PASSED  
          sts_serial|  11|    100000|     100|0.00464903|   WEAK   
          sts_serial|  12|    100000|     100|0.72601500|  PASSED  
          sts_serial|  12|    100000|     100|0.82880717|  PASSED  
          sts_serial|  13|    100000|     100|0.38451930|  PASSED  
          sts_serial|  13|    100000|     100|0.59932375|  PASSED  
          sts_serial|  14|    100000|     100|0.93233088|  PASSED  
          sts_serial|  14|    100000|     100|0.53939127|  PASSED  
          sts_serial|  15|    100000|     100|0.54421382|  PASSED  
          sts_serial|  15|    100000|     100|0.09386158|  PASSED  
          sts_serial|  16|    100000|     100|0.67601319|  PASSED  
          sts_serial|  16|    100000|     100|0.99157901|  PASSED  
         rgb_bitdist|   1|    100000|     100|0.68831114|  PASSED  
         rgb_bitdist|   2|    100000|     100|0.29599239|  PASSED  
         rgb_bitdist|   3|    100000|     100|0.68482641|  PASSED  
         rgb_bitdist|   4|    100000|     100|0.29531438|  PASSED  
         rgb_bitdist|   5|    100000|     100|0.36199702|  PASSED  
         rgb_bitdist|   6|    100000|     100|0.85112004|  PASSED  
         rgb_bitdist|   7|    100000|     100|0.25406273|  PASSED  
         rgb_bitdist|   8|    100000|     100|0.24452604|  PASSED  
         rgb_bitdist|   9|    100000|     100|0.12204142|  PASSED  
         rgb_bitdist|  10|    100000|     100|0.95557011|  PASSED  
         rgb_bitdist|  11|    100000|     100|0.87203787|  PASSED  
         rgb_bitdist|  12|    100000|     100|0.25414607|  PASSED  
rgb_minimum_distance|   2|     10000|    1000|0.00286754|   WEAK   
rgb_minimum_distance|   3|     10000|    1000|0.98122859|  PASSED  
rgb_minimum_distance|   4|     10000|    1000|0.29966674|  PASSED  
rgb_minimum_distance|   5|     10000|    1000|0.00224758|   WEAK   
    rgb_permutations|   2|    100000|     100|0.37834478|  PASSED  
    rgb_permutations|   3|    100000|     100|0.63410655|  PASSED  
    rgb_permutations|   4|    100000|     100|0.33457846|  PASSED  
    rgb_permutations|   5|    100000|     100|0.73405058|  PASSED  
      rgb_lagged_sum|   0|   1000000|     100|0.70778973|  PASSED  
      rgb_lagged_sum|   1|   1000000|     100|0.31584662|  PASSED  
      rgb_lagged_sum|   2|   1000000|     100|0.99119929|  PASSED  
      rgb_lagged_sum|   3|   1000000|     100|0.24455526|  PASSED  
      rgb_lagged_sum|   4|   1000000|     100|0.57205313|  PASSED  
      rgb_lagged_sum|   5|   1000000|     100|0.96929233|  PASSED  
      rgb_lagged_sum|   6|   1000000|     100|0.31476916|  PASSED  
      rgb_lagged_sum|   7|   1000000|     100|0.11363654|  PASSED  
      rgb_lagged_sum|   8|   1000000|     100|0.72520171|  PASSED  
      rgb_lagged_sum|   9|   1000000|     100|0.56682518|  PASSED  
      rgb_lagged_sum|  10|   1000000|     100|0.53850942|  PASSED  
      rgb_lagged_sum|  11|   1000000|     100|0.05592550|  PASSED  
      rgb_lagged_sum|  12|   1000000|     100|0.09519905|  PASSED  
      rgb_lagged_sum|  13|   1000000|     100|0.40604743|  PASSED  
      rgb_lagged_sum|  14|   1000000|     100|0.60477106|  PASSED  
      rgb_lagged_sum|  15|   1000000|     100|0.51825597|  PASSED  
      rgb_lagged_sum|  16|   1000000|     100|0.81622391|  PASSED  
      rgb_lagged_sum|  17|   1000000|     100|0.72188616|  PASSED  
      rgb_lagged_sum|  18|   1000000|     100|0.45494400|  PASSED  
      rgb_lagged_sum|  19|   1000000|     100|0.96782984|  PASSED  
      rgb_lagged_sum|  20|   1000000|     100|0.59045790|  PASSED  
      rgb_lagged_sum|  21|   1000000|     100|0.91417744|  PASSED  
      rgb_lagged_sum|  22|   1000000|     100|0.92933198|  PASSED  
      rgb_lagged_sum|  23|   1000000|     100|0.36063798|  PASSED  
      rgb_lagged_sum|  24|   1000000|     100|0.43677304|  PASSED  
      rgb_lagged_sum|  25|   1000000|     100|0.48113980|  PASSED  
      rgb_lagged_sum|  26|   1000000|     100|0.41156181|  PASSED  
      rgb_lagged_sum|  27|   1000000|     100|0.81953642|  PASSED  
      rgb_lagged_sum|  28|   1000000|     100|0.56145543|  PASSED  
      rgb_lagged_sum|  29|   1000000|     100|0.98435656|  PASSED  
      rgb_lagged_sum|  30|   1000000|     100|0.83933294|  PASSED  
      rgb_lagged_sum|  31|   1000000|     100|0.17762105|  PASSED  
      rgb_lagged_sum|  32|   1000000|     100|0.85582191|  PASSED  
     rgb_kstest_test|   0|     10000|    1000|0.92042417|  PASSED  
     dab_bytedistrib|   0|  51200000|       1|0.69039559|  PASSED  
             dab_dct| 256|     50000|       1|0.86951789|  PASSED  
Preparing to run test 207.  ntuple = 0
        dab_filltree|  32|  15000000|       1|0.72345504|  PASSED  
        dab_filltree|  32|  15000000|       1|0.20988777|  PASSED  
Preparing to run test 208.  ntuple = 0
       dab_filltree2|   0|   5000000|       1|0.83778128|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.30110871|  PASSED  
Preparing to run test 209.  ntuple = 0
        dab_monobit2|  12|  65000000|       1|0.98332237|  PASSED 
```

As can be seen, both runs have only 2 weak tests. As the dieharder tests are statistical tests, it is quite normal for a good pseudorandom generator to produce a few weak results, and even fail a couple.
For more information regarding the significance of the results refer to the dieharder documentation. 
