package collector

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type dataLine struct {
	val  float64
	desc string
}

type valueParser func(rawValue string) (float64, error)

// vps contains the strategies to parse a value from
// the nscd output
var vps = []valueParser{
	// parse int strategy
	func(rawValue string) (float64, error) {
		v, err := strconv.Atoi(rawValue)
		return float64(v), err
	},
	// parse boolean "yes"/"no"
	func(rawValue string) (float64, error) {
		if rawValue == "no" {
			return 0, nil
		} else if rawValue == "yes" {
			return 1, nil
		} else {
			return 0, errors.New("not a boolean value")
		}
	},
	// parse percentage value
	func(rawValue string) (float64, error) {
		if strings.HasSuffix(rawValue, "%") {
			perc := strings.TrimSuffix(rawValue, "%")
			v, err := strconv.Atoi(perc)
			return float64(v), err
		}
		return 0, errors.New("not a percentage value")
	},
	// parse duration
	func(rawValue string) (float64, error) {
		s := strings.Replace(rawValue, " ", "", -1)

		// determine days, because time.ParseDuration can't parse it
		regex, err := regexp.Compile("(\\d)*d")
		if err != nil {
			return 0, err
		}
		var (
			offset    int // additional days
			dayStrLen int
		)
		ds := regex.FindStringSubmatch(s)
		if len(ds) > 1 {
			days, err := strconv.Atoi(ds[1])
			if err != nil {
				return 0, err
			}
			dayStrLen = len(ds[0])
			offset += days * 24 * int(time.Hour.Seconds())
		}
		withoutDays := s[dayStrLen:]
		d, err := time.ParseDuration(withoutDays)
		return float64(int(d.Seconds()) + offset), err
	},
}

// Scrape calls nscd --statistics and parses the result.
func Scrape(nscdPath string) (map[string][]dataLine, error) {
	out, err := exec.Command(nscdPath, "--statistics").Output()
	if err != nil {
		return nil, err
	}
	return parse(out)
}

func parse(buf []byte) (map[string][]dataLine, error) {
	var (
		section string
		data    = map[string][]dataLine{}
		lines   = strings.Split(string(buf), "\n")
	)
	for _, line := range lines {
		// skip empty lines
		if len(line) == 0 {
			continue
		}
		// new section
		if !strings.HasPrefix(line, " ") {
			section = strings.TrimSuffix(line, ":")
			continue
		}

		// parse data line
		var (
			rawVal = strings.TrimSpace(line[:16])
			val    float64
			desc   = strings.TrimSpace(line[17:])
			err    error
		)

		for _, p := range vps {
			val, err = p(rawVal)
			if err == nil {
				break
			}
		}
		if err != nil {
			return nil, fmt.Errorf("could not parse line: %v", line)
		}

		dl := dataLine{
			val:  val,
			desc: desc,
		}
		data[section] = append(data[section], dl)
	}
	return data, nil
}
