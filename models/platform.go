package models

type GetSearchParam struct {
	Platform   Platform
	Keyword    string
	Page       int64
	Pagesize   int64
	CachePic   bool
	CacheMusic bool
	CacheLrc   bool
}

// GetSearchResp
// 统一搜索结果
type GetSearchResp struct {
	ID         string   `json:"id"`
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
	CachePic   bool
	CacheMusic bool
	CacheLrc   bool
}

type GetHotContentParam struct {
	Platform   Platform
	Identifier string
	CachePic   bool
	CacheMusic bool
	CacheLrc   bool
}

type GetHotContentResp struct {
	Platform   Platform
	Identifier string
}
