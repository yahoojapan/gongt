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

// Package gongt provides implementation of Go API for https://github.com/yahoojapan/NGT
package gongt

/*
#cgo LDFLAGS: -lngt
#include <NGT/Capi.h>
*/
import "C"

import (
	"errors"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type (
	// StrictSearchResult is struct with same type in NGT core
	StrictSearchResult struct {
		ID       uint32
		Distance float32
		Error    error
	}
	// SearchResult is struct for comfortable use in Go
	SearchResult struct {
		ID       int
		Distance float64
	}
	// NGT is gongt base struct
	NGT struct {
		prop   Property
		index  C.NGTIndex
		ospace C.NGTObjectSpace
		mu     *sync.RWMutex
		errs   []error
	}
	// Property includes parameters for NGT
	Property struct {
		Dimension           int
		CreationEdgeSize    int
		SearchEdgeSize      int
		ObjectType          ObjectType
		DistanceType        DistanceType
		IndexPath           string
		BulkInsertChunkSize int
	}
)

// ObjectType is alias of object type in NGT
type ObjectType int

// DistanceType is alias of distance type in NGT
type DistanceType int

const (
	// ObjectNone is unknown object type
	ObjectNone ObjectType = iota
	// Uint8 is 8bit unsigned integer
	Uint8
	// Float is 32bit floating point number
	Float

	// DistanceNone is unknown distance type
	DistanceNone DistanceType = iota - 1
	// L1 is l1 norm
	L1
	// L2 is l2 norm
	L2
	// Angle is angle distance
	Angle
	// Hamming is hamming distance
	Hamming
	// Cosine is cosine distance
	Cosine
	// NormalizedAngle is angle distance with normalization
	NormalizedAngle
	// NormalizedCosine is cosine distance with normalization
	NormalizedCosine

	// DefaultDimension is 0
	DefaultDimension = 0
	// DefaultCreationEdgeSize is 10
	DefaultCreationEdgeSize = 10
	// DefaultSearchEdgeSize is 10
	DefaultSearchEdgeSize = 40
	// DefaultObjectType is Float
	DefaultObjectType = Float
	// DefaultDistanceType is L2
	DefaultDistanceType = L2
	// DefaultEpsilon is 0.01
	DefaultEpsilon = 0.01
	// DefaultBulkInsertChunkSize is 100
	DefaultBulkInsertChunkSize = 100
	// DefaultPoolSize is 1
	DefaultPoolSize = 1

	// ErrorCode is false
	ErrorCode = C._Bool(false)
)

var (
	once = &sync.Once{}
	ngt  *NGT

	// ErrCAPINotImplemented raises using not implemented function in C API
	ErrCAPINotImplemented = errors.New("Not implemented in C API")
)

func init() {
	Get()
}

func newGoError(err C.NGTError) error {
	return errors.New(C.GoString(C.ngt_get_error_string(err)))
}

// Get returns singleton instance NGT
//	ngt := gongt.Get()
func Get() *NGT {
	once.Do(func() {
		ngt = New("/tmp/ngt-" + time.Now().Format(time.RFC3339))
	})
	return ngt
}

// New returns NGT instance
//	ngt := gongt.New("index Path")
func New(indexPath string) *NGT {
	return &NGT{
		mu: &sync.RWMutex{},
		prop: Property{
			BulkInsertChunkSize: DefaultBulkInsertChunkSize,
			CreationEdgeSize:    DefaultCreationEdgeSize,
			Dimension:           DefaultDimension,
			DistanceType:        DefaultDistanceType,
			IndexPath:           indexPath,
			ObjectType:          DefaultObjectType,
			SearchEdgeSize:      DefaultSearchEdgeSize,
		},
	}
}

// GetDim returns NGT dimension
//	dimension := gongt.GetDim()
func GetDim() int {
	return ngt.GetDim()
}

// GetDim returns NGT dimension
//	dimension := gongt.Get().GetDim()
//	dimension := gongt.New("Index Path").GetDim()
func (n NGT) GetDim() int {
	return n.prop.Dimension
}

// GetPath returns path to index directory
//	indexPath := gongt.GetPath()
func GetPath() string {
	return ngt.GetPath()
}

// GetPath returns path to index directory
//	indexPath := gongt.Get().GetPath()
//	indexPath := gongt.New("index path").GetPath()
func (n NGT) GetPath() string {
	return n.prop.IndexPath
}

// SetIndexPath sets path to index directory
//	gongt.SetIndexPath("index Path")
func SetIndexPath(path string) *NGT {
	return ngt.SetIndexPath(path)
}

// SetIndexPath sets path to index directory
//	gongt.Get().SetIndexPath("index Path")
//	gongt.New("").SetIndexPath("index Path")
func (n *NGT) SetIndexPath(path string) *NGT {
	if path != "" {
		n.mu.Lock()
		n.prop.IndexPath = path
		n.mu.Unlock()
	}
	return n
}

// SetDimension sets NGT feature dimension
//	gongt.SetDimension(10) // Dimension Setting
func SetDimension(dimension int) *NGT {
	return ngt.SetDimension(dimension)
}

// SetDimension sets NGT feature dimension
//	gongt.Get().SetDimension(10) // Dimension Setting
//	gongt.New("Index Path").SetDimension(10) // Dimension Setting
func (n *NGT) SetDimension(dimension int) *NGT {
	if dimension > 0 {
		n.mu.Lock()
		n.prop.Dimension = dimension
		n.mu.Unlock()
	}
	return n
}

// SetCreationEdgeSize sets creation edge size
//	gongt.SetCreationEdgeSize(10) // CreationEdgeSize Setting
func SetCreationEdgeSize(size int) *NGT {
	return ngt.SetCreationEdgeSize(size)
}

// SetCreationEdgeSize sets creation edge size
//	gongt.Get().SetCreationEdgeSize(10) // CreationEdgeSize Setting
//	gongt.New("").SetCreationEdgeSize(10) // CreationEdgeSize Setting
func (n *NGT) SetCreationEdgeSize(size int) *NGT {
	if size > 0 {
		n.mu.Lock()
		n.prop.CreationEdgeSize = size
		n.mu.Unlock()
	}
	return n
}

// SetSearchEdgeSize sets search edge size
//	gongt.SetSearchEdgeSize(10) // SearchEdgeSize Setting
func SetSearchEdgeSize(size int) *NGT {
	return ngt.SetSearchEdgeSize(size)
}

// SetSearchEdgeSize sets search edge size
//	gongt.Get().SetSearchEdgeSize(10) // SearchEdgeSize Setting
//	gongt.New("").SetSearchEdgeSize(10) // SearchEdgeSize Setting
func (n *NGT) SetSearchEdgeSize(size int) *NGT {
	if size >= 0 {
		n.mu.Lock()
		n.prop.SearchEdgeSize = size
		n.mu.Unlock()
	}
	return n
}

// SetObjectType sets object type
//	gongt.SetObjectType(gongt.Float) // ObjectType Setting
//	gongt.SetObjectType(gongt.Uint8) // ObjectType Setting
//	gongt.SetObjectType(gongt.ObjectNone) // ObjectType Setting
func SetObjectType(ot ObjectType) *NGT {
	return ngt.SetObjectType(ot)
}

// SetObjectType sets object type
//	gongt.Get().SetObjectType(gongt.Float) // ObjectType Setting
//	gongt.Get().SetObjectType(gongt.Uint8) // ObjectType Setting
//	gongt.Get().SetObjectType(gongt.ObjectNone) // ObjectType Setting
//	gongt.New("").SetObjectType(gongt.Float) // ObjectType Setting
//	gongt.New("").SetObjectType(gongt.Uint8) // ObjectType Setting
//	gongt.New("").SetObjectType(gongt.ObjectNone) // ObjectType Setting
func (n *NGT) SetObjectType(ot ObjectType) *NGT {
	n.mu.Lock()
	n.prop.ObjectType = ot
	n.mu.Unlock()

	return n
}

// SetDistanceType sets distanc
//	gongt.SetDistanceType(gongt.L1) // DistanceType Setting
//	gongt.SetDistanceType(gongt.L2) // DistanceType Setting
//	gongt.SetDistanceType(gongt.Hamming) // DistanceType Setting
func SetDistanceType(dt DistanceType) *NGT {
	return ngt.SetDistanceType(dt)
}

// SetDistanceType sets distance type
//	gongt.Get().SetDistanceType(gongt.L1) // DistanceType Setting
//	gongt.Get().SetDistanceType(gongt.L2) // DistanceType Setting
//	gongt.Get().SetDistanceType(gongt.Hamming) // DistanceType Setting
//	gongt.New("").SetDistanceType(gongt.L1) // DistanceType Setting
//	gongt.New("").SetDistanceType(gongt.L2) // DistanceType Setting
//	gongt.New("").SetDistanceType(gongt.Hamming) // DistanceType Setting
func (n *NGT) SetDistanceType(dt DistanceType) *NGT {
	n.mu.Lock()
	n.prop.DistanceType = dt
	n.mu.Unlock()
	return n
}

// SetBulkInsertChunkSize sets insert chunk size
func SetBulkInsertChunkSize(size int) *NGT {
	return ngt.SetBulkInsertChunkSize(size)
}

// SetBulkInsertChunkSize sets insert chunk size
func (n *NGT) SetBulkInsertChunkSize(size int) *NGT {
	n.mu.Lock()
	n.prop.BulkInsertChunkSize = size
	n.mu.Unlock()

	return n
}

// Open configures using Property and returns NGT instance
func Open() *NGT {
	return ngt.Open()
}

// Open configures using Property and returns NGT instance
func (n *NGT) Open() *NGT {
	n.mu.Lock()
	defer n.mu.Unlock()

	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	prop := C.ngt_create_property(ebuf)
	if prop == nil {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}
	defer C.ngt_destroy_property(prop)
	if C.ngt_set_property_dimension(prop, C.int32_t(n.prop.Dimension), ebuf) == ErrorCode {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}
	if C.ngt_set_property_edge_size_for_creation(prop, C.int16_t(n.prop.CreationEdgeSize), ebuf) == ErrorCode {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}
	if C.ngt_set_property_edge_size_for_search(prop, C.int16_t(n.prop.SearchEdgeSize), ebuf) == ErrorCode {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}

	switch n.prop.ObjectType {
	case Uint8:
		if C.ngt_set_property_object_type_integer(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	case Float:
		if C.ngt_set_property_object_type_float(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	default:
		n.errs = append(n.errs, errors.New("Illegal object type"))
		return n
	}

	switch n.prop.DistanceType {
	case L1:
		if C.ngt_set_property_distance_type_l1(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	case L2:
		if C.ngt_set_property_distance_type_l2(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	case Angle:
		if C.ngt_set_property_distance_type_angle(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	case Hamming:
		if C.ngt_set_property_distance_type_hamming(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	case Cosine:
		if C.ngt_set_property_distance_type_cosine(prop, ebuf) == ErrorCode {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	case NormalizedAngle:
		// TODO: not implemented in C API
		n.errs = append(n.errs, ErrCAPINotImplemented)
		return n
	case NormalizedCosine:
		// TODO: not implemented in C API
		n.errs = append(n.errs, ErrCAPINotImplemented)
		return n
	default:
		n.errs = append(n.errs, errors.New("Illegal distance type"))
		return n
	}

	n.index = C.ngt_open_index(C.CString(n.prop.IndexPath), ebuf)
	if n.index == nil {
		err := newGoError(ebuf)
		if strings.Contains(err.Error(), "PropertySet::load: Cannot load the property file ") || strings.Contains(err.Error(), "PropertSet::load: Cannot load the property file ") {
			n.index = C.ngt_create_graph_and_tree(C.CString(n.prop.IndexPath), prop, ebuf)
			if n.index == nil {
				n.errs = append(n.errs, newGoError(ebuf))
				return n
			}
			if C.ngt_save_index(n.index, C.CString(n.prop.IndexPath), ebuf) == ErrorCode {
				n.errs = append(n.errs, newGoError(ebuf))
				return n
			}
		} else {
			n.errs = append(n.errs, newGoError(ebuf))
			return n
		}
	}

	if C.ngt_get_property(n.index, prop, ebuf) == ErrorCode {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}
	n.prop.Dimension = int(C.ngt_get_property_dimension(prop, ebuf))
	if n.prop.Dimension == -1 {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}
	n.prop.ObjectType = ObjectType(C.ngt_get_property_object_type(prop, ebuf))
	if n.prop.ObjectType == -1 {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}

	n.ospace = C.ngt_get_object_space(n.index, ebuf)
	if n.ospace == nil {
		n.errs = append(n.errs, newGoError(ebuf))
		return n
	}

	return n
}

// StrictSearch is C type stricted search function
func StrictSearch(vec []float64, size int, epsilon, radius float32) ([]StrictSearchResult, error) {
	return ngt.StrictSearch(vec, size, epsilon, radius)
}

// StrictSearch is C type stricted search function
func (n *NGT) StrictSearch(vec []float64, size int, epsilon, radius float32) ([]StrictSearchResult, error) {
	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	results := C.ngt_create_empty_results(ebuf)
	defer C.ngt_destroy_results(results)
	if results == nil {
		return nil, newGoError(ebuf)
	}

	n.mu.RLock()
	ret := C.ngt_search_index(n.index, (*C.double)(&vec[0]), C.int32_t(n.prop.Dimension), C.size_t(size), C.float(epsilon), C.float(radius), results, ebuf)
	n.mu.RUnlock()
	if ret == ErrorCode {
		return nil, newGoError(ebuf)
	}
	rsize := int(C.ngt_get_size(results, ebuf))
	if rsize == -1 {
		return nil, newGoError(ebuf)
	}
	result := make([]StrictSearchResult, rsize)
	for i := 0; i < rsize; i++ {
		d := C.ngt_get_result(results, C.uint32_t(i), ebuf)
		if d.id == 0 && d.distance == 0 {
			result[i] = StrictSearchResult{0, 0, newGoError(ebuf)}
		} else {
			result[i] = StrictSearchResult{uint32(d.id), float32(d.distance), nil}
		}
	}

	return result, nil
}

// Search returns search result as []SearchResult
func Search(vec []float64, size int, epsilon float64) ([]SearchResult, error) {
	return ngt.Search(vec, size, epsilon)
}

// Search returns search result as []SearchResult
func (n *NGT) Search(vec []float64, size int, epsilon float64) ([]SearchResult, error) {
	res, err := n.StrictSearch(vec, size, float32(epsilon), -1.0)
	if err != nil {
		return nil, err
	}
	idx := 0
	result := make([]SearchResult, len(res))
	for _, val := range res {
		if val.Error == nil {
			result[idx] = SearchResult{int(val.ID), float64(val.Distance)}
			idx++
		}
	}
	return result[:idx], nil
}

// StrictInsert is C type stricted insert function
func StrictInsert(vec []float64) (uint, error) {
	return ngt.StrictInsert(vec)
}

// StrictInsert is C type stricted insert function
func (n *NGT) StrictInsert(vec []float64) (uint, error) {
	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	n.mu.Lock()
	id := C.ngt_insert_index(n.index, (*C.double)(&vec[0]), C.uint32_t(n.prop.Dimension), ebuf)
	n.mu.Unlock()
	if id == 0 {
		err := newGoError(ebuf)
		n.errs = append(n.errs, err)
		return 0, err
	}

	return uint(id), nil
}

// Insert returns NGT object id.
// This only stores not indexing, must execute CreateIndex and SaveIndex.
func Insert(vec []float64) (int, error) {
	return ngt.Insert(vec)
}

// Insert returns NGT object id.
// This only stores not indexing, you must call CreateIndex and SaveIndex.
func (n *NGT) Insert(vec []float64) (int, error) {
	id, err := n.StrictInsert(vec)
	return int(id), err
}

// InsertCommit returns NGT object id.
// This stores and indexes at the same time.
func InsertCommit(vec []float64, poolSize int) (int, error) {
	return ngt.InsertCommit(vec, poolSize)
}

// InsertCommit returns NGT object id.
// This stores and indexes at the same time.
func (n *NGT) InsertCommit(vec []float64, poolSize int) (int, error) {
	id, err := n.StrictInsert(vec)
	if err != nil {
		return int(id), err
	}

	err = n.CreateIndex(poolSize)
	if err != nil {
		return int(id), err
	}

	err = n.SaveIndex()
	if err != nil {
		return int(id), err
	}

	return int(id), nil
}

// BulkInsert returns NGT object ids.
// This only stores not indexing, you must call CreateIndex and SaveIndex.
func BulkInsert(vecs [][]float64) ([]int, []error) {
	return ngt.BulkInsert(vecs)
}

// BulkInsert returns NGT object ids.
// This only stores not indexing, you must call CreateIndex and SaveIndex.
func (n *NGT) BulkInsert(vecs [][]float64) ([]int, []error) {
	ids := make([]int, 0, len(vecs))
	errs := make([]error, 0, len(vecs))

	var id int
	var err error

	for _, vec := range vecs {
		if id, err = n.Insert(vec); err == nil {
			ids = append(ids, id)
		} else {
			errs = append(errs, err)
		}
	}

	return ids, errs
}

// BulkInsertCommit returns NGT object ids.
// This stores and indexes at the same time.
func BulkInsertCommit(vecs [][]float64, poolSize int) ([]int, []error) {
	return ngt.BulkInsertCommit(vecs, poolSize)
}

// BulkInsertCommit returns NGT object ids.
// This stores and indexes at the same time.
func (n *NGT) BulkInsertCommit(vecs [][]float64, poolSize int) ([]int, []error) {
	ids := make([]int, 0, len(vecs))
	errs := make([]error, 0, len(vecs))
	idx := 0
	var id int
	var err error
	for _, vec := range vecs {
		if id, err = n.Insert(vec); err == nil {
			ids = append(ids, id)
			idx++
			if idx >= n.prop.BulkInsertChunkSize {
				err = n.CreateAndSaveIndex(poolSize)
				if err != nil {
					errs = append(errs, err)
				}
				idx = 0
			}
		} else {
			errs = append(errs, err)
		}
	}
	err = n.CreateAndSaveIndex(poolSize)
	if err != nil {
		errs = append(errs, err)
	}
	return ids, errs
}

// CreateAndSaveIndex call  CreateIndex and SaveIndex in a row.
func CreateAndSaveIndex(poolSize int) error {
	return ngt.CreateAndSaveIndex(poolSize)
}

// CreateAndSaveIndex call  CreateIndex and SaveIndex in a row.
func (n *NGT) CreateAndSaveIndex(poolSize int) error {
	err := n.CreateIndex(poolSize)
	if err != nil {
		return err
	}
	return n.SaveIndex()
}

// CreateIndex creates NGT index.
func CreateIndex(poolSize int) error {
	return ngt.CreateIndex(poolSize)
}

// CreateIndex creates NGT index.
func (n *NGT) CreateIndex(poolSize int) error {
	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	n.mu.Lock()
	ret := C.ngt_create_index(n.index, C.uint32_t(poolSize), ebuf)
	n.mu.Unlock()
	if ret == ErrorCode {
		err := newGoError(ebuf)
		n.errs = append(n.errs, err)
		return err
	}

	return nil
}

// SaveIndex stores NGT index to storage.
func SaveIndex() error {
	return ngt.SaveIndex()
}

// SaveIndex stores NGT index to storage.
func (n *NGT) SaveIndex() error {
	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	n.mu.RLock()
	ret := C.ngt_save_index(n.index, C.CString(n.prop.IndexPath), ebuf)
	n.mu.RUnlock()

	if ret == ErrorCode {
		err := newGoError(ebuf)
		n.errs = append(n.errs, err)
		return err
	}

	return nil
}

// StrictRemove is C type stricted remove function
func StrictRemove(id uint) error {
	return ngt.StrictRemove(id)
}

// StrictRemove is C type stricted remove function
func (n *NGT) StrictRemove(id uint) error {
	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	n.mu.Lock()
	ret := C.ngt_remove_index(n.index, C.ObjectID(id), ebuf)
	n.mu.Unlock()
	if ret == ErrorCode {
		err := newGoError(ebuf)
		n.errs = append(n.errs, err)
		return err
	}

	return nil
}

// Remove removes from NGT index.
func Remove(id int) error {
	return ngt.Remove(id)
}

// Remove removes from NGT index.
func (n *NGT) Remove(id int) error {
	return n.StrictRemove(uint(id))
}

// GetStrictVector is C type stricted GetVector function.
func GetStrictVector(id uint) ([]float32, error) {
	return ngt.GetStrictVector(id)
}

// GetStrictVector is C type stricted GetVector function.
func (n *NGT) GetStrictVector(id uint) ([]float32, error) {
	ebuf := C.ngt_create_error_object()
	defer C.ngt_destroy_error_object(ebuf)

	ret := make([]float32, n.prop.Dimension)
	switch n.prop.ObjectType {
	case Float:
		n.mu.RLock()
		results := C.ngt_get_object_as_float(n.ospace, C.ObjectID(id), ebuf)
		n.mu.RUnlock()
		if results == nil {
			err := newGoError(ebuf)
			n.errs = append(n.errs, err)
			return nil, err
		}
		slice := (*[1 << 30]C.float)(unsafe.Pointer(results))[:n.prop.Dimension:n.prop.Dimension]
		for i := 0; i < n.prop.Dimension; i++ {
			ret[i] = float32(slice[i])
		}
	case Uint8:
		n.mu.RLock()
		results := C.ngt_get_object_as_integer(n.ospace, C.ObjectID(id), ebuf)
		n.mu.RUnlock()
		if results == nil {
			err := newGoError(ebuf)
			n.errs = append(n.errs, err)
			return nil, err
		}
		slice := (*[1 << 30]C.uchar)(unsafe.Pointer(results))[:n.prop.Dimension:n.prop.Dimension]
		for i := 0; i < n.prop.Dimension; i++ {
			ret[i] = float32(slice[i])
		}
	default:
		err := errors.New("Unsupported ObjectType")
		n.errs = append(n.errs, err)
		return nil, err
	}
	return ret, nil
}

// GetVector returns vector stored in NGT index.
func GetVector(id int) ([]float64, error) {
	return ngt.GetVector(id)
}

// GetVector returns vector stored in NGT index.
func (n *NGT) GetVector(id int) ([]float64, error) {
	v, err := n.GetStrictVector(uint(id))
	if err != nil {
		return nil, err
	}

	ret := make([]float64, len(v))
	for i, e := range v {
		ret[i] = float64(e)
	}
	return ret, nil
}

// Close NGT index.
func Close() {
	if ngt != nil {
		ngt.Close()
	}
}

// Close NGT index.
func (n *NGT) Close() {
	if n.index != nil {
		C.ngt_close_index(n.index)
		n.index = nil
	}
}

// GetErrors returns errors
func GetErrors() []error {
	return ngt.errs
}

// GetErrors returns errors
func (n *NGT) GetErrors() []error {
	return n.errs
}
