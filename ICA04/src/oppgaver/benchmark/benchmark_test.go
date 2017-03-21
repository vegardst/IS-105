package main

import (
	"../../oppgaver"
	"testing"
)

func Benchmark3c(b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StartTimer()
		oppgaver.Oppgave3c(1)
		b.StopTimer()

	}

}
