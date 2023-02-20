package main
import (
  "fmt"
  "sync"
)

func main() {

  // declaring the package
  var wg sync.WaitGroup


  // looping over the jobs
  for i := 0; i <= 5; i ++ {
    // adding the job to the wait group
    wg.Add(1)

    // calling go routine so that it won't stop the execution here
    go callMe(i, &wg)
  }

  // tells the program to wait let it finish all the go routine or wg.Count
  wg.Wait()

  
  fmt.Println("execution done")

  
}

func callMe(i int, wg *sync.WaitGroup) {
  // wg.Done to mark it as its done
  defer wg.Done()
  fmt.Println("function printed", i)
}
