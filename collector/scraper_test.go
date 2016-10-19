package collector

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("nscd_output_sample.txt")
	if err != nil {
		t.Fatalf("could not open sample file: %v", err)
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("could not read from file: %v", err)
	}

	data, err := parse(buf)
	if err != nil {
		t.Fatalf("parsing failed: %v", err)
	}

	if data["hosts cache"][0].val != 1 {
		t.Fail()
	}
	if data["hosts cache"][3].val != 211 {
		t.Fail()
	}
	if data["hosts cache"][12].val != 78 {
		t.Fail()
	}
}
