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

// beam-playground:
//   name: combine-fn
//   description: Combine-fn
//   multifile: false
//   context_line: 37
//   categories:
//     - Quickstart
//   complexity: ADVANCED
//   tags:
//     - hellobeam

package main

import (
	"context"

	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/log"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/x/beamx"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/x/debug"
)

func main() {
	ctx := context.Background()

	p, s := beam.NewPipelineWithRoot()

	input := beam.Create(s, 10, 20, 50, 70, 90)

	output := applyTransform(s, input)

	debug.Print(s, output)

	err := beamx.Run(ctx, p)

	if err != nil {
		log.Exitf(context.Background(), "Failed to execute job: %v", err)
	}
}

func applyTransform(s beam.Scope, input beam.PCollection) beam.PCollection {
	return beam.Combine(s, &averageFn{}, input)
}

type averageAccum struct {
	Count int
	Sum   int
}

type averageFn struct{}

func (c *averageFn) CreateAccumulator() averageAccum {
	return averageAccum{}
}

func (c *averageFn) AddInput(accum averageAccum, input int) averageAccum {
	accum.Count++
	accum.Sum += input
	return accum
}

func (c *averageFn) MergeAccumulators(accumA, accumB averageAccum) averageAccum {
	return averageAccum{
		Count: accumA.Count + accumB.Count,
		Sum:   accumA.Sum + accumB.Sum,
	}
}

func (c *averageFn) ExtractOutput(accum averageAccum) float64 {
	if accum.Count == 0 {
		return 0
	}
	return float64(accum.Sum) / float64(accum.Count)
}
