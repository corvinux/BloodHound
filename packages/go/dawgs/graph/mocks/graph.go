// Copyright 2025 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
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
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: graph.go
//
// Generated by this command:
//
//	mockgen -source=graph.go -copyright_file=../../../../LICENSE.header -destination=./mocks/graph.go -package=graph_mocks .
//

// Package graph_mocks is a generated GoMock package.
package graph_mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	graph "github.com/specterops/bloodhound/dawgs/graph"
	size "github.com/specterops/bloodhound/dawgs/util/size"
	gomock "go.uber.org/mock/gomock"
)

// MockPropertyValue is a mock of PropertyValue interface.
type MockPropertyValue struct {
	ctrl     *gomock.Controller
	recorder *MockPropertyValueMockRecorder
	isgomock struct{}
}

// MockPropertyValueMockRecorder is the mock recorder for MockPropertyValue.
type MockPropertyValueMockRecorder struct {
	mock *MockPropertyValue
}

// NewMockPropertyValue creates a new mock instance.
func NewMockPropertyValue(ctrl *gomock.Controller) *MockPropertyValue {
	mock := &MockPropertyValue{ctrl: ctrl}
	mock.recorder = &MockPropertyValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPropertyValue) EXPECT() *MockPropertyValueMockRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockPropertyValue) Any() any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Any")
	ret0, _ := ret[0].(any)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockPropertyValueMockRecorder) Any() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockPropertyValue)(nil).Any))
}

// Bool mocks base method.
func (m *MockPropertyValue) Bool() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bool")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bool indicates an expected call of Bool.
func (mr *MockPropertyValueMockRecorder) Bool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockPropertyValue)(nil).Bool))
}

// Float64 mocks base method.
func (m *MockPropertyValue) Float64() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float64")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64 indicates an expected call of Float64.
func (mr *MockPropertyValueMockRecorder) Float64() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64", reflect.TypeOf((*MockPropertyValue)(nil).Float64))
}

// IDSlice mocks base method.
func (m *MockPropertyValue) IDSlice() ([]graph.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDSlice")
	ret0, _ := ret[0].([]graph.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IDSlice indicates an expected call of IDSlice.
func (mr *MockPropertyValueMockRecorder) IDSlice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDSlice", reflect.TypeOf((*MockPropertyValue)(nil).IDSlice))
}

// Int mocks base method.
func (m *MockPropertyValue) Int() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int indicates an expected call of Int.
func (mr *MockPropertyValueMockRecorder) Int() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockPropertyValue)(nil).Int))
}

// Int64 mocks base method.
func (m *MockPropertyValue) Int64() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64 indicates an expected call of Int64.
func (mr *MockPropertyValueMockRecorder) Int64() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64", reflect.TypeOf((*MockPropertyValue)(nil).Int64))
}

// Int64Slice mocks base method.
func (m *MockPropertyValue) Int64Slice() ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64Slice")
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64Slice indicates an expected call of Int64Slice.
func (mr *MockPropertyValueMockRecorder) Int64Slice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64Slice", reflect.TypeOf((*MockPropertyValue)(nil).Int64Slice))
}

// IsNil mocks base method.
func (m *MockPropertyValue) IsNil() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNil")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNil indicates an expected call of IsNil.
func (mr *MockPropertyValueMockRecorder) IsNil() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNil", reflect.TypeOf((*MockPropertyValue)(nil).IsNil))
}

// String mocks base method.
func (m *MockPropertyValue) String() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// String indicates an expected call of String.
func (mr *MockPropertyValueMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockPropertyValue)(nil).String))
}

// StringSlice mocks base method.
func (m *MockPropertyValue) StringSlice() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringSlice")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StringSlice indicates an expected call of StringSlice.
func (mr *MockPropertyValueMockRecorder) StringSlice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringSlice", reflect.TypeOf((*MockPropertyValue)(nil).StringSlice))
}

// Time mocks base method.
func (m *MockPropertyValue) Time() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Time indicates an expected call of Time.
func (mr *MockPropertyValueMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockPropertyValue)(nil).Time))
}

// Uint64 mocks base method.
func (m *MockPropertyValue) Uint64() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint64")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uint64 indicates an expected call of Uint64.
func (mr *MockPropertyValueMockRecorder) Uint64() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint64", reflect.TypeOf((*MockPropertyValue)(nil).Uint64))
}

// MockEntity is a mock of Entity interface.
type MockEntity struct {
	ctrl     *gomock.Controller
	recorder *MockEntityMockRecorder
	isgomock struct{}
}

// MockEntityMockRecorder is the mock recorder for MockEntity.
type MockEntityMockRecorder struct {
	mock *MockEntity
}

