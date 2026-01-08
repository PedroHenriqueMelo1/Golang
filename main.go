package main

import ( // import Key reservada para fazer importações de pacote com () conseguimos importar multiplos pacotes
    "fmt" // fmt, pacote com um conjunto de funções designado para imprimir coisas no console
    steph "strings" // pacote nativo de tratamento de strings, podemos alterar também o nome do pacote de import colocando o nome desejavel na mesma linha e em seguida o nome do pacote
)

func main() { // função main 
  fmt.Println("Hello, World") // Printando Hello World
    
  fmt.Println(steph.Split("olá mundo", "")) // Tratando string, e printando.
}
