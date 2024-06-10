// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"time"
)

// Ref: #/components/schemas/App
type App struct {
	Name OptString `json:"name"`
	ID   OptString `json:"id"`
}

// GetName returns the value of Name.
func (s *App) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *App) GetID() OptString {
	return s.ID
}

// SetName sets the value of Name.
func (s *App) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *App) SetID(val OptString) {
	s.ID = val
}

// Ref: #/components/schemas/AppHistory
type AppHistory struct {
	Name      OptString   `json:"name"`
	ID        OptString   `json:"id"`
	Timestamp OptDateTime `json:"timestamp"`
}

// GetName returns the value of Name.
func (s *AppHistory) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *AppHistory) GetID() OptString {
	return s.ID
}

// GetTimestamp returns the value of Timestamp.
func (s *AppHistory) GetTimestamp() OptDateTime {
	return s.Timestamp
}

// SetName sets the value of Name.
func (s *AppHistory) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *AppHistory) SetID(val OptString) {
	s.ID = val
}

// SetTimestamp sets the value of Timestamp.
func (s *AppHistory) SetTimestamp(val OptDateTime) {
	s.Timestamp = val
}

// Ref: #/components/schemas/AppsHistory
type AppsHistory struct {
	History []AppHistory `json:"history"`
}

// GetHistory returns the value of History.
func (s *AppsHistory) GetHistory() []AppHistory {
	return s.History
}

// SetHistory sets the value of History.
func (s *AppsHistory) SetHistory(val []AppHistory) {
	s.History = val
}

func (*AppsHistory) getAppsHistoryRes() {}

// Ref: #/components/schemas/AppsResponse
type AppsResponse struct {
	Apps []App `json:"apps"`
}

// GetApps returns the value of Apps.
func (s *AppsResponse) GetApps() []App {
	return s.Apps
}

// SetApps sets the value of Apps.
func (s *AppsResponse) SetApps(val []App) {
	s.Apps = val
}

func (*AppsResponse) getAppsPerPartitionPerQueueRes() {}

// Ref: #/components/schemas/ContainerHistory
type ContainerHistory struct {
	Name      OptString   `json:"name"`
	ID        OptString   `json:"id"`
	Timestamp OptDateTime `json:"timestamp"`
}

// GetName returns the value of Name.
func (s *ContainerHistory) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *ContainerHistory) GetID() OptString {
	return s.ID
}

// GetTimestamp returns the value of Timestamp.
func (s *ContainerHistory) GetTimestamp() OptDateTime {
	return s.Timestamp
}

// SetName sets the value of Name.
func (s *ContainerHistory) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *ContainerHistory) SetID(val OptString) {
	s.ID = val
}

// SetTimestamp sets the value of Timestamp.
func (s *ContainerHistory) SetTimestamp(val OptDateTime) {
	s.Timestamp = val
}

// Ref: #/components/schemas/ContainersHistory
type ContainersHistory struct {
	History []ContainerHistory `json:"history"`
}

// GetHistory returns the value of History.
func (s *ContainersHistory) GetHistory() []ContainerHistory {
	return s.History
}

// SetHistory sets the value of History.
func (s *ContainersHistory) SetHistory(val []ContainerHistory) {
	s.History = val
}

func (*ContainersHistory) getContainersHistoryRes() {}

// Ref: #/components/schemas/Event
type Event struct {
	Type  OptString `json:"type"`
	Count OptInt    `json:"count"`
}

// GetType returns the value of Type.
func (s *Event) GetType() OptString {
	return s.Type
}

// GetCount returns the value of Count.
func (s *Event) GetCount() OptInt {
	return s.Count
}

// SetType sets the value of Type.
func (s *Event) SetType(val OptString) {
	s.Type = val
}

// SetCount sets the value of Count.
func (s *Event) SetCount(val OptInt) {
	s.Count = val
}

// Ref: #/components/schemas/EventStatistics
type EventStatistics struct {
	Events []Event `json:"events"`
}

// GetEvents returns the value of Events.
func (s *EventStatistics) GetEvents() []Event {
	return s.Events
}

// SetEvents sets the value of Events.
func (s *EventStatistics) SetEvents(val []Event) {
	s.Events = val
}

func (*EventStatistics) getEventStatisticsRes() {}

// GetAppsHistoryInternalServerError is response for GetAppsHistory operation.
type GetAppsHistoryInternalServerError struct{}

func (*GetAppsHistoryInternalServerError) getAppsHistoryRes() {}

// GetAppsPerPartitionPerQueueInternalServerError is response for GetAppsPerPartitionPerQueue operation.
type GetAppsPerPartitionPerQueueInternalServerError struct{}

func (*GetAppsPerPartitionPerQueueInternalServerError) getAppsPerPartitionPerQueueRes() {}

// GetContainersHistoryInternalServerError is response for GetContainersHistory operation.
type GetContainersHistoryInternalServerError struct{}

func (*GetContainersHistoryInternalServerError) getContainersHistoryRes() {}

// GetEventStatisticsInternalServerError is response for GetEventStatistics operation.
type GetEventStatisticsInternalServerError struct{}

