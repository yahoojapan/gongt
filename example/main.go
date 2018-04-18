//
// Copyright (C) 2017 Yahoo Japan Corporation:
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
	"runtime"

	"github.com/kpango/glg"
	"github.com/yahoojapan/gongt"
	"gonum.org/v1/hdf5"
)

func getVectors(path, key string) ([][]float64, error) {
	f, err := hdf5.OpenFile(path, hdf5.F_ACC_RDONLY)
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

func create(name, path string) {
	if _, err := os.Stat(name); err == nil {
		glg.Infof("[%s] %s exists", name, path)
		return
	}
	vectors, err := getVectors(path, "train")
	if err != nil {
		glg.Warn(err)
		return
	}
	glg.Infof("[%s] %d items", name, len(vectors))
	defer glg.Infof("[%s] done", name)

	n := gongt.New(name).SetObjectType(gongt.Float).SetDimension(len(vectors[0])).Open()
	defer n.Close()

	for _, v := range vectors {
		n.Insert(v)
	}
	if err := n.CreateAndSaveIndex(runtime.NumCPU()); err != nil {
		glg.Warn(err)
	}
}

func search(name, path string) {
	n := gongt.New(name).Open()
	defer n.Close()

	vectors, err := getVectors(path, "test")
	if err != nil {
		glg.Warn(err)
		return
	}
	glg.Infof("[%s] %d items", name, len(vectors))
	defer glg.Infof("[%s] done", name)

	for _, v := range vectors {
		n.Search(v, 10, gongt.DefaultEpsilon)
		// result, err := n.Search(v, 10, gongt.DefaultEpsilon) // do something using result and err
	}
}

func main() {
	c := flag.Bool("create", false, "run create")
	s := flag.Bool("search", false, "run search")

	n := flag.String("name", "", "dataset name")
	p := flag.String("path", "", "dataset path")

	flag.Parse()
	if *c {
		create(*n, *p)
	}
	if *s {
		search(*n, *p)
	}
}
