package common

/**
 * @export
 * @enum {string}
 */

type TransactionType string

const (
	FAILED                  TransactionType = "FAILED"
	DEBIT_PAYMENT                           = "DEBIT_PAYMENT"
	CREDIT_PAYMENT                          = "CREDIT_PAYMENT"
	CREDIT_DEPOSIT                          = "CREDIT_DEPOSIT"
	DEBIT_WITHDRAWAL                        = "DEBIT_WITHDRAWAL"
	CANCEL_WITHDRAWAL                       = "CANCEL_WITHDRAWAL"
	DEBIT_OUTGOING_PAYMENT                  = "DEBIT_OUTGOING_PAYMENT"
	CREDIT_INCOMING_PAYMENT                 = "CREDIT_INCOMING_PAYMENT"
	EXCHANGE_BUY                            = "EXCHANGE_BUY"
	EXCHANGE_SELL                           = "EXCHANGE_SELL"
)
