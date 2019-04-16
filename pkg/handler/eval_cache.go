package handler

import (
	"sync"
	"time"

	"github.com/checkr/flagr/pkg/config"
	"github.com/checkr/flagr/pkg/entity"
	"github.com/checkr/flagr/pkg/util"

	"github.com/sirupsen/logrus"
	"github.com/zhouzhuojie/withtimeout"
)

var (
	singletonEvalCache     *EvalCache
	singletonEvalCacheOnce sync.Once
)

type mapCache map[string]*entity.Flag

// EvalCache is the in-memory cache just for evaluation
type EvalCache struct {
	mapCacheLock sync.RWMutex
	idCache      mapCache
	keyCache     mapCache

	refreshTimeout  time.Duration
	refreshInterval time.Duration
}

// GetEvalCache gets the EvalCache
var GetEvalCache = func() *EvalCache {
	singletonEvalCacheOnce.Do(func() {
		ec := &EvalCache{
			idCache:         make(map[string]*entity.Flag),
			keyCache:        make(map[string]*entity.Flag),
			refreshTimeout:  config.Config.EvalCacheRefreshTimeout,
			refreshInterval: config.Config.EvalCacheRefreshInterval,
		}
		singletonEvalCache = ec
	})
	return singletonEvalCache
}

// Start starts the polling of EvalCache
func (ec *EvalCache) Start() {
	err := ec.reloadMapCache()
	if err != nil {
		panic(err)
	}
	go func() {
		for range time.Tick(ec.refreshInterval) {
			err := ec.reloadMapCache()
			if err != nil {
				logrus.WithField("err", err).Error("reload evaluation cache error")
			}
		}
	}()
}

// GetByFlagKeyOrID gets the flag by Key or ID
func (ec *EvalCache) GetByFlagKeyOrID(keyOrID interface{}) *entity.Flag {
	ec.mapCacheLock.RLock()
	defer ec.mapCacheLock.RUnlock()

	s := util.SafeString(keyOrID)
	f, ok := ec.idCache[s]
	if !ok {
		f = ec.keyCache[s]
	}
	return f
}

func (ec *EvalCache) reloadMapCache() error {
	if config.Config.NewRelicEnabled {
		defer config.Global.NewrelicApp.StartTransaction("eval_cache_reload", nil, nil).End()
	}

	_, _, err := withtimeout.Do(ec.refreshTimeout, func() (interface{}, error) {
		idCache, keyCache, err := ec.fetchAllFlags()
		if err != nil {
			return nil, err
		}

		ec.mapCacheLock.Lock()
		defer ec.mapCacheLock.Unlock()

		ec.idCache = idCache
		ec.keyCache = keyCache
		return nil, err
	})

	return err
}
