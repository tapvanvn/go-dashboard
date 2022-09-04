package common

import (
	"errors"

	"github.com/tapvanvn/godepsolver"
)

var ErrInvalidPath = errors.New("InvalidPath")

type MODULE string

const (
	MODULE_CHECK_DATABASE = MODULE("check_data") //check if database structure init yet
	MODULE_APIKEY         = MODULE("apikey")     //system secret key
	MODULE_JWT            = MODULE("jwt")        //jwt
	MODULE_SESSION        = MODULE("session")    //
	MODULE_AUTH           = MODULE("auth")       //authenication module for verify 3rd account
	MODULE_ACCOUNT        = MODULE("account")    //manage account
	MODULE_ADMIN          = MODULE("admin")
)

var ModuleDependencies map[string][]string = map[string][]string{
	string(MODULE_JWT): {
		string(MODULE_APIKEY),
	},
	string(MODULE_AUTH): {
		string(MODULE_SESSION),
	},
	string(MODULE_ACCOUNT): {
		string(MODULE_CHECK_DATABASE),
		string(MODULE_JWT),
		string(MODULE_ACCOUNT),
		string(MODULE_SESSION),
	},
	string(MODULE_ADMIN): {
		string(MODULE_ACCOUNT),
	},
}

var ModuleDependencySolver *godepsolver.GeneralSolver = godepsolver.NewGeneralSolver(ModuleDependencies)

var EmptyModules map[MODULE]bool = map[MODULE]bool{}
