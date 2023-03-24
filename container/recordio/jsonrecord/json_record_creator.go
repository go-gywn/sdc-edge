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
package jsonrecord

import (
	"bytes"
	"encoding/json"
	"sdc-edge/api"
)

type RecordCreator struct {
}

func (r *RecordCreator) CreateRecord(
	context api.StageContext,
	lineText string,
	messageId string,
	headers []*api.Field,
) (api.Record, error) {
	recordBuffer := bytes.NewBufferString(lineText)
	decoder := json.NewDecoder(recordBuffer)
	var recordValue interface{}
	err := decoder.Decode(&recordValue)
	if err != nil {
		return nil, err
	}
	return context.CreateRecord(messageId, recordValue)
}
