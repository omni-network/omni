// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

package promutil_test

import (
	"testing"

	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // This test uses global prometheus registry so concurrent tests are not safe.
func TestResetGaugeVec(t *testing.T) {
	const resetTest = "reset_test"

	var testResetGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Name: resetTest,
		Help: "",
	}, []string{"label0", "label1"})

	testResetGauge.WithLabelValues("1", "a").Set(0)
	assertVecLen(t, resetTest, 1)

	// Same labels, should not increase length
	testResetGauge.WithLabelValues("1", "a").Set(1)
	assertVecLen(t, resetTest, 1)

	testResetGauge.WithLabelValues("2", "b").Set(2)
	assertVecLen(t, resetTest, 2)

	testResetGauge.Reset()
	assertVecLen(t, resetTest, 0)

	testResetGauge.WithLabelValues("3", "c").Set(3)
	assertVecLen(t, resetTest, 1)

	testResetGauge.WithLabelValues("3", "d").Set(3)
	assertVecLen(t, resetTest, 2)

	testResetGauge.WithLabelValues("3", "e").Set(3)
	assertVecLen(t, resetTest, 3)

	testResetGauge.WithLabelValues("4", "z").Set(4)
	assertVecLen(t, resetTest, 4)

	testResetGauge.Reset("3", "c")
	assertVecLen(t, resetTest, 3)

	testResetGauge.Reset("3")
	assertVecLen(t, resetTest, 1)
}

func assertVecLen(t *testing.T, name string, l int) { //nolint:unparam // abstracting name is fine even though it is always currently constant
	t.Helper()

	metrics, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	for _, metricFam := range metrics {
		if metricFam.GetName() != name {
			continue
		}

		require.Len(t, metricFam.GetMetric(), l)

		return
	}

	if l == 0 {
		return
	}

	require.Fail(t, "metric not found")
}
