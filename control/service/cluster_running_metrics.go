package service

import (
	"laser-control/common"
	"laser-control/context"
	"laser-control/params"
	"time"
)

type ClusterRunningMetrics struct {
	Id                       int32   `gorm:"primary_key"`
	Qps                      int32   `gorm:"type:int(32);null"`
	Kps                      int32   `gorm:"type:int(32);null"`
	KpsWrite                 int32   `gorm:"type:int(32);null"`
	KpsRead                  int32   `gorm:"type:int(32);null"`
	Bps                      int64   `gorm:"type:bigint(20);null"`
	BpsWrite                 int64   `gorm:"type:bigint(20);null"`
	BpsRead                  int64   `gorm:"type:bigint(20);null"`
	DataSize                 int64   `gorm:"type:bigint(20);null"`
	P99                      float32 `gorm:"type:float;null"`
	P999                     float32 `gorm:"type:float;null"`
	JitterDuration           int32   `gorm:"type:int(32);null"`
	CapacityIdleRate         float32 `gorm:"type:float;null"`
	CapacityIdleRateAbsolute float32 `gorm:"type:float;null"`
	CollectTime              time.Time
	TimeType                 int32 `gorm:"type:int(32);null"` // 1-平峰，2-高峰
}

type ClusterRunningMetricsModel struct {
	ctx *context.Context
}

func NewClusterRunningMetricsModel(ctx *context.Context) *ClusterRunningMetricsModel {
	return &ClusterRunningMetricsModel{
		ctx: ctx,
	}
}

func (cluster *ClusterRunningMetricsModel) Insert(metricsInfo params.ClusterRunningMetricsInfo) *common.Status {
	addInfo := &ClusterRunningMetrics{
		Qps:                      metricsInfo.Qps,
		Kps:                      metricsInfo.Kps,
		KpsWrite:                 metricsInfo.KpsWrite,
		KpsRead:                  metricsInfo.KpsRead,
		Bps:                      metricsInfo.Bps,
		BpsWrite:                 metricsInfo.BpsWrite,
		BpsRead:                  metricsInfo.BpsRead,
		DataSize:                 metricsInfo.DataSize,
		P99:                      metricsInfo.P99,
		P999:                     metricsInfo.P999,
		JitterDuration:           metricsInfo.JitterDuration,
		CapacityIdleRate:         metricsInfo.CapacityIdleRate,
		CapacityIdleRateAbsolute: metricsInfo.CapacityIdleRateAbsolute,
		TimeType:                 metricsInfo.TimeType,
	}
	addInfo.CollectTime = time.Unix(metricsInfo.CollectTime, 0)

	if err := cluster.ctx.Db().Create(addInfo).Error; err != nil {
		return common.StatusWithError(err)
	}
	return common.StatusOk()
}

func (cluster *ClusterRunningMetricsModel) List(listParams params.ClusterRunningMetricsListParams) ([]ClusterRunningMetrics, uint32) {
	db := cluster.ctx.Db().Model(&ClusterRunningMetrics{}).Where(&ClusterRunningMetrics{TimeType: listParams.TimeType})

	if listParams.StartTime > 0 {
		startTimeUnix := time.Unix(listParams.StartTime, 0)
		db = db.Where("collect_time >= ?", startTimeUnix)
	}

	if listParams.EndTime > 0 {
		endTimeUnix := time.Unix(listParams.EndTime, 0)
		db = db.Where("collect_time <= ?", endTimeUnix)
	}

	var total uint32
	db.Count(&total)
	if listParams.Page > 0 && listParams.Limit > 0 {
		offset := (listParams.Page - 1) * listParams.Limit
		db = db.Offset(offset).Limit(listParams.Limit)
	}

	var clusterRunningMetrics []ClusterRunningMetrics
	db.Order("collect_time desc").Find(&clusterRunningMetrics)

	return clusterRunningMetrics, total
}
