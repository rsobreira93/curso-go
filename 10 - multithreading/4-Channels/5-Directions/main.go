package main

// ch chan<- string preencher canal
func recebe(name string, ch chan<- string) {
	ch <- name
}

// ch <-chan string ler canal
func ler(ch <-chan string) {
	<-ch
}

func main() {
	ch := make(chan string)
	go recebe("gopher", ch)
	ler(ch)
}
