package musicify

import (
	"context"
	"github.com/sirupsen/logrus"
	"poulo-music/models"
	"poulo-music/musicify"
)

var _ musicify.UseCase = (*UseCase)(nil)

type UseCase struct {
	log       *logrus.Logger
	musicRepo musicify.MusicRepository
}

func NewUseCase(log *logrus.Logger, musicRepo musicify.MusicRepository) *UseCase {
	return &UseCase{log: log, musicRepo: musicRepo}
}

func (u UseCase) SaveMusic(ctx context.Context, music *models.Music) (*models.Music, error) {
	return u.musicRepo.SaveMusic(ctx, music)
}

func (u UseCase) RemoveMusic(ctx context.Context, id int64) (*models.Music, error) {
	return u.musicRepo.RemoveMusic(ctx, id)
}

func (u UseCase) ModifyMusic(ctx context.Context, music *models.Music) (*models.Music, error) {
	return u.musicRepo.ModifyMusic(ctx, music)
}

func (u UseCase) GetMusicByID(ctx context.Context, id int64) (*models.Music, error) {
	return u.musicRepo.GetMusicByID(ctx, id)
}

func (u UseCase) ListMusicByBVId(ctx context.Context, bvid string) ([]*models.Music, error) {
	return u.musicRepo.ListMusicByBVId(ctx, bvid)
}

func (u UseCase) GetMusicByHash(ctx context.Context, hash string) (*models.Music, error) {
	return u.musicRepo.GetMusicByHash(ctx, hash)
}

func (u UseCase) ListAllMusic(ctx context.Context) ([]*models.Music, error) {
	return u.musicRepo.ListAllMusic(ctx)
}

func (u UseCase) ListMusicByPlaylistID(ctx context.Context, playlistId int64) ([]*models.Music, error) {
	return u.musicRepo.ListMusicByPlaylistID(ctx, playlistId)
}

func (u UseCase) SavePlaylist(ctx context.Context, playList *models.Playlist) (*models.Playlist, error) {
	return u.musicRepo.SavePlaylist(ctx, playList)
}

func (u UseCase) RemovePlaylist(ctx context.Context, id int64) (*models.Playlist, error) {
	return u.musicRepo.RemovePlaylist(ctx, id)
}

func (u UseCase) ModifyPlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	return u.musicRepo.ModifyPlaylist(ctx, playlist)
}

func (u UseCase) GetPlaylistByID(ctx context.Context, id int64) (*models.Playlist, error) {
	return u.musicRepo.GetPlaylistByID(ctx, id)
}

func (u UseCase) GetPlaylistByName(ctx context.Context, name string) (*models.Playlist, error) {
	return u.musicRepo.GetPlaylistByName(ctx, name)
}

func (u UseCase) SavePlaylistItem(ctx context.Context, playListItem *models.PlaylistItem) (*models.PlaylistItem, error) {
	return u.musicRepo.SavePlaylistItem(ctx, playListItem)
}

func (u UseCase) RemovePlaylistItem(ctx context.Context, id int64) (*models.PlaylistItem, error) {
	return u.musicRepo.RemovePlaylistItem(ctx, id)
}
