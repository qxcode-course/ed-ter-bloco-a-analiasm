package main
import("fmt"
"math")

func main(){
var a, b, c, s, area float64;

fmt.Scanf("%f ", &a)
fmt.Scanf("%f ", &b)
fmt.Scanf("%f ", &c)

s = (a+b+c)/2
area = math.Sqrt(s * (s - a) * (s - b) * (s - c))

fmt.Printf("%.2f\n", area)

}