package main
import "fmt"

const (
  Red             = (1<<iota)
  Green           = (1<<iota)
  Blue, ColorMask = (1<<iota), (1<<(iota+1))-1
)

func printf(str string, args ...interface{})(int, error) {
  _,err := fmt.Printf(str, args...)
  return len(args), err
}

func main() {
  count := 1
  closure := func(msg string) {
    printf("%d %s\n", count, msg)
    count++
  }
  closure("A Message")
  closure("Another Message")

  loops := 1
  for loops > 0 {

    printf( "\nNumber of loops?\n" )
    fmt.Scanf("%d",&loops)

    for i := 0; i < loops; i++ {
      printf("%d", i)
    }
  }

  for {
    break
  }

  for i := 0; i < 10; i++ {
    L:
    for {
      for {
        break L
      }
    }
    printf( "%d\n", i)
  }

}
