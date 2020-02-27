//+build ignore

package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	os.MkdirAll("dsp32/dsp", 0666)
	os.MkdirAll("dsp64/dsp", 0666)

	dsp, err := ioutil.ReadFile("dsp.go")
	check(err)
	dsp32 := strings.Replace(string(dsp), "FLOAT", "float32", -1)
	dsp64 := strings.Replace(string(dsp), "FLOAT", "float64", -1)
	check(ioutil.WriteFile("dsp32/dsp/dsp.go", []byte(dsp32), 0666))
	check(ioutil.WriteFile("dsp64/dsp/dsp.go", []byte(dsp64), 0666))

	dspTest, err := ioutil.ReadFile("dsp_test.go")
	check(err)
	dspTest32 := strings.Replace(string(dspTest), "FLOAT", "float32", -1)
	dspTest64 := strings.Replace(string(dspTest), "FLOAT", "float64", -1)
	check(ioutil.WriteFile("dsp32/dsp/dsp_test.go", []byte(dspTest32), 0666))
	check(ioutil.WriteFile("dsp64/dsp/dsp_test.go", []byte(dspTest64), 0666))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
