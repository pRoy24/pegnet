// Copyright (c) of parts are held by the various contributors (see the CLA)
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package polling

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pegnet/pegnet/common"

	"github.com/cenkalti/backoff"
	log "github.com/sirupsen/logrus"
	"github.com/zpatrick/go-config"
)

// OpenExchangeRatesDataSource is the datasource at "https://openexchangerates.org/"
type OpenExchangeRatesDataSource struct {
	config  *config.Config
	lastPeg PegAssets
}

func NewOpenExchangeRatesDataSource(config *config.Config) (*OpenExchangeRatesDataSource, error) {
	s := new(OpenExchangeRatesDataSource)
	s.config = config
	return s, nil
}

func (d *OpenExchangeRatesDataSource) Name() string {
	return "OpenExchangeRates"
}

func (d *OpenExchangeRatesDataSource) Url() string {
	return "https://openexchangerates.org/"
}

func (d *OpenExchangeRatesDataSource) SupportedPegs() []string {
	return common.MergeLists(common.CurrencyAssets, common.CommodityAssets, []string{"XBT"})
}

func (d *OpenExchangeRatesDataSource) FetchPegPrices() (peg PegAssets, err error) {
	resp, err := CallOpenExchangeRates(d.config)
	if err != nil {
		return nil, err
	}

	peg = make(map[string]PegItem)

	timestamp := time.Unix(resp.Timestamp, 0)
	for _, currencyISO := range d.SupportedPegs() {
		// Price is inverted
		if v, ok := resp.Rates[currencyISO]; ok {
			peg[currencyISO] = PegItem{Value: 1 / v, When: timestamp, WhenUnix: timestamp.Unix()}
		}

		// Special case for btc
		if currencyISO == "XBT" {
			if v, ok := resp.Rates["BTC"]; ok {
				peg[currencyISO] = PegItem{Value: 1 / v, When: timestamp, WhenUnix: timestamp.Unix()}
			}
		}
	}

	return
}

func (d *OpenExchangeRatesDataSource) FetchPegPrice(peg string) (i PegItem, err error) {
	return FetchPegPrice(peg, d.FetchPegPrices)
}

// --

type OpenExchangeRatesResponse struct {
	Disclaimer  string             `json:"disclaimer"`
	License     string             `json:"license"`
	Timestamp   int64              `json:"timestamp"`
	Base        string             `json:"base"`
	Error       bool               `json:"error"`
	Status      int64              `json:"status"`
	Message     string             `json:"message"`
	Description string             `json:"description"`
	Rates       map[string]float64 `json:"rates"`
}

func CallOpenExchangeRates(c *config.Config) (response OpenExchangeRatesResponse, err error) {
	var OpenExchangeRatesResponse OpenExchangeRatesResponse

	var apikey string
	{
		apikey, err = c.String("Oracle.OpenExchangeRatesKey")
		check(err)
	}

	operation := func() error {
		resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=" + apikey)
		if err != nil {
			log.WithError(err).Warning("Failed to get response from OpenExchangeRates")
			return err
		}

		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return err
		} else if err = json.Unmarshal(body, &OpenExchangeRatesResponse); err != nil {
			return err
		}
		return nil
	}

	err = backoff.Retry(operation, PollingExponentialBackOff())
	return OpenExchangeRatesResponse, err
}

func HandleOpenExchangeRates(response OpenExchangeRatesResponse, peg PegAssets) {

	// Handle Response Errors
	if response.Error {
		log.WithFields(log.Fields{
			"status":      response.Status,
			"message":     response.Message,
			"description": response.Description,
		}).Fatal("Failed to access OpenExchangeRates")
	}

	UpdatePegAssets(response.Rates, response.Timestamp, peg)
}

func OpenExchangeRatesInterface(config *config.Config, peg PegAssets) error {
	log.Debug("Pulling Asset data from OpenExchangesRates")
	OpenExchangeRatesResponse, err := CallOpenExchangeRates(config)
	if err != nil {
		return fmt.Errorf("failed to access OpenExchangesRates : %s", err.Error())
	}
	HandleOpenExchangeRates(OpenExchangeRatesResponse, peg)
	return nil
}
