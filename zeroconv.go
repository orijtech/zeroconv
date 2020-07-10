// Copyright 2020 Orijtech, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zeroconv

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

// LengthPrefixedBytesToString takes a byte slice whose first elements are the Uvarint
// encoded length of the string to be converted.
func LengthPrefixedBytesToString(blob []byte) (str string, err error) {
	size, n := binary.Uvarint(blob)
	if n <= 0 {
		return "", fmt.Errorf("bad length encoding n: %d", n)
	}
	return BytesToString(blob[n : n+int(size)]), nil
}

// BytesToString converts a byte slice into a string without incurring
// the overhead of []byte(string) allocations.
func BytesToString(blob []byte) (str string) {
	// TODO: Perform some string interning as well if deemed necessary.
	bHdr := (*reflect.SliceHeader)(unsafe.Pointer(&blob))
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sHdr.Data = bHdr.Data
	sHdr.Len = bHdr.Len
	return str
}
