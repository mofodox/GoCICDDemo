package main

import (
	"math"
	"testing"

	. "github.com/franela/goblin"
)

func TestCalc(t *testing.T) {
	gob := Goblin(t)
	gob.Describe("Calc File", func() {
		gob.It("should add two numbers ", func() {
			gob.Assert(Add(1, 2)).Equal(3)
			gob.Assert(Add(1, 0)).Equal(1)
			gob.Assert(Add(2, -2)).Equal(0)
		})

		gob.It("should subtract two numbers", func() {
			gob.Assert(Subtract(1, 1)).Equal(0)
			gob.Assert(Subtract(10, 5)).Equal(5)
			gob.Assert(Subtract(-5, -5)).Equal(0)
		})

		gob.It("should multiply two numbers", func() {
			gob.Assert(Multiply(1, 5)).Equal(5)
			gob.Assert(Multiply(5, 5)).Equal(25)
			gob.Assert(Multiply(100, 0)).Equal(0)

		})

		gob.It("should divide two numbers", func() {
			gob.Assert(Divide(100, 5)).Equal(float64(20))
			gob.Assert(Divide(100, 1000)).Equal(float64(0.1))
			gob.Assert(Divide(100, 0)).Equal(math.Inf(1))

		})
	})
}
