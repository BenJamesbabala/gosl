// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rnd

import (
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/utl"
)

func Test_int01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("int01. integers")

	Init(1234)

	nints := 10
	irange := utl.IntRange(nints) // integers; e.g. 0,1,2,3,4,5,6,7,8,9
	ifreqs := make([]int, nints)  // frequencies of each integer

	labels := make([]string, nints)
	for i := 0; i < nints; i++ {
		labels[i] = io.Sf("%3d", irange[i])
	}

	nsamples := 1000
	t0 := time.Now()
	for i := 0; i < nsamples; i++ {
		gen := IntRand(0, nints-1)
		for j, val := range irange {
			if gen == val {
				ifreqs[j]++
				break
			}
		}
	}
	io.Pforan("time elapsed = %v\n", time.Now().Sub(t0))

	io.Pf(TextHist(labels, ifreqs, 60))
}

func Test_int02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("int02. integers")

	Init(1234)

	nints := 10
	irange := utl.IntRange(nints) // integers; e.g. 0,1,2,3,4,5,6,7,8,9
	ifreqs := make([]int, nints)  // frequencies of each integer

	labels := make([]string, nints)
	for i := 0; i < nints; i++ {
		labels[i] = io.Sf("%3d", irange[i])
	}

	t0 := time.Now()
	nsamples := 1000
	samples := make([]int, nsamples)
	IntRands(samples, 0, nints-1)
	//io.Pforan("samples = %v\n", samples)
	for i := 0; i < nsamples; i++ {
		for j, val := range irange {
			if samples[i] == val {
				ifreqs[j]++
				break
			}
		}
	}
	io.Pforan("time elapsed = %v\n", time.Now().Sub(t0))

	io.Pf(TextHist(labels, ifreqs, 60))
}

func Test_flt01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("flt01. float64")

	Init(1234)

	xmin := 10.0
	xmax := 20.0

	nsamples := 10
	t0 := time.Now()
	for i := 0; i < nsamples; i++ {
		gen := DblRand(xmin, xmax)
		io.Pforan("gen = %v\n", gen)
	}
	io.Pforan("time elapsed = %v\n", time.Now().Sub(t0))

}