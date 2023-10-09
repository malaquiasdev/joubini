package orderItems

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetOrderItems(accessToken string, amazonOrderId string) ([]OrderItem, error) {
	urlBase := "https://sellingpartnerapi-na.amazon.com/orders/v0/orders/"
	url := urlBase + amazonOrderId + "/orderItems"

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

	return data.Payload.OrderItems, nil
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ProductInfo struct {
	NumberOfItems string `json:"NumberOfItems"`
}

type ItemTax struct {
	CurrencyCode string `json:"CurrencyCode"`
	Amount       string `json:"Amount"`
}

type ItemPrice struct {
	CurrencyCode string `json:"CurrencyCode"`
	Amount       string `json:"Amount"`
}

type OrderItem struct {
	ProductInfo          ProductInfo `json:"ProductInfo"`
	BuyerInfo            struct{}    `json:"BuyerInfo"`
	ItemTax              ItemTax     `json:"ItemTax"`
	QuantityShipped      int         `json:"QuantityShipped"`
	ItemPrice            ItemPrice   `json:"ItemPrice"`
	ASIN                 string      `json:"ASIN"`
	SellerSKU            string      `json:"SellerSKU"`
	Title                string      `json:"Title"`
	IsGift               string      `json:"IsGift"`
	IsTransparency       bool        `json:"IsTransparency"`
	QuantityOrdered      int         `json:"QuantityOrdered"`
	PromotionDiscountTax ItemTax     `json:"PromotionDiscountTax"`
	PromotionDiscount    ItemTax     `json:"PromotionDiscount"`
	OrderItemId          string      `json:"OrderItemId"`
}

type Payload struct {
	OrderItems    []OrderItem `json:"OrderItems"`
	AmazonOrderId string      `json:"AmazonOrderId"`
}

type Root struct {
	Payload Payload `json:"payload"`
}
