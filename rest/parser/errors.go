package handlers

import "errors"

var ErrSubscriptionFailure = errors.New("subscription failure")
var ErrTransactionsNotFound = errors.New("no transactions for this address found")
