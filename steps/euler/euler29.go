/**
  *
  * Problem 29

  Consider all integer combinations of a^b for 2 <= a <= 5 and 2 <=  b <=  5:

  2^2=4, 2^3=8, 2^4=16, 2^5=32
  3^2=9, 3^3=27, 3^4=81, 3^5=243
  4^2=16, 4^3=64, 4^4=256, 4^5=1024
  5^2=25, 5^3=125, 5^4=625, 5^5=3125

  If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

  4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

  How many distinct terms are in the sequence generated by a^b for 2 <= a <= 100 and 2 <=  b <= 100?

  Answer:

  **/

  package main

  import "math/big"
  import "fmt"
  import "time"
  
  /* import "math"*/
  
  const max = 100
  
  func main() {
    num := big.NewInt(1)
  
    t := time.Now()
    var start = t.UnixNano()
    fmt.Printf("Starting Euler #29 \n")
    elems := make(map[string]int)
  
    for a := 2; a <= max; a++ {
      for b := 2; b <= max; b++ {
        num.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
        elems[num.String()] = 1
      }
    }
  
    // count elements in num
    for k, _ := range elems {
      fmt.Println(k)
    }
  
    qed := len(elems)
  
    tt := time.Now()
    var end = tt.UnixNano()
    duration := (end - start) / 1000000
    fmt.Printf("QED: %d \n", qed)
    fmt.Printf("Calculated in %vms \n", duration)
  
  }
  