package consensus

import (
	"fmt"

	"github.com/forbole/bdjuno/v4/modules/actions/logging"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/v4/modules/utils"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "consensus").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Minute().Do(func() {
		utils.WatchMethod(m.updateBlockTimeInMinute)
	}); err != nil {
		return fmt.Errorf("error while setting up consensus periodic operation: %s", err)
	}

	if _, err := scheduler.Every(1).Hour().Do(func() {
		utils.WatchMethod(m.updateBlockTimeInHour)
	}); err != nil {
		return fmt.Errorf("error while setting up consensus periodic operation: %s", err)
	}

	if _, err := scheduler.Every(1).Day().Do(func() {
		utils.WatchMethod(m.updateBlockTimeInDay)
	}); err != nil {
		return fmt.Errorf("error while setting up consensus periodic operation: %s", err)
	}

	return nil
}

// updateBlockTimeInMinute insert average block time in the latest minute
func (m *Module) updateBlockTimeInMinute() error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in minutes")

	block, err := m.db.GetLastBlock()
	if err != nil {
		return fmt.Errorf("error while getting last block: %s", err)
	}

	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error while getting genesis: %s", err)
	}

	// Skip if the genesis does not exist
	if genesis == nil {
		return nil
	}

	// Check if the chain has been created at least a minute ago
	if block.Timestamp.Sub(genesis.Time).Minutes() < 0 {
		return nil
	}

	minute, err := m.db.GetBlockHeightTimeMinuteAgo(block.Timestamp)
	if err != nil {
		return fmt.Errorf("error while gettting block height a minute ago: %s", err)
	}
	newBlockTime := block.Timestamp.Sub(minute.Timestamp).Seconds() / float64(block.Height-minute.Height)

	updated, err := m.db.SaveAverageBlockTimePerMin(newBlockTime, block.Height)
	if err != nil {
		return err
	}
	if updated {
		logging.BlockTimeGauge.WithLabelValues("minute").Set(newBlockTime)
	}

	return nil
}

// updateBlockTimeInHour insert average block time in the latest hour
func (m *Module) updateBlockTimeInHour() error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in hours")

	block, err := m.db.GetLastBlock()
	if err != nil {
		return fmt.Errorf("error while getting last block: %s", err)
	}

	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error while getting genesis: %s", err)
	}

	// Skip if the genesis does not exist
	if genesis == nil {
		return nil
	}

	// Check if the chain has been created at least an hour ago
	if block.Timestamp.Sub(genesis.Time).Hours() < 0 {
		return nil
	}

	hour, err := m.db.GetBlockHeightTimeHourAgo(block.Timestamp)
	if err != nil {
		return fmt.Errorf("error while getting block height an hour ago: %s", err)
	}
	newBlockTime := block.Timestamp.Sub(hour.Timestamp).Seconds() / float64(block.Height-hour.Height)

	updated, err := m.db.SaveAverageBlockTimePerHour(newBlockTime, block.Height)
	if err != nil {
		return err
	}
	if updated {
		logging.BlockTimeGauge.WithLabelValues("hour").Set(newBlockTime)
	}

	return nil
}

// updateBlockTimeInDay insert average block time in the latest minute
func (m *Module) updateBlockTimeInDay() error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in days")

	block, err := m.db.GetLastBlock()
	if err != nil {
		return fmt.Errorf("error while getting last block: %s", err)
	}

	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error while getting genesis: %s", err)
	}

	// Skip if the genesis does not exist
	if genesis == nil {
		return nil
	}

	// Check if the chain has been created at least a days ago
	if block.Timestamp.Sub(genesis.Time).Hours() < 24 {
		return nil
	}

	day, err := m.db.GetBlockHeightTimeDayAgo(block.Timestamp)
	if err != nil {
		return fmt.Errorf("error while getting block time a day ago: %s", err)
	}
	newBlockTime := block.Timestamp.Sub(day.Timestamp).Seconds() / float64(block.Height-day.Height)

	updated, err := m.db.SaveAverageBlockTimePerDay(newBlockTime, block.Height)
	if err != nil {
		return err
	}
	if updated {
		logging.BlockTimeGauge.WithLabelValues("day").Set(newBlockTime)
	}

	return nil

}
