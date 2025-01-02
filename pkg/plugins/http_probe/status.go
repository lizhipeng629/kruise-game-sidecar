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

import "sync"

type HttpProbeStatus struct {
	status           string     // 记录当前状态
	err              error      // 记录最后一次发生的错误
	activeGoroutines int        // 当前活跃的 goroutine 数量
	mu               sync.Mutex // 用于保护状态字段的并发访问
}

func (h *HttpProbeStatus) setStatus(status string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.status = status
}

func (h *HttpProbeStatus) getStatus() string {
	return h.status
}

func (h *HttpProbeStatus) setError(err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.err = err
}

func (h *HttpProbeStatus) getError() error {
	return h.err
}

func (h *HttpProbeStatus) incrementGoroutines() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.activeGoroutines++
}

func (h *HttpProbeStatus) decrementGoroutines() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.activeGoroutines--
}

func (h *HttpProbeStatus) getActiveGoroutines() int {
	return h.activeGoroutines
}
