package stats

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/GeorgiyGusev/gtrk-back/gen/proto/stats/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repo interface {
	GetStatisticsForNews(ctx context.Context, newsID, source string, startDate, endDate time.Time, aggregationPeriod stats_gen_v1.AggregationPeriod) ([]*stats_gen_v1.ViewData, error)
	GetStatisticsForAllNews(ctx context.Context, source string, startDate, endDate time.Time, aggregationPeriod stats_gen_v1.AggregationPeriod) ([]*stats_gen_v1.ViewData, error)
}

type Xui struct {
	ID    string `validate:"required,uuid4"`
	Email string `validate:"required,email"`
}

type Handler struct {
	repo Repo
	stats_gen_v1.UnimplementedNewsStatisticsServiceServer
}

func RegisterHandlers(srv *grpc.Server, handlers *Handler) {
	stats_gen_v1.RegisterNewsStatisticsServiceServer(srv, handlers)
}

func NewHandler(repo Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetViewsStatisticsForNews(ctx context.Context, request *stats_gen_v1.GetViewsStatisticsForNewsRequest) (*stats_gen_v1.GetViewsStatisticsForNewsResponse, error) {
	if request.NewsId == "" {
		return nil, status.Error(codes.InvalidArgument, "news_id is required")
	}

	// Получаем даты начала и конца для временного диапазона
	startDate, endDate, err := calculateTimeRange(request.TimeRange)
	if err != nil {
		return nil, err
	}

	source, err := recognizeSource(request.Source)
	if err != nil {
		return nil, err
	}

	views, err := h.repo.GetStatisticsForNews(ctx, request.NewsId, source, startDate, endDate, request.Period)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics for news: %w", err)
	}

	slog.Info("lala", "newsId", request.NewsId)

	response := &stats_gen_v1.GetViewsStatisticsForNewsResponse{
		NewsId: request.NewsId,
		Views:  views,
	}

	return response, nil
}

func (h *Handler) GetViewsStatisticsForAllNews(ctx context.Context, request *stats_gen_v1.GetViewsStatisticsForAllNewsRequest) (*stats_gen_v1.GetViewsStatisticsForAllNewsResponse, error) {
	startDate, endDate, err := calculateTimeRange(request.TimeRange)
	if err != nil {
		return nil, err
	}

	source, err := recognizeSource(request.Source)
	if err != nil {
		return nil, err
	}

	views, err := h.repo.GetStatisticsForAllNews(ctx, source, startDate, endDate, request.Period)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics for all news: %w", err)
	}

	response := &stats_gen_v1.GetViewsStatisticsForAllNewsResponse{
		Views: views,
	}

	return response, nil
}

func calculateTimeRange(timeRange stats_gen_v1.TimeRange) (startDate time.Time, endDate time.Time, err error) {
	now := time.Now()

	switch timeRange {
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_24_HOURS:
		startDate = now.Add(-24 * time.Hour)
		endDate = now
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_7_DAYS:
		startDate = now.Add(-7 * 24 * time.Hour)
		endDate = now
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_30_DAYS:
		startDate = now.Add(-30 * 24 * time.Hour)
		endDate = now
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_90_DAYS:
		startDate = now.Add(-90 * 24 * time.Hour)
		endDate = now
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_180_DAYS:
		startDate = now.Add(-180 * 24 * time.Hour)
		endDate = now
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_1_YEAR:
		startDate = now.Add(-365 * 24 * time.Hour)
		endDate = now
	case stats_gen_v1.TimeRange_TIME_RANGE_LAST_1_HOUR:
		startDate = now.Add(-1 * time.Hour)
		endDate = now
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("unsupported time range: %v", timeRange)
	}

	return startDate, endDate, nil
}

func recognizeSource(source stats_gen_v1.Source) (string, error) {
	switch source {
	case stats_gen_v1.Source_SOURCE_TELEGRAM:
		return "telegram", nil
	case stats_gen_v1.Source_SOURCE_VK:
		return "vk", nil
	case stats_gen_v1.Source_SOURCE_SITE:
		return "site", nil
	case stats_gen_v1.Source_SOURCE_UNSPECIFIED:
		return "", errors.New("unspecified source")
	default:
		return "", errors.New("unspecified source")
	}
}
