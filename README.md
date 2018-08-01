# BitGenGo : Random Bit Generator 

A random bit generator written in Go, based on processor context switch

# Benchmark

The program was run against a dieharder suite and compared to the results of running dieharder 
against random bits generated from random.org, which is supposed to be truely random (as the source of randomness is atmospheric noise).
To get dieharder and run the benchmark yourself, install via: apt-get install dieharder.
The output of running ```dieharder -a -f output.txt``` where output is a file containing a stream of bits, and the machine is a VM - Intel(R) Core(TM) i5 CPU, 2.67GHz:
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
