// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Rafaela314/instituto-maternar/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/Rafaela314/instituto-maternar/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CountDoctorReviews mocks base method.
func (m *MockStore) CountDoctorReviews(arg0 context.Context, arg1 int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDoctorReviews", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDoctorReviews indicates an expected call of CountDoctorReviews.
func (mr *MockStoreMockRecorder) CountDoctorReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDoctorReviews", reflect.TypeOf((*MockStore)(nil).CountDoctorReviews), arg0, arg1)
}

// CountPlaceReviews mocks base method.
func (m *MockStore) CountPlaceReviews(arg0 context.Context, arg1 int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountPlaceReviews", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountPlaceReviews indicates an expected call of CountPlaceReviews.
func (mr *MockStoreMockRecorder) CountPlaceReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountPlaceReviews", reflect.TypeOf((*MockStore)(nil).CountPlaceReviews), arg0, arg1)
}

// CreateDoctor mocks base method.
func (m *MockStore) CreateDoctor(arg0 context.Context, arg1 db.CreateDoctorParams) (db.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDoctor", arg0, arg1)
	ret0, _ := ret[0].(db.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDoctor indicates an expected call of CreateDoctor.
func (mr *MockStoreMockRecorder) CreateDoctor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDoctor", reflect.TypeOf((*MockStore)(nil).CreateDoctor), arg0, arg1)
}

// CreatePlace mocks base method.
func (m *MockStore) CreatePlace(arg0 context.Context, arg1 db.CreatePlaceParams) (db.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlace", arg0, arg1)
	ret0, _ := ret[0].(db.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlace indicates an expected call of CreatePlace.
func (mr *MockStoreMockRecorder) CreatePlace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlace", reflect.TypeOf((*MockStore)(nil).CreatePlace), arg0, arg1)
}

// CreateReview mocks base method.
func (m *MockStore) CreateReview(arg0 context.Context, arg1 db.CreateReviewParams) (db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", arg0, arg1)
	ret0, _ := ret[0].(db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockStoreMockRecorder) CreateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockStore)(nil).CreateReview), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteDoctor mocks base method.
func (m *MockStore) DeleteDoctor(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDoctor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDoctor indicates an expected call of DeleteDoctor.
func (mr *MockStoreMockRecorder) DeleteDoctor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDoctor", reflect.TypeOf((*MockStore)(nil).DeleteDoctor), arg0, arg1)
}

// DeletePlace mocks base method.
func (m *MockStore) DeletePlace(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlace indicates an expected call of DeletePlace.
func (mr *MockStoreMockRecorder) DeletePlace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlace", reflect.TypeOf((*MockStore)(nil).DeletePlace), arg0, arg1)
}

// DeleteReview mocks base method.
func (m *MockStore) DeleteReview(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockStoreMockRecorder) DeleteReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockStore)(nil).DeleteReview), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetDoctor mocks base method.
func (m *MockStore) GetDoctor(arg0 context.Context, arg1 int64) (db.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDoctor", arg0, arg1)
	ret0, _ := ret[0].(db.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDoctor indicates an expected call of GetDoctor.
func (mr *MockStoreMockRecorder) GetDoctor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctor", reflect.TypeOf((*MockStore)(nil).GetDoctor), arg0, arg1)
}

// GetPlace mocks base method.
func (m *MockStore) GetPlace(arg0 context.Context, arg1 int64) (db.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlace", arg0, arg1)
	ret0, _ := ret[0].(db.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlace indicates an expected call of GetPlace.
func (mr *MockStoreMockRecorder) GetPlace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlace", reflect.TypeOf((*MockStore)(nil).GetPlace), arg0, arg1)
}

// GetReview mocks base method.
func (m *MockStore) GetReview(arg0 context.Context, arg1 int64) (db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", arg0, arg1)
	ret0, _ := ret[0].(db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockStoreMockRecorder) GetReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockStore)(nil).GetReview), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListDoctorReviews mocks base method.
func (m *MockStore) ListDoctorReviews(arg0 context.Context, arg1 int32) ([]db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDoctorReviews", arg0, arg1)
	ret0, _ := ret[0].([]db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDoctorReviews indicates an expected call of ListDoctorReviews.
func (mr *MockStoreMockRecorder) ListDoctorReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDoctorReviews", reflect.TypeOf((*MockStore)(nil).ListDoctorReviews), arg0, arg1)
}

// ListDoctors mocks base method.
func (m *MockStore) ListDoctors(arg0 context.Context, arg1 db.ListDoctorsParams) ([]db.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDoctors", arg0, arg1)
	ret0, _ := ret[0].([]db.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDoctors indicates an expected call of ListDoctors.
func (mr *MockStoreMockRecorder) ListDoctors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDoctors", reflect.TypeOf((*MockStore)(nil).ListDoctors), arg0, arg1)
}

// ListPlaces mocks base method.
func (m *MockStore) ListPlaces(arg0 context.Context) ([]db.Place, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaces", arg0)
	ret0, _ := ret[0].([]db.Place)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaces indicates an expected call of ListPlaces.
func (mr *MockStoreMockRecorder) ListPlaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaces", reflect.TypeOf((*MockStore)(nil).ListPlaces), arg0)
}

// ListReviews mocks base method.
func (m *MockStore) ListReviews(arg0 context.Context) ([]db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReviews", arg0)
	ret0, _ := ret[0].([]db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReviews indicates an expected call of ListReviews.
func (mr *MockStoreMockRecorder) ListReviews(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReviews", reflect.TypeOf((*MockStore)(nil).ListReviews), arg0)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0)
}

// ReviewTx mocks base method.
func (m *MockStore) ReviewTx(arg0 context.Context, arg1 db.CreateReviewParams) (db.ReviewTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReviewTx", arg0, arg1)
	ret0, _ := ret[0].(db.ReviewTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReviewTx indicates an expected call of ReviewTx.
func (mr *MockStoreMockRecorder) ReviewTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReviewTx", reflect.TypeOf((*MockStore)(nil).ReviewTx), arg0, arg1)
}

// UpdateDoctor mocks base method.
func (m *MockStore) UpdateDoctor(arg0 context.Context, arg1 db.UpdateDoctorParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDoctor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDoctor indicates an expected call of UpdateDoctor.
func (mr *MockStoreMockRecorder) UpdateDoctor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDoctor", reflect.TypeOf((*MockStore)(nil).UpdateDoctor), arg0, arg1)
}

// UpdatePlace mocks base method.
func (m *MockStore) UpdatePlace(arg0 context.Context, arg1 db.UpdatePlaceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlace indicates an expected call of UpdatePlace.
func (mr *MockStoreMockRecorder) UpdatePlace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlace", reflect.TypeOf((*MockStore)(nil).UpdatePlace), arg0, arg1)
}

// UpdateReview mocks base method.
func (m *MockStore) UpdateReview(arg0 context.Context, arg1 db.UpdateReviewParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockStoreMockRecorder) UpdateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockStore)(nil).UpdateReview), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
