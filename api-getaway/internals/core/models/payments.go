package models

import (
	"time"
)

// I have put the models here but a better approach is having them in a separate file and linking them here
type PaystackAccountTopUp struct {
	Event string `json:"event"`
	Data  struct {
	  ID                int64     `json:"id"`
	  Domain            string    `json:"domain"`
	  Status            string    `json:"status"`
	  Reference         string    `json:"reference"`
	  Amount            int       `json:"amount"`
	  Message           string    `json:"message"`
	  GatewayResponse   string    `json:"gateway_response"`
	  CreatedAt         time.Time `json:"created_at"`
	  Channel           string    `json:"channel"`
	  Currency          string    `json:"currency"`
	  IPAddress         string    `json:"ip_address"`
	  Metadata          struct {
		CustomFields []struct {
		  DisplayName  string `json:"display_name"`
		  VariableName string `json:"variable_name"`
		  Value        string `json:"value"`
		} `json:"custom_fields"`
		Referrer     string `json:"referrer"`
	  } `json:"metadata"`
	  FeesBreakdown    struct {
		Amount  string `json:"amount"`
		Formula string `json:"formula"`
		Type    string `json:"type"`
	  } `json:"fees_breakdown"`
	  Log              string `json:"log"`
	  Fees              int     `json:"fees"`
	  FeesSplit         string `json:"fees_split"`
	  Authorization    struct {
		AuthorizationCode    string `json:"authorization_code"`
		Bin                  string `json:"bin"`
		Last4                string `json:"last4"`
		ExpMonth             string `json:"exp_month"`
		ExpYear              string `json:"exp_year"`
		Channel              string `json:"channel"`
		CardType             string `json:"card_type"`
		Bank                 string `json:"bank"`
		CountryCode          string `json:"country_code"`
		Brand                string `json:"brand"`
		Reusable             bool   `json:"reusable"`
		Signature            string `json:"signature"`
		AccountName          string `json:"account_name"`
		ReceiverBankAccountNumber string `json:"receiver_bank_account_number"`
		ReceiverBank          string `json:"receiver_bank"`
	  } `json:"authorization"`
	  Customer          struct {
		ID                    int       `json:"id"`
		FirstName              string    `json:"first_name"`
		LastName               string    `json:"last_name"`
		Email                  string    `json:"email"`
		CustomerCode           string    `json:"customer_code"`
		Phone                  string    `json:"phone"`
		Metadata              string    `json:"metadata"`
		RiskAction             string    `json:"risk_action"`
		InternationalFormatPhone string    `json:"international_format_phone"`
	  } `json:"customer"`
	  OrderID              string    `json:"order_id"`
	  PaidAt                time.Time `json:"paid_at"` // Consistent naming convention
	  RequestedAmount       int       `json:"requested_amount"`
	  PosTransactionData    string    `json:"pos_transaction_data"`
	  Source                struct {
		Type           string `json:"type"`
		Source         string `json:"source"`
		EntryPoint     string `json:"entry_point"`
		Identifier     string `json:"identifier"`
	  } `json:"source"`
	} `json:"data"`
  } 
  
