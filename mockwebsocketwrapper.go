// Code generated by MockGen. DO NOT EDIT.
// Source: websocketwrapper.go

// Package websocketwrapper is a generated GoMock package.
package websocketwrapper

import (
	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
	io "io"
	net "net"
	http "net/http"
	reflect "reflect"
	time "time"
)

// MockWebsocketUpgrader is a mock of WebsocketUpgrader interface
type MockWebsocketUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockWebsocketUpgraderMockRecorder
}

// MockWebsocketUpgraderMockRecorder is the mock recorder for MockWebsocketUpgrader
type MockWebsocketUpgraderMockRecorder struct {
	mock *MockWebsocketUpgrader
}

// NewMockWebsocketUpgrader creates a new mock instance
func NewMockWebsocketUpgrader(ctrl *gomock.Controller) *MockWebsocketUpgrader {
	mock := &MockWebsocketUpgrader{ctrl: ctrl}
	mock.recorder = &MockWebsocketUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWebsocketUpgrader) EXPECT() *MockWebsocketUpgraderMockRecorder {
	return m.recorder
}

// Upgrade mocks base method
func (m *MockWebsocketUpgrader) Upgrade(w http.ResponseWriter, r *http.Request, h http.Header) (WebsocketConn, error) {
	ret := m.ctrl.Call(m, "Upgrade", w, r, h)
	ret0, _ := ret[0].(WebsocketConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade
func (mr *MockWebsocketUpgraderMockRecorder) Upgrade(w, r, h interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockWebsocketUpgrader)(nil).Upgrade), w, r, h)
}

// MockWebsocketConn is a mock of WebsocketConn interface
type MockWebsocketConn struct {
	ctrl     *gomock.Controller
	recorder *MockWebsocketConnMockRecorder
}

// MockWebsocketConnMockRecorder is the mock recorder for MockWebsocketConn
type MockWebsocketConnMockRecorder struct {
	mock *MockWebsocketConn
}

// NewMockWebsocketConn creates a new mock instance
func NewMockWebsocketConn(ctrl *gomock.Controller) *MockWebsocketConn {
	mock := &MockWebsocketConn{ctrl: ctrl}
	mock.recorder = &MockWebsocketConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWebsocketConn) EXPECT() *MockWebsocketConnMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockWebsocketConn) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWebsocketConnMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWebsocketConn)(nil).Close))
}

// CloseHandler mocks base method
func (m *MockWebsocketConn) CloseHandler() func(int, string) error {
	ret := m.ctrl.Call(m, "CloseHandler")
	ret0, _ := ret[0].(func(int, string) error)
	return ret0
}

// CloseHandler indicates an expected call of CloseHandler
func (mr *MockWebsocketConnMockRecorder) CloseHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseHandler", reflect.TypeOf((*MockWebsocketConn)(nil).CloseHandler))
}

// EnableWriteCompression mocks base method
func (m *MockWebsocketConn) EnableWriteCompression(enable bool) {
	m.ctrl.Call(m, "EnableWriteCompression", enable)
}

// EnableWriteCompression indicates an expected call of EnableWriteCompression
func (mr *MockWebsocketConnMockRecorder) EnableWriteCompression(enable interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableWriteCompression", reflect.TypeOf((*MockWebsocketConn)(nil).EnableWriteCompression), enable)
}

// LocalAddr mocks base method
func (m *MockWebsocketConn) LocalAddr() net.Addr {
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr
func (mr *MockWebsocketConnMockRecorder) LocalAddr() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockWebsocketConn)(nil).LocalAddr))
}

// NextReader mocks base method
func (m *MockWebsocketConn) NextReader() (int, io.Reader, error) {
	ret := m.ctrl.Call(m, "NextReader")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(io.Reader)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NextReader indicates an expected call of NextReader
func (mr *MockWebsocketConnMockRecorder) NextReader() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextReader", reflect.TypeOf((*MockWebsocketConn)(nil).NextReader))
}

