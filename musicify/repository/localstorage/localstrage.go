package localstorage

import (
	"context"
	"gorm.io/gorm"
	"poulo-music/models"
	"poulo-music/musicify"
)

var _ musicify.MusicRepository = (*MusicLocalstorage)(nil)

type MusicLocalstorage struct {
	db *gorm.DB
}

func NewMusicLocalstorage(db *gorm.DB) *MusicLocalstorage {
	return &MusicLocalstorage{db: db}
}

func (m *MusicLocalstorage) AutoMigrate() error {
	return m.db.AutoMigrate(models.Music{}, models.Playlist{}, models.PlaylistItem{})
}

func (m *MusicLocalstorage) SaveMusic(ctx context.Context, music *models.Music) (*models.Music, error) {
	return music, m.db.Save(&music).Error
}

func (m *MusicLocalstorage) RemoveMusic(ctx context.Context, id int64) (*models.Music, error) {
	return nil, m.db.Delete(&models.Music{ID: id}).Error
}

func (m *MusicLocalstorage) ModifyMusic(ctx context.Context, music *models.Music) (*models.Music, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MusicLocalstorage) GetMusicByID(ctx context.Context, id int64) (*models.Music, error) {
	var misic = models.Music{ID: id}
	return &misic, m.db.First(&misic).Error
}

func (m *MusicLocalstorage) ListMusicByBVId(ctx context.Context, bvid string) ([]*models.Music, error) {
	var musics []*models.Music
	return musics, m.db.Find(&musics, "bvid = ?", bvid).Error
}

func (m *MusicLocalstorage) GetMusicByHash(ctx context.Context, hash string) (*models.Music, error) {
	var misic = models.Music{}
	return &misic, m.db.First(&misic, "hash = ?", hash).Error
}

func (m *MusicLocalstorage) ListAllMusic(ctx context.Context) ([]*models.Music, error) {
	var musics []*models.Music
	return musics, m.db.Find(&musics).Error
}

func (m *MusicLocalstorage) ListMusicByPlaylistID(ctx context.Context, playlistId int64) ([]*models.Music, error) {
	var musics []*models.Music
	return musics, m.db.Table("musics m").
		Select("m.id,m.bv_id,m.name,m.time_skips_str,m.hash,m.created_at,m.updated_at").
		Joins("join playlist_items pi on pi.music_id = m.id and pi.playlist_id = ?", playlistId).
		Scan(&musics).Error
}

func (m *MusicLocalstorage) SavePlaylist(ctx context.Context, playList *models.Playlist) (*models.Playlist, error) {
	return playList, m.db.Save(&playList).Error
}

func (m *MusicLocalstorage) RemovePlaylist(ctx context.Context, id int64) (*models.Playlist, error) {
	return nil, m.db.Delete(&models.Playlist{ID: id}).Error
}

func (m *MusicLocalstorage) ModifyPlaylist(ctx context.Context, playlist *models.Playlist) (*models.Playlist, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MusicLocalstorage) GetPlaylistByID(ctx context.Context, id int64) (*models.Playlist, error) {
	var playlist = models.Playlist{ID: id}
	return &playlist, m.db.First(&playlist).Error
}

func (m *MusicLocalstorage) GetPlaylistByName(ctx context.Context, name string) (*models.Playlist, error) {
	var playlist models.Playlist
	return &playlist, m.db.First(&playlist, "name=?", name).Error
}

func (m *MusicLocalstorage) SavePlaylistItem(ctx context.Context, playListItem *models.PlaylistItem) (*models.PlaylistItem, error) {
	return playListItem, m.db.Save(&playListItem).Error
}

func (m *MusicLocalstorage) RemovePlaylistItem(ctx context.Context, id int64) (*models.PlaylistItem, error) {
	return nil, m.db.Delete(&models.PlaylistItem{ID: id}).Error
}
