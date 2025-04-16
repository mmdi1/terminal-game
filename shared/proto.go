package shared

import (
	"encoding/json"
)

// 消息类型
const (
	MsgTypeLogin      = "login"       // 玩家登录
	MsgTypeWordAssign = "word_assign" // 词语分配
	MsgTypeDescribe   = "describe"    // 玩家描述
	MsgTypeVote       = "vote"        // 玩家投票
	MsgTypeGameState  = "game_state"  // 游戏状态更新
	MsgTypeGameResult = "game_result" // 游戏结果
)

// 消息结构体
type Message struct {
	Type string      `json:"type"` // 消息类型
	Data interface{} `json:"data"` // 消息数据
}

// 序列化消息
func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

// 反序列化消息
func Deserialize(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}

// 登录请求
type LoginRequest struct {
	PlayerID string `json:"player_id"` // 玩家ID
}

// 词语分配
type WordAssign struct {
	PlayerID string `json:"player_id"` // 玩家ID
	Word     string `json:"word"`      // 分配的词语
}

// 游戏操作请求
type GameAction struct {
	PlayerID string      `json:"player_id"` // 玩家ID
	RoomID   string      `json:"room_id"`   // 房间ID
	Action   string      `json:"action"`    // 操作类型
	Data     interface{} `json:"data"`      // 操作数据
}

// 描述请求
type DescribeRequest struct {
	PlayerID    string `json:"player_id"`   // 玩家ID
	Description string `json:"description"` // 描述内容
}

// 投票请求
type VoteRequest struct {
	PlayerID       string `json:"player_id"`        // 投票者ID
	TargetPlayerID string `json:"target_player_id"` // 被投票玩家ID
}

// 游戏状态
type GameState struct {
	Stage        string         `json:"stage"`        // 当前阶段 (waiting, describing, voting, result)
	Players      []PlayerState  `json:"players"`      // 玩家状态
	Descriptions []Description  `json:"descriptions"` // 当前轮描述
	Votes        map[string]int `json:"votes"`        // 投票结果
	Round        int            `json:"round"`        // 当前轮次
}

type PlayerState struct {
	PlayerID string `json:"player_id"` // 玩家ID
	Alive    bool   `json:"alive"`     // 是否存活
}

type Description struct {
	PlayerID    string `json:"player_id"`   // 玩家ID
	Description string `json:"description"` // 描述内容
}

// 游戏结果
type GameResult struct {
	Winner string `json:"winner"` // 胜利方 (civilian/undercover)
}