// NewMockEntity creates a new mock instance.
func NewMockEntity(ctrl *gomock.Controller) *MockEntity {
	mock := &MockEntity{ctrl: ctrl}
	mock.recorder = &MockEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntity) EXPECT() *MockEntityMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockEntity) ID() graph.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(graph.ID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockEntityMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockEntity)(nil).ID))
}

// Properties mocks base method.
func (m *MockEntity) Properties() *graph.Properties {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Properties")
	ret0, _ := ret[0].(*graph.Properties)
	return ret0
}

// Properties indicates an expected call of Properties.
func (mr *MockEntityMockRecorder) Properties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Properties", reflect.TypeOf((*MockEntity)(nil).Properties))
}

// SizeOf mocks base method.
func (m *MockEntity) SizeOf() size.Size {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SizeOf")
	ret0, _ := ret[0].(size.Size)
	return ret0
}

// SizeOf indicates an expected call of SizeOf.
func (mr *MockEntityMockRecorder) SizeOf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SizeOf", reflect.TypeOf((*MockEntity)(nil).SizeOf))
}

// MockBatch is a mock of Batch interface.
type MockBatch struct {
	ctrl     *gomock.Controller
	recorder *MockBatchMockRecorder
	isgomock struct{}
}

// MockBatchMockRecorder is the mock recorder for MockBatch.
type MockBatchMockRecorder struct {
	mock *MockBatch
}

// NewMockBatch creates a new mock instance.
func NewMockBatch(ctrl *gomock.Controller) *MockBatch {
	mock := &MockBatch{ctrl: ctrl}
	mock.recorder = &MockBatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatch) EXPECT() *MockBatchMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockBatch) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockBatchMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockBatch)(nil).Commit))
}

// CreateNode mocks base method.
func (m *MockBatch) CreateNode(node *graph.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNode", node)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNode indicates an expected call of CreateNode.
func (mr *MockBatchMockRecorder) CreateNode(node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNode", reflect.TypeOf((*MockBatch)(nil).CreateNode), node)
}

// CreateRelationship mocks base method.
func (m *MockBatch) CreateRelationship(relationship *graph.Relationship) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelationship", relationship)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRelationship indicates an expected call of CreateRelationship.
func (mr *MockBatchMockRecorder) CreateRelationship(relationship any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelationship", reflect.TypeOf((*MockBatch)(nil).CreateRelationship), relationship)
}

// CreateRelationshipByIDs mocks base method.
func (m *MockBatch) CreateRelationshipByIDs(startNodeID, endNodeID graph.ID, kind graph.Kind, properties *graph.Properties) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelationshipByIDs", startNodeID, endNodeID, kind, properties)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRelationshipByIDs indicates an expected call of CreateRelationshipByIDs.
func (mr *MockBatchMockRecorder) CreateRelationshipByIDs(startNodeID, endNodeID, kind, properties any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelationshipByIDs", reflect.TypeOf((*MockBatch)(nil).CreateRelationshipByIDs), startNodeID, endNodeID, kind, properties)
}

// DeleteNode mocks base method.
func (m *MockBatch) DeleteNode(id graph.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNode", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNode indicates an expected call of DeleteNode.
func (mr *MockBatchMockRecorder) DeleteNode(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNode", reflect.TypeOf((*MockBatch)(nil).DeleteNode), id)
}

// DeleteRelationship mocks base method.
func (m *MockBatch) DeleteRelationship(id graph.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRelationship", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRelationship indicates an expected call of DeleteRelationship.
func (mr *MockBatchMockRecorder) DeleteRelationship(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRelationship", reflect.TypeOf((*MockBatch)(nil).DeleteRelationship), id)
}

// Nodes mocks base method.
func (m *MockBatch) Nodes() graph.NodeQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(graph.NodeQuery)
	return ret0
}

// Nodes indicates an expected call of Nodes.
func (mr *MockBatchMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockBatch)(nil).Nodes))
}

// Relationships mocks base method.
func (m *MockBatch) Relationships() graph.RelationshipQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relationships")
	ret0, _ := ret[0].(graph.RelationshipQuery)
	return ret0
}

// Relationships indicates an expected call of Relationships.
func (mr *MockBatchMockRecorder) Relationships() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relationships", reflect.TypeOf((*MockBatch)(nil).Relationships))
}

