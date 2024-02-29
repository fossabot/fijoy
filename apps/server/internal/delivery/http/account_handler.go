package handler

import (
	"database/sql"
	"encoding/json"
	"fijoy/internal/gen/postgres/railway/public/model"
	"fmt"
	"net/http"
	"time"

	. "fijoy/internal/gen/postgres/railway/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-playground/validator"
	"github.com/nrednav/cuid2"
	"github.com/shopspring/decimal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type accountHandler struct {
	tokenAuth *jwtauth.JWTAuth
	db        *sql.DB
}

func NewAccountHandler(r *chi.Mux, tokenAuth *jwtauth.JWTAuth, db *sql.DB) {
	handler := &accountHandler{tokenAuth, db}

	r.Route("/v1/accounts", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/", handler.getAccounts)
		r.Post("/", handler.createAccount)
		r.Delete("/{accountID}", handler.deleteAccount)
	})
}

func (ah *accountHandler) getAccounts(w http.ResponseWriter, r *http.Request) {
	// _, claims, _ := jwtauth.FromContext(r.Context())
	workspaceID := r.URL.Query().Get("workspace_id")

	stmt := SELECT(FijoyAccount.AllColumns).FROM(FijoyAccount).
		WHERE(FijoyAccount.WorkspaceID.EQ(String(workspaceID)))

	dest := []model.FijoyAccount{}

	err := stmt.Query(ah.db, &dest)
	if err != nil {
		panic(err)
	}

	jsonText, _ := json.MarshalIndent(dest, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonText))
}

type createAccount struct {
	Name        string          `json:"Name" validate:"required"`
	AccountType string          `json:"AccountType" validate:"required"`
	Institution string          `json:"Institution" validate:"required"`
	Balance     decimal.Decimal `json:"Balance" validate:"required"`
	Currency    string          `json:"Currency" validate:"required"`
}

func (ah *accountHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	workspaceID := r.URL.Query().Get("workspace_id")
	if workspaceID == "" {
		http.Error(w, "Invalid request: workspace_id is required", http.StatusBadRequest)
		return
	}

	// TODO: check if this user has access to this workspace
	var createAccount createAccount

	err := json.NewDecoder(r.Body).Decode(&createAccount)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(createAccount)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	account := model.FijoyAccount{
		ID:          "account_" + cuid2.Generate(),
		Name:        createAccount.Name,
		AccountType: getAccountType(createAccount.AccountType),
		Institution: createAccount.Institution,
		WorkspaceID: workspaceID,
		Balance:     createAccount.Balance.InexactFloat64(),
		Currency:    createAccount.Currency,
		UpdatedAt:   time.Now().UTC(),
	}

	stmt := FijoyAccount.INSERT(FijoyAccount.AllColumns).MODEL(account)

	_, err = stmt.Exec(ah.db)
	if err != nil {
		http.Error(w, "Failed to create account: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ah *accountHandler) deleteAccount(w http.ResponseWriter, r *http.Request) {
	workspaceID := r.URL.Query().Get("workspace_id")
	if workspaceID == "" {
		http.Error(w, "Invalid request: workspace_id is required", http.StatusBadRequest)
		return
	}

	accountID := chi.URLParam(r, "accountID")

	stmt := FijoyAccount.DELETE().WHERE(FijoyAccount.ID.EQ(String(accountID)))

	_, err := stmt.Exec(ah.db)
	if err != nil {
		http.Error(w, "Failed to delete account: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAccountType(accountType string) model.FijoyAccountType {
	switch accountType {
	case "checking":
		return model.FijoyAccountType_Chequing
	case "savings":
		return model.FijoyAccountType_Savings
	case "credit":
		return model.FijoyAccountType_Credit
	default:
		return ""
	}
}