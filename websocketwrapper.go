package websocketwrapper

import (
	"io"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WebsocketUpgrader is an interface implemented by websocketUpgraderWrapper to wrap websocket.Upgrader.
// Create a NewMockWebsocketUpgrader(*gomock.Controller) to write unit tests,
// or New(*websocket.Upgrader) for real use.
type WebsocketUpgrader interface {
	Upgrade(w http.ResponseWriter, r *http.Request, h http.Header) (WebsocketConn, error)
}

type websocketUpgraderWrapper struct {
	Upgrader *websocket.Upgrader
}

func (wuw websocketUpgraderWrapper) Upgrade(w http.ResponseWriter, r *http.Request, h http.Header) (WebsocketConn, error) {
	return wuw.Upgrader.Upgrade(w, r, h)
}

// New wraps a *websocket.Upgrader to implement WebsocketUpgrader
func New(u *websocket.Upgrader) WebsocketUpgrader {
	return websocketUpgraderWrapper{Upgrader: u}
}

// WebsocketConn is the interface implemented by *websocket.Conn
type WebsocketConn interface {
	Close() error
	CloseHandler() func(code int, text string) error
	EnableWriteCompression(enable bool)
	LocalAddr() net.Addr
	NextReader() (messageType int, r io.Reader, err error)
	NextWriter(messageType int) (io.WriteCloser, error)
	PingHandler() func(appData string) error
	PongHandler() func(appData string) error
	ReadJSON(v interface{}) error
	ReadMessage() (messageType int, p []byte, err error)
	RemoteAddr() net.Addr
	SetCloseHandler(h func(code int, text string) error)
	SetCompressionLevel(level int) error
	SetPingHandler(h func(appData string) error)
	SetPongHandler(h func(appData string) error)
	SetReadDeadline(t time.Time) error
	SetReadLimit(limit int64)
	SetWriteDeadline(t time.Time) error
	Subprotocol() string
	UnderlyingConn() net.Conn
	WriteControl(messageType int, data []byte, deadline time.Time) error
	WriteJSON(v interface{}) error
	WriteMessage(messageType int, data []byte) error
	WritePreparedMessage(pm *websocket.PreparedMessage) error
}

var _ WebsocketConn = &websocket.Conn{}
