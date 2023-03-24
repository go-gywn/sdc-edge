// Copyright 2018 StreamSets Inc.
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
package execution

import "sdc-edge/api/validation"

type PreviewOutput struct {
	PreviewStatus string              `json:"status"`
	Issues        *validation.Issues  `json:"issues"`
	Output        [][]StageOutputJson `json:"batchesOutput"`
	Message       string              `json:"message"`
}

func NewPreviewOutput(batchOutputs [][]StageOutput) ([][]StageOutputJson, error) {
	batchOutputsJson := make([][]StageOutputJson, len(batchOutputs))
	for batchIndex, batchOutput := range batchOutputs {
		batchOutputJson := make([]StageOutputJson, len(batchOutput))
		for stageIndex, stageOutput := range batchOutput {
			stageOutputJson, err := NewStageOutputJson(stageOutput)
			if err != nil {
				return batchOutputsJson, err
			}
			batchOutputJson[stageIndex] = *stageOutputJson
		}
		batchOutputsJson[batchIndex] = batchOutputJson
	}

	return batchOutputsJson, nil
}
