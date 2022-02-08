package watch

import (
	"jrasp-daemon/java_process"
	"jrasp-daemon/utils"
)

// HeartBeat 心跳信息
type HeartBeatInfo struct {
	Status map[int32]AgentInfo `json:"agentInfo"`
}

// java agent信息
type AgentInfo struct {
	Pid          int32                   `json:"pid"`       // 进程信息
	StartTime    string                  `json:"startTime"` // 启动时间
	InjectStatus java_process.InjectType `json:"status"`    // 注入状态
	// jdk版本
}

func NewHeartBeat() *HeartBeatInfo {
	return &HeartBeatInfo{
		Status: make(map[int32]AgentInfo),
	}
}

func (hb *HeartBeatInfo) Append(jp *java_process.JavaProcess) {
	agentInfo := AgentInfo{
		Pid:          jp.JavaPid,
		StartTime:    jp.StartTime,
		InjectStatus: jp.InjectedStatus,
	}
	hb.Status[jp.JavaPid] = agentInfo
}

// 转成json字符串
func (hb *HeartBeatInfo) toJsonString() string {
	return utils.ToString(hb.Status)
}
