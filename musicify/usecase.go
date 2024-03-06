package musicify

import (
	"context"
	"poulo-music/models"
)

type UseCase interface {
	SaveMusic(ctx context.Context, music *models.Music) (*models.Music, error)
	RemoveMusic(ctx context.Context, id int64) (*models.Music, error)
	ModifyMusic(ctx context.Context, music *models.Music) (*models.Music, error)
	GetMusicByID(ctx context.Context, id int64) (*models.Music, error)
	ListMusicByBVId(ctx context.Context, bvid string) ([]*models.Music, error)
	GetMusicByHash(ctx context.Context, hash string) (*models.Music, error)
	ListAllMusic(ctx context.Context) ([]*models.Music, error)
	ListMusicByPlaylistID(ctx context.Context, playlistId int64) ([]*models.Music, error)

	SavePlaylist(ctx context.Context, playList *models.Playlist) (*models.Playlist, error)
	RemovePlaylist(ctx context.Context, id int64) (*models.Playlist, error)
	ModifyPlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error)
	GetPlaylistByID(ctx context.Context, id int64) (*models.Playlist, error)
	GetPlaylistByName(ctx context.Context, name string) (*models.Playlist, error)

	SavePlaylistItem(ctx context.Context, playListItem *models.PlaylistItem) (*models.PlaylistItem, error)
	RemovePlaylistItem(ctx context.Context, id int64) (*models.PlaylistItem, error)
}
