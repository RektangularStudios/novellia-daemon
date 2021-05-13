package prometheus_monitoring

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
	"github.com/RektangularStudios/novellia/internal/config"
)

// https://prometheus.io/docs/guides/go-application/

const (
	namespace = "novellia"
	status_interval = 30 * time.Second
)

var (
	microserviceStatusMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name: "microservice_status",
		Help: "Health status indicator for the Novellia microservice",
	})
	cardanoStatusMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name: "cardano_status",
		Help: "Health status indicator for Cardano services such as GraphQL and cardano-node",
	})
	productIDsListedMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name: "products_ids_listed",
		Help: "Number of products IDs returned when accessing Novellia",
	})
	databaseConnectionFailedMetric = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name: "database_connection_failed",
		Help: "Incremented whenever the database connection fails",
	})
)

type statusIndicators struct {
	microserviceStatus float64
	cardanoStatus float64
}

func getStatus() (*statusIndicators, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config from env")
	}

	req, err := http.NewRequest("GET", config.Monitoring.StatusURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status health check failed: %+v", resp)
	}

	var respBody nvla.Status
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return nil, err
	}

	s := statusIndicators{
		microserviceStatus: 0,
		cardanoStatus: 0,
	}
	if respBody.Status == "UP" {
		s.microserviceStatus = 1
	}
	if respBody.Cardano.Initialized {
		s.cardanoStatus = float64(respBody.Cardano.SyncPercentage)
	}
	fmt.Printf("Checked Status, result: %+v\n", respBody)

	return &s, nil
}

func RecordNumberOfProductIDsListed(count int) {
	productIDsListedMetric.Set(float64(count))
}

func RecordDatabaseConnectionFailure() {
	databaseConnectionFailedMetric.Inc()
}

func RecordMetrics() {
	go func() {
		for {
			indicators, err := getStatus()
			if err != nil {
				indicators = &statusIndicators{
					microserviceStatus: 0,
					cardanoStatus: 0,
				}
				fmt.Printf("Checked status, got error: %+v\n", err)
			}
			
			microserviceStatusMetric.Set(indicators.microserviceStatus)
			cardanoStatusMetric.Set(indicators.cardanoStatus)

			time.Sleep(status_interval)
		}
	}()
}
