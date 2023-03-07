package test2

import (
	"fmt"
	"strconv"
)

type ThreadService interface {
	Run(numberThread int, totalThread int)
}


type Thread struct {
	totalThread int
	fin         chan bool
	thread ThreadService
}

func NewThread(totalThread int, thread ThreadService,)Thread {
	t:=Thread{}
	t.fin = make(chan bool)
	t.totalThread=totalThread
	t.thread=thread
	return t
}

func (t Thread) Start(){
	for i := 0; i < t.totalThread; i++ {
		go t.run(i, t.totalThread, t.fin)
	}
	count := 0
	for i := 0; i < t.totalThread; i++ {
		<-t.fin
		count++
		fmt.Println("threads run " + strconv.Itoa(count) + "/" + strconv.Itoa(t.totalThread))
	}
}

func (t Thread) run(numberThread int, totalThread int, fin chan bool){
		t.thread.Run(numberThread, totalThread)
		fin<-true
	
}
