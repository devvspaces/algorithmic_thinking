package main

import "fmt"

type Point struct {
	x int
	y int
}

var maze = []string{
	"x xxxxxxxx x",
	"x        x x",
	"xxxx    xx x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x          x",
	"xxxxxxxxxxxx",
}

var dirs = [][]int{
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, 0},
}

func Append(s *[]Point, p Point) {
	*s = append(*s, p)
}

func walk(maze []string, wall string, end Point, curr Point, seen [][]bool, path *[]Point) bool {

	if curr.x < 0 || curr.y < 0 || curr.x >= len(maze[0]) || curr.y >= len(maze) {
		return false
	}

	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	if curr.y == end.y && curr.x == end.x {
		Append(path, curr)
		return true
	}

	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	Append(path, curr)

	for _, value := range dirs {
		if walk(maze, wall, end, Point{x: curr.x + value[0], y: curr.y + value[1]}, seen, path) {
			return true
		}
	}

	*path = (*path)[0 : len(*path)-1]

	return false

}

func find_path(maze []string, wall string, start Point, end Point) []Point {

	seen := [][]bool{}

	for idx := range maze {
		row := []bool{}
		for i := 0; i < len(maze[idx]); i++ {
			row = append(row, false)
		}
		seen = append(seen, row)
	}

	path := []Point{}
	walk(maze, wall, end, start, seen, &path)
	return path

}

func main() {

	start := Point{x: 10, y: 0}
	end := Point{x: 1, y: 0}
	path := find_path(
		maze, "x", start, end,
	)
	fmt.Println(path)

}
