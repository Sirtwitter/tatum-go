package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/tatumio/tatum-go/model/response/common"
	"github.com/tatumio/tatum-go/model/response/xrp"
	"math/big"
	"net/url"
	"strconv"
)

type Xrp struct {
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetFee" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) XrpGetFee() uint32 {
	url := "/v3/xrp/fee"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	drops := result["drops"].(map[string]interface{})
	fee, err := strconv.ParseUint(fmt.Sprint(drops["base_fee"]), 10, 32)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return uint32(fee)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetAccountInfo" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) XrpGetAccountInfo(account string) *xrp.AccountInfo {

	url := "/v3/xrp/account/" + account
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	ledgerCurrentIndex := result["ledger_current_index"]
	accountData := result["account_data"].(map[string]interface{})
	sequence := accountData["Sequence"]

	return &xrp.AccountInfo{LedgerCurrentIndex: ledgerCurrentIndex.(uint32), Sequence: sequence.(uint32)}

}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpBroadcast" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) XrpBroadcast(txData string, signatureId string) *common.TransactionHash {
	url := "/v3/xrp/broadcast"

	payload := make(map[string]interface{})
	payload["txData"] = txData
	if len(signatureId) > 0 {
		payload["signatureId"] = signatureId
	}

	requestJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(requestJSON))

	var txHash common.TransactionHash
	var result map[string]interface{}
	res, err := sender.SendPost(url, requestJSON)
	if err == nil {
		json.Unmarshal([]byte(res), &result)
		txHash.TxId = fmt.Sprint(result["txId"])
	}
	return &txHash
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetLastClosedLedger" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) XrpGetCurrentLedger() uint64 {
	url := "/v3/xrp/info"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(res), &result)
	return result["ledger_index"].(uint64)
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetLedger" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) XrpGetLedger(i uint64) string {
	url := "/v3/xrp/ledger/" + string(i)
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetAccountBalance" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) xrpGetAccountBalance(address string) *big.Int {
	url := "/v3/xrp/account/" + address + "/balance"
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return big.NewInt(0)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(res), &result)
	balance, ok := new(big.Int).SetString(result["balance"].(string), 10)
	if !ok {
		fmt.Println(err.Error())
		return big.NewInt(0)
	}
	return balance
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetTransaction" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) xrpGetTransaction(hash string) string {
	url := "/v3/xrp/transaction/" + hash
	res, err := sender.SendGet(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res
}

/**
 * For more details, see <a href="https://tatum.io/apidoc#operation/XrpGetAccountTx" target="_blank">Tatum API documentation</a>
 */
func (x *Xrp) XrpGetAccountTransactions(address string, min uint32, marker string) string {
	url, _ := url.Parse("/v3/xrp/account/tx/" + address)
	q := url.Query()
	q.Add("min", strconv.FormatUint(uint64(min), 10))
	q.Add("marker", marker)
	url.RawQuery = q.Encode()
	fmt.Println(url.String())

	res, err := sender.SendGet(url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return res
}