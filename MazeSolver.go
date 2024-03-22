package main

type Point struct {
	x int
	y int
}

var directions = [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func walkMaze(maze []string, wall byte, curr Point, end Point, seen *[][]bool, path *[]Point) bool {
	if curr.x < 0 || curr.y < 0 || curr.x >= len(maze[0]) || curr.y >= len(maze) {
		return false
	}

	if maze[curr.y][curr.x] == wall {
		return false
	}

	if curr.x == end.x && curr.y == end.y {
		(*path) = append(*path, curr)
		return true
	}

	if (*seen)[curr.y][curr.x] {
		return false
	}

	(*seen)[curr.y][curr.x] = true
	(*path) = append(*path, curr)

	for i := 0; i < len(directions); i++ {
		testPoint := Point{x: curr.x + directions[i].x, y: curr.y + directions[i].y}
		if walkMaze(maze, wall, testPoint, end, seen, path) {
			return true
		}
	}

	(*path) = (*path)[:len(*path)-1]

	return false
}

func MazeSolver(maze []string, wall byte, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	path := []Point{}

	walkMaze(maze, wall, start, end, &seen, &path)

	return path
}
