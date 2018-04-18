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

package gongt_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/yahoojapan/gongt"
	"gonum.org/v1/hdf5"
)

type data struct {
	name string
	path string
}

var (
	fashionmnist = data{"Fashion-MNIST", "assets/bench/fashion-mnist-784-euclidean.hdf5"}
	glove25      = data{"GloVe-25", "assets/bench/glove-25-angular.hdf5"}
	glove50      = data{"GloVe-50", "assets/bench/glove-50-angular.hdf5"}
	glove100     = data{"GloVe-100", "assets/bench/glove-100-angular.hdf5"}
	glove200     = data{"GloVe-200", "assets/bench/glove-200-angular.hdf5"}
	mnist        = data{"MNIST", "assets/bench/mnist-784-euclidean.hdf5"}
	nytimes      = data{"NYTimes", "assets/bench/nytimes-256-angular.hdf5"}
	sift         = data{"SIFT", "assets/bench/sift-128-euclidean.hdf5"}
)

func BenchmarkFashionMNIST(b *testing.B) {
	benchmarkAll(b, fashionmnist)
}

func BenchmarkGlove25(b *testing.B) {
	benchmarkAll(b, glove25)
}

func BenchmarkGlove50(b *testing.B) {
	benchmarkAll(b, glove50)
}

func BenchmarkGlove100(b *testing.B) {
	benchmarkAll(b, glove100)
}

func BenchmarkGlove200(b *testing.B) {
	benchmarkAll(b, glove200)
}

func BenchmarkMIST(b *testing.B) {
	benchmarkAll(b, mnist)
}

func BenchmarkNYTimes(b *testing.B) {
	benchmarkAll(b, nytimes)
}

func BenchmarkSIFT(b *testing.B) {
	benchmarkAll(b, sift)
}

func benchmarkAll(b *testing.B, d data) {
	benchmarkInsert(b, d)
	benchmarkSearch(b, d)
}

func load(path, name string) ([][]float64, error) {
	f, err := hdf5.OpenFile(path, hdf5.F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dset, err := f.OpenDataset(name)
	if err != nil {
		return nil, err
	}
	defer dset.Close()
	space := dset.Space()
	defer space.Close()
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

func benchmarkInsert(b *testing.B, d data) {
	dataset, err := load(d.path, "train")
	if err != nil {
		b.Error(err)
	}
	b.Run("Insert", func(sb *testing.B) {
		tmpdir, err := ioutil.TempDir("", "tmpdir")
		if err != nil {
			sb.Error(err)
		}
		defer os.RemoveAll(tmpdir)

		n := gongt.New(tmpdir).SetObjectType(gongt.Float).SetDimension(len(dataset[0])).Open()
		defer n.Close()

		sb.ReportAllocs()
		sb.ResetTimer()
		sb.StartTimer()
		for i := 0; i < sb.N; i++ {
			n.Insert(dataset[i%len(dataset)])
		}
		sb.StopTimer()
	})

	b.Run("InsertParallel", func(sb *testing.B) {
		tmpdir, err := ioutil.TempDir("", "tmpdir")
		if err != nil {
			sb.Error(err)
		}
		defer os.RemoveAll(tmpdir)

		n := gongt.New(tmpdir).SetObjectType(gongt.Float).SetDimension(len(dataset[0])).Open()
		defer n.Close()

		sb.ReportAllocs()
		sb.ResetTimer()
		sb.StartTimer()
		sb.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				n.Insert(dataset[i%len(dataset)])
				i++
			}
		})
		sb.StopTimer()
	})
}

func benchmarkSearch(b *testing.B, d data) {
	dataset, err := load(d.path, "test")
	if err != nil {
		b.Error(err)
	}
	path := "assets/bench" + d.name
	n := gongt.New(path).Open()
	defer n.Close()
	size := 10
	b.Run("Search", func(sb *testing.B) {
		sb.ReportAllocs()
		sb.ResetTimer()
		sb.StartTimer()
		for i := 0; i < sb.N; i++ {
			n.Search(dataset[i%len(dataset)], size, gongt.DefaultEpsilon)
		}
		sb.StopTimer()
	})

	b.Run("SearchParallel", func(sb *testing.B) {
		sb.ReportAllocs()
		sb.ResetTimer()
		sb.StartTimer()
		sb.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				n.Search(dataset[i%len(dataset)], size, gongt.DefaultEpsilon)
				i++
			}
		})
		sb.StopTimer()
	})
}
