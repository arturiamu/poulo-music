package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Music struct {
	ID           uint       `json:"id" gorm:"primary_key"`        //音乐ID
	BVId         string     `json:"bvid" gorm:"column:bvid"`      //B站视频BV 号
	Name         string     `json:"name"`                         //音乐名称 B站视频名称
	Title        string     `json:"title"`                        //音乐标题 用户自定义
	Url          string     `json:"url"`                          //音乐地址
	Author       string     `json:"author"`                       //音乐作者
	Pic          string     `json:"pic"`                          //音乐封面
	Lrc          string     `json:"lrc"`                          //音乐歌词
	Theme        string     `json:"theme"`                        //音乐主题
	TimeSkips    []TimeSkip `json:"time_skips" gorm:"-"`          //时间点跳过
	TimeSkipsStr string     `json:"-"`                            //时间点跳过 db json字符串
	Hash         string     `json:"hash" gorm:"type:varchar(32)"` //音乐hash,由 BVId 和 TimeSkips 生成
	CreatedAt    time.Time  `json:"create_at" gorm:"AUTO_CREATE"` //创建时间
	UpdatedAt    time.Time  `json:"update_at" gorm:"AUTO_UPDATE"` //更新时间
}

func (m *Music) ToHash() string {
	data := m.BVId + m.TimeSkipsStr
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

type TimeSkip struct {
	Timestamp int64 `json:"timestamp"` //跳过的时间点 单位毫秒
	Time      int64 `json:"time"`      //跳过的时间 单位毫秒
}

type Playlist struct {
	ID        uint      `json:"id" gorm:"primary_key"`        //播放列表ID
	Name      string    `json:"name"`                         //播放列表名称
	CreatedAt time.Time `json:"create_at" gorm:"AUTO_CREATE"` //创建时间
	UpdatedAt time.Time `json:"update_at" gorm:"AUTO_UPDATE"` //更新时间
}

type PlaylistItem struct {
	ID         uint      `json:"id" gorm:"primary_key"`        //ID
	PlaylistID uint      `json:"playlist_id"`                  //播放列表项ID
	MusicID    uint      `json:"music_id"`                     //音乐ID
	CreatedAt  time.Time `json:"create_at" gorm:"AUTO_CREATE"` //创建时间 对应音乐加入播放列表的时间
}
