// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/progress/compileui/display.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	vscode "github.com/tensorchord/MIDI/pkg/editor/vscode"
	compileui "github.com/tensorchord/MIDI/pkg/progress/compileui"
)

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockWriter) Finish() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finish")
}

// Finish indicates an expected call of Finish.
func (mr *MockWriterMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockWriter)(nil).Finish))
}

// LogVSCodePlugin mocks base method.
func (m *MockWriter) LogVSCodePlugin(p vscode.Plugin, action compileui.Action, cached bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogVSCodePlugin", p, action, cached)
}

// LogVSCodePlugin indicates an expected call of LogVSCodePlugin.
func (mr *MockWriterMockRecorder) LogVSCodePlugin(p, action, cached interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogVSCodePlugin", reflect.TypeOf((*MockWriter)(nil).LogVSCodePlugin), p, action, cached)
}

// LogZSH mocks base method.
func (m *MockWriter) LogZSH(action compileui.Action, cached bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogZSH", action, cached)
}

// LogZSH indicates an expected call of LogZSH.
func (mr *MockWriterMockRecorder) LogZSH(action, cached interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogZSH", reflect.TypeOf((*MockWriter)(nil).LogZSH), action, cached)
}
