/*
* MIT License
*
* Copyright (c) 2017 SmartestEE Inc.
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/

/*
 * Revision History:
 *     Initial: 2017/07/09        Yang Zhengtian
 */
package main

import "fmt"

func permutation(a []int, k int, m int) {
	var i int
	var j int
	if k == m {
		for i = 0; i <= m; i++ {
			fmt.Print(a[i])
		}
		fmt.Println()
	} else {
		for j = k; j <= m; j++ {
			a[j], a[k] = a[k], a[j]
			permutation(a, k+1, m)
			a[j], a[k] = a[k], a[j]
		}
	}
}

func main() {
	var N int
	fmt.Print("N=")
	fmt.Scanln(&N)
	var a []int = make([]int, N)
	for p := 0; p < N; p++ {
		a[p] = p + 1
	}
	if N < 1 {
		fmt.Println("ERROR!")
		return
	} else {
		permutation(a, 0, N-1)
	}
}
