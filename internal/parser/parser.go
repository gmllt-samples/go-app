package parser

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(input string) (time.Duration, error) {
	if strings.Contains(input, ",") {
		parts := strings.Split(input, ",")
		idx, err := randInt(0, len(parts)-1)
		if err != nil {
			return 0, err
		}
		return time.ParseDuration(strings.TrimSpace(parts[idx]))
	}
	if strings.Contains(input, "-") {
		parts := strings.SplitN(input, "-", 2)
		min, err1 := time.ParseDuration(parts[0])
		max, err2 := time.ParseDuration(parts[1])
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("invalid duration range")
		}
		if min > max {
			min, max = max, min
		}
		diff := max - min
		nano, err := randInt(0, int(diff.Nanoseconds()))
		if err != nil {
			return 0, err
		}
		return min + time.Duration(nano), nil
	}
	return time.ParseDuration(input)
}

func ParseSize(input string) (int, error) {
	if strings.Contains(input, ",") {
		parts := strings.Split(input, ",")
		idx, err := randInt(0, len(parts)-1)
		if err != nil {
			return 0, err
		}
		return parseUnit(parts[idx])
	}
	if strings.Contains(input, "-") {
		parts := strings.SplitN(input, "-", 2)
		min, err1 := parseUnit(parts[0])
		max, err2 := parseUnit(parts[1])
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("invalid size range")
		}
		if min > max {
			min, max = max, min
		}
		return randInt(min, max)
	}
	return parseUnit(input)
}

func ParseStatus(input string) (int, error) {
	if strings.Contains(input, ",") {
		parts := strings.Split(input, ",")
		idx, err := randInt(0, len(parts)-1)
		if err != nil {
			return 0, err
		}
		return strconv.Atoi(strings.TrimSpace(parts[idx]))
	}
	if strings.Contains(input, "-") {
		parts := strings.SplitN(input, "-", 2)
		min, err1 := strconv.Atoi(parts[0])
		max, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("invalid status range")
		}
		if min > max {
			min, max = max, min
		}
		return randInt(min, max)
	}
	return strconv.Atoi(input)
}

func parseUnit(input string) (int, error) {
	units := map[string]int{
		"B": 1,
		"K": 1_000,
		"M": 1_000_000,
		"G": 1_000_000_000,
	}
	s := strings.ToUpper(strings.TrimSpace(input))
	split := len(s)
	for i, r := range s {
		if r < '0' || r > '9' {
			split = i
			break
		}
	}
	num := s[:split]
	unit := s[split:]

	val, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	if unit == "" {
		return val, nil
	}
	mult, ok := units[unit]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}
	return val * mult, nil
}

func randInt(min, max int) (int, error) {
	if min >= max {
		return min, nil
	}
	diff := max - min + 1
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	if err != nil {
		return 0, fmt.Errorf("secure random failed: %w", err)
	}
	return int(nBig.Int64()) + min, nil
}
