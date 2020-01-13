// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/cloudsonic/sonic-server/api"
	"github.com/cloudsonic/sonic-server/domain"
	"github.com/cloudsonic/sonic-server/engine"
	"github.com/cloudsonic/sonic-server/itunesbridge"
	"github.com/cloudsonic/sonic-server/persistence"
	"github.com/cloudsonic/sonic-server/persistence/db_ledis"
	"github.com/cloudsonic/sonic-server/persistence/db_sql"
	"github.com/cloudsonic/sonic-server/persistence/db_storm"
	"github.com/cloudsonic/sonic-server/scanner"
	"github.com/deluan/gomate"
	"github.com/deluan/gomate/ledis"
	"github.com/google/wire"
)

// Injectors from wire_injectors.go:

func CreateApp(musicFolder string, p persistence.ProviderIdentifier) *App {
	provider := createPersistenceProvider(p)
	checkSumRepository := provider.CheckSumRepository
	itunesScanner := scanner.NewItunesScanner(checkSumRepository)
	mediaFileRepository := provider.MediaFileRepository
	albumRepository := provider.AlbumRepository
	artistRepository := provider.ArtistRepository
	artistIndexRepository := provider.ArtistIndexRepository
	playlistRepository := provider.PlaylistRepository
	propertyRepository := provider.PropertyRepository
	db := newDB()
	search := engine.NewSearch(artistRepository, albumRepository, mediaFileRepository, db)
	importer := scanner.NewImporter(musicFolder, itunesScanner, mediaFileRepository, albumRepository, artistRepository, artistIndexRepository, playlistRepository, propertyRepository, search)
	app := NewApp(importer)
	return app
}

func CreateSubsonicAPIRouter(p persistence.ProviderIdentifier) *api.Router {
	provider := createPersistenceProvider(p)
	propertyRepository := provider.PropertyRepository
	mediaFolderRepository := provider.MediaFolderRepository
	artistIndexRepository := provider.ArtistIndexRepository
	artistRepository := provider.ArtistRepository
	albumRepository := provider.AlbumRepository
	mediaFileRepository := provider.MediaFileRepository
	browser := engine.NewBrowser(propertyRepository, mediaFolderRepository, artistIndexRepository, artistRepository, albumRepository, mediaFileRepository)
	cover := engine.NewCover(mediaFileRepository, albumRepository)
	nowPlayingRepository := provider.NowPlayingRepository
	listGenerator := engine.NewListGenerator(albumRepository, mediaFileRepository, nowPlayingRepository)
	itunesControl := itunesbridge.NewItunesControl()
	playlistRepository := provider.PlaylistRepository
	playlists := engine.NewPlaylists(itunesControl, playlistRepository, mediaFileRepository)
	ratings := engine.NewRatings(itunesControl, mediaFileRepository, albumRepository, artistRepository)
	scrobbler := engine.NewScrobbler(itunesControl, mediaFileRepository, nowPlayingRepository)
	db := newDB()
	search := engine.NewSearch(artistRepository, albumRepository, mediaFileRepository, db)
	router := api.NewRouter(browser, cover, listGenerator, playlists, ratings, scrobbler, search)
	return router
}

func createSQLProvider() *Provider {
	albumRepository := db_sql.NewAlbumRepository()
	artistRepository := db_sql.NewArtistRepository()
	checkSumRepository := db_sql.NewCheckSumRepository()
	artistIndexRepository := db_sql.NewArtistIndexRepository()
	mediaFileRepository := db_sql.NewMediaFileRepository()
	mediaFolderRepository := persistence.NewMediaFolderRepository()
	nowPlayingRepository := persistence.NewNowPlayingRepository()
	playlistRepository := db_ledis.NewPlaylistRepository()
	propertyRepository := db_sql.NewPropertyRepository()
	provider := &Provider{
		AlbumRepository:       albumRepository,
		ArtistRepository:      artistRepository,
		CheckSumRepository:    checkSumRepository,
		ArtistIndexRepository: artistIndexRepository,
		MediaFileRepository:   mediaFileRepository,
		MediaFolderRepository: mediaFolderRepository,
		NowPlayingRepository:  nowPlayingRepository,
		PlaylistRepository:    playlistRepository,
		PropertyRepository:    propertyRepository,
	}
	return provider
}

