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
	"fmt"

	"github.com/yahoojapan/gongt"
)

func ExampleGet() {
	// Fetch Singleton GoNGT instance
	ngt := gongt.Get()
	// Output:
	//
	_ = ngt
}

func ExampleNew() {
	// Instantiate GoNGT
	ngt := gongt.New("assets/example")
	// Output:
	//
	_ = ngt

}

func ExampleGetDim() {
	// Fetch Dimension Size
	gongt.SetIndexPath("assets/example").Open()
	dim := gongt.GetDim()
	fmt.Println(dim)
	// Output:
	// 128
}

func ExampleNGT_GetDim() {
	// Fetch Dimension Size
	gongt.SetIndexPath("assets/example").Open()
	dim := gongt.Get().GetDim()
	fmt.Println(dim)
	// Output:
	// 128
}

func ExampleGetPath() {
	// Fetch Path Location
	gongt.SetIndexPath("assets/example").Open()
	path := gongt.GetPath()
	fmt.Println(path)
	// Output:
	// assets/example
}

func ExampleNGT_GetPath() {
	// Fetch Path Location
	gongt.SetIndexPath("assets/example").Open()
	path := gongt.Get().GetPath()
	fmt.Println(path)
	// Output:
	// assets/example
}

func ExampleSetIndexPath() {
	// Fetch Path Location
	gongt.SetIndexPath("/tmp/index-path")
	// Output:
	//
}

func ExampleNGT_SetIndexPath() {
	// Fetch Path Location
	gongt.Get().SetIndexPath("/tmp/index-path")
	// Output:
	//
}

func ExampleSetDimension() {
	// Set Dimension
	gongt.SetDimension(128)
	// Output:
	//
}

func ExampleNGT_SetDimension() {
	// Set Dimension
	gongt.Get().SetDimension(128)
	// Output:
	//
}

func ExampleSetCreationEdgeSize() {
	// Set Creation Edge Size
	gongt.SetCreationEdgeSize(30)
	// Output:
	//
}

func ExampleNGT_SetCreationEdgeSize() {
	// Set Creation Edge Size
	gongt.Get().SetCreationEdgeSize(30)
	// Output:
	//

}

func ExampleSetSearchEdgeSize() {
	// Set Search Edge Size
	gongt.SetSearchEdgeSize(10)
	// Output:
	//
}

func ExampleNGT_SetSearchEdgeSize() {
	// Set Search Edge Size
	gongt.Get().SetSearchEdgeSize(10)
	// Output:
	//
}

func ExampleSetObjectType() {
	// Set Object Type
	// gongt.SetObjectType(gongt.Uint8) // ObjectType Setting
	// gongt.SetObjectType(gongt.ObjectNone) // ObjectType Setting
	gongt.SetObjectType(gongt.Float)
	// Output:
	//
}

func ExampleNGT_SetObjectType() {
	// Set Object Type
	// gongt.Get().SetObjectType(gongt.Uint8) // ObjectType Setting
	// gongt.Get().SetObjectType(gongt.ObjectNone) // ObjectType Setting
	gongt.Get().SetObjectType(gongt.Float) // ObjectType Setting
	// Output:
	//
}

func ExampleSetDistanceType() {
	// Set Distance Type
	gongt.SetDistanceType(gongt.L2)
	// Output:
	//
}

func ExampleNGT_SetDistanceType() {
	// Set Distance Type
	gongt.Get().SetDistanceType(gongt.Hamming)
	// Output:
	//
}

func ExampleSetBulkInsertChunkSize() {
	// Set Bulk Insert Chunk Size
	gongt.SetBulkInsertChunkSize(5)
	// Output:
	//
}

func ExampleNGT_SetBulkInsertChunkSize() {
	// Set Bulk Insert Chunk Size
	gongt.Get().SetBulkInsertChunkSize(5)
	// Output:
	//
}

func ExampleOpen() {
	// Set Bulk Insert Chunk Size
	ngt := gongt.Open()
	// Output:
	//
	_ = ngt
}

func ExampleNGT_Open() {
	// Set Bulk Insert Chunk Size
	ngt := gongt.Get().Open()
	// Output:
	//
	_ = ngt
}

func ExampleStrictSearch() {
	// Strict Vector Search
	vector := []float64{1, 0, 0, 0, 0, 0}
	res, err := gongt.StrictSearch(vector, 1, gongt.DefaultEpsilon)
	// Output:
	_, _ = res, err
}

func ExampleNGT_StrictSearch() {
	// Strict Vector Search
	vector := []float64{1, 0, 0, 0, 0, 0}
	res, err := gongt.Get().StrictSearch(vector, 1, gongt.DefaultEpsilon)
	// Output:
	//
	_, _ = res, err
}

func ExampleSearch() {
	// Vector Search
	vector := []float64{1, 0, 0, 0, 0, 0}
	res, err := gongt.Search(vector, 1, gongt.DefaultEpsilon)
	// Output:
	//
	_, _ = res, err
}

func ExampleNGT_Search() {
	// Vector Search
	vector := []float64{1, 0, 0, 0, 0, 0}
	res, err := gongt.Get().Search(vector, 1, gongt.DefaultEpsilon)
	// Output:
	//
	_, _ = res, err
}

