package handlers

import (
	constants "DC/FnO/constant"
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"DC/FnO/pkg/utils/conversions"
	"errors"

	"go.uber.org/zap"
)

// HandleLogonOut processes the LogonOut data and returns the number of streams.
func HandleLogonOut(data []byte, dts interface{}) (int, error) {
	
	if err := conversions.FromLittleEndian(data, dts, "signon-out"); err != nil {
		logger.Log().Error("Failed to parse the packet", zap.Error(err))
		return 0, err
	}

	// Ensure dts is of the correct type
	header, ok := dts.(*models.ExternalLogonOutHeader)
	if !ok {
		logger.Log().Error("Unexpected type for LogonOut response", zap.Any("response", dts))
		return 0, errors.New("unexpected response type")
	}

	// Check the TransactionCode
	if header.MsgData.MessageHeader.TransactionCode != constants.SIGNON_OUT {
		logger.Log().Error("received transactioncode is not valid", zap.Any("header", header))
		return 0, errors.New("unknown packet")
	}

	// Extract stream count from AlphaChar
	streamCount := int(header.MsgData.MessageHeader.AlphaChar[0])
	logger.Log().Info("no. of trades to download, ", zap.Int("streamCount", streamCount))

	return streamCount, nil
}
