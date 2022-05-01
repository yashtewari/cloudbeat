// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package transformer

import (
	"io/fs"

	"github.com/gofrs/uuid"
	"k8s.io/client-go/kubernetes"
)

type ResourceTypeMetadata struct {
	CycleMetadata
	Type string
}

type ResourceFields struct {
	ID   string      `json:"id"`
	Type string      `json:"type"`
	Raw  interface{} `json:"raw"`
}

type ResourceMetadata struct {
	ResourceTypeMetadata
	ResourceId string
}

type CycleMetadata struct {
	CycleId uuid.UUID
}

type CommonDataProvider struct {
	kubeClient kubernetes.Interface
	fsys fs.FS
}

type CommonData struct {
	clusterId string
	nodeId string
}

type CommonDataInterface interface {
	GetData() CommonData
	GetResourceId(string) string
}