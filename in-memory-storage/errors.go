package persistence

import "errors"

var ErrTransactionNotFound = errors.New("transaction not found")
var ErrTransactionAlreadyExist = errors.New("transaction already exist")
var ErrSubscriptionNotFound = errors.New("subscription not found")
