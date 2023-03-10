// Code generated by Kitex v0.4.4. DO NOT EDIT.

package accountservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	account "github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return accountServiceServiceInfo
}

var accountServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AccountService"
	handlerType := (*account.AccountService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateAccount":      kitex.NewMethodInfo(createAccountHandler, newCreateAccountArgs, newCreateAccountResult, false),
		"DeleteAccount":      kitex.NewMethodInfo(deleteAccountHandler, newDeleteAccountArgs, newDeleteAccountResult, false),
		"UpdateAccount":      kitex.NewMethodInfo(updateAccountHandler, newUpdateAccountArgs, newUpdateAccountResult, false),
		"MGetAccount":        kitex.NewMethodInfo(mGetAccountHandler, newMGetAccountArgs, newMGetAccountResult, false),
		"GetAccountByUserID": kitex.NewMethodInfo(getAccountByUserIDHandler, newGetAccountByUserIDArgs, newGetAccountByUserIDResult, false),
		"CreateBill":         kitex.NewMethodInfo(createBillHandler, newCreateBillArgs, newCreateBillResult, false),
		"DeleteBill":         kitex.NewMethodInfo(deleteBillHandler, newDeleteBillArgs, newDeleteBillResult, false),
		"UpdateBill":         kitex.NewMethodInfo(updateBillHandler, newUpdateBillArgs, newUpdateBillResult, false),
		"MGetBill":           kitex.NewMethodInfo(mGetBillHandler, newMGetBillArgs, newMGetBillResult, false),
		"GetBillByAccountID": kitex.NewMethodInfo(getBillByAccountIDHandler, newGetBillByAccountIDArgs, newGetBillByAccountIDResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "account",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.CreateAccountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).CreateAccount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateAccountArgs:
		success, err := handler.(account.AccountService).CreateAccount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateAccountResult)
		realResult.Success = success
	}
	return nil
}
func newCreateAccountArgs() interface{} {
	return &CreateAccountArgs{}
}

func newCreateAccountResult() interface{} {
	return &CreateAccountResult{}
}

type CreateAccountArgs struct {
	Req *account.CreateAccountRequest
}

