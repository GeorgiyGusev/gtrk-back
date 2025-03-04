package stats

import (
	"context"
	"fmt"
	"github.com/GeorgiyGusev/gtrk-back/pkg/logging"
	"log/slog"
	"time"

	"github.com/GeorgiyGusev/gtrk-back/gen/proto/stats/v1"
	"github.com/jmoiron/sqlx"
)

type RepoImpl struct {
	db *sqlx.DB
}

func NewRepoImpl(db *sqlx.DB) *RepoImpl {
	return &RepoImpl{db: db}
}

// GetStatisticsForNews получает статистику по просмотрам для конкретной новости с учетом агрегации
func (r *RepoImpl) GetStatisticsForNews(ctx context.Context, newsID, source string, startDate, endDate time.Time, aggregationPeriod stats_gen_v1.AggregationPeriod) ([]*stats_gen_v1.ViewData, error) {
	aggregationInterval := getAggregationInterval(aggregationPeriod)

	query := `
		SELECT
			toStartOfInterval(created_at, INTERVAL 1 %s) AS time,
			SUM(views_count) AS views_count
		FROM views_data
		WHERE news_id = $1 AND created_at BETWEEN $2 AND $3 AND source = $4
		GROUP BY time
		ORDER BY time ASC
	`

	query = fmt.Sprintf(query, aggregationInterval)

	var views []*stats_gen_v1.ViewData
	err := r.db.SelectContext(ctx, &views, query, newsID, startDate, endDate, source)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics for news %s: %w", newsID, err)
	}

	return views, nil
}

// GetStatisticsForAllNews получает статистику по просмотрам для всех новостей с учетом агрегации
func (r *RepoImpl) GetStatisticsForAllNews(ctx context.Context, source string, startDate, endDate time.Time, aggregationPeriod stats_gen_v1.AggregationPeriod) ([]*stats_gen_v1.ViewData, error) {
	aggregationInterval := getAggregationInterval(aggregationPeriod)

	query := `
		SELECT
			toStartOfInterval(created_at, INTERVAL 1 %s) AS time,
			SUM(views_count) AS views_count
		FROM views_data
		WHERE created_at BETWEEN $1 AND $2 AND source = $3
		GROUP BY time
		ORDER BY time ASC
	`

	query = fmt.Sprintf(query, aggregationInterval)

	var views []*stats_gen_v1.ViewData
	err := r.db.SelectContext(ctx, &views, query, startDate, endDate, source)
	if err != nil {
		slog.Error("failed to get statistics for all news", logging.ErrorField(err))
		return nil, fmt.Errorf("failed to get statistics for all news: %w", err)
	}

	return views, nil
}

// getAggregationInterval преобразует AggregationPeriod в строковое представление для SQL-запроса
func getAggregationInterval(aggregationPeriod stats_gen_v1.AggregationPeriod) string {
	switch aggregationPeriod {
	case stats_gen_v1.AggregationPeriod_AGGREGATION_PERIOD_MINUTE:
		return "MINUTE"
	case stats_gen_v1.AggregationPeriod_AGGREGATION_PERIOD_HOUR:
		return "HOUR"
	case stats_gen_v1.AggregationPeriod_AGGREGATION_PERIOD_DAY:
		return "DAY"
	case stats_gen_v1.AggregationPeriod_AGGREGATION_PERIOD_WEEK:
		return "WEEK"
	case stats_gen_v1.AggregationPeriod_AGGREGATION_PERIOD_MONTH:
		return "MONTH"
	case stats_gen_v1.AggregationPeriod_AGGREGATION_PERIOD_UNSPECIFIED:
		return "HOUR"
	default:
		return "HOUR" // по умолчанию агрегация по часам
	}
}
