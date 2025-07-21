package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
    ID               int    `json:"id"`
    AccountID        int    `json:"account_id"`
    Type             string `json:"type"`
    Amount           int    `json:"amount"`
    BalanceAfter     int    `json:"balance_after"`
    TransactionAt    string `json:"transaction_at"`
    Description      string `json:"description"`
    CounterAccountNo string `json:"counter_account_no"`
    CreatedAt        string `json:"created_at"`
    UpdatedAt        string `json:"updated_at"`
}

var transactions []Transaction

func loadTransactions() error {
    data, err := ioutil.ReadFile("db.json")
    if err != nil {
        return err
    }
    return json.Unmarshal(data, &transactions)
}

func getTransactionsByAccountID(c *gin.Context) {
    accountIDStr := c.Param("account_id")
    accountID, err := strconv.Atoi(accountIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account_id"})
        return
    }

    var filtered []Transaction
    for _, t := range transactions {
        if t.AccountID == accountID {
            filtered = append(filtered, t)
        }
    }

    if len(filtered) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No transactions found for account_id " + accountIDStr})
        return
    }

    c.JSON(http.StatusOK, filtered)
}

func main() {
    if err := loadTransactions(); err != nil {
        panic(err)
    }

    r := gin.Default()

    r.GET("/transactions/account/:account_id", getTransactionsByAccountID)

    r.Run(":3003")
}
