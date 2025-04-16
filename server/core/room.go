package core

import (
	"sync"
	"terminal/server/core/iface"
	"terminal/shared"
	"terminal/shared/types"
)

// 房间
type Room struct {
	ID      string
	OwnerID string
	Players map[string]*iface.Player
	Game    iface.Game
	Mutex   sync.Mutex
}

func NewRoom(ply *iface.Player) *Room {
	room := &Room{
		ID:      ply.ID,
		OwnerID: ply.ID,
		Players: map[string]*iface.Player{ply.ID: ply},
	}
	return room
}
func (r *Room) ProvideGame(game iface.Game) {
	r.Game = game
}
func (r *Room) Broadcast() {}
func (r *Room) Emit(event shared.GameEvent) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	switch event.Type {
	case types.EventTypeSendToPlayer:
		if player, ok := r.Players[event.PlayerID]; ok {
			player.Conn.SendBuffMsg(event.MsgType, event.Data)
		}
	case types.EvenetTypeBroadcast:
		for _, player := range r.Players {
			player.Conn.SendBuffMsg(event.MsgType, event.Data)
		}
	}
}
