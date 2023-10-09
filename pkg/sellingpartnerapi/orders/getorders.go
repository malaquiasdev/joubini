package orders

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetOrders(accessToken string, createdAfter string) ([]Order, error) {
	urlBase := "https://sellingpartnerapi-na.amazon.com/orders/v0/orders?MarketplaceIds=A2Q3Y263D00KWC"
	url := urlBase + "&CreatedAfter=" + createdAfter

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error could not create request:", err)
		return nil, err
	}

	req.Header.Set("x-amz-access-token", accessToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error could not execute request:", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error could not read response body: %s\n", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		var errResponse ErrorResponse
		if err := json.Unmarshal([]byte(body), &errResponse); err != nil {
			fmt.Println("error unmarshal JSON:", err)
			return nil, err
		}
		fmt.Printf("error request response: %s\n", body)
		return nil, errors.New(errResponse.Errors[0].Message)
	}

	var data Root
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		fmt.Println("error unmarshal JSON:", err)
		return nil, err
	}

	return data.Payload.Orders, nil
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type BuyerInfo struct {
	BuyerEmail  string `json:"BuyerEmail"`
	BuyerCounty string `json:"BuyerCounty"`
}

type ShippingAddress struct {
	StateOrRegion string `json:"StateOrRegion"`
	PostalCode    string `json:"PostalCode"`
	City          string `json:"City"`
	CountryCode   string `json:"CountryCode"`
}

type OrderTotal struct {
	CurrencyCode string `json:"CurrencyCode"`
	Amount       string `json:"Amount"`
}

type Order struct {
	BuyerInfo                    BuyerInfo       `json:"BuyerInfo"`
	AmazonOrderId                string          `json:"AmazonOrderId"`
	EarliestShipDate             string          `json:"EarliestShipDate"`
	SalesChannel                 string          `json:"SalesChannel"`
	OrderStatus                  string          `json:"OrderStatus"`
	NumberOfItemsShipped         int             `json:"NumberOfItemsShipped"`
	OrderType                    string          `json:"OrderType"`
	IsPremiumOrder               bool            `json:"IsPremiumOrder"`
	IsPrime                      bool            `json:"IsPrime"`
	FulfillmentChannel           string          `json:"FulfillmentChannel"`
	NumberOfItemsUnshipped       int             `json:"NumberOfItemsUnshipped"`
	HasRegulatedItems            bool            `json:"HasRegulatedItems"`
	IsReplacementOrder           string          `json:"IsReplacementOrder"`
	IsSoldByAB                   bool            `json:"IsSoldByAB"`
	LatestShipDate               string          `json:"LatestShipDate"`
	ShipServiceLevel             string          `json:"ShipServiceLevel"`
	IsISPU                       bool            `json:"IsISPU"`
	MarketplaceId                string          `json:"MarketplaceId"`
	PurchaseDate                 string          `json:"PurchaseDate"`
	ShippingAddress              ShippingAddress `json:"ShippingAddress"`
	IsAccessPointOrder           bool            `json:"IsAccessPointOrder"`
	SellerOrderId                string          `json:"SellerOrderId"`
	PaymentMethod                string          `json:"PaymentMethod"`
	IsBusinessOrder              bool            `json:"IsBusinessOrder"`
	OrderTotal                   OrderTotal      `json:"OrderTotal"`
	PaymentMethodDetails         []string        `json:"PaymentMethodDetails"`
	IsGlobalExpressEnabled       bool            `json:"IsGlobalExpressEnabled"`
	LastUpdateDate               string          `json:"LastUpdateDate"`
	ShipmentServiceLevelCategory string          `json:"ShipmentServiceLevelCategory"`
}

type Payload struct {
	Orders        []Order `json:"Orders"`
	NextToken     string  `json:"NextToken"`
	CreatedBefore string  `json:"CreatedBefore"`
}

type Root struct {
	Payload Payload `json:"payload"`
}