func (p *CreateAccountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.CreateAccountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateAccountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateAccountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateAccountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateAccountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateAccountArgs) Unmarshal(in []byte) error {
	msg := new(account.CreateAccountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateAccountArgs_Req_DEFAULT *account.CreateAccountRequest

func (p *CreateAccountArgs) GetReq() *account.CreateAccountRequest {
	if !p.IsSetReq() {
		return CreateAccountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateAccountArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateAccountResult struct {
	Success *account.CreateAccountResponse
}

var CreateAccountResult_Success_DEFAULT *account.CreateAccountResponse

func (p *CreateAccountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.CreateAccountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateAccountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateAccountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateAccountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateAccountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateAccountResult) Unmarshal(in []byte) error {
	msg := new(account.CreateAccountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateAccountResult) GetSuccess() *account.CreateAccountResponse {
	if !p.IsSetSuccess() {
		return CreateAccountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateAccountResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.CreateAccountResponse)
}

func (p *CreateAccountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.DeleteAccountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).DeleteAccount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteAccountArgs:
		success, err := handler.(account.AccountService).DeleteAccount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteAccountResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteAccountArgs() interface{} {
	return &DeleteAccountArgs{}
}

func newDeleteAccountResult() interface{} {
	return &DeleteAccountResult{}
}

type DeleteAccountArgs struct {
	Req *account.DeleteAccountRequest
}

func (p *DeleteAccountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.DeleteAccountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteAccountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteAccountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteAccountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteAccountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteAccountArgs) Unmarshal(in []byte) error {
	msg := new(account.DeleteAccountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteAccountArgs_Req_DEFAULT *account.DeleteAccountRequest

func (p *DeleteAccountArgs) GetReq() *account.DeleteAccountRequest {
	if !p.IsSetReq() {
		return DeleteAccountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteAccountArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteAccountResult struct {
	Success *account.DeleteAccountResponse
}

var DeleteAccountResult_Success_DEFAULT *account.DeleteAccountResponse

func (p *DeleteAccountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.DeleteAccountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteAccountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteAccountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteAccountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteAccountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteAccountResult) Unmarshal(in []byte) error {
	msg := new(account.DeleteAccountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteAccountResult) GetSuccess() *account.DeleteAccountResponse {
	if !p.IsSetSuccess() {
		return DeleteAccountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteAccountResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.DeleteAccountResponse)
}

func (p *DeleteAccountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updateAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.UpdateAccountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).UpdateAccount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateAccountArgs:
		success, err := handler.(account.AccountService).UpdateAccount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateAccountResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateAccountArgs() interface{} {
	return &UpdateAccountArgs{}
}

func newUpdateAccountResult() interface{} {
	return &UpdateAccountResult{}
}

type UpdateAccountArgs struct {
	Req *account.UpdateAccountRequest
}

func (p *UpdateAccountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.UpdateAccountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateAccountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateAccountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateAccountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateAccountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateAccountArgs) Unmarshal(in []byte) error {
	msg := new(account.UpdateAccountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateAccountArgs_Req_DEFAULT *account.UpdateAccountRequest

func (p *UpdateAccountArgs) GetReq() *account.UpdateAccountRequest {
	if !p.IsSetReq() {
		return UpdateAccountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateAccountArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateAccountResult struct {
	Success *account.UpdateAccountResponse
}

var UpdateAccountResult_Success_DEFAULT *account.UpdateAccountResponse

func (p *UpdateAccountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.UpdateAccountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateAccountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateAccountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateAccountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateAccountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateAccountResult) Unmarshal(in []byte) error {
	msg := new(account.UpdateAccountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateAccountResult) GetSuccess() *account.UpdateAccountResponse {
	if !p.IsSetSuccess() {
		return UpdateAccountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateAccountResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.UpdateAccountResponse)
}

func (p *UpdateAccountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.MGetAccountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).MGetAccount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetAccountArgs:
		success, err := handler.(account.AccountService).MGetAccount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetAccountResult)
		realResult.Success = success
	}
	return nil
}
func newMGetAccountArgs() interface{} {
	return &MGetAccountArgs{}
}

func newMGetAccountResult() interface{} {
	return &MGetAccountResult{}
}

type MGetAccountArgs struct {
	Req *account.MGetAccountRequest
}

func (p *MGetAccountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.MGetAccountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MGetAccountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MGetAccountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MGetAccountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetAccountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetAccountArgs) Unmarshal(in []byte) error {
	msg := new(account.MGetAccountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetAccountArgs_Req_DEFAULT *account.MGetAccountRequest

func (p *MGetAccountArgs) GetReq() *account.MGetAccountRequest {
	if !p.IsSetReq() {
		return MGetAccountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetAccountArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetAccountResult struct {
	Success *account.MGetAccountResponse
}

var MGetAccountResult_Success_DEFAULT *account.MGetAccountResponse

func (p *MGetAccountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.MGetAccountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MGetAccountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MGetAccountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MGetAccountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetAccountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetAccountResult) Unmarshal(in []byte) error {
	msg := new(account.MGetAccountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetAccountResult) GetSuccess() *account.MGetAccountResponse {
	if !p.IsSetSuccess() {
		return MGetAccountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetAccountResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.MGetAccountResponse)
}

func (p *MGetAccountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getAccountByUserIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.GetAccountByUserIDRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).GetAccountByUserID(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetAccountByUserIDArgs:
		success, err := handler.(account.AccountService).GetAccountByUserID(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetAccountByUserIDResult)
		realResult.Success = success
	}
	return nil
}
func newGetAccountByUserIDArgs() interface{} {
	return &GetAccountByUserIDArgs{}
}

func newGetAccountByUserIDResult() interface{} {
	return &GetAccountByUserIDResult{}
}

type GetAccountByUserIDArgs struct {
	Req *account.GetAccountByUserIDRequest
}

func (p *GetAccountByUserIDArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.GetAccountByUserIDRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetAccountByUserIDArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetAccountByUserIDArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetAccountByUserIDArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetAccountByUserIDArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetAccountByUserIDArgs) Unmarshal(in []byte) error {
	msg := new(account.GetAccountByUserIDRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetAccountByUserIDArgs_Req_DEFAULT *account.GetAccountByUserIDRequest

func (p *GetAccountByUserIDArgs) GetReq() *account.GetAccountByUserIDRequest {
	if !p.IsSetReq() {
		return GetAccountByUserIDArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetAccountByUserIDArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetAccountByUserIDResult struct {
	Success *account.GetAccountByUserIDResponse
}

var GetAccountByUserIDResult_Success_DEFAULT *account.GetAccountByUserIDResponse

func (p *GetAccountByUserIDResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.GetAccountByUserIDResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetAccountByUserIDResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetAccountByUserIDResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetAccountByUserIDResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetAccountByUserIDResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetAccountByUserIDResult) Unmarshal(in []byte) error {
	msg := new(account.GetAccountByUserIDResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetAccountByUserIDResult) GetSuccess() *account.GetAccountByUserIDResponse {
	if !p.IsSetSuccess() {
		return GetAccountByUserIDResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetAccountByUserIDResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.GetAccountByUserIDResponse)
}

func (p *GetAccountByUserIDResult) IsSetSuccess() bool {
	return p.Success != nil
}

func createBillHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.CreateBillRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).CreateBill(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateBillArgs:
		success, err := handler.(account.AccountService).CreateBill(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateBillResult)
		realResult.Success = success
	}
	return nil
}
func newCreateBillArgs() interface{} {
	return &CreateBillArgs{}
}

func newCreateBillResult() interface{} {
	return &CreateBillResult{}
}

type CreateBillArgs struct {
	Req *account.CreateBillRequest
}

func (p *CreateBillArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.CreateBillRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateBillArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateBillArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateBillArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateBillArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateBillArgs) Unmarshal(in []byte) error {
	msg := new(account.CreateBillRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateBillArgs_Req_DEFAULT *account.CreateBillRequest

func (p *CreateBillArgs) GetReq() *account.CreateBillRequest {
	if !p.IsSetReq() {
		return CreateBillArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateBillArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateBillResult struct {
	Success *account.CreateBillResponse
}

var CreateBillResult_Success_DEFAULT *account.CreateBillResponse

func (p *CreateBillResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.CreateBillResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateBillResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateBillResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateBillResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateBillResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateBillResult) Unmarshal(in []byte) error {
	msg := new(account.CreateBillResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateBillResult) GetSuccess() *account.CreateBillResponse {
	if !p.IsSetSuccess() {
		return CreateBillResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateBillResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.CreateBillResponse)
}

func (p *CreateBillResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteBillHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.DeleteBillRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).DeleteBill(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteBillArgs:
		success, err := handler.(account.AccountService).DeleteBill(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteBillResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteBillArgs() interface{} {
	return &DeleteBillArgs{}
}

func newDeleteBillResult() interface{} {
	return &DeleteBillResult{}
}

type DeleteBillArgs struct {
	Req *account.DeleteBillRequest
}

func (p *DeleteBillArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.DeleteBillRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteBillArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteBillArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteBillArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteBillArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteBillArgs) Unmarshal(in []byte) error {
	msg := new(account.DeleteBillRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteBillArgs_Req_DEFAULT *account.DeleteBillRequest

func (p *DeleteBillArgs) GetReq() *account.DeleteBillRequest {
	if !p.IsSetReq() {
		return DeleteBillArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteBillArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteBillResult struct {
	Success *account.DeleteBillResponse
}

var DeleteBillResult_Success_DEFAULT *account.DeleteBillResponse

func (p *DeleteBillResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.DeleteBillResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteBillResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteBillResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteBillResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteBillResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteBillResult) Unmarshal(in []byte) error {
	msg := new(account.DeleteBillResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteBillResult) GetSuccess() *account.DeleteBillResponse {
	if !p.IsSetSuccess() {
		return DeleteBillResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteBillResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.DeleteBillResponse)
}

func (p *DeleteBillResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updateBillHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.UpdateBillRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).UpdateBill(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateBillArgs:
		success, err := handler.(account.AccountService).UpdateBill(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateBillResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateBillArgs() interface{} {
	return &UpdateBillArgs{}
}

func newUpdateBillResult() interface{} {
	return &UpdateBillResult{}
}

type UpdateBillArgs struct {
	Req *account.UpdateBillRequest
}

func (p *UpdateBillArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.UpdateBillRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateBillArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateBillArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateBillArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateBillArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateBillArgs) Unmarshal(in []byte) error {
	msg := new(account.UpdateBillRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateBillArgs_Req_DEFAULT *account.UpdateBillRequest

func (p *UpdateBillArgs) GetReq() *account.UpdateBillRequest {
	if !p.IsSetReq() {
		return UpdateBillArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateBillArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateBillResult struct {
	Success *account.UpdateBillResponse
}

var UpdateBillResult_Success_DEFAULT *account.UpdateBillResponse

func (p *UpdateBillResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.UpdateBillResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateBillResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateBillResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateBillResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateBillResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateBillResult) Unmarshal(in []byte) error {
	msg := new(account.UpdateBillResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateBillResult) GetSuccess() *account.UpdateBillResponse {
	if !p.IsSetSuccess() {
		return UpdateBillResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateBillResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.UpdateBillResponse)
}

func (p *UpdateBillResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetBillHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.MGetBillRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).MGetBill(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetBillArgs:
		success, err := handler.(account.AccountService).MGetBill(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetBillResult)
		realResult.Success = success
	}
	return nil
}
func newMGetBillArgs() interface{} {
	return &MGetBillArgs{}
}

func newMGetBillResult() interface{} {
	return &MGetBillResult{}
}

type MGetBillArgs struct {
	Req *account.MGetBillRequest
}

func (p *MGetBillArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.MGetBillRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MGetBillArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MGetBillArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MGetBillArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetBillArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetBillArgs) Unmarshal(in []byte) error {
	msg := new(account.MGetBillRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetBillArgs_Req_DEFAULT *account.MGetBillRequest

func (p *MGetBillArgs) GetReq() *account.MGetBillRequest {
	if !p.IsSetReq() {
		return MGetBillArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetBillArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetBillResult struct {
	Success *account.MGetBillResponse
}

var MGetBillResult_Success_DEFAULT *account.MGetBillResponse

func (p *MGetBillResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.MGetBillResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MGetBillResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MGetBillResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MGetBillResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetBillResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetBillResult) Unmarshal(in []byte) error {
	msg := new(account.MGetBillResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetBillResult) GetSuccess() *account.MGetBillResponse {
	if !p.IsSetSuccess() {
		return MGetBillResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetBillResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.MGetBillResponse)
}

func (p *MGetBillResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getBillByAccountIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(account.GetBillByAccountIDRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(account.AccountService).GetBillByAccountID(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetBillByAccountIDArgs:
		success, err := handler.(account.AccountService).GetBillByAccountID(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetBillByAccountIDResult)
		realResult.Success = success
	}
	return nil
}
func newGetBillByAccountIDArgs() interface{} {
	return &GetBillByAccountIDArgs{}
}

func newGetBillByAccountIDResult() interface{} {
	return &GetBillByAccountIDResult{}
}

type GetBillByAccountIDArgs struct {
	Req *account.GetBillByAccountIDRequest
}

func (p *GetBillByAccountIDArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(account.GetBillByAccountIDRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetBillByAccountIDArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetBillByAccountIDArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetBillByAccountIDArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetBillByAccountIDArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetBillByAccountIDArgs) Unmarshal(in []byte) error {
	msg := new(account.GetBillByAccountIDRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetBillByAccountIDArgs_Req_DEFAULT *account.GetBillByAccountIDRequest

func (p *GetBillByAccountIDArgs) GetReq() *account.GetBillByAccountIDRequest {
	if !p.IsSetReq() {
		return GetBillByAccountIDArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetBillByAccountIDArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetBillByAccountIDResult struct {
	Success *account.GetBillByAccountIDResponse
}

var GetBillByAccountIDResult_Success_DEFAULT *account.GetBillByAccountIDResponse

func (p *GetBillByAccountIDResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(account.GetBillByAccountIDResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetBillByAccountIDResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetBillByAccountIDResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetBillByAccountIDResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetBillByAccountIDResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetBillByAccountIDResult) Unmarshal(in []byte) error {
	msg := new(account.GetBillByAccountIDResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetBillByAccountIDResult) GetSuccess() *account.GetBillByAccountIDResponse {
	if !p.IsSetSuccess() {
		return GetBillByAccountIDResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetBillByAccountIDResult) SetSuccess(x interface{}) {
	p.Success = x.(*account.GetBillByAccountIDResponse)
}

func (p *GetBillByAccountIDResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateAccount(ctx context.Context, Req *account.CreateAccountRequest) (r *account.CreateAccountResponse, err error) {
	var _args CreateAccountArgs
	_args.Req = Req
	var _result CreateAccountResult
	if err = p.c.Call(ctx, "CreateAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteAccount(ctx context.Context, Req *account.DeleteAccountRequest) (r *account.DeleteAccountResponse, err error) {
	var _args DeleteAccountArgs
	_args.Req = Req
	var _result DeleteAccountResult
	if err = p.c.Call(ctx, "DeleteAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAccount(ctx context.Context, Req *account.UpdateAccountRequest) (r *account.UpdateAccountResponse, err error) {
	var _args UpdateAccountArgs
	_args.Req = Req
	var _result UpdateAccountResult
	if err = p.c.Call(ctx, "UpdateAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetAccount(ctx context.Context, Req *account.MGetAccountRequest) (r *account.MGetAccountResponse, err error) {
	var _args MGetAccountArgs
	_args.Req = Req
	var _result MGetAccountResult
	if err = p.c.Call(ctx, "MGetAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAccountByUserID(ctx context.Context, Req *account.GetAccountByUserIDRequest) (r *account.GetAccountByUserIDResponse, err error) {
	var _args GetAccountByUserIDArgs
	_args.Req = Req
	var _result GetAccountByUserIDResult
	if err = p.c.Call(ctx, "GetAccountByUserID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateBill(ctx context.Context, Req *account.CreateBillRequest) (r *account.CreateBillResponse, err error) {
	var _args CreateBillArgs
	_args.Req = Req
	var _result CreateBillResult
	if err = p.c.Call(ctx, "CreateBill", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteBill(ctx context.Context, Req *account.DeleteBillRequest) (r *account.DeleteBillResponse, err error) {
	var _args DeleteBillArgs
	_args.Req = Req
	var _result DeleteBillResult
	if err = p.c.Call(ctx, "DeleteBill", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateBill(ctx context.Context, Req *account.UpdateBillRequest) (r *account.UpdateBillResponse, err error) {
	var _args UpdateBillArgs
	_args.Req = Req
	var _result UpdateBillResult
	if err = p.c.Call(ctx, "UpdateBill", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetBill(ctx context.Context, Req *account.MGetBillRequest) (r *account.MGetBillResponse, err error) {
	var _args MGetBillArgs
	_args.Req = Req
	var _result MGetBillResult
	if err = p.c.Call(ctx, "MGetBill", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetBillByAccountID(ctx context.Context, Req *account.GetBillByAccountIDRequest) (r *account.GetBillByAccountIDResponse, err error) {
	var _args GetBillByAccountIDArgs
	_args.Req = Req
	var _result GetBillByAccountIDResult
	if err = p.c.Call(ctx, "GetBillByAccountID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
