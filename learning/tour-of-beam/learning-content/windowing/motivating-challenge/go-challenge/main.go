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
//   name: WindowingChallenge
//   description: Windowing motivating challenge.
//   multifile: false
//   context_line: 43
//   categories:
//     - Quickstart
//   complexity: ADVANCED
//   tags:
//     - hellobeam

package main

import (
	"context"
	"strconv"
	"strings"
	_ "time"

	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam"
	_ "github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/core/graph/window"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/io/textio"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/log"
	_ "github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/transforms/stats"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/x/beamx"
	"github.com/cd-paliv/beam-fork/sdks/v3/go/pkg/beam/x/debug"
)

func main() {
	ctx := context.Background()

	p, s := beam.NewPipelineWithRoot()

	input := textio.Read(s, "gs://apache-beam-samples/nyc_taxi/misc/sample1000.csv")

	// Extract cost from PCollection
	cost := ExtractCostFromFile(s, input)

	debug.Printf(s, "Above pCollection output", cost)

	err := beamx.Run(ctx, p)

	if err != nil {
		log.Exitf(context.Background(), "Failed to execute job: %v", err)
	}
}

func ExtractCostFromFile(s beam.Scope, input beam.PCollection) beam.PCollection {
	return beam.ParDo(s, func(line string) float64 {
		taxi := strings.Split(strings.TrimSpace(line), ",")
		if len(taxi) > 16 {
			cost, _ := strconv.ParseFloat(taxi[16], 64)
			return cost
		}
		return 0.0
	}, input)
}