func ExampleStrictInsert() {
	// Strict Vector Insert
	vector := []float64{1, 0, 0, 0, 0, 0}
	id, err := gongt.StrictInsert(vector)
	// Output:
	//
	_, _ = id, err
}

func ExampleNGT_StrictInsert() {
	// Strict Vector Insert
	vector := []float64{1, 0, 0, 0, 0, 0}
	id, err := gongt.Get().StrictInsert(vector)
	// Output:
	//
	_, _ = id, err
}

func ExampleInsert() {
	// Vector Insert
	vector := []float64{1, 0, 0, 0, 0, 0}
	id, err := gongt.Insert(vector)
	// Output:
	//
	_, _ = id, err
}

func ExampleNGT_Insert() {
	// Vector Insert
	vector := []float64{1, 0, 0, 0, 0, 0}
	id, err := gongt.Get().Insert(vector)
	// Output:
	//
	_, _ = id, err
}

func ExampleInsertCommit() {
	// Vector Insert
	vector := []float64{1, 0, 0, 0, 0, 0}
	id, err := gongt.InsertCommit(vector, 10)
	// Output:
	//
	_, _ = id, err
}

func ExampleNGT_InsertCommit() {
	// Vector Insert
	vector := []float64{1, 0, 0, 0, 0, 0}
	id, err := gongt.Get().InsertCommit(vector, 10)
	// Output:
	//
	_, _ = id, err
}

func ExampleBulkInsert() {
	// Vector Bulk Insert
	vectors := [][]float64{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0, 0},
	}
	ids, errs := gongt.BulkInsert(vectors)
	// Output:
	//
	_, _ = ids, errs
}

func ExampleNGT_BulkInsert() {
	// Vector Bulk Insert
	vectors := [][]float64{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0, 0},
	}
	ids, errs := gongt.Get().BulkInsert(vectors)
	// Output:
	//
	_, _ = ids, errs
}

func ExampleBulkInsertCommit() {
	// Vector Bulk Insert And Commit
	vectors := [][]float64{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0, 0},
	}
	ids, errs := gongt.BulkInsertCommit(vectors, gongt.DefaultPoolSize)
	// Output:
	//
	_, _ = ids, errs
}

func ExampleNGT_BulkInsertCommit() {
	// Vector Bulk Insert And Commit
	vectors := [][]float64{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{1, 1, 0, 0, 0, 0},
	}
	ids, errs := gongt.Get().BulkInsertCommit(vectors, gongt.DefaultPoolSize)
	// Output:
	//
	_, _ = ids, errs
}

func ExampleCreateAndSaveIndex() {
	// Create And Save Index
	gongt.CreateAndSaveIndex(10)
	// Output:
	//
}

func ExampleNGT_CreateAndSaveIndex() {
	// Create And Save Index
	gongt.Get().CreateAndSaveIndex(10)
	// Output:
	//
}

func ExampleCreateIndex() {
	// Create Index
	gongt.CreateIndex(10)
	// Output:
	//
}

func ExampleNGT_CreateIndex() {
	// Create Index
	gongt.Get().CreateIndex(10)
	// Output:
	//
}

func ExampleSaveIndex() {
	// Save Index
	gongt.SaveIndex()
	// Output:
	//
}

func ExampleNGT_SaveIndex() {
	// Save Index
	gongt.Get().SaveIndex()
	// Output:
	//
}

func ExampleStrictRemove() {
	// Remove Vector
	gongt.StrictRemove(8)
	// Output:
	//
}

func ExampleNGT_StrictRemove() {
	// Remove Vector
	gongt.Get().StrictRemove(8)
	// Output:
	//
}

func ExampleRemove() {
	// Remove Vector
	gongt.Remove(8)
	// Output:
	//
}

func ExampleNGT_Remove() {
	// Remove Vector
	gongt.Get().Remove(8)
	// Output:
	//
}

func ExampleGetStrictVector() {
	// Get Vector
	vec, err := gongt.GetStrictVector(1)
	// Output:
	//
	_, _ = vec, err
}

func ExampleNGT_GetStrictVector() {
	// Get Vector
	vec, err := gongt.Get().GetStrictVector(1)
	// Output:
	//
	_, _ = vec, err
}

func ExampleGetVector() {
	// Get Vector
	vec, err := gongt.GetVector(1)
	// Output:
	//
	_, _ = vec, err
}

func ExampleNGT_GetVector() {
	// Get Vector
	vec, err := gongt.Get().GetVector(1)
	// Output:
	//
	_, _ = vec, err
}

func ExampleClose() {
	// Close NGT
	gongt.Close()
	// Output:
	//
}

func ExampleNGT_Close() {
	// Close NGT
	gongt.Get().Close()
	// Output:
	//
}

func ExampleGetErrors() {
	// Close NGT
	errs := gongt.GetErrors()
	// Output:
	//
	_ = errs
}

func ExampleNGT_GetErrors() {
	// Close NGT
	errs := gongt.Get().GetErrors()
	// Output:
	//
	_ = errs
}
