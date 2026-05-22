package http

import "errors"

var ErrPriceIsRequired = errors.New("order price is required")
var ErrOrderIDInvalid = errors.New("order id must have a valid uuid format")
