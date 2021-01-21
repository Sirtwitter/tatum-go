package ledger

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/request"
	"github.com/tatumio/tatum-go/model/response/ledger"
	"net/url"
	"strconv"
)

type Transaction struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactionsByReference" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) GetTransactionsByReference(reference string) *[]ledger.Transaction {
	url, _ := url.Parse("/v3/ledger/transaction/reference/" + reference)

	var txs []ledger.Transaction
	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = json.Unmarshal([]byte(res), &txs)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &txs
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/sendTransaction" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) StoreTransaction(transaction request.CreateTransaction) *ledger.Reference {
	//await validateOrReject(transaction);

	url, _ := url.Parse("/v3/ledger/transaction")

	requestJSON, err := json.Marshal(transaction)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(url.String(), requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var reference ledger.Reference
	err = json.Unmarshal([]byte(res), &reference)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &reference
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactionsByAccountId" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) GetTransactionsByAccount(filter request.TransactionFilter, pageSize uint16, offset uint16) *[]ledger.Transaction {
	//await validateOrReject(filter);
	return getTransactions("/v3/ledger/transaction/account", filter, pageSize, offset)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactionsByCustomerId" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) GetTransactionsByCustomer(filter request.TransactionFilter, pageSize uint16, offset uint16) *[]ledger.Transaction {
	//await validateOrReject(filter);
	return getTransactions("/v3/ledger/transaction/customer", filter, pageSize, offset)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactions" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) GetTransactionsByLedger(filter request.TransactionFilter, pageSize uint16, offset uint16) *[]ledger.Transaction {
	//await validateOrReject(filter);
	return getTransactions("/v3/ledger/transaction/ledger", filter, pageSize, offset)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactionsByAccountId" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) CountTransactionsByAccount(filter request.TransactionFilter) uint64 {
	//await validateOrReject(filter);
	return countTransactions("/v3/ledger/transaction/account?count=true", filter)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactionsByCustomerId" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) CountTransactionsByCustomer(filter request.TransactionFilter) uint64 {
	//await validateOrReject(filter);
	return countTransactions("/v3/ledger/transaction/customer?count=true", filter)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/getTransactions" target="_blank">Tatum API documentation</a>
 */
func (t *Transaction) CountTransactionsByLedger(filter request.TransactionFilter) uint64 {
	//await validateOrReject(filter);
	return countTransactions("/v3/ledger/transaction/ledger?count=true", filter)
}

func getTransactions(urlStr string, filter request.TransactionFilter, pageSize uint16, offset uint16) *[]ledger.Transaction {
	url, _ := url.Parse(urlStr)
	q := url.Query()
	q.Add("offset", strconv.FormatUint(uint64(offset), 10))
	q.Add("pageSize", strconv.FormatUint(uint64(pageSize), 10))
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	requestJSON, err := json.Marshal(filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	res, err := sender.SendPost(url.String(), requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var txs []ledger.Transaction
	err = json.Unmarshal([]byte(res), &txs)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &txs
}

func countTransactions(url string, filter request.TransactionFilter) uint64 {
	requestJSON, err := json.Marshal(filter)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	res, err := sender.SendPost(url, requestJSON)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	count, err := strconv.ParseUint(res, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return count
}