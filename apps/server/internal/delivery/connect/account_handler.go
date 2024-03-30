package handler

import (
	"context"
	"database/sql"
	"errors"
	"fijoy/internal/entity"
	"fijoy/internal/gen/postgres/model"
	"fijoy/internal/gen/proto/fijoy/v1/fijoyv1connect"
	"fijoy/internal/util"
	"time"

	. "fijoy/internal/gen/postgres/table"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/nrednav/cuid2"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	fijoyv1 "fijoy/internal/gen/proto/fijoy/v1"
)

type AccountServer struct {
	db *sql.DB
}

func NewAccountHandler(r *chi.Mux, tokenAuth *jwtauth.JWTAuth, db *sql.DB) {
	accountServer := &AccountServer{db: db}

	path, handler := fijoyv1connect.NewAccountServiceHandler(accountServer)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Mount(path, handler)
	})
}

func jetAccountToConnectAccount(a *entity.FijoyAccount) *fijoyv1.Account {
	return &fijoyv1.Account{
		Id:          a.ID,
		WorkspaceId: a.WorkspaceID,
		Name:        a.Name,
		AccountType: util.JetAccountTypeToConnectAccountType[a.AccountType],
		Balance:     util.DecimalToMoney(a.Balance),
		Currency:    a.Currency,
		Institution: a.Institution,
		CreatedAt:   timestamppb.New(a.CreatedAt),
		UpdatedAt:   timestamppb.New(a.UpdatedAt),
	}
}

func (s *AccountServer) CreateAccount(
	ctx context.Context,
	req *connect.Request[fijoyv1.CreateAccountRequest],
) (*connect.Response[fijoyv1.Account], error) {
	workspaceId, err := util.ExtractWorkspaceIdFromHeader(req.Header())
	if err != nil {
		return nil, err
	}

	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	workspaceUser, err := util.GetWorkspaceUserPermission(s.db, userId, workspaceId)
	if err != nil {
		return nil, err
	}

	if !util.HasEditPermission(&workspaceUser) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("user does not have edit permission"))
	}

	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	if err = v.Validate(req.Msg); err != nil {
		return nil, err
	}

	account := entity.FijoyAccount{
		FijoyAccount: model.FijoyAccount{
			ID:          "account_" + cuid2.Generate(),
			Name:        req.Msg.Name,
			AccountType: util.ConnectAccountTypeToJetAccountType[req.Msg.AccountType],
			Institution: req.Msg.Institution,
			WorkspaceID: workspaceId,
			Currency:    req.Msg.Currency,
			UpdatedAt:   time.Now().UTC(),
		},
		Balance: decimal.NewFromInt(0),
	}

	stmt := FijoyAccount.INSERT(FijoyAccount.AllColumns).MODEL(account).
		RETURNING(FijoyAccount.AllColumns)

	dest := entity.FijoyAccount{}

	err = stmt.QueryContext(ctx, s.db, &dest)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(jetAccountToConnectAccount(&dest))

	return res, nil
}

func (s *AccountServer) GetAccounts(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[fijoyv1.Accounts], error) {
	workspaceId, err := util.ExtractWorkspaceIdFromHeader(req.Header())
	if err != nil {
		return nil, err
	}

	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	workspaceUser, err := util.GetWorkspaceUserPermission(s.db, userId, workspaceId)
	if err != nil {
		return nil, err
	}

	if !util.HasViewPermission(&workspaceUser) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("user does not have view permission"))
	}

	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	if err = v.Validate(req.Msg); err != nil {
		return nil, err
	}

	stmt := SELECT(FijoyAccount.AllColumns).FROM(FijoyAccount).
		WHERE(FijoyAccount.WorkspaceID.EQ(String(workspaceId)))

	dest := []entity.FijoyAccount{}

	err = stmt.QueryContext(ctx, s.db, &dest)
	if err != nil {
		return nil, err
	}

	accounts := make([]*fijoyv1.Account, len(dest))
	for i, w := range dest {
		accounts[i] = jetAccountToConnectAccount(&w)
	}

	res := connect.NewResponse(&fijoyv1.Accounts{
		Accounts: accounts,
	})

	return res, nil
}

func (s *AccountServer) GetAccountById(
	ctx context.Context,
	req *connect.Request[fijoyv1.GetAccountByIdRequest],
) (*connect.Response[fijoyv1.Account], error) {
	workspaceId, err := util.ExtractWorkspaceIdFromHeader(req.Header())
	if err != nil {
		return nil, err
	}

	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	workspaceUser, err := util.GetWorkspaceUserPermission(s.db, userId, workspaceId)
	if err != nil {
		return nil, err
	}

	if !util.HasViewPermission(&workspaceUser) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("user does not have view permission"))
	}

	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	if err = v.Validate(req.Msg); err != nil {
		return nil, err
	}

	stmt := SELECT(FijoyAccount.AllColumns).FROM(FijoyAccount).
		WHERE(AND(
			FijoyAccount.ID.EQ(String(req.Msg.Id)),
			FijoyAccount.WorkspaceID.EQ(String(workspaceId)),
		))

	dest := entity.FijoyAccount{}

	err = stmt.QueryContext(ctx, s.db, &dest)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(jetAccountToConnectAccount(&dest))

	return res, nil
}

func (s *AccountServer) DeleteAccountById(
	ctx context.Context,
	req *connect.Request[fijoyv1.DeleteAccountByIdRequest],
) (*connect.Response[fijoyv1.Account], error) {
	workspaceId, err := util.ExtractWorkspaceIdFromHeader(req.Header())
	if err != nil {
		return nil, err
	}

	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	workspaceUser, err := util.GetWorkspaceUserPermission(s.db, userId, workspaceId)
	if err != nil {
		return nil, err
	}

	if !util.HasEditPermission(&workspaceUser) {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("user does not have edit permission"))
	}

	v, err := protovalidate.New()
	if err != nil {
		return nil, err
	}

	if err = v.Validate(req.Msg); err != nil {
		return nil, err
	}

	stmt := FijoyAccount.DELETE().WHERE(FijoyAccount.ID.EQ(String(req.Msg.Id))).RETURNING(FijoyAccount.AllColumns)

	dest := entity.FijoyAccount{}

	err = stmt.QueryContext(ctx, s.db, &dest)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(jetAccountToConnectAccount(&dest))

	return res, nil
}
