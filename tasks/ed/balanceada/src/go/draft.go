package main

import (
	"bufio"
	"fmt"
	"os"
)

func estaBalanceada(s string) bool{

    pares := map[rune]rune{
        ')' : '(',
        ']' : '[',
    }
    
    var pilha []rune

    for _, c := range s {

        if c == '(' || c == '['{

            pilha = append(pilha, c)

        }else if c == ')' || c == ']'{

            if len(pilha) == 0 || pilha[len(pilha)-1] != pares[c]{

                return false
            }

            pilha = pilha[:len(pilha)-1]
        }
    }

    return len(pilha) == 0
}
func main() {

    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan(){
        entrada := scanner.Text()

        if estaBalanceada(entrada){
            fmt.Println("balanceado")
        }else{
            fmt.Println("nao balanceado")
        }
    }

}