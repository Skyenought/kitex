// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"os"
	"testing"

	"github.com/cloudwego/kitex/internal/test"
)

func TestCombineOutputPath(t *testing.T) {
	ns := "aaa.bbb.ccc"
	path1 := "kitex_path/code"
	output1 := CombineOutputPath(path1, ns)
	test.Assert(t, output1 == "kitex_path/code/aaa/bbb/ccc")
	path2 := "kitex_path/{namespace}/code"
	output2 := CombineOutputPath(path2, ns)
	test.Assert(t, output2 == "kitex_path/aaa/bbb/ccc/code")
	path3 := "kitex_path/{namespaceUnderscore}/code"
	output3 := CombineOutputPath(path3, ns)
	test.Assert(t, output3 == "kitex_path/aaa_bbb_ccc/code")
}

func TestGetGOPATH(t *testing.T) {
	orig := os.Getenv("GOPATH")
	defer func() {
		os.Setenv("GOPATH", orig)
	}()

	os.Setenv("GOPATH", "/usr/bin/go:/usr/local/bin/go")
	test.Assert(t, GetGOPATH() == "/usr/bin/go")
	os.Setenv("GOPATH", "")
	test.Assert(t, GetGOPATH() != "")
}

func TestIDLName(t *testing.T) {
	tests := []struct {
		filename string
		want     string
	}{
		{"test.thrift", "test"},
		{"Test.thrift", "test"},
		{"testP.thrift", "test_p"},
		{"../../test_p.thrift", "test_p"},
		{"C:\\\\Users\\Username\\Documents\\file.txt", "file"},
		{"a.b.c.thrift", "a_b_c"},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := IDLName(tt.filename); got != tt.want {
				t.Errorf("IDLName() = %v, want %v", got, tt.want)
			}
		})
	}
}
