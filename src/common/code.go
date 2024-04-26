package common

import "errors"

type RespCode int64

const (
	SvcCtx = "svc_ctx"
)

const (
	CodeSuccess RespCode = 1000 + iota
	CodeInvalidParam
	CodeUserExists
	CodeUserNotFound
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthNotFound
	CodeInvalidAuth
	CodeInvalidToken
	CodeUserNotLogin
	CodeCommentDeleted
	CodeVoteTimeExpired
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "invalid parameters",
	CodeUserExists:      "user exists",
	CodeUserNotFound:    "user not found",
	CodeInvalidPassword: "invalid password",
	CodeServerBusy:      "server busy",
	CodeAuthNotFound:    "authorization not found",
	CodeInvalidAuth:     "invalid format of authorization",
	CodeInvalidToken:    "invalid token",
	CodeUserNotLogin:    "user not login",
	CodeCommentDeleted:  "comment has been deleted",
	CodeVoteTimeExpired: "cant vote for this post because the post was too old",
}

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

func (c RespCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

var (
	ErrorUserExist       = errors.New("user exists")
	ErrorUserNotFound    = errors.New("user not found")
	ErrorInvalidPassword = errors.New("invalid password")
	ErrorInvalidID       = errors.New("invalid id")
	ErrorCommentDeleted  = errors.New("comment has been deleted")
	ErrorUserNotLogin    = errors.New("user not login")

	ErrorRowsAffected = errors.New("rows affected != 1")
)

var (
	ErrorVoteTimeExpired = errors.New("vote time expired")
)
