package main

import (
	"fmt"
	"sync"
)

func ExecuteExample1() {

	var chatChannel chan string
	var wg sync.WaitGroup

	numMsgs := 10
	chatChannel = make(chan string)

	wg.Add(2)
	go sendMessages(numMsgs, chatChannel, &wg)
	go receiveMessage(chatChannel, &wg)

	wg.Wait()

}

func sendMessages(numMsgs int, chatChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(chatChannel)

	for i := 0; i < numMsgs; i++ {
		chatChannel <- fmt.Sprintf("message %d", i)
	}

}

func receiveMessage(chatChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range chatChannel {
		fmt.Println("Message read from channel", msg)
	}
}
