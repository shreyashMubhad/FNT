package messages

import (
	"DC/FnO/constant"
	"DC/FnO/pkg/logger"
	"DC/FnO/pkg/models"
	"DC/FnO/pkg/utils"
	"context"
	"errors"
	"fmt"
	"net"

	"go.uber.org/zap"
)


func dcDownload(streamNum int) *models.DcDownloadRequest {
	logger.Log().Debug("Initializing DC Download Request")

	// Initialize the DC Download Request
	dcHeader := &models.DcDownloadRequest{
		MessageHeader: models.InternalHeader{
			TransactionCode: constants.DC_DOWNLOAD,
			LogTime:         0,
			AlphaChar:       [2]byte{byte(streamNum)}, 
			UserId:          32106,
			ErrorCode:       0,
			TimeStamp:       [8]byte{},
			TimeStamp1:      [8]byte{},
			TimeStamp2:      [8]byte{},
			MessageLength:   constants.DownloadRequestLen,
		},
		SeqNum: float64(utils.GetSequenceNumber()),
	}

	// logger.Log().Info("DC Download Request", zap.Any("dcHeader", dcHeader))

	return dcHeader
}

func getDcDownloadRequest(streamNum int, seqNum uint32) models.ExternalDcDownloadHeader {
	logger.Log().Debug("Creating DC download request")

	dcDownload := dcDownload(streamNum)
	checksum := utils.GetCheckSum(dcDownload)
	request := models.ExternalDcDownloadHeader{
		BodyLen:  constants.DownloadRequestLen,
		SeqNum:   seqNum,
		CheckSum: [16]byte(checksum),
		MsgData:  *dcDownload,
	}
	logger.Log().Debug("SeqNum is : ", zap.Uint32("------------------------------------------------------------------", seqNum))
	return request
}



// Sends trade requests and pushes responses to the trade channel
func TradeRequest(c context.Context, conn net.Conn, tradeChan *chan models.ExtractTrade, streamnum int, lastSeqSendToExcng *uint32) error {
	if conn == nil {
		logger.Log().Error("Connection object is nil")
		return errors.New("connection is nil")
	}
	
	// Initialize lastSeqSendToExcng to 0 if it's the first time
	if *lastSeqSendToExcng == 0 {
		logger.Log().Info("Initializing lastSeqSendToExcng to 0")
	}
	// logger.Log().Info("lastSeqSendToExcng : ", zap.Uint32("",lastSeqSendToExcng))
	// Create the request with the current value of lastSeqSendToExcng
	request := getDcDownloadRequest(streamnum, *lastSeqSendToExcng)

	logger.Log().Info("sending dc_download request")
	if err := utils.Send(conn, request, constants.DC_DOWNLOAD_REQUEST); err != nil {
		logger.Log().Error("Failed to send the request", zap.Error(err))
		return fmt.Errorf("failed to send request: %w", err)
	}

	// Receive the response
	response, _, _, seqNum, err := utils.Recv(conn)
	if err != nil {
		logger.Log().Error("Failed to read the header", zap.Error(err))
		return fmt.Errorf("failed to receive response: %w", err)
	}

	
	// Update lastSeqSendToExcng with the sequence number from the response
	*lastSeqSendToExcng = seqNum
	logger.Log().Info("Received response, updated lastSeqSendToExcng", zap.Uint32("seqNum", *lastSeqSendToExcng))

	// Extract trade from the response
	trade, ok := response.(models.ExtractTrade)
	if !ok {
		logger.Log().Error("Failed to extract trade from response", zap.Any("response", response))
		return fmt.Errorf("unexpected response type: %T", response)
	}

	// Send the trade to the tradeChan channel
	logger.Log().Info("passing trade to tradeChan channel.")
	*tradeChan <- trade

	return nil
}
