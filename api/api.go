package api

import (
	"fmt"
	"sync"

	"net/http"

	"github.com/gin-gonic/gin"
)

type HistoryEntry struct {
	EntryType string  `json:"entry_type"`
	Amount    float64 `json:"amount"`
	Balance   float64 `json:"balance"`
}

type API struct {
	router  *gin.Engine
	history []HistoryEntry
	lock    *sync.Mutex
	balance float64
}

type postRequest struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func (a *API) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	a.router.ServeHTTP(w, req)
}

func New() *API {
	a := &API{
		history: make([]HistoryEntry, 0),
		balance: float64(0),
		lock:    &sync.Mutex{},
	}

	r := gin.Default()

	grp1 := r.Group("/api")
	{
		grp1.POST("/transaction", a.Post)
		grp1.GET("/getBalance", a.GET)
		grp1.GET("/getHistory", a.GetBalance)
	}

	a.router = r

	return a
}

func (a *API) Post(c *gin.Context) {
	var req postRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid Json body: %s", err.Error()),
		})
		return
	}

	if req.Type != "d" && req.Type != "c" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Should be 'c' for credit or 'd' for debit",
		})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "amount can not be zero or negative",
		})
		return
	}

	a.lock.Lock()
	defer a.lock.Lock()

	switch req.Type {
	case "d":
		if a.balance-req.Amount < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "insuficient saving",
			})
			return
		}
	case "c":
		a.balance += req.Amount
	}

	a.history = append(a.history, HistoryEntry{
		EntryType: req.Type,
		Amount:    req.Amount,
		Balance:   a.balance,
	})

	c.JSONP(http.StatusOK, gin.H{
		"new_balance": a.balance,
	})

	return

}

func (a *API) GET(c *gin.Context) {
	a.lock.Lock()
	defer a.lock.Unlock()

	c.JSONP(http.StatusOK, gin.H{
		"balance": a.balance,
	})
}

func (a *API) GetBalance(c *gin.Context) {
	a.lock.Lock()
	defer a.lock.Unlock()
	c.JSONP(http.StatusOK, a.history)
}

func (a *API) Run() error {
	return a.router.Run(":8080")
}
