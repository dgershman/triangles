package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestLegA(t *testing.T) {
	g := Goblin(t)
	g.Describe("Get Leg A", func() {
		g.It("Should equal mathematical value", func() {
			g.Assert(Float64(LegA(18, 96)).Round(.1, 1)).Equal(94.3)
		})
	})
}

func TestLegB(t *testing.T) {
	g := Goblin(t)
	g.Describe("Get Leg B", func() {
		g.It("Should equal mathematical value", func() {
			g.Assert(Float64(LegB(94.3, 96)).Round(.1, 1)).Equal(18.0)
		})
	})
}

func TestLegC(t *testing.T) {
	g := Goblin(t)
	g.Describe("Get Leg B", func() {
		g.It("Should equal mathematical value", func() {
			g.Assert(Float64(LegC(18, 94.3)).Round(.1, 1)).Equal(96.0)
		})
	})
}

func TestAngleA(t *testing.T) {
	g := Goblin(t)
	g.Describe("Get Angle A", func() {
		g.It("Should equal mathematical value", func() {
			g.Assert(Float64(AngleA(94.3, 96)).Round(.1, 1)).Equal(10.8)
		})
	})
}

func TestAngleB(t *testing.T) {
	g := Goblin(t)
	g.Describe("Get Angle B", func() {
		g.It("Should equal mathematical value", func() {
			g.Assert(Float64(AngleB()).Round(.1, 1)).Equal(90.0)
		})
	})
}

func TestAngleC(t *testing.T) {
	g := Goblin(t)
	g.Describe("Get Angle C", func() {
		g.It("Should equal mathematical value", func() {
			g.Assert(Float64(AngleC(18, 96)).Round(.1, 1)).Equal(79.2)
		})
	})
}