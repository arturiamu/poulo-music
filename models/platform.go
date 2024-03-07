package models

type GetSearchParam struct {
	Keyword      string
	Page         int64
	Pagesize     int64
	SaveFile     bool
	SaveFilePath string
}

// GetSearchResp
// 统一搜索结果
type GetSearchResp struct {
	ID         int64    `json:"id"`
	Platform   Platform `json:"platform"`   // 平台
	Identifier string   `json:"identifier"` // 平台标识符
	Title      string   `json:"title"`      // 内容标题
	Name       string   `json:"name"`       // 内容名称
	Describe   string   `json:"describe"`   // 内容描述
	Type       string   `json:"type"`       // 内容类型
	Cover      string   `json:"cover"`      // 内容封面
}

type GetMusicParam struct {
	Platform   Platform
	Identifier string
}

type GetHotContentParam struct {
	Platform   Platform
	Identifier string
}

type GetHotContentResp struct {
	Platform   Platform
	Identifier string
}