func createLedisDBProvider() *Provider {
	albumRepository := db_ledis.NewAlbumRepository()
	artistRepository := db_ledis.NewArtistRepository()
	checkSumRepository := db_ledis.NewCheckSumRepository()
	artistIndexRepository := db_ledis.NewArtistIndexRepository()
	mediaFileRepository := db_ledis.NewMediaFileRepository()
	mediaFolderRepository := persistence.NewMediaFolderRepository()
	nowPlayingRepository := db_ledis.NewNowPlayingRepository()
	playlistRepository := db_ledis.NewPlaylistRepository()
	propertyRepository := db_ledis.NewPropertyRepository()
	provider := &Provider{
		AlbumRepository:       albumRepository,
		ArtistRepository:      artistRepository,
		CheckSumRepository:    checkSumRepository,
		ArtistIndexRepository: artistIndexRepository,
		MediaFileRepository:   mediaFileRepository,
		MediaFolderRepository: mediaFolderRepository,
		NowPlayingRepository:  nowPlayingRepository,
		PlaylistRepository:    playlistRepository,
		PropertyRepository:    propertyRepository,
	}
	return provider
}

func createStormProvider() *Provider {
	albumRepository := db_storm.NewAlbumRepository()
	artistRepository := db_storm.NewArtistRepository()
	checkSumRepository := db_storm.NewCheckSumRepository()
	artistIndexRepository := db_storm.NewArtistIndexRepository()
	mediaFileRepository := db_storm.NewMediaFileRepository()
	mediaFolderRepository := persistence.NewMediaFolderRepository()
	nowPlayingRepository := persistence.NewNowPlayingRepository()
	playlistRepository := db_storm.NewPlaylistRepository()
	propertyRepository := db_storm.NewPropertyRepository()
	provider := &Provider{
		AlbumRepository:       albumRepository,
		ArtistRepository:      artistRepository,
		CheckSumRepository:    checkSumRepository,
		ArtistIndexRepository: artistIndexRepository,
		MediaFileRepository:   mediaFileRepository,
		MediaFolderRepository: mediaFolderRepository,
		NowPlayingRepository:  nowPlayingRepository,
		PlaylistRepository:    playlistRepository,
		PropertyRepository:    propertyRepository,
	}
	return provider
}

// wire_injectors.go:

type Provider struct {
	AlbumRepository       domain.AlbumRepository
	ArtistRepository      domain.ArtistRepository
	CheckSumRepository    scanner.CheckSumRepository
	ArtistIndexRepository domain.ArtistIndexRepository
	MediaFileRepository   domain.MediaFileRepository
	MediaFolderRepository domain.MediaFolderRepository
	NowPlayingRepository  domain.NowPlayingRepository
	PlaylistRepository    domain.PlaylistRepository
	PropertyRepository    domain.PropertyRepository
}

var allProviders = wire.NewSet(itunesbridge.NewItunesControl, engine.Set, scanner.Set, newDB, api.NewRouter, wire.FieldsOf(new(*Provider), "AlbumRepository", "ArtistRepository", "CheckSumRepository",
	"ArtistIndexRepository", "MediaFileRepository", "MediaFolderRepository", "NowPlayingRepository",
	"PlaylistRepository", "PropertyRepository"), createPersistenceProvider,
)

func createPersistenceProvider(provider persistence.ProviderIdentifier) *Provider {
	switch provider {
	case "sql":
		return createSQLProvider()
	case "storm":
		return createStormProvider()
	default:
		return createLedisDBProvider()
	}
}

func newDB() gomate.DB {
	return ledis.NewEmbeddedDB(db_ledis.Db())
}
