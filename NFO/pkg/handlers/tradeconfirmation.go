package handlers

import (
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"DC/FnO/pkg/utils/conversions"
	"bytes"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func HandleTradeConfirmation(data []byte) (*models.ExtractTrade, error) {
	// Create an instance of ExternalTradeConfirmation to unmarshal the data
	var externalResponse models.ExternalTradeConfirmation
	// lastsend :=
	if err := conversions.FromLittleEndian(data, &externalResponse, "trade-confirmation"); err != nil {
		logger.Log().Error("Failed to parse the packet", zap.Error(err))
		return nil, err
	}

	// logger.Log().Info("trd:", zap.Any("----------------------",externalResponse))
	logger.Log().Info("processing TradeConfirmation response")

	dst := externalResponse.MsgData

	activityType := string(bytes.Trim(dst.ActivityType[:], "\x00"))
	
	trade := &models.ExtractTrade{
		ResponseOrderNumber:   dst.ResponseOrderNumber,
		BrokerId:              conversions.BytesToString(dst.BrokerId[:]),
		TradeNum:              dst.TradeNum,
		AccountNum:            conversions.BytesToString(dst.AccountNum[:]),
		BuySell:               conversions.BytesToString(dst.BuySell[:]),
		OriginalVol:           dst.OriginalVol,
		DisclosedVol:          dst.DisclosedVol,
		RemainingVol:          dst.RemainingVol,
		DisclosedVolRemaining: dst.DisclosedVolRemaining,
		Price:                 dst.Price,
		StOrderFlags:          dst.StOrderFlags,
		Gtd:                   dst.Gtd,
		FillNumber:            dst.FillNumber,
		FillQty:               dst.FillQty,
		FillPrice:             int64(dst.FillPrice),
		VolFilledToday:        dst.VolFilledToday,
		ActivityType:          activityType,
		ActivityTime:          dst.ActivityTime,
		OpOrderNum:            dst.OpOrderNum,
		OpBrokerId:            conversions.BytesToString(dst.OpBrokerId[:]),
		InstrumentName:        conversions.BytesToString(dst.ContrectDesc.InstrumentName[:]),
		Symbol:                conversions.BytesToString(dst.ContrectDesc.Symbol[:]),
		ExpiryDate:            time.Unix(int64(dst.ContrectDesc.ExpiryDate), 0),
		StrikePrice:           dst.ContrectDesc.StrikePrice,
		OptionType:            conversions.BytesToString(dst.ContrectDesc.OptionType[:]),
		CALevel:               conversions.BytesToString(dst.ContrectDesc.CALevel[:]),
		OpenClose:             string(dst.OpenClose),
		OldOpenClose:          string(dst.OldOpenClose),
		AdditionalOrderFlags:  strconv.Itoa(int(dst.AdditionalOrderFlags)),
		GiveUpTrade:           "",
		PAN:                   conversions.BytesToString(dst.PAN[:]),
		OldPan:                conversions.BytesToString(dst.OldPan[:]),
		AlgoId:                dst.AlgoId,
		AlgoCategory:          conversions.BytesToString(dst.AlgoCategory[:]),
		LastActivityReference: strconv.Itoa(int(dst.LastActivityReference)),
		NnfField:              strconv.Itoa(int(dst.NnfField)),
	}

	logger.Log().Info("Extracted Trade details", zap.Any("ExtractTrade", trade))
	logger.Log().Info("trade details extracted sucessfully")
	return trade, nil
}
