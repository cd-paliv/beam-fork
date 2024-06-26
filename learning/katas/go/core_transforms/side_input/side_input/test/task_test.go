// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"testing"

	"beam.apache.org/learning/katas/core_transforms/side_input/side_input/pkg/task"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/testing/passert"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/testing/ptest"
)

func TestApplyTransform(t *testing.T) {
	p, s := beam.NewPipelineWithRoot()

	tests := []struct {
		args struct {
			citiesAndCountriesKV beam.PCollection
			persons              beam.PCollection
		}
		want []interface{}
	}{
		{
			args: struct {
				citiesAndCountriesKV beam.PCollection
				persons              beam.PCollection
			}{
				citiesAndCountriesKV: beam.ParDo(
					s.Scope("Cities and Countries"),
					func(_ []byte, emit func(string, string)) {
						emit("Beijing", "China")
						emit("London", "United Kingdom")
						emit("San Francisco", "United States")
						emit("Singapore", "Singapore")
						emit("Sydney", "Australia")
					},
					beam.Impulse(s)),
				persons: beam.Create(s,
					task.Person{Name: "Henry", City: "Singapore"},
					task.Person{Name: "Jane", City: "San Francisco"},
					task.Person{Name: "Lee", City: "Beijing"},
					task.Person{Name: "John", City: "Sydney"},
					task.Person{Name: "Alfred", City: "London"},
				),
			},
			want: []interface{}{
				task.Person{
					Name:    "Henry",
					City:    "Singapore",
					Country: "Singapore",
				},
				task.Person{
					Name:    "Jane",
					City:    "San Francisco",
					Country: "United States",
				},
				task.Person{
					Name:    "Lee",
					City:    "Beijing",
					Country: "China",
				},
				task.Person{
					Name:    "John",
					City:    "Sydney",
					Country: "Australia",
				},
				task.Person{
					Name:    "Alfred",
					City:    "London",
					Country: "United Kingdom",
				},
			},
		},
	}
	for _, tt := range tests {
		got := task.ApplyTransform(s, tt.args.persons, tt.args.citiesAndCountriesKV)
		passert.Equals(s, got, tt.want...)
		if err := ptest.Run(p); err != nil {
			t.Error(err)
		}
	}
}
