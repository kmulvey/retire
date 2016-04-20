package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetire(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 8.574727266290296e+06, Retire(30, 67, 90, 100000, .0322, .1, .75), "should be ~8574727")
}

func BenchmarkRetire(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Retire(30, 67, 90, 100000, .0322, .1, .75)
	}
}

func TestFV(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 258771.02257936544, futureValue(100000, .0322, 30), "should be ~258771.02")
}

func BenchmarkFV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		futureValue(100000, .0322, 30)
	}
}

func TestPV(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 38644.20328181447, presentValue(100000, .0322, 30), "should be ~38644.20")
}

func BenchmarkPV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		presentValue(100000, .0322, 30)
	}
}
