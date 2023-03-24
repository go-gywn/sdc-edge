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
package toevent

import (
	"sdc-edge/api"
	"sdc-edge/api/validation"
	"sdc-edge/container/common"
	"sdc-edge/stages/stagelibrary"
)

const (
	Library   = "streamsets-datacollector-dev-lib"
	StageName = "com_streamsets_pipeline_stage_destination_ToEventTarget"
)

type Destination struct {
	*common.BaseStage
}

func init() {
	stagelibrary.SetCreator(Library, StageName, func() api.Stage {
		return &Destination{BaseStage: &common.BaseStage{}}
	})
}

func (d *Destination) Init(stageContext api.StageContext) []validation.Issue {
	return d.BaseStage.Init(stageContext)
}

func (d *Destination) Write(batch api.Batch) error {
	counter := 1
	for _, record := range batch.GetRecords() {
		rootField, _ := record.Get()
		recordId := common.CreateRecordId("event-target", counter)
		if eventRecord, err := d.GetStageContext().CreateEventRecord(
			recordId,
			nil,
			"event-target",
			1,
		); err == nil {
			eventRecord.Set(rootField)
			d.GetStageContext().ToEvent(eventRecord)
		} else {
			d.GetStageContext().ToError(err, eventRecord)
		}
		counter++
	}
	return nil
}