// NextWriter mocks base method
func (m *MockWebsocketConn) NextWriter(messageType int) (io.WriteCloser, error) {
	ret := m.ctrl.Call(m, "NextWriter", messageType)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextWriter indicates an expected call of NextWriter
func (mr *MockWebsocketConnMockRecorder) NextWriter(messageType interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextWriter", reflect.TypeOf((*MockWebsocketConn)(nil).NextWriter), messageType)
}

// PingHandler mocks base method
func (m *MockWebsocketConn) PingHandler() func(string) error {
	ret := m.ctrl.Call(m, "PingHandler")
	ret0, _ := ret[0].(func(string) error)
	return ret0
}

// PingHandler indicates an expected call of PingHandler
func (mr *MockWebsocketConnMockRecorder) PingHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingHandler", reflect.TypeOf((*MockWebsocketConn)(nil).PingHandler))
}

// PongHandler mocks base method
func (m *MockWebsocketConn) PongHandler() func(string) error {
	ret := m.ctrl.Call(m, "PongHandler")
	ret0, _ := ret[0].(func(string) error)
	return ret0
}

// PongHandler indicates an expected call of PongHandler
func (mr *MockWebsocketConnMockRecorder) PongHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PongHandler", reflect.TypeOf((*MockWebsocketConn)(nil).PongHandler))
}

// ReadJSON mocks base method
func (m *MockWebsocketConn) ReadJSON(v interface{}) error {
	ret := m.ctrl.Call(m, "ReadJSON", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadJSON indicates an expected call of ReadJSON
func (mr *MockWebsocketConnMockRecorder) ReadJSON(v interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadJSON", reflect.TypeOf((*MockWebsocketConn)(nil).ReadJSON), v)
}

// ReadMessage mocks base method
func (m *MockWebsocketConn) ReadMessage() (int, []byte, error) {
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage
func (mr *MockWebsocketConnMockRecorder) ReadMessage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockWebsocketConn)(nil).ReadMessage))
}

// RemoteAddr mocks base method
func (m *MockWebsocketConn) RemoteAddr() net.Addr {
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockWebsocketConnMockRecorder) RemoteAddr() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockWebsocketConn)(nil).RemoteAddr))
}

// SetCloseHandler mocks base method
func (m *MockWebsocketConn) SetCloseHandler(h func(int, string) error) {
	m.ctrl.Call(m, "SetCloseHandler", h)
}

// SetCloseHandler indicates an expected call of SetCloseHandler
func (mr *MockWebsocketConnMockRecorder) SetCloseHandler(h interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCloseHandler", reflect.TypeOf((*MockWebsocketConn)(nil).SetCloseHandler), h)
}

