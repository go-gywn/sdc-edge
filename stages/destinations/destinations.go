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
package destinations

import (
	_ "sdc-edge/stages/destinations/coap"
	_ "sdc-edge/stages/destinations/firehose"
	_ "sdc-edge/stages/destinations/http"
	_ "sdc-edge/stages/destinations/influxdb"
	_ "sdc-edge/stages/destinations/kafka"
	_ "sdc-edge/stages/destinations/kinesis"
	_ "sdc-edge/stages/destinations/mqtt"
	_ "sdc-edge/stages/destinations/s3"
	_ "sdc-edge/stages/destinations/toerror"
	_ "sdc-edge/stages/destinations/toevent"
	_ "sdc-edge/stages/destinations/trash"
	_ "sdc-edge/stages/destinations/websocket"
)
