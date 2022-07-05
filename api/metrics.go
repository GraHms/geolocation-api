package api

import (
	"github.com/grahms/geolocationservice/util"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func setGauge() {
	gaugeCsvRows := &ginmetrics.Metric{
		Type:        ginmetrics.Counter,
		Name:        "csv_rows",
		Description: "Total amount of valid and invalid csv rows",
		Labels:      []string{"VALID_ROWS", "INVALID_ROWS", util.IPADDRKEY, util.CITY, util.COORDS, util.COUNTRYNAMEKEY, util.COUNTRYCODEKEY},
	}

	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(gaugeCsvRows)
}
