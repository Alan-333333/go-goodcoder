package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {

	f, err := os.Create("cpu.pb.gz")
	if err != nil {
		log.Fatal("could not create cpu profile: ", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	f, err = os.Create("mem.pb.gz")
	if err != nil {
		log.Fatal("could not create mem profile: ", err)
	}
	defer f.Close()

	pprof.WriteHeapProfile(f)

	ch := make(chan int)
	for i := 0; i < 100000; i++ {
		go func() {
			var buf [1024]byte
			buf[0] = '1'
			<-ch
		}()
	}

}
