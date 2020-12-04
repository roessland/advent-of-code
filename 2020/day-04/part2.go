package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type DistanceUnit string
const Centimeter DistanceUnit = "cm"
const Inch DistanceUnit = "in"
var ErrInvalidDistanceUnit = errors.New("invalid distance unit")

func NewDistanceUnit(unit string) (DistanceUnit, error) {
	if unit == string(Centimeter) {
		return DistanceUnit(unit), nil
	}
	if unit == string(Inch) {
		return DistanceUnit(unit), nil
	}
	return DistanceUnit(""), ErrInvalidDistanceUnit
}

type Height struct {
	Value int
	Unit DistanceUnit
}

func NewHeight(height string) (Height, error) {
	r := regexp.MustCompile(`^(?P<Val>\d+)(?P<Unit>.+)$`)
	match := r.FindStringSubmatch(height)
	results := map[string]string{}
	for i, name := range match {
		results[r.SubexpNames()[i]] = name
	}

	unit, err := NewDistanceUnit(results["Unit"])
	if err != nil {
		return Height{}, err
	}
	val, err := strconv.Atoi(results["Val"])
	if err != nil {
		return Height{}, errors.New("invalid number " + results["Val"])
	}
	return Height{Value: val, Unit: unit}, nil
}

type HairColor string

func NewHairColor(color string) (HairColor, error) {
	r := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if !r.MatchString(color) {
		return "", errors.New("invalid hair color " + color)
	}
	return HairColor(color), nil
}

type EyeColor string

func NewEyeColor(color string) (EyeColor, error) {
	for _, validColor := range []EyeColor{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if color == string(validColor) {
			return validColor, nil
		}
	}
	return "", errors.New("invalid eye color " + color)
}

type PassportID string

func NewPassportID(id string) (PassportID, error) {
	r := regexp.MustCompile(`^\d{9}$`)
	if !r.MatchString(id) {
		return "", errors.New("invalid passport ID " + id)
	}
	return PassportID(id), nil
}

type Pass struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func (p Pass) ToValidated() (ValidatedPass, error) {
	var v ValidatedPass
	var err error

	v.Byr, err = strconv.Atoi(p.Byr)
	if err != nil || v.Byr < 1920 || 2002 < v.Byr {
		return v, errors.New("invalid birthyear")
	}

	v.Iyr, err = strconv.Atoi(p.Iyr)
	if err != nil || v.Iyr < 2010 || 2020 < v.Iyr {
		return v, errors.New("invalid issue year")
	}

	v.Eyr, err = strconv.Atoi(p.Eyr)
	if err != nil || v.Eyr < 2020 || 2030 < v.Eyr {
		return v, errors.New("invalid expiration year")
	}

	v.Hgt, err = NewHeight(p.Hgt)
	if err != nil ||
		(v.Hgt.Unit == Inch && (v.Hgt.Value < 59 || 76 < v.Hgt.Value)) ||
		(v.Hgt.Unit == Centimeter && (v.Hgt.Value < 150 || 193 < v.Hgt.Value)) {
		return v, errors.New("invalid height")
	}

	v.Hcl, err = NewHairColor(p.Hcl)
	if err != nil {
		return v, err
	}

	v.Ecl, err = NewEyeColor(p.Ecl)
	if err != nil {
		return v, err
	}

	v.Pid, err = NewPassportID(p.Pid)
	if err != nil {
		return v, err
	}

	v.Cid = p.Cid

	return v, nil
}

type ValidatedPass struct {
	Byr int
	Iyr int
	Eyr int
	Hgt Height
	Hcl HairColor
	Ecl EyeColor
	Pid PassportID
	Cid string
}

func (p *Pass) AddInfos(infos string) {
	infos = strings.ReplaceAll(infos, "\n", " ")
	parts := strings.Split(infos, " ")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		keyVal := strings.Split(part, ":")
		p.AddInfo(keyVal[0], strings.Trim(keyVal[1], " "))
	}
	return
}

func (p *Pass) AddInfo(key, val string) {
	switch key {
	case "iyr":
		p.Iyr = val
	case "cid":
		p.Cid = val
	case "eyr":
		p.Eyr = val
	case "hcl":
		p.Hcl = val
	case "hgt":
		p.Hgt = val
	case "ecl":
		p.Ecl = val
	case "byr":
		p.Byr = val
	case "pid":
		p.Pid = val
	default:
		log.Fatal("unknown key " + key)
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var ps []Pass
	var p Pass
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ps = append(ps, p)
			p = Pass{}
		}
		p.AddInfos(line)
	}
	ps = append(ps, p) // the final one

	numValid := 0
	for _, p := range ps {
		_, err := p.ToValidated()
		if err == nil {
			numValid++
		}
	}
	fmt.Println(numValid)
}