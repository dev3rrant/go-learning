package maths

import (
	"fmt"
	"io"
	"math"
	"time"
)

const secondHandLength = 90
const minuteHandLength = 80
const hourHandLength = 50
const clockCenterX = 150
const clockCenterY = 150

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func SvgWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	Secondhand(w, t)
	Minutehand(w, t)
	Hourhand(w, t)
	io.WriteString(w, svgEnd)
}

type Point struct {
	X float64
	Y float64
}

func Secondhand(w io.Writer, t time.Time) {
	p := makehand(SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func Minutehand(w io.Writer, t time.Time) {
	p := makehand(MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func Hourhand(w io.Writer, t time.Time) {
	p := makehand(HourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makehand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCenterX, p.Y + clockCenterY}
}

func angleToPoints(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SecondsToRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func MinutesToRadians(t time.Time) float64 {
	return (SecondsToRadians(t) / secondsInClock) + (math.Pi / (secondsInHalfClock / float64(t.Minute())))
}

func HoursToRadians(t time.Time) float64 {
	return (MinutesToRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64((t.Hour() % hoursInClock))))
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoints(SecondsToRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoints(MinutesToRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoints(HoursToRadians(t))
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
