package game

import (
	"math/rand"
	"sync"
	"terminal/server/core/iface"
	"terminal/shared"
	"terminal/shared/types"
)

// 词语对
type WordPair struct {
	CivilianWord   string // 平民词语
	UndercoverWord string // 卧底词语
}

type whoWdGameExtension struct {
	Word string
}

// 谁是卧底
type WhoWdGame struct {
	maxPlayer    int
	players      map[string]*whoWdGameExtension
	wordPairs    []WordPair
	currentRound int
	currentStage string
	descriptions []shared.Description
	votes        map[string]int
	eventBus     shared.EventBus
	mutex        sync.Mutex
}

func NewWhoWdGame(maxPlayer int) *WhoWdGame {
	return &WhoWdGame{
		wordPairs: []WordPair{
			{"苹果", "香蕉"},
			{"猫", "狗"},
			{"汽车", "摩托车"},
			{"书", "笔记本"},
		},
		currentRound: 0,
		currentStage: "waiting",
		maxPlayer:    maxPlayer,
		votes:        make(map[string]int),
	}
}

func (g *WhoWdGame) Init(eventBus shared.EventBus, players map[string]*iface.Player) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.eventBus = eventBus
	if len(players) == g.maxPlayer {
		g.currentStage = "describing"
		g.currentRound = 1
		g.assignWords()
	}
}

// 分配词语
func (gm *WhoWdGame) assignWords() {
	wordPair := gm.wordPairs[rand.Intn(len(gm.wordPairs))]
	undercoverIndex := rand.Intn(gm.maxPlayer)
	playerIDs := make([]string, 0, len(gm.players))
	for id := range gm.players {
		playerIDs = append(playerIDs, id)
	}
	for i, id := range playerIDs {
		word := wordPair.CivilianWord
		if i == undercoverIndex {
			word = wordPair.UndercoverWord
		}
		gm.players[id].Word = word
		msg := shared.Message{
			Type: shared.MsgTypeWordAssign,
			Data: shared.WordAssign{PlayerID: id, Word: word},
		}
		data, _ := msg.Serialize()
		gm.eventBus.Emit(shared.GameEvent{
			Type:     "",
			PlayerID: id,
			MsgType:  types.R_InfoMsg,
			Data:     data,
		})
	}
}

// 处理操作
func (g *WhoWdGame) HandleAction(action shared.GameAction) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch action.Action {
	case "describe":
		var desc string
		if data, ok := action.Data.(map[string]interface{}); ok {
			desc = data["description"].(string)
		}
		g.descriptions = append(g.descriptions, shared.Description{
			PlayerID:    action.PlayerID,
			Description: desc,
		})
	case "vote":
		var targetID string
		if data, ok := action.Data.(map[string]interface{}); ok {
			targetID = data["target_player_id"].(string)
		}
		g.votes[targetID]++
	}
}

// 获取状态
func (g *WhoWdGame) GetState() shared.GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// state := UndercoverState{
	// 	Stage:        g.currentStage,
	// 	Players:      make([]shared.PlayerState, 0),
	// 	Descriptions: g.descriptions,
	// 	Votes:        g.votes,
	// 	Round:        g.currentRound,
	// }
	return shared.GameState{
		// GameType: shared.GameTypeUndercover,
		// Stage:    g.currentStage,
		// Data:     state,
	}
}

// 广播状态
func (g *WhoWdGame) BroadcastState() {
	state := g.GetState()
	msg := shared.Message{
		Type: shared.MsgTypeGameState,
		Data: state,
	}
	data, _ := msg.Serialize()
	g.eventBus.Emit(shared.GameEvent{
		Type:    types.EvenetTypeBroadcast,
		MsgType: types.R_InfoMsg,
		Data:    data,
	})
}
