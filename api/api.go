package api

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ekzjuperi/request-saver/models"
	"github.com/google/uuid"
	"github.com/jsimonetti/berkeleydb"
	"github.com/valyala/fasthttp"
)

// API struct contains api dependencies.
type API struct {
	dbConn *berkeleydb.Db
	port   string
}

// NewAPI creates new Api instance.
func NewAPI(
	dbConn *berkeleydb.Db,
	port string,
) *API {
	return &API{
		dbConn: dbConn,
		port:   port,
	}
}

// Start runs API.
func (a *API) Start() {
	err := fasthttp.ListenAndServe(fmt.Sprintf(":%v", a.port), a.SaveRequest())
	if err != nil {
		log.Fatalf("fasthttp.ListenAndServe() err: %v", err)
	}
}

// SaveRequest func save request to bd.
func (a *API) SaveRequest() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		r := &ctx.Request

		req := models.Request{
			TS:     time.Now().Unix(),
			Header: r.Header.String(),
			Body:   string(r.Body()),
			URI:    r.URI().String(),
		}

		b, err := json.Marshal(req)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)

			log.Printf("json.Marshal() err: %v", err)

			return
		}

		err = a.dbConn.Put(uuid.NewString(), string(b))
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)

			log.Printf("p.dbConn.Put() err: %v", err)

			return
		}

		log.Printf("request: %v saved to db", req)

		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}
