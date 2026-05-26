package main
import "fmt"
func main() {
   
   var qtdpessoas, Escolhido int
   fmt.Scan(&qtdpessoas)
   fmt.Scan(&Escolhido)
   
   vivos := []int{}
   pos := -1

   for i := 1; i <= qtdpessoas; i++{
      vivos = append(vivos, i)
      if i == i {
         if i == Escolhido{
             pos = i - 1
         }
      }
   }

   for{

      fmt.Printf("[ ")

         for i, v := range vivos{
            fmt.Print(v)
            if i == pos {
               fmt.Print(">")
                  }
                  fmt.Print(" ")
            }

         fmt.Println("]")

         if len(vivos) == 1{
            break
         }

         alvo := (pos + 1) % len(vivos)

         vivos = append(vivos[:alvo], vivos[alvo + 1:]...)

         pos = alvo % len(vivos)
      }

     }

