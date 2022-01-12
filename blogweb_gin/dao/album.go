package dao

import (
	"blogweb_gin/models"
	"blogweb_gin/tools"
)

type AlbumDao struct {
	*tools.Orm
}

func NewAlbum() *AlbumDao {
	return &AlbumDao{tools.DbEngine}
}

// InsertAlbum -------插入图片---------------
func (ad *AlbumDao)InsertAlbum(album *models.Album) (int64, error) {
	return ad.Insert(album)
}

// FindAllAlbums --------查询图片----------
func (ad *AlbumDao)FindAllAlbums() ([]*models.Album, error) {
	album := new(models.Album)
	var albums []*models.Album
	rows, err := ad.Where(" id > ? ", 0).Rows(album)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(album)
		album := models.Album{Id: album.Id, FilePath: album.FilePath, FileName: album.FileName, Status: album.Status, CreateTime: album.CreateTime}
		albums = append(albums, &album)
	}
	return albums, nil
}

