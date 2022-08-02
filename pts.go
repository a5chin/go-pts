package pts

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type PointsOperator interface {
	Describe() string
	Mean() (float64 float64)
	Std() (float64 float64)
	String() string
	Var() (float64 float64)
}

type PointOperator interface {
	Norm() (float64, float64)
	String() string
}

type Points struct {
	Array    []Point
	Path     string
	N_points int
}

type Point struct {
	X, Y float64
}

func (point *Point) Norm() (x float64, y float64) {
	r := math.Sqrt(
		math.Pow(point.X, 2) + math.Pow(point.Y, 2),
	)
	x, y = point.X/r, point.Y/r

	return x, y
}

func (point *Point) String() (s string) {
	s = fmt.Sprintf(
		"x: %v, y: %v", point.X, point.Y,
	)

	return s
}

func (points *Points) Describe() (s string) {
	x_mean, y_mean := points.Mean()
	x_var, y_var := points.Var()
	x_std, y_std := points.Std()

	s = fmt.Sprintf(
		`
	x	y
Mean:	%v	%v
Var:	%v	%v
Std:	%v	%v
`, x_mean, y_mean, x_var, y_var, x_std, y_std,
	)

	return s
}

func (points *Points) Mean() (float64, float64) {
	var (
		x, y     float64
		n_points float64 = float64(points.N_points)
	)
	for _, point := range points.Array {
		x += point.X
		y += point.Y
	}

	return x / n_points, y / n_points
}

func (points *Points) Var() (float64, float64) {
	var (
		x, y float64
	)
	x_mean, y_mean := points.Mean()

	for _, point := range points.Array {
		x += math.Pow(point.X-x_mean, 2)
		y += math.Pow(point.Y-y_mean, 2)
	}

	return x / float64(points.N_points), y / float64(points.N_points)
}

func (points *Points) Std() (float64, float64) {
	x_var, y_var := points.Var()

	return math.Sqrt(x_var), math.Sqrt(y_var)
}

func (points *Points) String() (s string) {
	s = fmt.Sprintf(
		"The file named %v has %v points.",
		points.Path,
		points.N_points,
	)

	return s
}

func ReadPTS(path string) (points Points) {
	path, _ = filepath.Abs(path)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}

	var (
		n_points int
		arr      []Point
	)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		if i == 1 {
			n_points = findNumber(text)
		} else if 2 < i && i < n_points+3 {
			x, y := createPoint(text)
			arr = append(arr, Point{x, y})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	points = Points{
		Array:    arr,
		Path:     path,
		N_points: n_points,
	}

	return points
}

func findNumber(text string) (n_points int) {
	num := regexp.MustCompile(`[0-9]+`)
	n := num.FindString(text)
	n_points, _ = strconv.Atoi(n)

	return n_points
}

func createPoint(str string) (x, y float64) {
	array := strings.Split(str, " ")
	x, _ = strconv.ParseFloat(array[0], 64)
	y, _ = strconv.ParseFloat(array[1], 64)
	return x, y
}
