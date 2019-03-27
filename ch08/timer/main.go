package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {

	i := 10

	go (func(i *int) {
		r := bufio.NewReader(os.Stdin)
		for {
			_, err := r.ReadString('\n')
			if err != nil {
				break
			}
			*i = 10
		}
	})(&i)

	for ; i > 0; i-- {
		log.Println(i)
		time.Sleep(1 * time.Second)
	}

	log.Println("Time is up")
}
