// Copyright 2021 FerretDB Inc.
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

package bson

import (
	"github.com/FerretDB/FerretDB/pkg/util/must"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

	"github.com/AlekSi/pointer"
)

var decimalTestCases = []testCase{{
	name: "1",
	v:    pointer.To(decimalType(must.NotFail(primitive.ParseDecimal128("1")))),
	b:    []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x30},
}, {
	name: "zero",
	v:    pointer.To(decimalType(must.NotFail(primitive.ParseDecimal128("0")))),
	b:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x30},
}}

func TestDecimal(t *testing.T) {
	t.Parallel()
	testBinary(t, decimalTestCases, func() bsontype { return new(decimalType) })
}

func FuzzDecimal(f *testing.F) {
	fuzzBinary(f, decimalTestCases, func() bsontype { return new(decimalType) })
}

func BenchmarkDecimal(b *testing.B) {
	benchmark(b, decimalTestCases, func() bsontype { return new(decimalType) })
}
