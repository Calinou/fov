// fov: Calculate horizontal or vertical FOV values for a given aspect ratio
//
// Copyright © 2018-2019 Hugo Locurcio and contributors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The `fov` binary must be built before running tests

func TestOneAspectRatio(t *testing.T) {
	out, err := exec.Command("./fov", "90h", "4:3").Output()

	assert.Nil(t, err)
	assert.True(
		t,
		(strings.Contains(string(out), "90.00°") &&
			strings.Contains(string(out), "73.74°") &&
			strings.Contains(string(out), "4:3")),
		"Expected values are present in output")
}

func TestTwoAspectRatios(t *testing.T) {
	out, err := exec.Command("./fov", "90h", "4:3", "16:9").Output()

	assert.Nil(t, err)
	assert.True(
		t,
		(strings.Contains(string(out), "90.00°") &&
			strings.Contains(string(out), "73.74°") &&
			strings.Contains(string(out), "4:3") &&
			strings.Contains(string(out), "106.26°") &&
			strings.Contains(string(out), "73.74°") &&
			strings.Contains(string(out), "16:9")),
		"Expected values are present in output")
}
