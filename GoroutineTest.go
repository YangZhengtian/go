/*
 * Revision History:
 *     Initial: 2017/08/20        Yang Zhengtian
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Printf("%d ", n)

			wg.Done()
		}(i)
	}

	wg.Wait()
}
