package messages

import (
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ProcessTrade(db *gorm.DB, tradeChan *chan models.ExtractTrade, errorChan *chan error, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	for trade := range *tradeChan {
		logger.Log().Info("Inserting record into DB")
		err := insertToDatabase(db, trade) // Pass the DB instance
		if err != nil {
			*errorChan <- fmt.Errorf("failed to process trade %+v: %w", trade, err)
		}
	}
}


func insertToDatabase(db *gorm.DB, trade models.ExtractTrade) error {
    // Define the SQL query for insertion
    query := `
    INSERT INTO "extract_trades" (
        "response_order_number", "broker_id", "trade_num", "account_num", "buy_sell", "original_vol", 
        "disclosed_vol", "remaining_vol", "disclosed_vol_remaining", "price", "st_order_flags", "gtd", 
        "fill_number", "fill_qty", "fill_price", "vol_filled_today", "activity_type", "activity_time", 
        "op_order_num", "op_broker_id", "token", "instrument_name", "symbol", "expiry_date", "strike_price", 
        "option_type", "ca_level", "open_close", "old_open_close", "additional_order_flags", "give_up_trade", 
        "pan", "old_pan", "algo_id", "algo_category", "last_activity_reference", "nnf_field"
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Execute the raw SQL query with values from the `trade` struct
	if err := db.Exec(query, 
	    trade.ResponseOrderNumber, trade.BrokerId, trade.TradeNum, trade.AccountNum, trade.BuySell, 
	    trade.OriginalVol, trade.DisclosedVol, trade.RemainingVol, trade.DisclosedVolRemaining, trade.Price, 
	    trade.StOrderFlags, trade.Gtd, trade.FillNumber, trade.FillQty, trade.FillPrice, trade.VolFilledToday, 
	    trade.ActivityType, trade.ActivityTime, trade.OpOrderNum, trade.OpBrokerId, trade.Token, trade.InstrumentName, 
	    trade.Symbol, trade.ExpiryDate, trade.StrikePrice, trade.OptionType, trade.CALevel, trade.OpenClose, 
	    trade.OldOpenClose, trade.AdditionalOrderFlags, trade.GiveUpTrade, trade.PAN, trade.OldPan, 
	    trade.AlgoId, trade.AlgoCategory, trade.LastActivityReference, trade.NnfField,
	).Error; err != nil {
	    logger.Log().Error("Failed to insert trade into database using raw query", zap.Error(err))
	    return err
	}


    logger.Log().Info("Trade successfully inserted into database using raw query")
    return nil
}


// Logs errors from the error channel
func ProcessError(errorChan *chan error) {
	for err := range *errorChan {
		logger.Log().Error("Error while processing trades", zap.Error(err))
	}
}
