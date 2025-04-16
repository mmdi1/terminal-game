package core

import (
	"sync"
	"terminal/shared/types"
)

// 房间管理器
type roomManager struct {
	rooms      map[string]*Room
	plyaerRoom sync.Map //记录玩家的room id
	sync.RWMutex
}

func NewRoomManager() *roomManager {
	return &roomManager{
		rooms: make(map[string]*Room),
	}
}

// 检查是否已在房间
func (r *roomManager) CheckInRoom(pid string) *Room {
	r.RLock()
	defer r.RUnlock()
	rid, ok := r.plyaerRoom.Load(pid)
	if ok {
		return r.rooms[rid.(string)]
	}
	return nil
}

// 新增房间
func (r *roomManager) Get(pid string) *Room {
	r.RLock()
	defer r.RUnlock()
	rid, ok := r.plyaerRoom.Load(pid)
	if ok {
		return r.rooms[rid.(string)]
	}
	return nil
}

// 新增房间
func (r *roomManager) Add(room *Room) {
	r.Lock()
	defer r.Unlock()
	r.rooms[room.OwnerID] = room
	r.plyaerRoom.Store(room.OwnerID, room.ID)
}

// 解散房间
func (r *roomManager) Dissolution(rid string) {
	r.Lock()
	defer r.Unlock()
	room, ok := r.rooms[rid]
	if ok {
		for _, v := range room.Players {
			v.Conn.SendBuffMsg(types.R_InfoMsg, []byte("房间已解散."))
			r.plyaerRoom.Delete(v.ID)
		}
		delete(r.rooms, rid)
	}
}
