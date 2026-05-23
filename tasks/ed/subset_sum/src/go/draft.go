package main
import (
    "bufio"
    "fmt"
    "os"
    "sort")

func canFindSum(index int, target int, arr []int) bool {

    if target == 0 {
		return true
	}

    if target < 0 || index >= len(arr) {
		return false
	}

    if arr[len(arr)-1] > target {
		return false
	}

    if canFindSum(index+1, target-arr[index], arr) {
		return true
	}

    return canFindSum(index+1, target, arr)
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    var n, k int
	fmt.Fscan(reader, &n, &k)

    arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

    sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

    if canFindSum(0, k, arr) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	

}
