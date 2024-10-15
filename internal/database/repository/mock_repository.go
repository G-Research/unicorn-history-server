// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/G-Research/yunikorn-history-server/internal/database/repository (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -destination=mock_repository.go -package=repository github.com/G-Research/yunikorn-history-server/internal/database/repository Repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	dao "github.com/G-Research/yunikorn-core/pkg/webservice/dao"
	model "github.com/G-Research/yunikorn-history-server/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddQueues mocks base method.
func (m *MockRepository) AddQueues(arg0 context.Context, arg1 *string, arg2 []*dao.PartitionQueueDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddQueues", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddQueues indicates an expected call of AddQueues.
func (mr *MockRepositoryMockRecorder) AddQueues(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddQueues", reflect.TypeOf((*MockRepository)(nil).AddQueues), arg0, arg1, arg2)
}

// DeleteInactivePartitions mocks base method.
func (m *MockRepository) DeleteInactivePartitions(arg0 context.Context, arg1 []*dao.PartitionInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInactivePartitions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInactivePartitions indicates an expected call of DeleteInactivePartitions.
func (mr *MockRepositoryMockRecorder) DeleteInactivePartitions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInactivePartitions", reflect.TypeOf((*MockRepository)(nil).DeleteInactivePartitions), arg0, arg1)
}

// DeleteQueues mocks base method.
func (m *MockRepository) DeleteQueues(arg0 context.Context, arg1 []*model.PartitionQueueDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQueues", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQueues indicates an expected call of DeleteQueues.
func (mr *MockRepositoryMockRecorder) DeleteQueues(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQueues", reflect.TypeOf((*MockRepository)(nil).DeleteQueues), arg0, arg1)
}

// GetActivePartitions mocks base method.
func (m *MockRepository) GetActivePartitions(arg0 context.Context) ([]*model.PartitionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivePartitions", arg0)
	ret0, _ := ret[0].([]*model.PartitionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivePartitions indicates an expected call of GetActivePartitions.
func (mr *MockRepositoryMockRecorder) GetActivePartitions(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivePartitions", reflect.TypeOf((*MockRepository)(nil).GetActivePartitions), arg0)
}

// GetAllApplications mocks base method.
func (m *MockRepository) GetAllApplications(arg0 context.Context, arg1 ApplicationFilters) ([]*model.ApplicationDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllApplications", arg0, arg1)
	ret0, _ := ret[0].([]*model.ApplicationDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllApplications indicates an expected call of GetAllApplications.
func (mr *MockRepositoryMockRecorder) GetAllApplications(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllApplications", reflect.TypeOf((*MockRepository)(nil).GetAllApplications), arg0, arg1)
}

// GetAllPartitions mocks base method.
func (m *MockRepository) GetAllPartitions(arg0 context.Context, arg1 PartitionFilters) ([]*model.PartitionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPartitions", arg0, arg1)
	ret0, _ := ret[0].([]*model.PartitionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPartitions indicates an expected call of GetAllPartitions.
func (mr *MockRepositoryMockRecorder) GetAllPartitions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPartitions", reflect.TypeOf((*MockRepository)(nil).GetAllPartitions), arg0, arg1)
}

// GetAllQueues mocks base method.
func (m *MockRepository) GetAllQueues(arg0 context.Context) ([]*model.PartitionQueueDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllQueues", arg0)
	ret0, _ := ret[0].([]*model.PartitionQueueDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllQueues indicates an expected call of GetAllQueues.
func (mr *MockRepositoryMockRecorder) GetAllQueues(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllQueues", reflect.TypeOf((*MockRepository)(nil).GetAllQueues), arg0)
}

// GetApplicationsHistory mocks base method.
func (m *MockRepository) GetApplicationsHistory(arg0 context.Context, arg1 HistoryFilters) ([]*dao.ApplicationHistoryDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationsHistory", arg0, arg1)
	ret0, _ := ret[0].([]*dao.ApplicationHistoryDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationsHistory indicates an expected call of GetApplicationsHistory.
func (mr *MockRepositoryMockRecorder) GetApplicationsHistory(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationsHistory", reflect.TypeOf((*MockRepository)(nil).GetApplicationsHistory), arg0, arg1)
}

// GetAppsPerPartitionPerQueue mocks base method.
func (m *MockRepository) GetAppsPerPartitionPerQueue(arg0 context.Context, arg1, arg2 string, arg3 ApplicationFilters) ([]*model.ApplicationDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppsPerPartitionPerQueue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*model.ApplicationDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppsPerPartitionPerQueue indicates an expected call of GetAppsPerPartitionPerQueue.
func (mr *MockRepositoryMockRecorder) GetAppsPerPartitionPerQueue(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppsPerPartitionPerQueue", reflect.TypeOf((*MockRepository)(nil).GetAppsPerPartitionPerQueue), arg0, arg1, arg2, arg3)
}

// GetContainersHistory mocks base method.
func (m *MockRepository) GetContainersHistory(arg0 context.Context, arg1 HistoryFilters) ([]*dao.ContainerHistoryDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainersHistory", arg0, arg1)
	ret0, _ := ret[0].([]*dao.ContainerHistoryDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainersHistory indicates an expected call of GetContainersHistory.
func (mr *MockRepositoryMockRecorder) GetContainersHistory(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainersHistory", reflect.TypeOf((*MockRepository)(nil).GetContainersHistory), arg0, arg1)
}

// GetNodeUtilizations mocks base method.
func (m *MockRepository) GetNodeUtilizations(arg0 context.Context) ([]*dao.PartitionNodesUtilDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeUtilizations", arg0)
	ret0, _ := ret[0].([]*dao.PartitionNodesUtilDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeUtilizations indicates an expected call of GetNodeUtilizations.
func (mr *MockRepositoryMockRecorder) GetNodeUtilizations(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeUtilizations", reflect.TypeOf((*MockRepository)(nil).GetNodeUtilizations), arg0)
}

// GetNodesPerPartition mocks base method.
func (m *MockRepository) GetNodesPerPartition(arg0 context.Context, arg1 string) ([]*dao.NodeDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodesPerPartition", arg0, arg1)
	ret0, _ := ret[0].([]*dao.NodeDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodesPerPartition indicates an expected call of GetNodesPerPartition.
func (mr *MockRepositoryMockRecorder) GetNodesPerPartition(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodesPerPartition", reflect.TypeOf((*MockRepository)(nil).GetNodesPerPartition), arg0, arg1)
}

// GetQueue mocks base method.
func (m *MockRepository) GetQueue(arg0 context.Context, arg1, arg2 string) (*model.PartitionQueueDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueue", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.PartitionQueueDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueue indicates an expected call of GetQueue.
func (mr *MockRepositoryMockRecorder) GetQueue(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueue", reflect.TypeOf((*MockRepository)(nil).GetQueue), arg0, arg1, arg2)
}

// GetQueuesPerPartition mocks base method.
func (m *MockRepository) GetQueuesPerPartition(arg0 context.Context, arg1 string) ([]*model.PartitionQueueDAOInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueuesPerPartition", arg0, arg1)
	ret0, _ := ret[0].([]*model.PartitionQueueDAOInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueuesPerPartition indicates an expected call of GetQueuesPerPartition.
func (mr *MockRepositoryMockRecorder) GetQueuesPerPartition(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueuesPerPartition", reflect.TypeOf((*MockRepository)(nil).GetQueuesPerPartition), arg0, arg1)
}

// InsertNodeUtilizations mocks base method.
func (m *MockRepository) InsertNodeUtilizations(arg0 context.Context, arg1 []*dao.PartitionNodesUtilDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNodeUtilizations", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNodeUtilizations indicates an expected call of InsertNodeUtilizations.
func (mr *MockRepositoryMockRecorder) InsertNodeUtilizations(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNodeUtilizations", reflect.TypeOf((*MockRepository)(nil).InsertNodeUtilizations), arg0, arg1)
}

// UpdateHistory mocks base method.
func (m *MockRepository) UpdateHistory(arg0 context.Context, arg1 []*dao.ApplicationHistoryDAOInfo, arg2 []*dao.ContainerHistoryDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHistory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHistory indicates an expected call of UpdateHistory.
func (mr *MockRepositoryMockRecorder) UpdateHistory(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHistory", reflect.TypeOf((*MockRepository)(nil).UpdateHistory), arg0, arg1, arg2)
}

// UpdateQueue mocks base method.
func (m *MockRepository) UpdateQueue(arg0 context.Context, arg1 *dao.PartitionQueueDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQueue", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQueue indicates an expected call of UpdateQueue.
func (mr *MockRepositoryMockRecorder) UpdateQueue(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQueue", reflect.TypeOf((*MockRepository)(nil).UpdateQueue), arg0, arg1)
}

// UpsertApplications mocks base method.
func (m *MockRepository) UpsertApplications(arg0 context.Context, arg1 []*dao.ApplicationDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertApplications", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertApplications indicates an expected call of UpsertApplications.
func (mr *MockRepositoryMockRecorder) UpsertApplications(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertApplications", reflect.TypeOf((*MockRepository)(nil).UpsertApplications), arg0, arg1)
}

// UpsertNodes mocks base method.
func (m *MockRepository) UpsertNodes(arg0 context.Context, arg1 []*dao.NodeDAOInfo, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNodes indicates an expected call of UpsertNodes.
func (mr *MockRepositoryMockRecorder) UpsertNodes(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNodes", reflect.TypeOf((*MockRepository)(nil).UpsertNodes), arg0, arg1, arg2)
}

// UpsertPartitions mocks base method.
func (m *MockRepository) UpsertPartitions(arg0 context.Context, arg1 []*dao.PartitionInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPartitions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPartitions indicates an expected call of UpsertPartitions.
func (mr *MockRepositoryMockRecorder) UpsertPartitions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPartitions", reflect.TypeOf((*MockRepository)(nil).UpsertPartitions), arg0, arg1)
}

// UpsertQueues mocks base method.
func (m *MockRepository) UpsertQueues(arg0 context.Context, arg1 []*dao.PartitionQueueDAOInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertQueues", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertQueues indicates an expected call of UpsertQueues.
func (mr *MockRepositoryMockRecorder) UpsertQueues(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertQueues", reflect.TypeOf((*MockRepository)(nil).UpsertQueues), arg0, arg1)
}
