package service

import (
	"github.com/ideal-rucksack/workflow-scheduler/pkg/errors"
	"github.com/ideal-rucksack/workflow-scheduler/pkg/util"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/api/requests"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/api/response"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/middleware"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"time"
)

type AccountService struct {
	repo *repo.AccountRepo
}

func (s AccountService) Signup(payload requests.Signup) (*response.Signup, error) {

	var (
		err               error
		accountRepository = s.repo
	)

	usernameAccount, err := accountRepository.QueryByUsername(payload.Username, false)
	if err != nil {
		return nil, err
	}

	if usernameAccount != nil {
		return nil, errors.NewIllegalArgumentError("username already exists")
	}

	emailAccount, err := accountRepository.QueryByEmail(payload.Email, false)
	if err != nil {
		return nil, err
	}

	if emailAccount != nil {
		return nil, errors.NewIllegalArgumentError("email already exists")
	}

	now := time.Now()
	deleted := false
	secret := util.GenerateRandomString(16)
	account := entities.Account{
		GlobalEntity: entities.GlobalEntity{
			Id:       nil,
			CreateAt: &now,
			ModifyAt: &now,
		},
		Nickname: &payload.Nickname,
		Code:     payload.Code,
		Username: &payload.Username,
		Secret:   &secret,
		Email:    &payload.Email,
		Status:   &payload.Status,
		Deleted:  &deleted,
	}

	err = accountRepository.Insert(account)
	if err != nil {
		return nil, err
	}

	return &response.Signup{
		Nickname: payload.Nickname,
		Username: payload.Username,
		Email:    payload.Email,
		Status:   payload.Status,
	}, err
}

// Verify 验证账户
// FIXME 这里验证账户暂时有问题 应该是通过邮箱验证并且 用户在没有通过验证即为没有注册成功 为了方便先暂时这样
func (s AccountService) Verify(payload requests.Verify) (*response.Verify, error) {
	var (
		err               error
		accountRepository = s.repo
	)

	account, err := accountRepository.QueryByUsername(payload.Username, false)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.NewIllegalArgumentError("account not exists")
	}

	if account.Password != nil {
		return nil, errors.NewIllegalArgumentError("account already verified")
	}

	// TODO verify code write to cache
	if payload.VerifyCode != "12345" {
		return nil, errors.NewIllegalArgumentError("verify code error")
	}

	active := entities.ACTIVE
	account.Status = &active

	now := time.Now()
	account.GlobalEntity.ModifyAt = &now
	hash := util.MD5SaltHash(payload.Password, *account.Secret)
	account.Password = &hash

	refreshToken, err := util.GenerateToken(*account.Id, *account.Secret, nil)
	if err != nil {
		return nil, err
	}

	account.RefreshToken = &refreshToken
	err = accountRepository.Update(*account)
	if err != nil {
		return nil, err
	}

	var expiresAt = time.Minute
	accessToken, err := util.GenerateToken(*account.Id, *account.Secret, &expiresAt)
	if err != nil {
		return nil, err
	}

	middleware.TokenStoreCache.SetToken(accessToken, *account.Secret)

	return &response.Verify{
		AccessToken:  accessToken,
		RefreshToken: &refreshToken,
	}, nil
}

func (s AccountService) SignIn(payload requests.SignIn) (*response.SignIn, error) {
	var (
		err               error
		accountRepository = s.repo
	)

	usernameAccount, err := accountRepository.QueryByUsername(payload.Username, false)
	if err != nil {
		return nil, err
	}

	if usernameAccount == nil {
		return nil, errors.NewIllegalArgumentError("username or password error")
	}

	var password = util.MD5SaltHash(payload.Password, *usernameAccount.Secret)

	account, err := accountRepository.QuerySignIn(payload.Username, password, false)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.NewIllegalArgumentError("username or password error")
	}

	hash := util.MD5SaltHash(payload.Password, *account.Secret)
	if *account.Password != hash {
		return nil, errors.NewIllegalArgumentError("password error")
	}

	var expiresAt = time.Minute
	accessToken, err := util.GenerateToken(*account.Id, *account.Secret, &expiresAt)
	if err != nil {
		return nil, err
	}

	refreshToken, err := util.GenerateToken(*account.Id, *account.Secret, nil)
	if err != nil {
		return nil, err
	}

	account.RefreshToken = &refreshToken
	err = accountRepository.Update(*account)
	if err != nil {
		return nil, err
	}

	middleware.TokenStoreCache.SetToken(accessToken, *account.Secret)

	return &response.SignIn{
		AccessToken:  accessToken,
		RefreshToken: account.RefreshToken,
	}, nil
}

func (s AccountService) SignOut(id int64) error {
	var (
		err               error
		accountRepository = s.repo
	)

	account, err := accountRepository.QueryById(id, false)
	if err != nil {
		return err
	}

	if account == nil {
		return nil
	}

	account.RefreshToken = nil
	return accountRepository.Update(*account)
}

func NewAccountService(repo *repo.AccountRepo) *AccountService {
	return &AccountService{repo: repo}
}
