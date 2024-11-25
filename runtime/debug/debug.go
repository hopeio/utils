/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package debug

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// go tool pprof ./cpu.pprof
func CpuPprof() func() {
	f, err := os.Create("./cpu.pprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	// StartCPUProfile为当前进程开启CPU profile。
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	// StopCPUProfile会停止当前的CPU profile（如果有）
	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

func MemPprof(fun func()) func() {
	f, err := os.Create("./mem.pprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	runtime.GC()
	fun()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	return func() {
		f.Close()
	}
}

func Pprof(opt string) func() {
	f, err := os.Create("./mem.pprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	runtime.GC()

	if err := pprof.Lookup(opt).WriteTo(f, 1); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	return func() {
		f.Close()
	}
}

func PrintMemoryUsage(flag any) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v TotalAlloc = %.2f MiB,HeapAlloc = %.2f MiB,Sys = %.2f MiB,HeapSys = %.2f MiB,StackSys = %.2f MiB,HeapInuse = %.2f MiB,StackInuse = %.2f MiB,Mallocs = %d,Frees = %d,NumGC = %d", flag, bToMb(m.TotalAlloc), bToMb(m.HeapAlloc), bToMb(m.Sys), bToMb(m.HeapSys), bToMb(m.StackSys), bToMb(m.HeapInuse), bToMb(m.StackInuse), m.Mallocs, m.Frees, m.NumGC)
}

func PrintStack(flag any) {
	// 创建一个 1MB 的缓冲区来存储堆栈信息
	buf := make([]byte, 1<<20) // 1MB 缓冲区
	// 获取当前 Goroutine 的堆栈信息
	stackLen := runtime.Stack(buf, false)
	log.Printf("%v Stack:\n%s", flag, buf[:stackLen])
}

func bToMb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}
