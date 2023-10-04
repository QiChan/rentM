package logic

import (
	"context"
	"time"

	"rentM/app/usercenter/cmd/rpc/internal/svc"
	"rentM/app/usercenter/cmd/rpc/pb"
	"rentM/common/ctxdata"
	"rentM/common/xerr"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	tn := time.Now().Unix()
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.AuthKey)
	if err != nil {
		return nil, err
	}
	uid := user.Id
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, tn, l.svcCtx.Config.JwtAuth.AccessExpire, uid)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResp{
		AccessToken:  jwtToken,
		AccessExpire: tn + l.svcCtx.Config.JwtAuth.AccessExpire,
		RefreshAfter: tn + l.svcCtx.Config.JwtAuth.AccessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
