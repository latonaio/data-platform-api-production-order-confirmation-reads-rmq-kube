package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-production-order-confirmation-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-confirmation-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header: header,
	}
	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"

	where = fmt.Sprintf("%s\nAND ProductionOrder = %d ", where, input.Header.ProductionOrder)

	where = fmt.Sprintf("%s\nAND ProductionOrderItem = %v", where, input.Header.ProductionOrderItem)

	where = fmt.Sprintf("%s\nAND Operations = %v", where, input.Header.Operations)

	where = fmt.Sprintf("%s\nAND OperationsItem = %v", where, input.Header.OperationsItem)

	where = fmt.Sprintf("%s\nAND OperationID = %v", where, input.Header.OperationID)

	where = fmt.Sprintf("%s\nAND ConfirmationCountingID = %v", where, input.Header.ConfirmationCountingID)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_production_order_confirmation_header_data
		` + where + ` ;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
