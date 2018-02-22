//
// Copyright (C) 2017 Yahoo Japan Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"flag"
	"os"

	"github.com/kpango/glg"
	"github.com/yahoojapan/gongt"
	"gonum.org/v1/hdf5"
)

type dataset struct {
	Name string
	Path string
}

var datasets = []*dataset{
	{"Fashion-MNIST", "assets/bench/fashion-mnist-784-euclidean.hdf5"},
	{"GloVe-25", "assets/bench/glove-25-angular.hdf5"},
	{"GloVe-50", "assets/bench/glove-50-angular.hdf5"},
	{"GloVe-100", "assets/bench/glove-100-angular.hdf5"},
	{"GloVe-200", "assets/bench/glove-200-angular.hdf5"},
	{"MNIST", "assets/bench/mnist-784-euclidean.hdf5"},
	{"NYTimes", "assets/bench/nytimes-256-angular.hdf5"},
	{"SIFT", "assets/bench/sift-128-euclidean.hdf5"},
}

func (d *dataset) getVectors(key string) ([][]float64, error) {
	f, err := hdf5.OpenFile(d.Path, hdf5.F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	dset, err := f.OpenDataset(key)
	if err != nil {
		return nil, err
	}
	space := dset.Space()
	dims, _, err := space.SimpleExtentDims()
	if err != nil {
		return nil, err
	}
	v := make([]float32, space.SimpleExtentNPoints())
	if err := dset.Read(&v); err != nil {
		return nil, err
	}

	row := int(dims[0])
	col := int(dims[1])

	vec := make([][]float64, row)
	for i := 0; i < row; i++ {
		vec[i] = make([]float64, col)
		for j := 0; j < col; j++ {
			vec[i][j] = float64(v[i*col+j])
		}
	}
	return vec, nil
}

func (d *dataset) create() {
	path := "assets/bench/" + d.Name
	if _, err := os.Stat(path); err == nil {
		glg.Infof("[%s] %s exists", d.Name, path)
		return
	}
	vectors, err := d.getVectors("train")
	if err != nil {
		glg.Warn(err)
		return
	}
	glg.Infof("[%s] %d items", d.Name, len(vectors))
	defer glg.Infof("[%s] done", d.Name)

	n := gongt.New(path).SetObjectType(gongt.Float).SetDimension(len(vectors[0])).Open()
	defer n.Close()

	for _, v := range vectors {
		n.Insert(v)
	}
	if err := n.CreateAndSaveIndex(4); err != nil {
		glg.Warn(err)
	}
}

func (d *dataset) search() {
	path := "assets/bench/" + d.Name
	n := gongt.New(path).Open()
	defer n.Close()

	vectors, err := d.getVectors("test")
	if err != nil {
		glg.Warn(err)
		return
	}
	glg.Infof("[%s] %d items", d.Name, len(vectors))
	defer glg.Infof("[%s] done", d.Name)

	for _, v := range vectors {
		n.Search(v, 10, gongt.DefaultEpsilon)
		// result, err := n.Search(v, 10, gongt.DefaultEpsilon) // do something using result and err
	}
}

func main() {
	c := flag.Bool("create", false, "run create")
	s := flag.Bool("search", false, "run search")

	flag.Parse()
	for _, d := range datasets {
		if *c {
			d.create()
		}
		if *s {
			d.search()
		}
	}
}
