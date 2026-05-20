package main

import (
	"fmt"
	"math"
)
func main() {
    var a, b, c, s, area float64

    fmt.Scanf("%f %f %f ", &a, &b, &c)

    s = (a+b+c)/2
    
    area = math.Sqrt(s * (s - a) * (s - b) * (s - c))

    fmt.Printf("%.2f\n", area)
}
