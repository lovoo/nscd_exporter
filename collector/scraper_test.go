package collector

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
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

func TestValueParser(t *testing.T) {
	type testData struct {
		raw          string
		shouldReturn float64
	}

	td := []testData{
		{"52%", 52},
		{"127", 127},
		{"-31", -31},
		{"16s", 16 * time.Second.Seconds()},
		{"255h 32s", 255*time.Hour.Seconds() + 32*time.Second.Seconds()},
		{"1d 20h  4m 33s", 44*time.Hour.Seconds() + 4*time.Minute.Seconds() + 33*time.Second.Seconds()},
	}

	for _, d := range td {
		t.Run(d.raw, func(t *testing.T) {
			var (
				r   float64
				err error
			)
			for _, vp := range vps {
				r, err = vp(d.raw)
				if err == nil {
					break
				}
			}
			if r != d.shouldReturn || err != nil {
				t.Fail()
			}
		})
	}
}