// UpdateNodeBy mocks base method.
func (m *MockBatch) UpdateNodeBy(update graph.NodeUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNodeBy", update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNodeBy indicates an expected call of UpdateNodeBy.
func (mr *MockBatchMockRecorder) UpdateNodeBy(update any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodeBy", reflect.TypeOf((*MockBatch)(nil).UpdateNodeBy), update)
}

// UpdateRelationshipBy mocks base method.
func (m *MockBatch) UpdateRelationshipBy(update graph.RelationshipUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRelationshipBy", update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRelationshipBy indicates an expected call of UpdateRelationshipBy.
func (mr *MockBatchMockRecorder) UpdateRelationshipBy(update any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRelationshipBy", reflect.TypeOf((*MockBatch)(nil).UpdateRelationshipBy), update)
}

// WithGraph mocks base method.
func (m *MockBatch) WithGraph(graphSchema graph.Graph) graph.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithGraph", graphSchema)
	ret0, _ := ret[0].(graph.Batch)
	return ret0
}

// WithGraph indicates an expected call of WithGraph.
func (mr *MockBatchMockRecorder) WithGraph(graphSchema any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithGraph", reflect.TypeOf((*MockBatch)(nil).WithGraph), graphSchema)
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
	isgomock struct{}
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTransaction) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit))
}

// CreateNode mocks base method.
func (m *MockTransaction) CreateNode(properties *graph.Properties, kinds ...graph.Kind) (*graph.Node, error) {
	m.ctrl.T.Helper()
	varargs := []any{properties}
	for _, a := range kinds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNode", varargs...)
	ret0, _ := ret[0].(*graph.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNode indicates an expected call of CreateNode.
func (mr *MockTransactionMockRecorder) CreateNode(properties any, kinds ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{properties}, kinds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNode", reflect.TypeOf((*MockTransaction)(nil).CreateNode), varargs...)
}

// CreateRelationshipByIDs mocks base method.
func (m *MockTransaction) CreateRelationshipByIDs(startNodeID, endNodeID graph.ID, kind graph.Kind, properties *graph.Properties) (*graph.Relationship, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelationshipByIDs", startNodeID, endNodeID, kind, properties)
	ret0, _ := ret[0].(*graph.Relationship)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRelationshipByIDs indicates an expected call of CreateRelationshipByIDs.
func (mr *MockTransactionMockRecorder) CreateRelationshipByIDs(startNodeID, endNodeID, kind, properties any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelationshipByIDs", reflect.TypeOf((*MockTransaction)(nil).CreateRelationshipByIDs), startNodeID, endNodeID, kind, properties)
}

// GraphQueryMemoryLimit mocks base method.
func (m *MockTransaction) GraphQueryMemoryLimit() size.Size {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphQueryMemoryLimit")
	ret0, _ := ret[0].(size.Size)
	return ret0
}

// GraphQueryMemoryLimit indicates an expected call of GraphQueryMemoryLimit.
func (mr *MockTransactionMockRecorder) GraphQueryMemoryLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphQueryMemoryLimit", reflect.TypeOf((*MockTransaction)(nil).GraphQueryMemoryLimit))
}

// Nodes mocks base method.
func (m *MockTransaction) Nodes() graph.NodeQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes")
	ret0, _ := ret[0].(graph.NodeQuery)
	return ret0
}

// Nodes indicates an expected call of Nodes.
func (mr *MockTransactionMockRecorder) Nodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockTransaction)(nil).Nodes))
}

// Query mocks base method.
func (m *MockTransaction) Query(query string, parameters map[string]any) graph.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", query, parameters)
	ret0, _ := ret[0].(graph.Result)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockTransactionMockRecorder) Query(query, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTransaction)(nil).Query), query, parameters)
}

// Raw mocks base method.
func (m *MockTransaction) Raw(query string, parameters map[string]any) graph.Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw", query, parameters)
	ret0, _ := ret[0].(graph.Result)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockTransactionMockRecorder) Raw(query, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockTransaction)(nil).Raw), query, parameters)
}

// Relationships mocks base method.
func (m *MockTransaction) Relationships() graph.RelationshipQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relationships")
	ret0, _ := ret[0].(graph.RelationshipQuery)
	return ret0
}

// Relationships indicates an expected call of Relationships.
func (mr *MockTransactionMockRecorder) Relationships() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relationships", reflect.TypeOf((*MockTransaction)(nil).Relationships))
}

// UpdateNode mocks base method.
func (m *MockTransaction) UpdateNode(node *graph.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNode", node)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNode indicates an expected call of UpdateNode.
func (mr *MockTransactionMockRecorder) UpdateNode(node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNode", reflect.TypeOf((*MockTransaction)(nil).UpdateNode), node)
}

// UpdateRelationship mocks base method.
func (m *MockTransaction) UpdateRelationship(relationship *graph.Relationship) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRelationship", relationship)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRelationship indicates an expected call of UpdateRelationship.
func (mr *MockTransactionMockRecorder) UpdateRelationship(relationship any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRelationship", reflect.TypeOf((*MockTransaction)(nil).UpdateRelationship), relationship)
}

// WithGraph mocks base method.
func (m *MockTransaction) WithGraph(graphSchema graph.Graph) graph.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithGraph", graphSchema)
	ret0, _ := ret[0].(graph.Transaction)
	return ret0
}

