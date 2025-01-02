/*
Copyright 2024

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httpprobe

import "github.com/magicsong/kidecar/pkg/store"

type EndpointConfig struct {
	URL                string                `json:"url"`                // 目标 URL
	Method             string                `json:"method"`             // HTTP 方法
	Headers            map[string]string     `json:"headers"`            // 请求头
	Timeout            int                   `json:"timeout"`            // 超时时间（秒）
	ExpectedStatusCode int                   `json:"expectedStatusCode"` // 预期的 HTTP 状态码
	StorageConfig      store.StorageConfig   `json:"storageConfig"`      // 存储配置
	JSONPathConfig     *store.JSONPathConfig `json:"jsonPathConfig"`     // JSONPath 配置
}

type HttpProbeConfig struct {
	StartDelaySeconds    int              `json:"startDelaySeconds"`    // 延迟启动时间（秒）
	Endpoints            []EndpointConfig `json:"endpoints,omitempty"`  // 多个端点的配置
	ProbeIntervalSeconds int              `json:"probeIntervalSeconds"` // 探测间隔时间（秒）
}
