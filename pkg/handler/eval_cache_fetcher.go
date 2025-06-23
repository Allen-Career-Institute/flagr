package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/Allen-Career-Institute/flagr/pkg/config"
	"github.com/Allen-Career-Institute/flagr/pkg/entity"
	"github.com/Allen-Career-Institute/flagr/pkg/util"
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
)

// EvalCacheJSON is the JSON serialization format of EvalCache's flags
type EvalCacheJSON struct {
	Flags []entity.Flag
}

func (ec *EvalCache) export() EvalCacheJSON {
	idCache := ec.cache.Load().(*cacheContainer).idCache
	fs := make([]entity.Flag, 0, len(idCache))
	for _, f := range idCache {
		ff := *f
		fs = append(fs, ff)
	}
	return EvalCacheJSON{Flags: fs}
}

func (ec *EvalCache) fetchAllFlags() (idCache map[string]*entity.Flag, keyCache map[string]*entity.Flag, tagCache map[string]map[uint]*entity.Flag, err error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"operation": "fetch_all_flags_processing",
			"duration_us": duration.Microseconds(),
		}).Info("fetch all flags processing completed")
	}()

	fs, err := fetchAllFlags()
	if err != nil {
		return nil, nil, nil, err
	}

	idCache = make(map[string]*entity.Flag)
	keyCache = make(map[string]*entity.Flag)
	tagCache = make(map[string]map[uint]*entity.Flag)

	for i := range fs {
		f := &fs[i]
		if err := f.PrepareEvaluation(); err != nil {
			return nil, nil, nil, err
		}

		if f.ID != 0 {
			idCache[util.SafeString(f.ID)] = f
		}
		if f.Key != "" {
			keyCache[f.Key] = f
		}
		if f.Tags != nil {
			for _, s := range f.Tags {
				if tagCache[s.Value] == nil {
					tagCache[s.Value] = make(map[uint]*entity.Flag)
				}
				tagCache[s.Value][f.ID] = f
			}
		}
	}
	return idCache, keyCache, tagCache, nil
}

type evalCacheFetcher interface {
	fetch() ([]entity.Flag, error)
}

func newFetcher() (evalCacheFetcher, error) {
	if !config.Config.EvalOnlyMode {
		return &dbFetcher{db: getDB()}, nil
	}

	switch config.Config.DBDriver {
	case "json_file":
		return &jsonFileFetcher{filePath: config.Config.DBConnectionStr}, nil
	case "json_http":
		return &jsonHTTPFetcher{url: config.Config.DBConnectionStr}, nil
	default:
		return nil, fmt.Errorf(
			"failed to create evaluation cache fetcher. DBDriver:%s is not supported",
			config.Config.DBDriver,
		)
	}
}

var fetchAllFlags = func() ([]entity.Flag, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"operation": "fetch_all_flags",
			"duration_us": duration.Microseconds(),
		}).Info("fetch all flags operation completed")
	}()

	fetcher, err := newFetcher()
	if err != nil {
		return nil, err
	}
	return fetcher.fetch()
}

type jsonFileFetcher struct {
	filePath string
}

func (ff *jsonFileFetcher) fetch() ([]entity.Flag, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"operation": "json_file_fetch",
			"duration_us": duration.Microseconds(),
			"file_path": ff.filePath,
		}).Info("JSON file fetch completed")
	}()

	var json = jsoniter.ConfigFastest

	b, err := os.ReadFile(ff.filePath)
	if err != nil {
		return nil, err
	}
	ecj := &EvalCacheJSON{}
	err = json.Unmarshal(b, ecj)
	if err != nil {
		return nil, err
	}
	return ecj.Flags, nil
}

type jsonHTTPFetcher struct {
	url string
}

func (hf *jsonHTTPFetcher) fetch() ([]entity.Flag, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"operation": "json_http_fetch",
			"duration_us": duration.Microseconds(),
			"url": hf.url,
		}).Info("JSON HTTP fetch completed")
	}()

	var json = jsoniter.ConfigFastest

	client := http.Client{Timeout: config.Config.EvalCacheRefreshTimeout}
	res, err := client.Get(hf.url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ecj := &EvalCacheJSON{}
	err = json.Unmarshal(b, ecj)
	if err != nil {
		return nil, err
	}
	return ecj.Flags, nil
}

type dbFetcher struct {
	db *gorm.DB
}

func (df *dbFetcher) fetch() ([]entity.Flag, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"operation": "database_fetch",
			"duration_us": duration.Microseconds(),
		}).Info("database fetch completed")
	}()

	// Use eager loading to avoid N+1 problem
	// doc: http://jinzhu.me/gorm/crud.html#preloading-eager-loading
	fs := []entity.Flag{}
	err := entity.PreloadSegmentsVariantsTags(df.db).Find(&fs).Error
	return fs, err
}
