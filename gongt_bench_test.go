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
	"os"
	"io/ioutil"

	"gonum.org/v1/hdf5"
	
	"testing"

	"github.com/yahoojapan/gongt"
)

type data struct {
	name string
	path string
}

var (
	fashionmnist = data{"Fashion-MNIST","assets/bench/fashion-mnist-784-euclidean.hdf5"}
	glove25 = data{"GloVe-25", "assets/bench/glove-25-angular.hdf5"}
	glove50 = data{"GloVe-50", "assets/bench/glove-50-angular.hdf5"}
	glove100 = data{"GloVe-100", "assets/bench/glove-100-angular.hdf5"}
	glove200 = data{"GloVe-200", "assets/bench/glove-200-angular.hdf5"}
	mnist =	data{"MNIST", "assets/bench/mnist-784-euclidean.hdf5"}
	nytimes = data{"NYTimes", "assets/bench/nytimes-256-angular.hdf5"}
	sift = data{"SIFT", "assets/bench/sift-128-euclidean.hdf5"}
)

func load(d data, name string) ([][]float64, error) {
	f, err := hdf5.OpenFile(d.path, hdf5.F_ACC_RDONLY)
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
      vec[i][j] = float64(v[i * col + j])
    }
  }
	return vec, nil
}

func benchmarkInsert(b *testing.B, d data) {
	dataset, err := load(d, "train")
	if err != nil {
		b.Error(err)
	}

	tmpdir, err := ioutil.TempDir("", "tmpdir")
	if err != nil {
		b.Error(err)
	}
	defer os.RemoveAll(tmpdir)
	
	n := gongt.New(tmpdir).SetObjectType(gongt.Float).SetDimension(len(dataset[0])).Open()
	defer n.Close()
	
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n.Insert(dataset[i % len(dataset)])
	}
	b.StopTimer()
}

func BenchmarkInsertFashionMNIST(b *testing.B) {
	benchmarkInsert(b, fashionmnist)
}

func BenchmarkInsertGlove25(b *testing.B) {
	benchmarkInsert(b, glove25)
}

func BenchmarkInsertGlove50(b *testing.B) {
	benchmarkInsert(b, glove50)
}

func BenchmarkInsertGlove100(b *testing.B) {
	benchmarkInsert(b, glove100)
}

func BenchmarkInsertGlove200(b *testing.B) {
	benchmarkInsert(b, glove200)
}

func BenchmarkInsertMNIST(b *testing.B) {
	benchmarkInsert(b, mnist)
}

func BenchmarkInsertNYTimes(b *testing.B) {
	benchmarkInsert(b, nytimes)
}

func BenchmarkInsertSIFT(b *testing.B) {
	benchmarkInsert(b, sift)
}

func benchmarkSearch(b *testing.B, d data) {
	dataset, err := load(d, "test")
	if err != nil {
		b.Error(err)
	}

	path := "assets/bench"+d.name
	n := gongt.New(path).Open()
	defer n.Close()
	size := 10

	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n.Search(dataset[i % len(dataset)], size, gongt.DefaultEpsilon)
	}
	b.StopTimer()
}

func BenchmarkSearchFashionMNIST(b *testing.B) {
	benchmarkSearch(b, fashionmnist)
}

func BenchmarkSearchGlove25(b *testing.B) {
	benchmarkSearch(b, glove25)
}

func BenchmarkSearchGlove50(b *testing.B) {
	benchmarkSearch(b, glove50)
}

func BenchmarkSearchGlove100(b *testing.B) {
	benchmarkSearch(b, glove100)
}

func BenchmarkSearchGlove200(b *testing.B) {
	benchmarkSearch(b, glove200)
}

func BenchmarkSearchMNIST(b *testing.B) {
	benchmarkSearch(b, mnist)
}

func BenchmarkSearchNYTimes(b *testing.B) {
	benchmarkSearch(b, nytimes)
}

func BenchmarkSearchSIFT(b *testing.B) {
	benchmarkSearch(b, sift)
}
