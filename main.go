package main
import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func main() {

var totalSize = uintptr(104857600)
var currentSize = uintptr(0)

for currentSize < totalSize{
	rand.Seed(time.Now().UTC().UnixNano())
	c := fmt.Sprintf("%c", randInt(1, 8000))
	fmt.Printf(c)
	currentSize += unsafe.Sizeof(c)
 }
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}
