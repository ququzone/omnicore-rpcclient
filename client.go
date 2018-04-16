package rpcclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
)

var (
	httpClient *http.Client
)

// Client ...
type Client struct {
	*rpcclient.Client
	User     string
	Password string
	URL      string
}

// New return new rpc client
func New(connect string, port int, user, password string) (*Client, error) {
	conn := &rpcclient.ConnConfig{
		Host:         fmt.Sprintf("%s:%d", connect, port),
		User:         user,
		Pass:         password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	c, err := rpcclient.New(conn, nil)
	if err != nil {
		return nil, err
	}
	return &Client{c, user, password, fmt.Sprintf("http://%s:%d", connect, port)}, nil
}

// OmniGettransaction ...
func (c *Client) OmniGettransaction(txHash string) (*Transaction, error) {
	cmd := NewGetTransactionCmd(txHash)
	resBytes, err := c.sendCmd(cmd)
	if err != nil {
		return nil, err
	}

	var result Transaction
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// OmniListBlockTransactions ...
func (c *Client) OmniListBlockTransactions(index int64) ([]string, error) {
	cmd := NewListBlockTransactionsCmd(index)
	resBytes, err := c.sendCmd(cmd)
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OmniGetBalance ...
func (c *Client) OmniGetBalance(address string, propertyID int64) (*Balance, error) {
	cmd := NewGetBalanceCmd(address, propertyID)
	resBytes, err := c.sendCmd(cmd)
	if err != nil {
		return nil, err
	}

	var result Balance
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) sendCmd(cmd interface{}) ([]byte, error) {
	// Get the method associated with the command.
	method, err := btcjson.CmdMethod(cmd)
	if err != nil {
		return nil, err
	}

	// Marshal the command.
	id := c.NextID()
	marshalledJSON, err := btcjson.MarshalCmd(id, cmd)
	if err != nil {
		return nil, err
	}

	jReq := &jsonRequest{
		id:             id,
		method:         method,
		cmd:            cmd,
		marshalledJSON: marshalledJSON,
	}
	return c.sendRequest(jReq)
}

func (c *Client) sendRequest(jReq *jsonRequest) ([]byte, error) {
	bodyReader := bytes.NewReader(jReq.marshalledJSON)
	httpReq, err := http.NewRequest("POST", c.URL, bodyReader)
	if err != nil {
		return nil, err
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")

	// Configure basic access authorization.
	httpReq.SetBasicAuth(c.User, c.Password)

	return c.sendPostRequest(httpReq, jReq)
}

func (c *Client) sendPostRequest(httpReq *http.Request, jReq *jsonRequest) ([]byte, error) {
	httpResponse, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		return nil, err
	}

	var res Response
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return nil, err
	}
	if res.Error.Code != 0 || res.Error.Message != "" {
		return nil, errors.New(res.Error.Message)
	}

	return res.Result, nil
}
