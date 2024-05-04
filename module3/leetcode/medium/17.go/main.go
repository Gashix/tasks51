package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type QuadTreeNode struct {
	TopLeft, BottomRight Point
	Points               []Point
	Children             [4]*QuadTreeNode
}

type QuadTree struct {
	Root *QuadTreeNode
}

func NewQuadTree(topLeft, bottomRight Point) *QuadTree {
	return &QuadTree{
		Root: &QuadTreeNode{
			TopLeft:     topLeft,
			BottomRight: bottomRight,
			Points:      make([]Point, 0),
			Children:    [4]*QuadTreeNode{},
		},
	}
}

func (qt *QuadTree) Insert(p Point) {
	insertPoint(qt.Root, p)
}

func insertPoint(node *QuadTreeNode, p Point) {
	if node == nil {
		return
	}

	if node.TopLeft.X <= p.X && p.X <= node.BottomRight.X &&
		node.TopLeft.Y <= p.Y && p.Y <= node.BottomRight.Y {
		node.Points = append(node.Points, p)
	} else {
		midX, midY := (node.TopLeft.X+node.BottomRight.X)/2, (node.TopLeft.Y+node.BottomRight.Y)/2

		if p.X <= midX {
			if p.Y <= midY {
				if node.Children[0] == nil {
					node.Children[0] = &QuadTreeNode{
						TopLeft:     node.TopLeft,
						BottomRight: Point{midX, midY},
						Points:      make([]Point, 0),
						Children:    [4]*QuadTreeNode{},
					}
				}
				insertPoint(node.Children[0], p)
			} else {
				if node.Children[2] == nil {
					node.Children[2] = &QuadTreeNode{
						TopLeft:     Point{node.TopLeft.X, midY + 1},
						BottomRight: Point{midX, node.BottomRight.Y},
						Points:      make([]Point, 0),
						Children:    [4]*QuadTreeNode{},
					}
				}
				insertPoint(node.Children[2], p)
			}
		} else {
			if p.Y <= midY {
				if node.Children[1] == nil {
					node.Children[1] = &QuadTreeNode{
						TopLeft:     Point{midX + 1, node.TopLeft.Y},
						BottomRight: Point{node.BottomRight.X, midY},
						Points:      make([]Point, 0),
						Children:    [4]*QuadTreeNode{},
					}
				}
				insertPoint(node.Children[1], p)
			} else {
				if node.Children[3] == nil {
					node.Children[3] = &QuadTreeNode{
						TopLeft:     Point{midX + 1, midY + 1},
						BottomRight: node.BottomRight,
						Points:      make([]Point, 0),
						Children:    [4]*QuadTreeNode{},
					}
				}
				insertPoint(node.Children[3], p)
			}
		}
	}
}

func (qt *QuadTree) QueryCircle(center Point, radius int) int {
	return queryCircle(qt.Root, center, radius)
}

func queryCircle(node *QuadTreeNode, center Point, radius int) int {
	if node == nil {
		return 0
	}

	if isInsideCircle(center, radius, node.TopLeft) &&
		isInsideCircle(center, radius, node.BottomRight) {
		return len(node.Points)
	}

	count := 0
	for _, p := range node.Points {
		if isInsideCircle(center, radius, p) {
			count++
		}
	}

	for _, child := range node.Children {
		count += queryCircle(child, center, radius)
	}

	return count
}

func isInsideCircle(center Point, radius int, p Point) bool {
	return math.Pow(float64(p.X-center.X), 2)+math.Pow(float64(p.Y-center.Y), 2) <= float64(radius*radius)
}

func main() {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}
	fmt.Println(countPoints(points, queries)) // Output: [3 2 2]
}

func countPoints(points [][]int, queries [][]int) []int {
	qt := NewQuadTree(Point{0, 0}, Point{500, 500})
	for _, p := range points {
		qt.Insert(Point{p[0], p[1]})
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		center := Point{q[0], q[1]}
		radius := q[2]
		ans[i] = qt.QueryCircle(center, radius)
	}
	return ans
}
