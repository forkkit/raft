// +build batchtest

package raft

// NOTE: This is exposed for middleware testing purposes and is not a stable API
func (m *MockFSM) ApplyBatch(logs []*Log) []interface{} {
	m.Lock()
	defer m.Unlock()

	ret := make([]interface{}, len(logs))
	for i, log := range logs {
		m.logs = append(m.logs, log.Data)
		ret[i] = len(m.logs)
	}

	return ret
}
