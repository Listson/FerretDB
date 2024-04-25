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
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/FerretDB/FerretDB/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/FerretDB/FerretDB/pkg/util/lazyerrors"
)

// decimal holds decimal128 BSON values.
type decimalType types.Decimal

func (i *decimalType) bsontype() {}

// ReadFrom implements bsontype interface.
func (i *decimalType) ReadFrom(r *bufio.Reader) error {
	var h, l uint64

	if err := binary.Read(r, binary.LittleEndian, &l); err != nil {
		return lazyerrors.Errorf("bson.Decimal128.l.ReadFrom (binary.Read): %w", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &h); err != nil {
		return lazyerrors.Errorf("bson.Decimal128.h.ReadFrom (binary.Read): %w", err)
	}

	*i = decimalType(primitive.NewDecimal128(h, l))

	return nil
}

// WriteTo implements bsontype interface.
func (i decimalType) WriteTo(w *bufio.Writer) error {
	v, err := i.MarshalBinary()
	if err != nil {
		return lazyerrors.Errorf("bson.Decimal128.WriteTo: %w", err)
	}

	_, err = w.Write(v)
	if err != nil {
		return lazyerrors.Errorf("bson.Decimal128.WriteTo: %w", err)
	}

	return nil
}

// MarshalBinary implements bsontype interface.
func (i decimalType) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	high, low := primitive.Decimal128(i).GetBytes()

	binary.Write(&buf, binary.LittleEndian, low)
	binary.Write(&buf, binary.LittleEndian, high)

	return buf.Bytes(), nil
}

// check interfaces
var (
	_ bsontype = (*decimalType)(nil)
)
