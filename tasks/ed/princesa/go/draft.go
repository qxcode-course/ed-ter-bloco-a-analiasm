package main
import "fmt"
func main() {
   
   var n_pessoas, E int
   fmt.Scan(&n_pessoas)
   fmt.Scan(&E)
   
   vivos := []int{}
   pos := -1

   for i := 1; i <= n_pessoas; i++{
      vivos = append(vivos, i)
      if i == i {
         if i == E{
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

