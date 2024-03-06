package platform

type MusicPlatform interface {
	GetPopularContent() (any, error)
	GetRecommendations() (any, error)
	GetSearchResults(keyword string) (any, error)
	GetMusicInfo(identifier string) (any, error)
}