func (*GetEventStatisticsInternalServerError) getEventStatisticsRes() {}

// GetNodeUtilizationsInternalServerError is response for GetNodeUtilizations operation.
type GetNodeUtilizationsInternalServerError struct{}

func (*GetNodeUtilizationsInternalServerError) getNodeUtilizationsRes() {}

// GetNodesPerPartitionInternalServerError is response for GetNodesPerPartition operation.
type GetNodesPerPartitionInternalServerError struct{}

func (*GetNodesPerPartitionInternalServerError) getNodesPerPartitionRes() {}

type GetNodesPerPartitionOKApplicationJSON []Node

func (*GetNodesPerPartitionOKApplicationJSON) getNodesPerPartitionRes() {}

// GetPartitionsInternalServerError is response for GetPartitions operation.
type GetPartitionsInternalServerError struct{}

func (*GetPartitionsInternalServerError) getPartitionsRes() {}

// GetQueuesPerPartitionInternalServerError is response for GetQueuesPerPartition operation.
type GetQueuesPerPartitionInternalServerError struct{}

func (*GetQueuesPerPartitionInternalServerError) getQueuesPerPartitionRes() {}

// Ref: #/components/schemas/Node
type Node struct {
	Name OptString `json:"name"`
	ID   OptString `json:"id"`
}

// GetName returns the value of Name.
func (s *Node) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *Node) GetID() OptString {
	return s.ID
}

// SetName sets the value of Name.
func (s *Node) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *Node) SetID(val OptString) {
	s.ID = val
}

// Ref: #/components/schemas/NodeUtilization
type NodeUtilization struct {
	Utilization []NodeUtilizationItem `json:"utilization"`
}

// GetUtilization returns the value of Utilization.
func (s *NodeUtilization) GetUtilization() []NodeUtilizationItem {
	return s.Utilization
}

// SetUtilization sets the value of Utilization.
func (s *NodeUtilization) SetUtilization(val []NodeUtilizationItem) {
	s.Utilization = val
}

func (*NodeUtilization) getNodeUtilizationsRes() {}

// Ref: #/components/schemas/NodeUtilizationItem
type NodeUtilizationItem struct {
	Name   OptString  `json:"name"`
	ID     OptString  `json:"id"`
	CPU    OptFloat64 `json:"cpu"`
	Memory OptFloat64 `json:"memory"`
}

// GetName returns the value of Name.
func (s *NodeUtilizationItem) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *NodeUtilizationItem) GetID() OptString {
	return s.ID
}

// GetCPU returns the value of CPU.
func (s *NodeUtilizationItem) GetCPU() OptFloat64 {
	return s.CPU
}

// GetMemory returns the value of Memory.
func (s *NodeUtilizationItem) GetMemory() OptFloat64 {
	return s.Memory
}

// SetName sets the value of Name.
func (s *NodeUtilizationItem) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *NodeUtilizationItem) SetID(val OptString) {
	s.ID = val
}

// SetCPU sets the value of CPU.
func (s *NodeUtilizationItem) SetCPU(val OptFloat64) {
	s.CPU = val
}

// SetMemory sets the value of Memory.
func (s *NodeUtilizationItem) SetMemory(val OptFloat64) {
	s.Memory = val
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Partition
type Partition struct {
	Name OptString `json:"name"`
	ID   OptString `json:"id"`
}

// GetName returns the value of Name.
func (s *Partition) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *Partition) GetID() OptString {
	return s.ID
}

// SetName sets the value of Name.
func (s *Partition) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *Partition) SetID(val OptString) {
	s.ID = val
}

// Ref: #/components/schemas/PartitionsResponse
type PartitionsResponse struct {
	Partitions []Partition `json:"partitions"`
}

// GetPartitions returns the value of Partitions.
func (s *PartitionsResponse) GetPartitions() []Partition {
	return s.Partitions
}

// SetPartitions sets the value of Partitions.
func (s *PartitionsResponse) SetPartitions(val []Partition) {
	s.Partitions = val
}

func (*PartitionsResponse) getPartitionsRes() {}

// Ref: #/components/schemas/Queue
type Queue struct {
	Name OptString `json:"name"`
	ID   OptString `json:"id"`
}

// GetName returns the value of Name.
func (s *Queue) GetName() OptString {
	return s.Name
}

// GetID returns the value of ID.
func (s *Queue) GetID() OptString {
	return s.ID
}

// SetName sets the value of Name.
func (s *Queue) SetName(val OptString) {
	s.Name = val
}

// SetID sets the value of ID.
func (s *Queue) SetID(val OptString) {
	s.ID = val
}

// Ref: #/components/schemas/QueuesResponse
type QueuesResponse struct {
	Queues []Queue `json:"queues"`
}

// GetQueues returns the value of Queues.
func (s *QueuesResponse) GetQueues() []Queue {
	return s.Queues
}

// SetQueues sets the value of Queues.
func (s *QueuesResponse) SetQueues(val []Queue) {
	s.Queues = val
}

func (*QueuesResponse) getQueuesPerPartitionRes() {}
