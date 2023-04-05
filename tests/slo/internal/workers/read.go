package workers

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"slo/internal/generator"
	"slo/internal/metrics"
)

func Read(ctx context.Context, st ReadWriter, rl *rate.Limiter, m *metrics.Metrics, logger *zap.Logger,
	en generator.Entries, idsPtr *[]generator.EntryID, mu *sync.RWMutex,
) {
	for {
		err := rl.Wait(ctx)
		if err != nil {
			return
		}

		ids := *idsPtr

		if len(ids) == 0 {
			continue
		}

		i := rand.Intn(len(ids))
		id := ids[i]

		metricID := m.StartJob(metrics.JobRead)

		e, err := st.Read(ctx, id)
		if err != nil {
			logger.Error(fmt.Errorf("get entry error: %w", err).Error())
			m.StopJob(metricID, false)
			continue
		}

		mu.RLock()
		original := en[id]
		mu.RUnlock()

		if reflect.DeepEqual(e, original) {
			m.StopJob(metricID, true)
			continue
		}
		logger.Info("payload does not match")
		m.StopJob(metricID, false)
	}
}
