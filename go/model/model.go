package model

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type Model struct {
	verts []Vec3f;
	faces [][]int;
}

func LoadFromFile(filename string) (*Model, error) {
	file, err := os.Open(filename);
	if err != nil {
		return nil, err
	}

	m := &Model{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var text string = scanner.Text()
		fmt.Println(text)

		if text == "" {
			continue
		}

		switch text[0:2] {
		case "v ":
			var x, y, z float64;
			n, err := fmt.Sscanf(text, "v %f %f %f", &x, &y, &z)
			if n != 3 {
				return nil, fmt.Errorf("Expected vertice containing 3 number, got %d", n)
			}
			if err != nil {
				return nil, err
			}

			v := NewVec3f(x, y, z)

			m.verts = append(m.verts, v)
		case "f ":
			var face []int = make([]int, 0)
			for _, t := range strings.Split(text[2:], " ") {
				s_index := strings.Split(t, "/")[0]
				idx, err := strconv.ParseInt(s_index, 10, 32)

				if err != nil {
					return nil, err
				}

				idx -= 1;

				face = append(face, int(idx))
			}

			m.faces = append(m.faces, face)
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m Model) Nverts() int {
	return len(m.verts)
}

func (m Model) Nfaces() int {
	return len(m.faces)
}

func (m Model) Vertice(i int) Vec3f {
	return m.verts[i]
}

func (m Model) Face( i int) []int {
	return m.faces[i]
}
