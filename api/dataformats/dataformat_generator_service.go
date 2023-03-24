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
package dataformats

import (
	"sdc-edge/api"
	"io"
)

const (
	DataFormatGeneratorServiceName = "com.streamsets.pipeline.api.service.dataformats.DataFormatGeneratorService"
)

type DataFormatGeneratorService interface {
	GetGenerator(writer io.Writer) (RecordWriter, error)
	IsWholeFileFormat() bool
	GetWholeFileName(record api.Record) (string, error)
	GetWholeFileExistsAction() string
	GetIncludeChecksumInTheEvents() bool
	GetChecksumAlgorithm() string
}

type RecordWriter interface {
	WriteRecord(r api.Record) error
	Flush() error
	Close() error
}
