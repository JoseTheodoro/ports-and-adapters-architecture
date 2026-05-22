package approveorder

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrOrderAlredyApproved = errors.New("order already approved")
