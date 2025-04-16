package room

import (
	"terminal/server/core"
	"terminal/server/core/game"
	"terminal/shared/types"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type CreateRoomRouter struct {
	znet.BaseRouter
}

func (cr *CreateRoomRouter) Handle(req ziface.IRequest) {
	ply := core.NewPlayer(req.GetConnection())
	room := core.GRoomManager.CheckInRoom(ply.ID)
	if room != nil {
		req.GetConnection().SendBuffMsg(types.R_WarningMsg, []byte("已有房间."))
	}
	room = core.NewRoom(ply)
	game := game.NewWhoWdGame(7)
	game.Init(room, room.Players)
	core.GRoomManager.Add(room)
}
