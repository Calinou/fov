// fov: Calculate horizontal or vertical FOV values for a given aspect ratio
//
// Copyright © 2018-2020 Hugo Locurcio and contributors
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
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// The `fov` binary must be built before running tests

type TestSuite struct {
	suite.Suite
	execPath string // Path to the binary (different on Windows)
}

// Entry point function for all tests
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	if runtime.GOOS == "windows" {
		suite.execPath = "fov.exe"
	} else {
		suite.execPath = "./fov"
	}
}

func (suite *TestSuite) TestOneAspectRatioHorizontal() {
	out, err := exec.Command(suite.execPath, "90h", "4:3").Output()

	assert.Nil(suite.T(), err)
	assert.True(
		suite.T(),
		(strings.Contains(string(out), "90.00°") &&
			strings.Contains(string(out), "73.74°") &&
			strings.Contains(string(out), "4:3")),
		"Expected values are present in output")
}

func (suite *TestSuite) TestOneAspectRatioVertical() {
	out, err := exec.Command(suite.execPath, "70v", "4:3").Output()

	assert.Nil(suite.T(), err)
	assert.True(
		suite.T(),
		(strings.Contains(string(out), "86.07°") &&
			strings.Contains(string(out), "70.00°") &&
			strings.Contains(string(out), "4:3")),
		"Expected values are present in output")
}

func (suite *TestSuite) TestTwoAspectRatiosHorizontal() {
	out, err := exec.Command(suite.execPath, "90h", "4:3", "16:9").Output()

	assert.Nil(suite.T(), err)
	assert.True(
		suite.T(),
		(strings.Contains(string(out), "90.00°") &&
			strings.Contains(string(out), "73.74°") &&
			strings.Contains(string(out), "4:3") &&
			strings.Contains(string(out), "106.26°") &&
			strings.Contains(string(out), "73.74°") &&
			strings.Contains(string(out), "16:9")),
		"Expected values are present in output")
}

func (suite *TestSuite) TestTwoAspectRatiosVertical() {
	out, err := exec.Command(suite.execPath, "70v", "4:3", "16:9").Output()

	assert.Nil(suite.T(), err)
	assert.True(
		suite.T(),
		(strings.Contains(string(out), "86.07°") &&
			strings.Contains(string(out), "70.00°") &&
			strings.Contains(string(out), "4:3") &&
			strings.Contains(string(out), "102.45°") &&
			strings.Contains(string(out), "70.00°") &&
			strings.Contains(string(out), "16:9")),
		"Expected values are present in output")
}
