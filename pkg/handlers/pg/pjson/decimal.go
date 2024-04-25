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

package pjson

import (
	"fmt"
	"github.com/FerretDB/FerretDB/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// decimalType represents BSON decimal128 data type.
type decimalType types.Decimal

func (bin *decimalType) pjsontype() {}

// MarshalJSON implements fjsontype interface.
func (bin *decimalType) MarshalJSON() ([]byte, error) {
	res := fmt.Sprintf(`{"$numberDecimal":"%s"}`, primitive.Decimal128(*bin).String())
	return []byte(res), nil
}

// check interfaces
var (
	_ pjsontype = (*decimalType)(nil)
)
