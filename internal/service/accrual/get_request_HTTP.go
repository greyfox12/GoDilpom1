package accrualServ

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

// Отправить Запрос http

func (s *Service) ExecuteGetRequestHTTP(ctx context.Context, orderNum string, accrualURL string) (*domain.TAccrualReq, error) {
	var bk domain.TAccrualReq

	client := &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", accrualURL+"/api/orders/"+orderNum, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//	fmt.Printf("Head response: %v\n", response.Header)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request status: %v", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &bk)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body: %v", string(body))
	}

	return &bk, nil

}
