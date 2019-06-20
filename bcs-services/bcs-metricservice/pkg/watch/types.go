/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package watch

import (
	btypes "bk-bcs/bcs-common/common/types"
	"bk-bcs/bcs-services/bcs-metricservice/pkg/types"
)

type EventType int

const (
	EventMetricUpd EventType = iota
	EventMetricDel
	EventDynamicUpd
)

func (et EventType) String() string {
	return eventTypeNames[et]
}

var (
	eventTypeNames = map[EventType]string{
		EventMetricUpd:  "MetricUpdate",
		EventMetricDel:  "MetricDelete",
		EventDynamicUpd: "DynamicUpdate",
	}
)

type MetricEvent struct {
	ID     string
	Type   EventType
	Metric *types.Metric
	First  bool
	Last   bool
	Meta   map[string]btypes.ObjectMeta
}