// WithGraph indicates an expected call of WithGraph.
func (mr *MockTransactionMockRecorder) WithGraph(graphSchema any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithGraph", reflect.TypeOf((*MockTransaction)(nil).WithGraph), graphSchema)
}

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
	isgomock struct{}
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// AssertSchema mocks base method.
func (m *MockDatabase) AssertSchema(ctx context.Context, dbSchema graph.Schema) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssertSchema", ctx, dbSchema)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssertSchema indicates an expected call of AssertSchema.
func (mr *MockDatabaseMockRecorder) AssertSchema(ctx, dbSchema any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssertSchema", reflect.TypeOf((*MockDatabase)(nil).AssertSchema), ctx, dbSchema)
}

// BatchOperation mocks base method.
func (m *MockDatabase) BatchOperation(ctx context.Context, batchDelegate graph.BatchDelegate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchOperation", ctx, batchDelegate)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchOperation indicates an expected call of BatchOperation.
func (mr *MockDatabaseMockRecorder) BatchOperation(ctx, batchDelegate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchOperation", reflect.TypeOf((*MockDatabase)(nil).BatchOperation), ctx, batchDelegate)
}

// Close mocks base method.
func (m *MockDatabase) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaseMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabase)(nil).Close), ctx)
}

// FetchKinds mocks base method.
func (m *MockDatabase) FetchKinds(ctx context.Context) (graph.Kinds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchKinds", ctx)
	ret0, _ := ret[0].(graph.Kinds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchKinds indicates an expected call of FetchKinds.
func (mr *MockDatabaseMockRecorder) FetchKinds(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchKinds", reflect.TypeOf((*MockDatabase)(nil).FetchKinds), ctx)
}

// ReadTransaction mocks base method.
func (m *MockDatabase) ReadTransaction(ctx context.Context, txDelegate graph.TransactionDelegate, options ...graph.TransactionOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, txDelegate}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadTransaction indicates an expected call of ReadTransaction.
func (mr *MockDatabaseMockRecorder) ReadTransaction(ctx, txDelegate any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, txDelegate}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTransaction", reflect.TypeOf((*MockDatabase)(nil).ReadTransaction), varargs...)
}

// RefreshKinds mocks base method.
func (m *MockDatabase) RefreshKinds(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshKinds", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshKinds indicates an expected call of RefreshKinds.
func (mr *MockDatabaseMockRecorder) RefreshKinds(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshKinds", reflect.TypeOf((*MockDatabase)(nil).RefreshKinds), ctx)
}

// Run mocks base method.
func (m *MockDatabase) Run(ctx context.Context, query string, parameters map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, query, parameters)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockDatabaseMockRecorder) Run(ctx, query, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockDatabase)(nil).Run), ctx, query, parameters)
}

// SetBatchWriteSize mocks base method.
func (m *MockDatabase) SetBatchWriteSize(interval int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBatchWriteSize", interval)
}

// SetBatchWriteSize indicates an expected call of SetBatchWriteSize.
func (mr *MockDatabaseMockRecorder) SetBatchWriteSize(interval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBatchWriteSize", reflect.TypeOf((*MockDatabase)(nil).SetBatchWriteSize), interval)
}

// SetDefaultGraph mocks base method.
func (m *MockDatabase) SetDefaultGraph(ctx context.Context, graphSchema graph.Graph) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefaultGraph", ctx, graphSchema)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultGraph indicates an expected call of SetDefaultGraph.
func (mr *MockDatabaseMockRecorder) SetDefaultGraph(ctx, graphSchema any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultGraph", reflect.TypeOf((*MockDatabase)(nil).SetDefaultGraph), ctx, graphSchema)
}

// SetWriteFlushSize mocks base method.
func (m *MockDatabase) SetWriteFlushSize(interval int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetWriteFlushSize", interval)
}

// SetWriteFlushSize indicates an expected call of SetWriteFlushSize.
func (mr *MockDatabaseMockRecorder) SetWriteFlushSize(interval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteFlushSize", reflect.TypeOf((*MockDatabase)(nil).SetWriteFlushSize), interval)
}

// WriteTransaction mocks base method.
func (m *MockDatabase) WriteTransaction(ctx context.Context, txDelegate graph.TransactionDelegate, options ...graph.TransactionOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, txDelegate}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteTransaction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTransaction indicates an expected call of WriteTransaction.
func (mr *MockDatabaseMockRecorder) WriteTransaction(ctx, txDelegate any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, txDelegate}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTransaction", reflect.TypeOf((*MockDatabase)(nil).WriteTransaction), varargs...)
}
