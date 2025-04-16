package core

import (
	"fmt"
	"terminal/server/core/iface"

	"github.com/aceld/zinx/ziface"
)

func NewPlayer(conn ziface.IConnection) *iface.Player {
	return &iface.Player{
		ID:   fmt.Sprintf(`ply:%v`, conn.GetConnID()),
		Conn: conn,
	}
}
