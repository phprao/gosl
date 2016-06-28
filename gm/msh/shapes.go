// Copyright 2015 Dorival Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

// ShapeFunction computes the shape function and derivatives
type ShapeFunction func(S []float64, dSdR [][]float64, R []float64, derivs bool)

var (
	NumVerts       map[string]int         // number of vertices on shape
	GeomNdim       map[string]int         // geometry number of space dimensions
	EdgeLocalVerts map[string][][]int     // local indices of vertices on edges of shape
	FaceLocalVerts map[string][][]int     // local indices of vertices on faces of shape
	NatCoords      map[string][][]float64 // natural coordinates of vertices on shape
)

var Functions = make(map[string]ShapeFunction) // shape functions and derivatives

func init() {

	NumVerts = map[string]int{
		"lin2":  2,
		"lin3":  3,
		"lin4":  4,
		"lin5":  5,
		"tri3":  3,
		"tri6":  6,
		"tri10": 10,
		"tri15": 15,
		"qua4":  4,
		"qua8":  8,
		"qua9":  9,
		"qua12": 12,
		"qua16": 16,
		"tet4":  4,
		"tet10": 10,
		"hex8":  8,
		"hex20": 20,
	}

	GeomNdim = map[string]int{
		"lin2":  1,
		"lin3":  1,
		"lin4":  1,
		"lin5":  1,
		"tri3":  2,
		"tri6":  2,
		"tri10": 2,
		"tri15": 2,
		"qua4":  2,
		"qua8":  2,
		"qua9":  2,
		"qua12": 2,
		"qua16": 2,
		"tet4":  3,
		"tet10": 3,
		"hex8":  3,
		"hex20": 3,
	}

	EdgeLocalVerts = map[string][][]int{
		"lin2":  [][]int{},
		"lin3":  [][]int{},
		"lin4":  [][]int{},
		"lin5":  [][]int{},
		"tri3":  [][]int{{0, 1}, {1, 2}, {2, 0}},
		"tri6":  [][]int{{0, 1, 3}, {1, 2, 4}, {2, 0, 5}},
		"tri10": [][]int{{0, 1, 3, 6}, {1, 2, 4, 7}, {2, 0, 5, 8}},
		"tri15": [][]int{{0, 1, 3, 6, 7}, {1, 2, 4, 8, 9}, {2, 0, 5, 10, 11}},
		"qua4":  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
		"qua8":  [][]int{{0, 1, 4}, {1, 2, 5}, {2, 3, 6}, {3, 0, 7}},
		"qua9":  [][]int{{0, 1, 4}, {1, 2, 5}, {2, 3, 6}, {3, 0, 7}},
		"qua12": [][]int{{0, 1, 4, 8}, {1, 2, 5, 9}, {2, 3, 6, 10}, {3, 0, 7, 11}},
		"qua16": [][]int{{0, 1, 4, 8}, {1, 2, 5, 9}, {2, 3, 6, 10}, {3, 0, 7, 11}},
		"tet4":  [][]int{{0, 1}, {1, 2}, {2, 0}, {0, 3}, {1, 3}, {2, 3}},
		"tet10": [][]int{{0, 1, 4}, {1, 2, 5}, {2, 0, 6}, {0, 3, 7}, {1, 3, 8}, {2, 3, 9}},
		"hex8":  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 5}, {5, 6}, {6, 7}, {7, 4}, {0, 4}, {1, 5}, {2, 6}, {3, 7}},
		"hex20": [][]int{{0, 1, 8}, {1, 2, 9}, {2, 3, 10}, {3, 0, 11}, {4, 5, 12}, {5, 6, 13}, {6, 7, 14}, {7, 4, 15}, {0, 4, 16}, {1, 5, 17}, {2, 6, 18}, {3, 7, 19}},
	}

	FaceLocalVerts = map[string][][]int{
		"tet4":  [][]int{{0, 3, 2}, {0, 1, 3}, {0, 2, 1}, {1, 2, 3}},
		"tet10": [][]int{{0, 3, 2, 7, 9, 6}, {0, 1, 3, 4, 8, 7}, {0, 2, 1, 6, 5, 4}, {1, 2, 3, 5, 9, 8}},
		"hex8":  [][]int{{0, 4, 7, 3}, {1, 2, 6, 5}, {0, 1, 5, 4}, {2, 3, 7, 6}, {0, 3, 2, 1}, {4, 5, 6, 7}},
		"hex20": [][]int{{0, 4, 7, 3, 16, 15, 19, 11}, {1, 2, 6, 5, 9, 18, 13, 17}, {0, 1, 5, 4, 8, 17, 12, 16}, {2, 3, 7, 6, 10, 19, 14, 18}, {0, 3, 2, 1, 11, 10, 9, 8}, {4, 5, 6, 7, 12, 13, 14, 15}},
	}

	NatCoords = map[string][][]float64{
		"lin2": [][]float64{
			{-1, 1},
		},
		"lin3": [][]float64{
			{-1, 1, 0},
		},
		"lin4": [][]float64{
			{-1, 1, -1.0 / 3.0, 1.0 / 3.0},
		},
		"lin5": [][]float64{
			{-1, 1, 0, -0.5, 0.5},
		},
		"tri3": [][]float64{
			{0, 1, 0},
			{0, 0, 1},
		},
		"tri6": [][]float64{
			{0, 1, 0, 0.5, 0.5, 0},
			{0, 0, 1, 0, 0.5, 0.5},
		},
		"tri10": [][]float64{
			{0, 1, 0, 1.0 / 3.0, 2.0 / 3.0, 0, 2.0 / 3.0, 1.0 / 3.0, 0, 1.0 / 3.0},
			{0, 0, 1, 0, 1.0 / 3.0, 2.0 / 3.0, 0, 2.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0},
		},
		"tri15": [][]float64{
			{0, 1, 0, 0.5, 0.5, 0, 0.25, 0.75, 0.75, 0.25, 0, 0, 0.25, 0.5, 0.25},
			{0, 0, 1, 0, 0.5, 0.5, 0, 0, 0.25, 0.75, 0.75, 0.25, 0.25, 0.25, 0.5},
		},
		"qua4": [][]float64{
			{-1, 1, 1, -1},
			{-1, -1, 1, 1},
		},
		"qua8": [][]float64{
			{-1, 1, 1, -1, 0, 1, 0, -1},
			{-1, -1, 1, 1, -1, 0, 1, 0},
		},
		"qua9": [][]float64{
			{-1, 1, 1, -1, 0, 1, 0, -1, 0},
			{-1, -1, 1, 1, -1, 0, 1, 0, 0},
		},
		"qua12": [][]float64{
			{-1, 1, 1, -1, -1.0 / 3.0, 1, 1.0 / 3.0, -1, 1.0 / 3.0, 1, -1.0 / 3.0, -1},
			{-1, -1, 1, 1, -1, -1.0 / 3.0, 1, 1.0 / 3.0, -1, 1.0 / 3.0, 1, -1.0 / 3.0},
		},
		"qua16": [][]float64{
			{-1, 1, 1, -1, -1.0 / 3.0, 1, 1.0 / 3.0, -1, 1.0 / 3.0, 1, -1.0 / 3.0, -1, -1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0, -1.0 / 3.0},
			{-1, -1, 1, 1, -1, -1.0 / 3.0, 1, 1.0 / 3.0, -1, 1.0 / 3.0, 1, -1.0 / 3.0, -1.0 / 3.0, -1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0},
		},
		"tet4": [][]float64{
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
		"tet10": [][]float64{
			{0, 1, 0, 0, 0.5, 0.5, 0, 0, 0.5, 0},
			{0, 0, 1, 0, 0, 0.5, 0.5, 0, 0, 0.5},
			{0, 0, 0, 1, 0, 0, 0, 0.5, 0.5, 0.5},
		},
		"hex8": [][]float64{
			{-1, 1, 1, -1, -1, 1, 1, -1},
			{-1, -1, 1, 1, -1, -1, 1, 1},
			{-1, -1, -1, -1, 1, 1, 1, 1},
		},
		"hex20": [][]float64{
			{-1, 1, 1, -1, -1, 1, 1, -1, 0, 1, 0, -1, 0, 1, 0, -1, -1, 1, 1, -1},
			{-1, -1, 1, 1, -1, -1, 1, 1, -1, 0, 1, 0, -1, 0, 1, 0, -1, -1, 1, 1},
			{-1, -1, -1, -1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 0, 0, 0, 0},
		},
	}
}