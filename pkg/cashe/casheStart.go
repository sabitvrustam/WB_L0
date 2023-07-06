package cashe

import (
	"github.com/dgraph-io/ristretto"
	"github.com/sirupsen/logrus"
)

func StartCashe(log *logrus.Logger) (cashe *ristretto.Cache, err error) {
	cashe, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal("не удалось запустить кэширование", err)
	}

	log.Info("кэширование запущено")

	return cashe, err

}
