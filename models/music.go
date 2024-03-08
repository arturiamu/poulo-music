package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

type Music struct {
	ID         int64    `json:"id" gorm:"primary_key"` // 数据库存储ID
	Platform   Platform `json:"platform"`              // 音乐平台,bili,qq,netease,kugou,kuwo,migu...
	Identifier string   `json:"identifier"`            // 音乐平台标识符,如B站视频BV号,QQ音乐歌曲ID,网易云音乐歌曲ID...

	Name   string `json:"name"`   // 音乐名称
	Artist string `json:"artist"` // 音乐作者
	Url    string `json:"url"`    // 音乐地址 本地文件或网络地址
	Cover  string `json:"cover"`  // 音乐封面 播放器封面图片
	Lrc    string `json:"lrc"`    // 音乐歌词 本地文件或网络地址
	Theme  string `json:"theme"`  // 音乐主题

	Album    string `json:"album"`    // 音乐专辑
	Release  string `json:"release"`  // 音乐发行时间
	Language string `json:"language"` // 音乐语言
	Type     string `json:"type"`     // 音乐类型

	TimeSkips    []TimeSkip `json:"time_skips" gorm:"-"` // 时间点跳过
	TimeSkipsStr string     `json:"-"`                   // 时间点跳过 db json字符串

	Hash      string    `json:"hash" gorm:"type:varchar(32)"` // 音乐hash,由 Platform, Identifier 和 TimeSkips 生成
	CreatedAt time.Time `json:"create_at" gorm:"AUTO_CREATE"` // 创建时间
	UpdatedAt time.Time `json:"update_at" gorm:"AUTO_UPDATE"` // 更新时间
}

type TimeSkip struct {
	Timestamp int64 `json:"timestamp"` //跳过的时间点 单位毫秒
	Time      int64 `json:"time"`      //跳过的时间 单位毫秒
}

func (m *Music) ToHash() string {
	timeSkipsStr, _ := json.Marshal(m.TimeSkips)
	data := fmt.Sprintf("%s%s%s", m.Platform, m.Identifier, cast.ToString(timeSkipsStr))
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (m *Music) BeforeSave(*gorm.DB) (err error) {
	bytes, err := json.Marshal(m.TimeSkips)
	if err != nil {
		return err
	}

	m.TimeSkipsStr = string(bytes)

	m.Hash = m.ToHash()
	return
}

func (m *Music) AfterFind(*gorm.DB) (err error) {
	return json.Unmarshal([]byte(m.TimeSkipsStr), &m.TimeSkips)
}

type Playlist struct {
	ID        int64     `json:"id" gorm:"primary_key"`        //播放列表ID
	Name      string    `json:"name"`                         //播放列表名称
	Describe  string    `json:"describe"`                     //播放列表描述
	CreatedAt time.Time `json:"create_at" gorm:"AUTO_CREATE"` //创建时间
	UpdatedAt time.Time `json:"update_at" gorm:"AUTO_UPDATE"` //更新时间
}

type PlaylistItem struct {
	ID         int64     `json:"id" gorm:"primary_key"`        //ID
	PlaylistID int64     `json:"playlist_id"`                  //播放列表项ID
	MusicID    int64     `json:"music_id"`                     //音乐ID
	CreatedAt  time.Time `json:"create_at" gorm:"AUTO_CREATE"` //创建时间 对应音乐加入播放列表的时间
}
