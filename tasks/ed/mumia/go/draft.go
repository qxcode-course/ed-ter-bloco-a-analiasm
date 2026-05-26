package main
import "fmt"

func main() {
var pessoa string
var idade int

fmt.Scanf("%s", &pessoa)  
fmt.Scanf("%d", &idade)

if idade < 12{
   fmt.Printf("%s eh crianca\n", pessoa)
}else if idade >= 12 && idade < 18{
	fmt.Printf("%s eh jovem\n", pessoa)
}else if idade >= 18 && idade < 65{
	fmt.Printf("%s eh adulto\n", pessoa)
}else if idade >= 65 && idade < 1000{
	fmt.Printf("%s eh idoso\n", pessoa)
}else{
	fmt.Printf("%s eh mumia\n", pessoa)
}



}
