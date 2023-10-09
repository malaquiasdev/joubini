package main

import (
	"fmt"
	"joubini/pkg/aws"
	orderItems "joubini/pkg/sellingpartnerapi/orderitems"
	"joubini/pkg/sellingpartnerapi/orders"
	"joubini/pkg/utils"
)

func main() {
	clientId := utils.GetEnv("LWA_ID", "")
	clientSecret := utils.GetEnv("LWA_SECRET", "")
	refreshToken := utils.GetEnv("LWA_REFRESH_TOKEN", "")

	token := aws.GetToken(clientId, clientSecret, refreshToken)

	fmt.Println(token.AccessToken)

	yesterday := utils.DateNowSubtractFormated("2006-01-02", 1)
	fmt.Println("today", yesterday)

	order, _ := orders.GetOrders(token.AccessToken, yesterday)
	amazonOrderId := order[len(order)-1].AmazonOrderId
	orderItems, _ := orderItems.GetOrderItems(token.AccessToken, amazonOrderId)
	fmt.Println("order", order)
	fmt.Println("orderItems", orderItems)
}
