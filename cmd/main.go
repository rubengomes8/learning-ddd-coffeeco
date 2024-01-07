package main

import (
	"context"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	coffeeco "github.com/rubengomes8/learning-ddd-coffeeco/internal"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/payment"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/purchase"
	"github.com/rubengomes8/learning-ddd-coffeeco/internal/store"
)

func main() {

	ctx := context.Background()

	// This is the test key from Stripe's documentation. Feel free to use it, no charges will actually be made.
	stripeTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"

	// infrastructure service - payment service
	stripeSvc, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	mongoConString := "mongodb://root:example@localhost:27017"
	// infrastructure service - purchase repo
	purchaseRepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := purchaseRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// infrastructure service - store repo
	storeRepo, err := store.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err := storeRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	storeSvc := store.NewService(storeRepo)
	purchaseSvc := purchase.NewService(stripeSvc, purchaseRepo, storeSvc)

	// build purchase to test
	storeID := uuid.New()
	// This is a test token from Stripe's documentation.
	cardToken := "tok_visa"
	p := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: storeID,
		},
		ProductsToPurchase: []coffeeco.Product{{
			ItemName:  "item1",
			BasePrice: *money.New(3300, "USD"),
		}},
		PaymentMeans: payment.MEANS_CARD,
	}
	if err := purchaseSvc.CompletePurchase(ctx, storeID, p, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("purchase was successful")
}