// SetCompressionLevel mocks base method
func (m *MockWebsocketConn) SetCompressionLevel(level int) error {
	ret := m.ctrl.Call(m, "SetCompressionLevel", level)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCompressionLevel indicates an expected call of SetCompressionLevel
func (mr *MockWebsocketConnMockRecorder) SetCompressionLevel(level interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCompressionLevel", reflect.TypeOf((*MockWebsocketConn)(nil).SetCompressionLevel), level)
}

// SetPingHandler mocks base method
func (m *MockWebsocketConn) SetPingHandler(h func(string) error) {
	m.ctrl.Call(m, "SetPingHandler", h)
}

// SetPingHandler indicates an expected call of SetPingHandler
func (mr *MockWebsocketConnMockRecorder) SetPingHandler(h interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPingHandler", reflect.TypeOf((*MockWebsocketConn)(nil).SetPingHandler), h)
}

// SetPongHandler mocks base method
func (m *MockWebsocketConn) SetPongHandler(h func(string) error) {
	m.ctrl.Call(m, "SetPongHandler", h)
}

// SetPongHandler indicates an expected call of SetPongHandler
func (mr *MockWebsocketConnMockRecorder) SetPongHandler(h interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPongHandler", reflect.TypeOf((*MockWebsocketConn)(nil).SetPongHandler), h)
}

// SetReadDeadline mocks base method
func (m *MockWebsocketConn) SetReadDeadline(t time.Time) error {
	ret := m.ctrl.Call(m, "SetReadDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockWebsocketConnMockRecorder) SetReadDeadline(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockWebsocketConn)(nil).SetReadDeadline), t)
}

// SetReadLimit mocks base method
func (m *MockWebsocketConn) SetReadLimit(limit int64) {
	m.ctrl.Call(m, "SetReadLimit", limit)
}

// SetReadLimit indicates an expected call of SetReadLimit
func (mr *MockWebsocketConnMockRecorder) SetReadLimit(limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadLimit", reflect.TypeOf((*MockWebsocketConn)(nil).SetReadLimit), limit)
}

// SetWriteDeadline mocks base method
func (m *MockWebsocketConn) SetWriteDeadline(t time.Time) error {
	ret := m.ctrl.Call(m, "SetWriteDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockWebsocketConnMockRecorder) SetWriteDeadline(t interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockWebsocketConn)(nil).SetWriteDeadline), t)
}

// Subprotocol mocks base method
func (m *MockWebsocketConn) Subprotocol() string {
	ret := m.ctrl.Call(m, "Subprotocol")
	ret0, _ := ret[0].(string)
	return ret0
}

// Subprotocol indicates an expected call of Subprotocol
func (mr *MockWebsocketConnMockRecorder) Subprotocol() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subprotocol", reflect.TypeOf((*MockWebsocketConn)(nil).Subprotocol))
}

// UnderlyingConn mocks base method
func (m *MockWebsocketConn) UnderlyingConn() net.Conn {
	ret := m.ctrl.Call(m, "UnderlyingConn")
	ret0, _ := ret[0].(net.Conn)
	return ret0
}

// UnderlyingConn indicates an expected call of UnderlyingConn
func (mr *MockWebsocketConnMockRecorder) UnderlyingConn() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnderlyingConn", reflect.TypeOf((*MockWebsocketConn)(nil).UnderlyingConn))
}

// WriteControl mocks base method
func (m *MockWebsocketConn) WriteControl(messageType int, data []byte, deadline time.Time) error {
	ret := m.ctrl.Call(m, "WriteControl", messageType, data, deadline)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteControl indicates an expected call of WriteControl
func (mr *MockWebsocketConnMockRecorder) WriteControl(messageType, data, deadline interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteControl", reflect.TypeOf((*MockWebsocketConn)(nil).WriteControl), messageType, data, deadline)
}

// WriteJSON mocks base method
func (m *MockWebsocketConn) WriteJSON(v interface{}) error {
	ret := m.ctrl.Call(m, "WriteJSON", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteJSON indicates an expected call of WriteJSON
func (mr *MockWebsocketConnMockRecorder) WriteJSON(v interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteJSON", reflect.TypeOf((*MockWebsocketConn)(nil).WriteJSON), v)
}

// WriteMessage mocks base method
func (m *MockWebsocketConn) WriteMessage(messageType int, data []byte) error {
	ret := m.ctrl.Call(m, "WriteMessage", messageType, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage
func (mr *MockWebsocketConnMockRecorder) WriteMessage(messageType, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockWebsocketConn)(nil).WriteMessage), messageType, data)
}

// WritePreparedMessage mocks base method
func (m *MockWebsocketConn) WritePreparedMessage(pm *websocket.PreparedMessage) error {
	ret := m.ctrl.Call(m, "WritePreparedMessage", pm)
	ret0, _ := ret[0].(error)
	return ret0
}

// WritePreparedMessage indicates an expected call of WritePreparedMessage
func (mr *MockWebsocketConnMockRecorder) WritePreparedMessage(pm interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritePreparedMessage", reflect.TypeOf((*MockWebsocketConn)(nil).WritePreparedMessage), pm)
}
