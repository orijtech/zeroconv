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
	"runtime"
	"runtime/debug"
	"testing"
)

func TestBytesToString(t *testing.T) {
	src := []byte("this one is that one or those ones")

	var ms runtime.MemStats
	var allocMem uint64
	runtime.GC()

	// Now turn off all garbage collection temporarily.
	defer debug.SetGCPercent(debug.SetGCPercent(0))

	runtime.ReadMemStats(&ms)
	allocMem = ms.TotalAlloc
	got := BytesToString(src)
	runtime.ReadMemStats(&ms)
	afterAllocMem := ms.TotalAlloc
	if got != string(src) {
		t.Fatalf("Mismatch:\nGot:  %q\nWant: %q\n", got, src)
	}
	if allocMem != afterAllocMem {
		t.Fatalf("Unfortunately some allocations were performed, before: %d\n after:  %d",
			allocMem, afterAllocMem)
	}
}

func TestLengthBytesPrefixedToString(t *testing.T) {
	buf := make([]byte, 1000)
	src := "this one is that one or those ones"
	n := binary.PutUvarint(buf, uint64(len(src)))
	copy(buf[n:], src)

	got, err := LengthPrefixedBytesToString(buf)
	if err != nil {
		t.Fatal(err)
	}
	if got != src {
		t.Fatalf("Mismatch:\nGot:  %q\nWant: %q\n", got, src)
	}
}
