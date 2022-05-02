package clockface

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

// Help from https://www.onlinetool.io/xmltogo/
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		input    time.Time
		expected Line
	}{
		// Test case 1
		{
			input:    simpleTime(0, 0, 0),
			expected: Line{150, 150, 150, 60},
		},
		// Test case 2
		{
			input:    simpleTime(0, 0, 30),
			expected: Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.input), func(t *testing.T) {

			b := bytes.Buffer{}
			SVGWriter(&b, c.input)
			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.expected, svg.Line) {
				t.Errorf(
					"Expected to find the second hand line %+v, in the SVG lines %+v",
					c.expected,
					svg.Line,
				)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		input    time.Time
		expected Line
	}{
		// Test case 1
		{
			input:    simpleTime(0, 0, 0),
			expected: Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.input), func(t *testing.T) {

			b := bytes.Buffer{}
			SVGWriter(&b, c.input)
			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.expected, svg.Line) {
				t.Errorf(
					"Expected to find the minute hand line %+v, in the SVG lines %+v",
					c.expected,
					svg.Line,
				)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		input    time.Time
		expected Line
	}{
		// Test case 1
		{
			simpleTime(6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.input), func(t *testing.T) {

			b := bytes.Buffer{}
			SVGWriter(&b, c.input)
			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.expected, svg.Line) {
				t.Errorf(
					"Expected to find the hour hand line %+v, in the SVG lines %+v",
					c.expected,
					svg.Line,
				)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}
