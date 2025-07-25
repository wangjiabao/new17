package data

import (
	"context"
	"dhb/app/app/internal/biz"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID                     int64     `gorm:"primarykey;type:int"`
	Address                string    `gorm:"type:varchar(100)"`
	Password               string    `gorm:"type:varchar(100)"`
	AddressTwo             string    `gorm:"type:varchar(100)"`
	PrivateKey             string    `gorm:"type:varchar(200)"`
	Last                   uint64    `gorm:"type:bigint;not null"`
	LastBiw                uint64    `gorm:"type:bigint;not null"`
	Amount                 uint64    `gorm:"type:bigint;not null"`
	AmountBiw              uint64    `gorm:"type:bigint;not null"`
	AmountUsdt             float64   `gorm:"type:decimal(65,20);not null"`
	MyTotalAmount          float64   `gorm:"type:decimal(65,20);not null"`
	AmountUsdtGet          float64   `gorm:"type:decimal(65,20);not null"`
	AmountUsdtOrigin       float64   `gorm:"type:decimal(65,20);not null"`
	AmountRecommendUsdtGet float64   `gorm:"type:decimal(65,20);not null"`
	AddressThree           string    `gorm:"type:varchar(100)"`
	PrivateKeyThree        string    `gorm:"type:varchar(400)"`
	WordThree              string    `gorm:"type:varchar(200)"`
	Undo                   int64     `gorm:"type:int;not null"`
	Vip                    int64     `gorm:"type:int;not null"`
	VipAdmin               int64     `gorm:"type:int;not null"`
	LockReward             int64     `gorm:"type:int;not null"`
	RecommendLevel         int64     `gorm:"type:int;not null"`
	OutRate                int64     `gorm:"type:int;not null"`
	Lock                   int64     `gorm:"type:int;not null"`
	CreatedAt              time.Time `gorm:"type:datetime;not null"`
	UpdatedAt              time.Time `gorm:"type:datetime;not null"`
	RecommendUserReward    int64     `gorm:"type:int;not null"`
	RecommendUser          int64     `gorm:"type:int;not null"`
	RecommendUserH         int64     `gorm:"type:int;not null"`
	AmountFour             float64   `gorm:"type:decimal(65,20);not null"`
	AmountFourGet          float64   `gorm:"type:decimal(65,20);not null"`
}

type Stake struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Status    int64     `gorm:"type:int;not null"`
	Day       int64     `gorm:"type:int;not null"`
	Amount    float64   `gorm:"type:decimal(65,20);not null"`
	Reward    float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserArea struct {
	ID         int64     `gorm:"primarykey;type:int"`
	UserId     int64     `gorm:"type:int;not null"`
	Level      int64     `gorm:"type:int;not null"`
	Amount     int64     `gorm:"type:bigint;not null"`
	SelfAmount int64     `gorm:"type:bigint;not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

type UserInfo struct {
	ID               int64     `gorm:"primarykey;type:int"`
	UserId           int64     `gorm:"type:int;not null"`
	Vip              int64     `gorm:"type:int;not null"`
	HistoryRecommend int64     `gorm:"type:int;not null"`
	LockVip          int64     `gorm:"type:int;not null"`
	UseVip           int64     `gorm:"type:int;not null"`
	TeamCsdBalance   int64     `gorm:"type:bigint;not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type UserRecommend struct {
	ID            int64     `gorm:"primarykey;type:int"`
	UserId        int64     `gorm:"type:int;not null"`
	RecommendCode string    `gorm:"type:varchar(10000);not null"`
	Total         int64     `gorm:"type:bigint;not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type UserCurrentMonthRecommend struct {
	ID              int64     `gorm:"primarykey;type:int"`
	UserId          int64     `gorm:"type:int;not null"`
	RecommendUserId int64     `gorm:"type:int;not null"`
	Date            time.Time `gorm:"type:datetime;not null"`
	CreatedAt       time.Time `gorm:"type:datetime;not null"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null"`
}

type Config struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(45);not null"`
	KeyName   string    `gorm:"type:varchar(45);not null"`
	Value     string    `gorm:"type:varchar(1000);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PriceChange struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Origin    int64     `gorm:"type:bigint"`
	Price     int64     `gorm:"type:bigint"`
	Status    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalance struct {
	ID                     int64     `gorm:"primarykey;type:int"`
	UserId                 int64     `gorm:"type:int"`
	BalanceUsdt            int64     `gorm:"type:bigint"`
	BalanceUsdtNew         int64     `gorm:"type:bigint"`
	BalanceUsdtFloat       float64   `gorm:"type:decimal(65,20);not null"`
	BalanceRawFloat        float64   `gorm:"type:decimal(65,20);not null"`
	BalanceKsdtFloat       float64   `gorm:"type:decimal(65,20);not null"`
	BalanceDhb             int64     `gorm:"type:bigint"`
	BalanceC               int64     `gorm:"type:bigint"`
	CreatedAt              time.Time `gorm:"type:datetime;not null"`
	UpdatedAt              time.Time `gorm:"type:datetime;not null"`
	AreaTotalFloat         float64   `gorm:"type:decimal(65,20);not null"`
	AreaTotalFloatTwo      float64   `gorm:"type:decimal(65,20);not null"`
	AreaTotalFloatThree    float64   `gorm:"type:decimal(65,20);not null"`
	RecommendTotalFloat    float64   `gorm:"type:decimal(65,20);not null"`
	RecommendLevelFloat    float64   `gorm:"type:decimal(65,20);not null"`
	RecommendTotalFloatTwo float64   `gorm:"type:decimal(65,20);not null"`
	AllFloat               float64   `gorm:"type:decimal(65,20);not null"`
	LocationTotalFloat     float64   `gorm:"type:decimal(65,20);not null"`
}

type UserRecommendArea struct {
	ID            int64     `gorm:"primarykey;type:int"`
	RecommendCode string    `gorm:"type:varchar(10000);not null"`
	Version       int64     `gorm:"type:int;not null"`
	Num           int64     `gorm:"type:int;not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type Withdraw struct {
	ID              int64     `gorm:"primarykey;type:int"`
	UserId          int64     `gorm:"type:int"`
	Amount          int64     `gorm:"type:bigint"`
	RelAmount       int64     `gorm:"type:bigint"`
	Status          string    `gorm:"type:varchar(45);not null"`
	Type            string    `gorm:"type:varchar(45);not null"`
	Address         string    `gorm:"type:varchar(45);not null"`
	BalanceRecordId int64     `gorm:"type:int"`
	CreatedAt       time.Time `gorm:"type:datetime;not null"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null"`
	AmountNew       float64   `gorm:"type:decimal(65,20);not null"`
	RelAmountNew    float64   `gorm:"type:decimal(65,20);not null"`
}

type Trade struct {
	ID           int64     `gorm:"primarykey;type:int"`
	UserId       int64     `gorm:"type:int"`
	AmountCsd    int64     `gorm:"type:bigint"`
	RelAmountCsd int64     `gorm:"type:bigint"`
	AmountHbs    int64     `gorm:"type:bigint"`
	RelAmountHbs int64     `gorm:"type:bigint"`
	CsdReward    int64     `gorm:"type:bigint"`
	Status       string    `gorm:"type:varchar(45);not null"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecord struct {
	ID           int64     `gorm:"primarykey;type:int"`
	UserId       int64     `gorm:"type:int"`
	Balance      int64     `gorm:"type:bigint"`
	Amount       int64     `gorm:"type:bigint"`
	Type         string    `gorm:"type:varchar(45);not null"`
	CoinType     string    `gorm:"type:varchar(45);not null"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
	BalanceNew   float64   `gorm:"type:decimal(65,20);not null"`
	AmountNew    float64   `gorm:"type:decimal(65,20);not null"`
	AmountNewTwo float64   `gorm:"type:decimal(65,20);not null"`
}

type Reward struct {
	ID               int64     `gorm:"primarykey;type:int"`
	UserId           int64     `gorm:"type:int;not null"`
	Amount           int64     `gorm:"type:bigint;not null"`
	AmountB          int64     `gorm:"type:bigint;not null"`
	BalanceRecordId  int64     `gorm:"type:int;not null"`
	Type             string    `gorm:"type:varchar(45);not null"`
	Address          string    `gorm:"type:varchar(45);not null"`
	TypeRecordId     int64     `gorm:"type:int;not null"`
	Reason           string    `gorm:"type:varchar(45);not null"`
	AmountNew        float64   `gorm:"type:decimal(65,20);not null"`
	AmountNewTwo     float64   `gorm:"type:decimal(65,20);not null"`
	ReasonLocationId int64     `gorm:"type:int;not null"`
	Status           int64     `gorm:"type:int;not null"`
	LocationType     string    `gorm:"type:varchar(45);not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type Total struct {
	ID        int64     `gorm:"primarykey;type:int"`
	One       float64   `gorm:"type:decimal(65,20);not null"`
	Two       float64   `gorm:"type:decimal(65,20);not null"`
	Three     float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Admin struct {
	ID       int64  `gorm:"primarykey;type:int"`
	Account  string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Type     string `gorm:"type:varchar(40);not null"`
}

type Auth struct {
	ID   int64  `gorm:"primarykey;type:int"`
	Name string `gorm:"type:varchar(100);not null"`
	Path string `gorm:"type:varchar(200);not null"`
	Url  string `gorm:"type:varchar(200);not null"`
}

type BalanceReward struct {
	ID             int64     `gorm:"primarykey;type:int"`
	UserId         int64     `gorm:"type:int;not null"`
	Amount         int64     `gorm:"type:bigint;not null"`
	Status         int64     `gorm:"type:int;not null"`
	SetDate        time.Time `gorm:"type:datetime;not null"`
	LastRewardDate time.Time `gorm:"type:datetime;not null"`
	CreatedAt      time.Time `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null"`
}

type BuyRecord struct {
	ID          int64     `gorm:"primarykey;type:int"`
	UserId      int64     `gorm:"type:int;not null"`
	Status      int64     `gorm:"type:int;not null"`
	LastUpdated int64     `gorm:"type:int;not null"`
	Amount      float64   `gorm:"type:decimal(65,20);not null"`
	AmountGet   float64   `gorm:"type:decimal(65,20);not null"`
	CreatedAt   time.Time `gorm:"type:datetime;not null"`
	UpdatedAt   time.Time `gorm:"type:datetime;not null"`
}

type AdminAuth struct {
	ID      int64 `gorm:"primarykey;type:int"`
	AdminId int64 `gorm:"type:int"`
	AuthId  int64 `gorm:"type:int"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

type ConfigRepo struct {
	data *Data
	log  *log.Helper
}

type UserInfoRepo struct {
	data *Data
	log  *log.Helper
}

type UserRecommendRepo struct {
	data *Data
	log  *log.Helper
}

type UserCurrentMonthRecommendRepo struct {
	data *Data
	log  *log.Helper
}

type UserBalanceRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserInfoRepo(data *Data, logger log.Logger) biz.UserInfoRepo {
	return &UserInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewConfigRepo(data *Data, logger log.Logger) biz.ConfigRepo {
	return &ConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserBalanceRepo(data *Data, logger log.Logger) biz.UserBalanceRepo {
	return &UserBalanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserRecommendRepo(data *Data, logger log.Logger) biz.UserRecommendRepo {
	return &UserRecommendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserCurrentMonthRecommendRepo(data *Data, logger log.Logger) biz.UserCurrentMonthRecommendRepo {
	return &UserCurrentMonthRecommendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetUserByAddress .
func (u *UserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where(&User{Address: address}).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:      user.ID,
		Address: user.Address,
	}, nil
}

// GetUserByAddressTwo .
func (u *UserRepo) GetUserByAddressTwo(ctx context.Context, address string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where("address=?", address).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:      user.ID,
		Address: user.Address,
		Amount:  user.Amount,
	}, nil
}

// GetConfigByKeys .
func (c *ConfigRepo) GetConfigByKeys(ctx context.Context, keys ...string) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := c.data.db.Where("key_name IN (?)", keys).Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("CONFIG_NOT_FOUND", "config not found")
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// GetConfigs .
func (c *ConfigRepo) GetConfigs(ctx context.Context) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := c.data.db.Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("CONFIG_NOT_FOUND", "config not found")
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// UpdateConfig .
func (c *ConfigRepo) UpdateConfig(ctx context.Context, id int64, value string) (bool, error) {
	var config Config
	config.Value = value

	res := c.data.DB(ctx).Table("config").Where("id=?", id).Updates(&config)
	if res.Error != nil {
		return false, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return true, nil
}

// GetUserById .
func (u *UserRepo) GetUserById(ctx context.Context, Id int64) (*biz.User, error) {
	var user User
	if err := u.data.db.Where(&User{ID: Id}).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:             user.ID,
		Address:        user.Address,
		RecommendLevel: user.RecommendLevel,
		OutRate:        user.OutRate,
		AmountFourGet:  user.AmountFourGet,
		AmountFour:     user.AmountFour,
	}, nil
}

// GetAdminByAccount .
func (u *UserRepo) GetAdminByAccount(ctx context.Context, account string, password string) (*biz.Admin, error) {
	var admin Admin
	if err := u.data.db.Where("account=? and password=?", account, password).Table("admin").First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ADMIN_NOT_FOUND", "admin not found")
		}

		return nil, errors.New(500, "ADMIN ERROR", err.Error())
	}

	return &biz.Admin{
		ID:       admin.ID,
		Password: admin.Password,
		Account:  admin.Account,
		Type:     admin.Type,
	}, nil
}

// GetAdminById .
func (u *UserRepo) GetAdminById(ctx context.Context, id int64) (*biz.Admin, error) {
	var admin Admin
	if err := u.data.db.Where("id=?", id).Table("admin").First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ADMIN_NOT_FOUND", "admin not found")
		}

		return nil, errors.New(500, "ADMIN ERROR", err.Error())
	}

	return &biz.Admin{
		ID:       admin.ID,
		Password: admin.Password,
		Account:  admin.Account,
		Type:     admin.Type,
	}, nil
}

// GetUserInfoByUserId .
func (ui *UserInfoRepo) GetUserInfoByUserId(ctx context.Context, userId int64) (*biz.UserInfo, error) {
	var userInfo UserInfo
	if err := ui.data.db.Where(&UserInfo{UserId: userId}).Table("user_info").First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USERINFO_NOT_FOUND", "userinfo not found")
		}

		return nil, errors.New(500, "USERINFO ERROR", err.Error())
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              userInfo.Vip,
		HistoryRecommend: userInfo.HistoryRecommend,
		UseVip:           userInfo.UseVip,
	}, nil
}

// GetUserByAddresses .
func (u *UserRepo) GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Where("address IN (?)", Addresses).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make(map[string]*biz.User, 0)
	for _, item := range users {
		res[item.Address] = &biz.User{
			ID:      item.ID,
			Address: item.Address,
		}
	}
	return res, nil
}

// GetUserCount .
func (u *UserRepo) GetUserCount(ctx context.Context) (int64, error) {
	var count int64
	if err := u.data.db.Table("user").Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return count, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return count, errors.New(500, "USER ERROR", err.Error())
	}

	return count, nil
}

// GetUserCountToday .
func (u *UserRepo) GetUserCountToday(ctx context.Context) (int64, error) {
	var count int64
	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := u.data.db.Table("user").
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return count, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return count, errors.New(500, "USER ERROR", err.Error())
	}

	return count, nil
}

// GetUserByUserIds .
func (u *UserRepo) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make(map[int64]*biz.User, 0)
	for _, item := range users {
		res[item.ID] = &biz.User{
			ID:               item.ID,
			Address:          item.Address,
			AddressTwo:       item.AddressTwo,
			AddressThree:     item.AddressThree,
			OutRate:          item.OutRate,
			AmountUsdtGet:    item.AmountUsdtGet,
			AmountUsdt:       item.AmountUsdt,
			AmountUsdtOrigin: item.AmountUsdtOrigin,
			MyTotalAmount:    item.MyTotalAmount,
		}
	}
	return res, nil
}

// GetAllUserBalance .
func (u *UserRepo) GetAllUserBalance(ctx context.Context) ([]*biz.UserBalance, error) {
	var userBalances []*UserBalance
	if err := u.data.db.Table("user_balance").Order("id asc").Find(&userBalances).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make([]*biz.UserBalance, 0)
	for _, userBalance := range userBalances {
		res = append(res, &biz.UserBalance{
			ID:                     userBalance.ID,
			UserId:                 userBalance.UserId,
			BalanceRawFloat:        userBalance.BalanceRawFloat,
			BalanceUsdtFloat:       userBalance.BalanceUsdtFloat,
			BalanceKsdtFloat:       userBalance.BalanceKsdtFloat,
			AreaTotalFloat:         userBalance.AreaTotalFloat,
			AreaTotalFloatTwo:      userBalance.AreaTotalFloatTwo,
			AreaTotalFloatThree:    userBalance.AreaTotalFloatThree,
			RecommendTotalFloat:    userBalance.RecommendTotalFloat,
			RecommendTotalFloatTwo: userBalance.RecommendTotalFloatTwo,
			LocationTotalFloat:     userBalance.LocationTotalFloat,
			RecommendLevelFloat:    userBalance.RecommendLevelFloat,
			AllFloat:               userBalance.AllFloat,
		})
	}

	return res, nil
}

// UpdateUserVip .
func (u *UserRepo) UpdateUserVip(ctx context.Context, userId int64, vip int64) error {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"vip": vip})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// GetUserByUserIdsTwo .
func (u *UserRepo) GetUserByUserIdsTwo(ctx context.Context, userIds []int64) (map[int64]*biz.User, error) {
	var users []*User

	res := make(map[int64]*biz.User, 0)
	if err := u.data.DB(ctx).Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, item := range users {
		res[item.ID] = &biz.User{
			ID:                     item.ID,
			Address:                item.Address,
			AmountUsdtGet:          item.AmountUsdtGet,
			AmountUsdt:             item.AmountUsdt,
			AmountRecommendUsdtGet: item.AmountRecommendUsdtGet,
			MyTotalAmount:          item.MyTotalAmount,
			Vip:                    item.Vip,
		}
	}
	return res, nil
}

// GetBuyRecord .
func (u *UserRepo) GetBuyRecord(ctx context.Context, day int) ([]*biz.BuyRecord, error) {
	res := make([]*biz.BuyRecord, 0)

	var buyRecord []*BuyRecord
	if err := u.data.db.Table("buy_record").Where("status=?", 1).
		Order("id asc").Find(&buyRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "buy_record ERROR", err.Error())
	}

	for _, v := range buyRecord {
		res = append(res, &biz.BuyRecord{
			ID:          v.ID,
			UserId:      v.UserId,
			Status:      v.Status,
			Amount:      v.Amount,
			AmountGet:   v.AmountGet,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			LastUpdated: v.LastUpdated,
		})
	}

	return res, nil
}

// GetAllUsersOrderAmountBiw .
func (u *UserRepo) GetAllUsersOrderAmountBiw(ctx context.Context) ([]*biz.User, error) {
	var users []*User

	res := make([]*biz.User, 0)
	if err := u.data.db.Table("user").Order("amount_biw desc").Limit(3).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, item := range users {
		res = append(res, &biz.User{
			ID:                     item.ID,
			Address:                item.Address,
			Lock:                   item.Lock,
			AddressTwo:             item.AddressTwo,
			AddressThree:           item.AddressThree,
			Amount:                 item.Amount,
			AmountUsdt:             item.AmountUsdt,
			AmountUsdtGet:          item.AmountUsdtGet,
			MyTotalAmount:          item.MyTotalAmount,
			AmountRecommendUsdtGet: item.AmountRecommendUsdtGet,
			Last:                   item.Last,
			Vip:                    item.Vip,
			VipAdmin:               item.VipAdmin,
			OutRate:                item.OutRate,
			UpdatedAt:              item.UpdatedAt,
			LockReward:             item.LockReward,
		})
	}
	return res, nil
}

// GetAllUsersRecommendOrder .
func (u *UserRepo) GetAllUsersRecommendOrder(ctx context.Context) ([]*biz.User, error) {
	var users []*User
	res := make([]*biz.User, 0)
	if err := u.data.db.Table("user").Order("amount_recommend_usdt_get desc").Limit(3).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	for _, item := range users {
		res = append(res, &biz.User{
			ID:                     item.ID,
			Address:                item.Address,
			Lock:                   item.Lock,
			AddressTwo:             item.AddressTwo,
			AddressThree:           item.AddressThree,
			Amount:                 item.Amount,
			AmountUsdt:             item.AmountUsdt,
			AmountUsdtGet:          item.AmountUsdtGet,
			MyTotalAmount:          item.MyTotalAmount,
			AmountRecommendUsdtGet: item.AmountRecommendUsdtGet,
			Last:                   item.Last,
			Vip:                    item.Vip,
			VipAdmin:               item.VipAdmin,
			OutRate:                item.OutRate,
			UpdatedAt:              item.UpdatedAt,
			LockReward:             item.LockReward,
		})
	}
	return res, nil
}

// GetAllUsersB .
func (ub *UserBalanceRepo) GetAllUsersB(ctx context.Context) ([]*biz.User, error) {
	var users []*User
	if err := ub.data.db.Table("user").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, item := range users {
		res = append(res, &biz.User{
			ID:                     item.ID,
			Address:                item.Address,
			Lock:                   item.Lock,
			AddressTwo:             item.AddressTwo,
			AddressThree:           item.AddressThree,
			Amount:                 item.Amount,
			AmountUsdt:             item.AmountUsdt,
			AmountUsdtGet:          item.AmountUsdtGet,
			MyTotalAmount:          item.MyTotalAmount,
			AmountRecommendUsdtGet: item.AmountRecommendUsdtGet,
			Last:                   item.Last,
			LastBiw:                item.LastBiw,
			Vip:                    item.Vip,
			VipAdmin:               item.VipAdmin,
			OutRate:                item.OutRate,
			UpdatedAt:              item.UpdatedAt,
			LockReward:             item.LockReward,
			CreatedAt:              item.CreatedAt,
			RecommendUserH:         item.RecommendUserH,
			AmountFourGet:          item.AmountFourGet,
			AmountFour:             item.AmountFour,
		})
	}
	return res, nil
}

// GetAllUsers .
func (u *UserRepo) GetAllUsers(ctx context.Context) ([]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Order("id asc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, item := range users {
		res = append(res, &biz.User{
			ID:                     item.ID,
			Address:                item.Address,
			Lock:                   item.Lock,
			AddressTwo:             item.AddressTwo,
			AddressThree:           item.AddressThree,
			Amount:                 item.Amount,
			AmountUsdt:             item.AmountUsdt,
			AmountUsdtGet:          item.AmountUsdtGet,
			MyTotalAmount:          item.MyTotalAmount,
			AmountRecommendUsdtGet: item.AmountRecommendUsdtGet,
			Last:                   item.Last,
			LastBiw:                item.LastBiw,
			Vip:                    item.Vip,
			VipAdmin:               item.VipAdmin,
			OutRate:                item.OutRate,
			UpdatedAt:              item.UpdatedAt,
			LockReward:             item.LockReward,
			CreatedAt:              item.CreatedAt,
			RecommendUserH:         item.RecommendUserH,
			AmountFourGet:          item.AmountFourGet,
			AmountFour:             item.AmountFour,
		})
	}
	return res, nil
}

// GetAllUserInfos .
func (u *UserRepo) GetAllUserInfos(ctx context.Context) ([]*biz.UserInfo, error) {
	var users []*UserInfo
	if err := u.data.db.Table("user_info").Where("use_vip=?", 1).Order("id desc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make([]*biz.UserInfo, 0)
	for _, item := range users {
		res = append(res, &biz.UserInfo{
			ID:               item.ID,
			UserId:           item.UserId,
			Vip:              item.Vip,
			HistoryRecommend: item.HistoryRecommend,
			TeamCsdBalance:   item.TeamCsdBalance,
			LockVip:          item.LockVip,
		})
	}
	return res, nil
}

// GetAdmins .
func (u *UserRepo) GetAdmins(ctx context.Context) ([]*biz.Admin, error) {
	var admins []*Admin
	if err := u.data.db.Table("admin").Find(&admins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("AdMIN_NOT_FOUND", "admin not found")
		}

		return nil, errors.New(500, "ADMIN ERROR", err.Error())
	}

	res := make([]*biz.Admin, 0)
	for _, item := range admins {
		res = append(res, &biz.Admin{
			ID:      item.ID,
			Account: item.Account,
		})
	}
	return res, nil
}

// GetAuths .
func (u *UserRepo) GetAuths(ctx context.Context) ([]*biz.Auth, error) {
	var auths []*Auth
	if err := u.data.db.Table("auth").Find(&auths).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("AUTH_NOT_FOUND", "auth not found")
		}

		return nil, errors.New(500, "AUTH ERROR", err.Error())
	}

	res := make([]*biz.Auth, 0)
	for _, item := range auths {
		res = append(res, &biz.Auth{
			ID:   item.ID,
			Name: item.Name,
			Path: item.Path,
		})
	}
	return res, nil
}

// GetAuthByIds .
func (u *UserRepo) GetAuthByIds(ctx context.Context, ids ...int64) (map[int64]*biz.Auth, error) {
	var auths []*Auth
	if err := u.data.db.Table("auth").Where("id IN (?)", ids).Find(&auths).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("AUTH_NOT_FOUND", "auth not found")
		}

		return nil, errors.New(500, "AUTH ERROR", err.Error())
	}

	res := make(map[int64]*biz.Auth, 0)
	for _, item := range auths {
		res[item.ID] = &biz.Auth{
			ID:   item.ID,
			Name: item.Name,
			Path: item.Path,
		}
	}
	return res, nil
}

// GetAdminAuth .
func (u *UserRepo) GetAdminAuth(ctx context.Context, adminId int64) ([]*biz.AdminAuth, error) {
	var auths []*AdminAuth
	if err := u.data.db.Table("admin_auth").Where("admin_id=?", adminId).Find(&auths).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("AUTH_NOT_FOUND", "auth not found")
		}

		return nil, errors.New(500, "AUTH ERROR", err.Error())
	}

	res := make([]*biz.AdminAuth, 0)
	for _, item := range auths {
		res = append(res, &biz.AdminAuth{
			ID:      item.ID,
			AdminId: item.AdminId,
			AuthId:  item.AuthId,
		})
	}
	return res, nil
}

// GetUsers .
func (u *UserRepo) GetUsers(ctx context.Context, b *biz.Pagination, address string, isLocation bool, vip int64) ([]*biz.User, error, int64) {
	var (
		users []*User
		count int64
	)
	instance := u.data.db.Table("user")
	if "" != address {
		instance = instance.Where("address=?", address)
	}

	if isLocation {
		instance = instance.Joins("inner join location_new on user.id = location_new.user_id").Group("user.id")
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found"), 0
		}

		return nil, errors.New(500, "USER ERROR", err.Error()), 0
	}

	res := make([]*biz.User, 0)
	for _, item := range users {
		res = append(res, &biz.User{
			ID:               item.ID,
			Address:          item.Address,
			Password:         item.Password,
			CreatedAt:        item.CreatedAt,
			Amount:           item.Amount,
			AmountUsdt:       item.AmountUsdt,
			AmountBiw:        item.AmountBiw,
			OutRate:          item.OutRate,
			RecommendLevel:   item.RecommendLevel,
			Lock:             item.Lock,
			AmountUsdtOrigin: item.AmountUsdtOrigin,
			AmountUsdtGet:    item.AmountUsdtGet,
			MyTotalAmount:    item.MyTotalAmount,
			Vip:              item.Vip,
			VipAdmin:         item.VipAdmin,
			LockReward:       item.LockReward,
			AmountFourGet:    item.AmountFourGet,
			AmountFour:       item.AmountFour,
		})
	}
	return res, nil, count
}

// CreateUser .
func (u *UserRepo) CreateUser(ctx context.Context, uc *biz.User) (*biz.User, error) {
	var user User
	user.Address = uc.Address
	res := u.data.DB(ctx).Table("user").Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "用户创建失败")
	}

	return &biz.User{
		ID:      user.ID,
		Address: user.Address,
	}, nil
}

// CreateAdmin .
func (u *UserRepo) CreateAdmin(ctx context.Context, a *biz.Admin) (*biz.Admin, error) {
	var admin Admin
	admin.Account = a.Account
	admin.Password = a.Password
	res := u.data.DB(ctx).Table("admin").Create(&admin)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ADMIN_ERROR", "用户创建失败")
	}

	return &biz.Admin{
		ID:       admin.ID,
		Password: admin.Password,
		Account:  admin.Account,
		Type:     admin.Type,
	}, nil
}

// CreateAdminAuth .
func (u *UserRepo) CreateAdminAuth(ctx context.Context, adminId int64, authId int64) (bool, error) {
	var adminAuth AdminAuth
	adminAuth.AdminId = adminId
	adminAuth.AuthId = authId
	res := u.data.DB(ctx).Table("admin_auth").Create(&adminAuth)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_ADMIN_ERROR", "记录创建失败")
	}

	return true, nil
}

// DeleteAdminAuth .
func (u *UserRepo) DeleteAdminAuth(ctx context.Context, adminId int64, authId int64) (bool, error) {
	var adminAuth AdminAuth
	adminAuth.AdminId = adminId
	adminAuth.AuthId = authId
	res := u.data.DB(ctx).Table("admin_auth").Where("admin_id=? and auth_id=?", adminId, authId).Delete(&adminAuth)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_ADMIN_ERROR", "记录删除失败")
	}

	return true, nil
}

// CreateUserInfo .
func (ui *UserInfoRepo) CreateUserInfo(ctx context.Context, u *biz.User) (*biz.UserInfo, error) {
	var userInfo UserInfo
	userInfo.UserId = u.ID

	res := ui.data.DB(ctx).Table("user_info").Create(&userInfo)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_INFO_ERROR", "用户信息创建失败")
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              0,
		HistoryRecommend: 0,
	}, nil
}

// UpdateUserPassword .
func (ui *UserInfoRepo) UpdateUserPassword(ctx context.Context, userId int64, password string) (*biz.User, error) {
	var user User
	user.Password = password
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).Updates(&user)
	if res.Error != nil {
		return nil, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return nil, nil
}

// UpdateUserInfo .
func (ui *UserInfoRepo) UpdateUserInfo(ctx context.Context, u *biz.UserInfo) (*biz.UserInfo, error) {
	var userInfo UserInfo
	userInfo.Vip = u.Vip
	userInfo.HistoryRecommend = u.HistoryRecommend

	res := ui.data.DB(ctx).Table("user_info").Where("user_id=?", u.UserId).Updates(&userInfo)
	if res.Error != nil {
		return nil, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              userInfo.Vip,
		HistoryRecommend: userInfo.HistoryRecommend,
	}, nil
}

// UpdateUserInfo2 .
func (ui *UserInfoRepo) UpdateUserInfo2(ctx context.Context, u *biz.UserInfo) (*biz.UserInfo, error) {
	var (
		userInfo UserInfo
		lockVip  = 1
		vip      = u.Vip
	)

	if -1 == vip {
		lockVip = 0
		vip = 0
	}

	res := ui.data.DB(ctx).Table("user_info").Where("user_id=?", u.UserId).Updates(
		map[string]interface{}{"vip": vip, "lock_vip": lockVip})

	if res.Error != nil {
		return nil, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              userInfo.Vip,
		HistoryRecommend: userInfo.HistoryRecommend,
	}, nil
}

// UpdateUserInfoVip .
func (ui *UserInfoRepo) UpdateUserInfoVip(ctx context.Context, userId, vip int64) (*biz.UserInfo, error) {
	var userInfo UserInfo
	res := ui.data.DB(ctx).Table("user_info").Where("user_id=?", userId).Updates(
		map[string]interface{}{"vip": vip, "use_vip": 1})
	if res.Error != nil {
		return nil, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              userInfo.Vip,
		HistoryRecommend: userInfo.HistoryRecommend,
	}, nil
}

// UpdateBalance .
func (ub *UserBalanceRepo) UpdateBalance(ctx context.Context, userId int64, amount int64) (bool, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": amount}).Error; nil != err {
		return false, errors.NotFound("user balance err", "user balance not found")
	}

	return true, nil
}

// CreateUserArea .
func (ur *UserRecommendRepo) CreateUserArea(ctx context.Context, u *biz.User) (bool, error) {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var userArea UserArea
	userArea.UserId = u.ID
	res := ur.data.DB(ctx).Table("user_area").Create(&userArea)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AREA_ERROR", "用户区信息创建失败")
	}

	return true, nil
}

// UpdateAdminPassword .
func (u *UserRepo) UpdateAdminPassword(ctx context.Context, account string, password string) (*biz.Admin, error) {
	var admin Admin
	admin.Password = password
	res := u.data.DB(ctx).Table("admin").Where("account=?", account).Updates(&admin)
	if res.Error != nil {
		return nil, errors.New(500, "UPDATE_ADMIN_ERROR", "用户信息修改失败")
	}

	return &biz.Admin{
		ID:       admin.ID,
		Password: admin.Password,
		Account:  admin.Account,
		Type:     admin.Type,
	}, nil
}

// GetUserRecommendByUserId .
func (ur *UserRecommendRepo) GetUserRecommendByUserId(ctx context.Context, userId int64) (*biz.UserRecommend, error) {
	var userRecommend UserRecommend
	if err := ur.data.db.Where(&UserRecommend{UserId: userId}).Table("user_recommend").First(&userRecommend).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return &biz.UserRecommend{
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
	}, nil
}

// GetUserRecommendByCode .
func (ur *UserRecommendRepo) GetUserRecommendByCode(ctx context.Context, code string) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)
	res := make([]*biz.UserRecommend, 0)

	instance := ur.data.db.Table("user_recommend").Where("recommend_code=?", code)

	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
		})
	}

	return res, nil
}

// GetUserRecommendLikeCode .
func (ur *UserRecommendRepo) GetUserRecommendLikeCode(ctx context.Context, code string) ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := ur.data.db.Where("recommend_code Like ?", code+"%").Table("user_recommend").Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
		})
	}

	return res, nil
}

// GetUserRecommends .
func (ur *UserRecommendRepo) GetUserRecommends(ctx context.Context) ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := ur.data.db.Table("user_recommend").Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			ID:            userRecommend.ID,
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			Total:         userRecommend.Total,
		})
	}

	return res, nil
}

// GetUserRecommendsFour .
func (ur *UserRecommendRepo) GetUserRecommendsFour(ctx context.Context) ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := ur.data.db.Table("user_recommend").Order("total desc").Limit(4).Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			Total:         userRecommend.Total,
		})
	}

	return res, nil
}

// GetUserRecommendLowAreas .
func (ur *UserRecommendRepo) GetUserRecommendLowAreas(ctx context.Context) ([]*biz.UserRecommendArea, error) {
	var firstRecommendArea *UserRecommendArea
	if err := ur.data.db.Order("num desc").Table("user_recommend_area").First(&firstRecommendArea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER RECOMMEND NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	var userRecommendAreas []*UserRecommendArea
	res := make([]*biz.UserRecommendArea, 0)
	if err := ur.data.db.Table("user_recommend_area").Find(&userRecommendAreas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommendArea := range userRecommendAreas {
		if firstRecommendArea.ID == userRecommendArea.ID {
			continue
		}
		res = append(res, &biz.UserRecommendArea{
			RecommendCode: userRecommendArea.RecommendCode,
			Num:           userRecommendArea.Num,
		})
	}

	return res, nil
}

// CreateUserRecommend .
func (ur *UserRecommendRepo) CreateUserRecommend(ctx context.Context, u *biz.User, recommendUser *biz.UserRecommend) (*biz.UserRecommend, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatInt(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var userRecommend UserRecommend
	userRecommend.UserId = u.ID
	userRecommend.RecommendCode = tmpRecommendCode

	res := ur.data.DB(ctx).Table("user_recommend").Create(&userRecommend)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_RECOMMEND_ERROR", "用户推荐关系创建失败")
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
	}, nil
}

// CreateUserRecommendArea .
func (ur *UserRecommendRepo) CreateUserRecommendArea(ctx context.Context, recommendAreas []*biz.UserRecommendArea) (bool, error) {

	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var userRecommendArea []*UserRecommendArea
	for _, v := range recommendAreas {
		userRecommendArea = append(userRecommendArea, &UserRecommendArea{
			RecommendCode: v.RecommendCode,
			Num:           v.Num,
		})
	}
	res := ur.data.DB(ctx).Table("user_recommend_area").Create(&userRecommendArea)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_RECOMMEND_AREA_ERROR", "用户推荐关系链路创建失败")
	}

	return true, nil
}

// UpdateUserAreaAmount .
func (ur *UserRecommendRepo) UpdateUserAreaAmount(ctx context.Context, userId int64, amount int64) (bool, error) {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var err error
	if err = ur.data.DB(ctx).Table("user_area").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"amount": gorm.Expr("amount + ?", amount)}).Error; nil != err {
		return false, errors.NotFound("user balance err", "user area not found")
	}

	return true, nil
}

// UpdateUserAreaSelfAmount .
func (ur *UserRecommendRepo) UpdateUserAreaSelfAmount(ctx context.Context, userId int64, amount int64) (bool, error) {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var err error
	if err = ur.data.DB(ctx).Table("user_area").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"self_amount": gorm.Expr("self_amount + ?", amount)}).Error; nil != err {
		return false, errors.NotFound("user balance err", "user area not found")
	}

	return true, nil
}

// UpdateUserAreaLevel .
func (ur *UserRecommendRepo) UpdateUserAreaLevel(ctx context.Context, userId int64, level int64) (bool, error) {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var err error
	if err = ur.data.DB(ctx).Table("user_area").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"level": level}).Error; nil != err {
		return false, errors.NotFound("user balance err", "user area not found")
	}

	return true, nil
}

// UpdateUserRecommendTotal .
func (ur *UserRecommendRepo) UpdateUserRecommendTotal(ctx context.Context, userId int64, total int64) error {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var err error
	if err = ur.data.DB(ctx).Table("user_recommend").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"total": gorm.Expr("total + ?", total)}).Error; nil != err {
		return err
	}

	return nil
}

// UpdateUserAreaLevelUp .
func (ur *UserRecommendRepo) UpdateUserAreaLevelUp(ctx context.Context, userId int64, level int64) (bool, error) {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var err error
	if err = ur.data.DB(ctx).Table("user_area").
		Where("user_id=?", userId).
		Where("level<?", level).
		Updates(map[string]interface{}{"level": level}).Error; nil != err {
		return false, errors.NotFound("user balance err", "user area not found")
	}

	return true, nil
}

// UndoUser .
func (u *UserRepo) UndoUser(ctx context.Context, userId int64, undo int64) (bool, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Updates(map[string]interface{}{"undo": undo})
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_ERROR", "用户修改失败")
	}

	return true, nil
}

// LockUser .
func (u *UserRepo) LockUser(ctx context.Context, userId int64, lock int64) (bool, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Updates(map[string]interface{}{"lock": lock})
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_ERROR", "用户修改失败")
	}

	return true, nil
}

func (u *UserRepo) LockUserReward(ctx context.Context, userId int64, lock int64) (bool, error) {
	res := u.data.DB(ctx).Table("user").Where("id=?", userId).Updates(map[string]interface{}{"lock_reward": lock})
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_ERROR", "用户修改失败")
	}

	return true, nil
}

// GetUserAreas .
func (ur *UserRecommendRepo) GetUserAreas(ctx context.Context, userIds []int64) ([]*biz.UserArea, error) {

	var userAreas []*UserArea
	if err := ur.data.DB(ctx).Where("user_id in (?)", userIds).Table("user_area").Find(&userAreas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER AREA NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER AREA ERROR", err.Error())
	}

	res := make([]*biz.UserArea, 0)
	for _, v := range userAreas {
		res = append(res, &biz.UserArea{
			ID:         v.ID,
			UserId:     v.UserId,
			Amount:     v.Amount,
			SelfAmount: v.SelfAmount,
			Level:      v.Level,
		})
	}

	return res, nil
}

// GetUserArea .
func (ur *UserRecommendRepo) GetUserArea(ctx context.Context, userId int64) (*biz.UserArea, error) {

	var userArea *UserArea
	if err := ur.data.db.Where("user_id=?", userId).Table("user_area").First(&userArea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER AREA NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER AREA ERROR", err.Error())
	}

	return &biz.UserArea{
		ID:         userArea.ID,
		UserId:     userArea.UserId,
		Amount:     userArea.Amount,
		SelfAmount: userArea.SelfAmount,
		Level:      userArea.Level,
	}, nil
}

// CreateUserBalance .
func (ub UserBalanceRepo) CreateUserBalance(ctx context.Context, u *biz.User) (*biz.UserBalance, error) {
	var userBalance UserBalance
	userBalance.UserId = u.ID
	res := ub.data.DB(ctx).Table("user_balance").Create(&userBalance)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BALANCE_ERROR", "用户余额信息创建失败")
	}

	return &biz.UserBalance{
		ID:          userBalance.ID,
		UserId:      userBalance.UserId,
		BalanceUsdt: userBalance.BalanceUsdt,
		BalanceDhb:  userBalance.BalanceDhb,
	}, nil
}

// GetUserBalance .
func (ub UserBalanceRepo) GetUserBalance(ctx context.Context, userId int64) (*biz.UserBalance, error) {
	var userBalance UserBalance
	if err := ub.data.DB(ctx).Where("user_id=?", userId).Table("user_balance").First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return &biz.UserBalance{
		ID:               userBalance.ID,
		UserId:           userBalance.UserId,
		BalanceUsdt:      userBalance.BalanceUsdt,
		BalanceDhb:       userBalance.BalanceDhb,
		BalanceUsdtFloat: userBalance.BalanceUsdtFloat,
	}, nil
}

// LocationReward .
func (ub *UserBalanceRepo) LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "location" // 给我分红的理由
	reward.ReasonLocationId = myLocationId
	reward.LocationType = locationType
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// WithdrawReward .
func (ub *UserBalanceRepo) WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}

	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "location" // 给我分红的理由
	reward.ReasonLocationId = myLocationId
	reward.LocationType = locationType
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// Deposit .
func (ub *UserBalanceRepo) Deposit(ctx context.Context, userId int64, amount int64, dhbAmount int64) (int64, error) {
	var err error
	//if err = ub.data.DB(ctx).Table("user_balance").
	//	Where("user_id=?", userId).
	//	Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", dhbAmount)}).Error; nil != err {
	//	return 0, errors.NotFound("user balance err", "user balance not found")
	//}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateLocationAgain .
func (ub *UserBalanceRepo) UpdateLocationAgain(ctx context.Context, locations []*biz.LocationNew) error {
	var (
		err error
	)
	for _, vLocations := range locations {
		res := ub.data.DB(ctx).Table("location_new").
			Where("id=?", vLocations.ID).
			Updates(map[string]interface{}{"stop_location_again": "1"})
		if 0 == res.RowsAffected || res.Error != nil {
			return err
		}
	}

	return nil
}

// UpdateLocationAgain2 .
func (ub *UserBalanceRepo) UpdateLocationAgain2(ctx context.Context, locations []*biz.LocationNew) error {
	var (
		err error
	)
	for _, vLocations := range locations {
		res := ub.data.DB(ctx).Table("location_new_2").
			Where("id=?", vLocations.ID).
			Updates(map[string]interface{}{"stop_location_again": "1"})
		if 0 == res.RowsAffected || res.Error != nil {
			return err
		}
	}

	return nil
}

// DepositLastNew2 .
func (ub *UserBalanceRepo) DepositLastNew2(ctx context.Context, userId int64, lastAmount int64) (int64, error) {
	var (
		err error
	)
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_new": gorm.Expr("balance_usdt_new + ?", lastAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward_new_2"
	userBalanceRecode.Amount = lastAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}
	//
	//var reward Reward
	//reward.UserId = userBalance.UserId
	//reward.Amount = lastAmount
	//reward.BalanceRecordId = userBalanceRecode.ID
	//reward.Type = "last"          // 本次分红的行为类型
	//reward.Reason = "last_reward" // 给我分红的理由
	//err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	//if err != nil {
	//	return 0, err
	//}

	return userBalanceRecode.ID, nil
}

// DepositLastNew .
func (ub *UserBalanceRepo) DepositLastNew(ctx context.Context, userId int64, lastAmount int64) (int64, error) {
	var (
		err error
	)
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", lastAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward_new"
	userBalanceRecode.Amount = lastAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}
	//
	//var reward Reward
	//reward.UserId = userBalance.UserId
	//reward.Amount = lastAmount
	//reward.BalanceRecordId = userBalanceRecode.ID
	//reward.Type = "last"          // 本次分红的行为类型
	//reward.Reason = "last_reward" // 给我分红的理由
	//err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	//if err != nil {
	//	return 0, err
	//}

	return userBalanceRecode.ID, nil
}

// DepositLastNewDhb .
func (ub *UserBalanceRepo) DepositLastNewDhb(ctx context.Context, userId int64, lastCoinAmount int64) error {
	var (
		err error
	)
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", lastCoinAmount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}
	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = 0
	userBalanceRecode.UserId = userId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.CoinType = "HBS"
	userBalanceRecode.Amount = lastCoinAmount
	res := ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode)
	if res.Error != nil {
		return errors.New(500, "CREATE_LOCATION_ERROR", "占位信息创建失败")
	}

	return nil
}

// DepositLastNewCsd .
func (ub *UserBalanceRepo) DepositLastNewCsd(ctx context.Context, userId int64, lastCoinAmount int64, tmpRecommendUserIdsInt []int64) error {
	var (
		err error
	)
	if len(tmpRecommendUserIdsInt) > 0 {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", lastCoinAmount)}).Error; nil != err {
			return errors.NotFound("user balance err", "user balance not found")
		}
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", lastCoinAmount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = 0
	userBalanceRecode.UserId = userId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.CoinType = "CSD"
	userBalanceRecode.Amount = lastCoinAmount
	res := ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode)
	if res.Error != nil {
		return errors.New(500, "CREATE_LOCATION_ERROR", "占位信息创建失败")
	}

	return nil
}

// UserDailyLocationReward .
func (ub *UserBalanceRepo) UserDailyLocationReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, coinAmount int64, status string, locationId int64) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", coinAmount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}

	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_daily"     // 本次分红的行为类型
	reward.Reason = "location_daily_reward" // 给我分红的理由
	reward.ReasonLocationId = locationId
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// DepositLast .
func (ub *UserBalanceRepo) DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error) {
	var (
		err error
	)
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", lastAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = lastAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = lastAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "last"          // 本次分红的行为类型
	reward.Reason = "last_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	res := ub.data.db.Table("location").
		Where("id=?", locationId).
		Updates(map[string]interface{}{"stop_location_again": "1"})
	if 0 == res.RowsAffected || res.Error != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// DepositDhb .
func (ub *UserBalanceRepo) DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// WithdrawUsdt .
func (ub *UserBalanceRepo) WithdrawUsdt(ctx context.Context, userId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_usdt>=?", userId, amount).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "withdraw"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// WithdrawDhb .
func (ub *UserBalanceRepo) WithdrawDhb(ctx context.Context, userId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_dhb>=?", userId, amount).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "withdraw"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// GreateWithdraw .
func (ub *UserBalanceRepo) GreateWithdraw(ctx context.Context, userId int64, amount int64, coinType string) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.UserId = userId
	withdraw.Amount = amount
	withdraw.Type = coinType
	res := ub.data.DB(ctx).Table("withdraw").Create(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录创建失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// UpdateWithdraw .
func (ub *UserBalanceRepo) UpdateWithdraw(ctx context.Context, id int64, status string) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.Status = status
	res := ub.data.DB(ctx).Table("withdraw").Where("id=?", id).Updates(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// UpdateWithdrawAmount .
func (ub *UserBalanceRepo) UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.Status = status
	withdraw.RelAmount = amount
	res := ub.data.DB(ctx).Table("withdraw").Where("id=?", id).Updates(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// UpdateTrade .
func (ub *UserBalanceRepo) UpdateTrade(ctx context.Context, id int64, status string) (*biz.Trade, error) {
	var trade Trade
	trade.Status = status
	res := ub.data.DB(ctx).Table("trade").Where("id=?", id).Updates(&trade)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Trade{
		ID:           trade.ID,
		UserId:       trade.UserId,
		AmountCsd:    trade.AmountCsd,
		RelAmountCsd: trade.RelAmountCsd,
		AmountHbs:    trade.AmountHbs,
		RelAmountHbs: trade.RelAmountHbs,
		Status:       trade.Status,
		CreatedAt:    trade.CreatedAt,
	}, nil
}

// GetTradeOk .
func (ub *UserBalanceRepo) GetTradeOk(ctx context.Context) (*biz.Trade, error) {
	var trade *Trade
	if err := ub.data.db.Where("status=?", "ok").Table("trade").First(&trade).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	return &biz.Trade{
		ID:           trade.ID,
		UserId:       trade.UserId,
		AmountCsd:    trade.AmountCsd,
		RelAmountCsd: trade.RelAmountCsd,
		AmountHbs:    trade.AmountHbs,
		RelAmountHbs: trade.RelAmountHbs,
		Status:       trade.Status,
		CreatedAt:    trade.CreatedAt,
	}, nil
}

// GetTradeOkkCsd .
func (ub *UserBalanceRepo) GetTradeOkkCsd(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Where("status=?", "okk").Table("trade").Select("sum(amount_csd) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetTradeOkkHbs .
func (ub *UserBalanceRepo) GetTradeOkkHbs(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Where("status=?", "okk").Table("trade").Select("sum(amount_hbs) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetTradeNotDeal .
func (ub *UserBalanceRepo) GetTradeNotDeal(ctx context.Context) ([]*biz.Trade, error) {
	var trades []*Trade
	res := make([]*biz.Trade, 0)
	if err := ub.data.db.Where("status=?", "default").Table("trade").Find(&trades).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, trade := range trades {
		res = append(res, &biz.Trade{
			ID:           trade.ID,
			UserId:       trade.UserId,
			AmountCsd:    trade.AmountCsd,
			RelAmountCsd: trade.RelAmountCsd,
			AmountHbs:    trade.AmountHbs,
			CsdReward:    trade.CsdReward,
			RelAmountHbs: trade.RelAmountHbs,
			Status:       trade.Status,
			CreatedAt:    trade.CreatedAt,
		})
	}
	return res, nil
}

// UpdateWithdrawPass .
func (ub *UserBalanceRepo) UpdateWithdrawPass(ctx context.Context, id int64) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.Status = "pass"
	res := ub.data.DB(ctx).Table("withdraw").Where("id=?", id).Where("status=?", "rewarded").Updates(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// GetWithdrawByUserId .
func (ub *UserBalanceRepo) GetWithdrawByUserId(ctx context.Context, userId int64) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Where("user_id=?", userId).Table("withdraw").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

// GetWithdrawNotDeal .
func (ub *UserBalanceRepo) GetWithdrawNotDeal(ctx context.Context) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Where("status=?", "").Table("withdraw").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

// GetWithdrawByUserIds .
func (ub *UserBalanceRepo) GetWithdrawByUserIds(ctx context.Context, userIds []int64) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Where("status=?", "success").Where("user_id in(?)", userIds).Table("withdraw").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

func (ub *UserBalanceRepo) GetWithdrawById(ctx context.Context, id int64) (*biz.Withdraw, error) {
	var withdraw *Withdraw
	if err := ub.data.db.Where("id=?", id).Table("withdraw").First(&withdraw).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}
	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// GetWithdraws .
func (ub *UserBalanceRepo) GetWithdraws(ctx context.Context, b *biz.Pagination, userId int64, withdrawType string) ([]*biz.Withdraw, error, int64) {
	var (
		withdraws []*Withdraw
		count     int64
	)
	res := make([]*biz.Withdraw, 0)

	instance := ub.data.db.Table("withdraw")

	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	if "" != withdrawType {
		instance = instance.Where("type=?", withdrawType)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found"), 0
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error()), 0
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
			AmountNew:       withdraw.AmountNew,
			RelAmountNew:    withdraw.RelAmountNew,
		})
	}
	return res, nil, count
}

// GetWithdrawPassOrRewarded .
func (ub *UserBalanceRepo) GetWithdrawPassOrRewarded(ctx context.Context) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Table("withdraw").Where("status=? or status=?", "pass", "rewarded").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

// GetWithdrawPassOrRewardedFirst .
func (ub *UserBalanceRepo) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*biz.Withdraw, error) {
	var withdraw *Withdraw
	if err := ub.data.db.Table("withdraw").Where("status=?", "rewarded").First(&withdraw).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
		Address:         withdraw.Address,
		AmountNew:       withdraw.AmountNew,
		RelAmountNew:    withdraw.RelAmountNew,
	}, nil
}

// RecommendReward .
func (ub *UserBalanceRepo) RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_vip" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateBalanceRewardLastRewardDate .
func (ub *UserBalanceRepo) UpdateBalanceRewardLastRewardDate(ctx context.Context, id int64) error {
	if res := ub.data.DB(ctx).Table("balance_reward").
		Where("id=?", id).
		Updates(map[string]interface{}{"last_reward_date": time.Now().UTC()}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	return nil
}

// FourRewardBiw .
func (ub *UserBalanceRepo) FourRewardBiw(ctx context.Context, userId int64, rewardAmount int64, num int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", rewardAmount), "four_total": gorm.Expr("four_total + ?", rewardAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_four_daily" // 本次分红的行为类型
	reward.Reason = "four"                   // 给我分红的理由
	reward.ReasonLocationId = num
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// FourRewardBYes .
func (ub *UserBalanceRepo) FourRewardYes(ctx context.Context, rewardAmount int64) error {
	var (
		err    error
		reward Reward
	)
	reward.UserId = 999999
	reward.Amount = rewardAmount
	reward.Type = "four_yes" // 本次分红的行为类型
	reward.Reason = "four_yes"
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// GetRewardFourYes .
func (ub *UserBalanceRepo) GetRewardFourYes(ctx context.Context) (*biz.Reward, error) {
	var reward *Reward
	if err := ub.data.db.Where("user_id=? and reason=?", 999999, "four_yes").Order("id desc").Table("reward").First(&reward).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}
	return &biz.Reward{
		ID:               reward.ID,
		UserId:           reward.UserId,
		Amount:           reward.Amount,
		BalanceRecordId:  reward.BalanceRecordId,
		Type:             reward.Type,
		TypeRecordId:     reward.TypeRecordId,
		Reason:           reward.Reason,
		ReasonLocationId: reward.ReasonLocationId,
		LocationType:     reward.LocationType,
		CreatedAt:        reward.CreatedAt,
	}, nil
}

// ExchangeBiw .
func (ub *UserBalanceRepo) ExchangeBiw(ctx context.Context, userId int64, currentMaxNew int64, feeRate int64) (int64, error) {
	var userBalance UserBalance
	var err error
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	if 0 >= userBalance.BalanceDhb {
		return 0, nil
	}

	tmp := currentMaxNew
	tmp -= tmp * feeRate / 1000
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_dhb": 0, "balance_usdt": gorm.Expr("balance_usdt + ?", tmp)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "exchange"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = tmp
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, nil
	}

	var (
		reward Reward
	)

	reward.UserId = userId
	reward.Amount = userBalance.BalanceDhb
	reward.AmountB = tmp
	reward.Type = "exchange_system" // 本次分红的行为类型
	reward.TypeRecordId = userBalanceRecode.ID
	reward.Reason = "exchange_2" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, nil
	}

	return 0, nil
}

// RecommendLocationRewardBiw .
func (ub *UserBalanceRepo) RecommendLocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", rewardAmount), "recommend_total": gorm.Expr("recommend_location_total + ?", rewardAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	if "stop" == stop {
		if 0 < userBalance.BalanceDhb {
			tmp := tmpMaxNew
			tmp -= tmp * feeRate / 1000
			if err = ub.data.DB(ctx).Table("user_balance").
				Where("user_id=?", userId).
				Updates(map[string]interface{}{"balance_dhb": 0, "balance_usdt": gorm.Expr("balance_usdt + ?", tmp)}).Error; nil != err {
				return 0, errors.NotFound("user balance err", "user balance not found")
			}

			var userBalanceRecode UserBalanceRecord
			userBalanceRecode.Balance = userBalance.BalanceDhb
			userBalanceRecode.UserId = userBalance.UserId
			userBalanceRecode.Type = "exchange"
			userBalanceRecode.CoinType = "dhb"
			userBalanceRecode.Amount = tmp
			err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
			if err != nil {
				return 0, nil
			}

			var (
				reward Reward
			)

			reward.UserId = userId
			reward.Amount = userBalance.BalanceDhb
			reward.AmountB = tmp
			reward.Type = "exchange_system" // 本次分红的行为类型
			reward.TypeRecordId = userBalanceRecode.ID
			reward.Reason = "exchange_2" // 给我分红的理由
			err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
			if err != nil {
				return 0, nil
			}
		}
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_recommend_location" // 本次分红的行为类型
	reward.Reason = "recommend_location"             // 给我分红的理由
	reward.ReasonLocationId = recommendNum
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// RecommendRewardBiw .
func (ub *UserBalanceRepo) RecommendRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64, userIdTwo int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", rewardAmount), "recommend_total": gorm.Expr("recommend_total + ?", rewardAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	//if "stop" == stop {
	//	if 0 < userBalance.BalanceDhb {
	//		tmp := tmpMaxNew
	//		tmp -= tmp * feeRate / 1000
	//		if err = ub.data.DB(ctx).Table("user_balance").
	//			Where("user_id=?", userId).
	//			Updates(map[string]interface{}{"balance_dhb": 0, "balance_usdt": gorm.Expr("balance_usdt + ?", tmp)}).Error; nil != err {
	//			return 0, errors.NotFound("user balance err", "user balance not found")
	//		}
	//
	//		var userBalanceRecode UserBalanceRecord
	//		userBalanceRecode.Balance = userBalance.BalanceDhb
	//		userBalanceRecode.UserId = userBalance.UserId
	//		userBalanceRecode.Type = "exchange"
	//		userBalanceRecode.CoinType = "dhb"
	//		userBalanceRecode.Amount = tmp
	//		err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	//		if err != nil {
	//			return 0, nil
	//		}
	//
	//		var (
	//			reward Reward
	//		)
	//
	//		reward.UserId = userId
	//		reward.Amount = userBalance.BalanceDhb
	//		reward.AmountB = tmp
	//		reward.Type = "exchange_system" // 本次分红的行为类型
	//		reward.TypeRecordId = userBalanceRecode.ID
	//		reward.Reason = "exchange_2" // 给我分红的理由
	//		err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	//		if err != nil {
	//			return 0, nil
	//		}
	//	}
	//}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.TypeRecordId = userIdTwo
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_recommend_daily" // 本次分红的行为类型
	reward.Reason = "recommend"                   // 给我分红的理由
	reward.ReasonLocationId = recommendNum
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserNewTwoNewTwo .
func (ui *UserInfoRepo) UpdateUserNewTwoNewTwo(ctx context.Context, userId int64, amount uint64, amountIspay float64) error {
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"amount": gorm.Expr("amount + ?", amount)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	res = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_raw_float": gorm.Expr("balance_raw_float + ?", amountIspay)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}

	res = ui.data.DB(ctx).Table("total").Where("id=?", 1).
		Updates(map[string]interface{}{"one": gorm.Expr("one + ?", amount)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}

	//now := time.Now().UTC()
	//var lasUpdated time.Time
	//if 16 <= now.Hour() {
	//	lasUpdated = now.AddDate(0, 0, 1)
	//} else {
	//	lasUpdated = now
	//}
	//todayStart := time.Date(lasUpdated.Year(), lasUpdated.Month(), lasUpdated.Day(), 16, 0, 0, 0, time.UTC)

	var buyRecord BuyRecord
	buyRecord.UserId = userId
	buyRecord.Amount = float64(amount)
	buyRecord.AmountGet = 0
	buyRecord.Status = 1
	buyRecord.LastUpdated = time.Now().UTC().Unix()

	res = ui.data.DB(ctx).Table("buy_record").Create(&buyRecord)
	if res.Error != nil {
		return errors.New(500, "CREATE_LOCATION_ERROR", "占位信息创建失败")
	}

	var (
		err    error
		reward Reward
	)

	reward.UserId = userId
	reward.AmountNew = float64(amount)
	reward.AmountNewTwo = float64(amount)
	reward.Type = "USDT"  // 本次分红的行为类型
	reward.Reason = "buy" // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return errors.New(500, "CREATE_LOCATION_ERROR", "占位信息创建失败")
	}

	var (
		rewardTwo Reward
	)

	rewardTwo.UserId = userId
	rewardTwo.AmountNew = amountIspay
	rewardTwo.AmountNewTwo = amountIspay
	rewardTwo.Type = "ISPAY"  // 本次分红的行为类型
	rewardTwo.Reason = "send" // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&rewardTwo).Error
	if err != nil {
		return errors.New(500, "CREATE_LOCATION_ERROR", "占位信息创建失败")
	}

	return nil
}

// UpdateUserMyTotalAmountAdd .
func (ui *UserInfoRepo) UpdateUserMyTotalAmountAdd(ctx context.Context, userId int64, amountUsdt, myTotal float64) error {

	if 0 < myTotal {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{
				"my_total_amount": gorm.Expr("my_total_amount + ?", amountUsdt),
				"last":            gorm.Expr("last + ?", 1),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_raw_float": gorm.Expr("balance_raw_float + ?", myTotal)})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
		}

		var (
			rewardTwo Reward
		)

		rewardTwo.UserId = userId
		rewardTwo.AmountNew = myTotal
		rewardTwo.AmountNewTwo = myTotal
		rewardTwo.Type = "ISPAY" // 本次分红的行为类型
		rewardTwo.TypeRecordId = int64(myTotal)
		rewardTwo.Reason = "send" // 给我分红的理由
		err := ui.data.DB(ctx).Table("reward").Create(&rewardTwo).Error
		if err != nil {
			return errors.New(500, "CREATE_LOCATION_ERROR", "占位信息创建失败")
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"my_total_amount": gorm.Expr("my_total_amount + ?", amountUsdt)})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	return nil
}

// UpdateUserMyTotalAmount .
func (ui *UserInfoRepo) UpdateUserMyTotalAmount(ctx context.Context, userId int64, amountUsdt float64) error {
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"my_total_amount": gorm.Expr("my_total_amount - ?", amountUsdt)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateUserMyTotalAmountSub .
func (ui *UserInfoRepo) UpdateUserMyTotalAmountSub(ctx context.Context, userId int64, amountUsdt float64) error {
	return nil
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"my_total_amount": gorm.Expr("my_total_amount - ?", amountUsdt)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateUserRewardRecommend2 .
func (ui *UserInfoRepo) UpdateUserRewardRecommend2(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string) error {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{
				"amount":     gorm.Expr("amount - ?", amountOrigin),
				"out_rate":   gorm.Expr("out_rate + ?", 1),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
			Updates(map[string]interface{}{
				"status":     2,
				"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	} else {
		if 0 < usdt {
			res := ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
				Updates(map[string]interface{}{
					"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				})
			if res.Error != nil {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}
		}
	}

	if 0 < usdt {
		if err = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{
				"balance_usdt_float":    gorm.Expr("balance_usdt_float + ?", usdt),
				"balance_raw_float":     gorm.Expr("balance_raw_float + ?", raw),
				"recommend_total_float": gorm.Expr("recommend_total_float + ?", usdtOrigin),
			}).Error; nil != err {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var reward Reward
		reward.UserId = userId
		reward.AmountNew = usdt
		reward.AmountNewTwo = raw
		reward.Address = address
		reward.Reason = "recommend" // 直推
		err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateUserSubBuyRecord .
func (ui *UserInfoRepo) UpdateUserSubBuyRecord(ctx context.Context, id, userId int64, amountOrigin float64) error {
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{
			"amount":     gorm.Expr("amount - ?", amountOrigin),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
		Updates(map[string]interface{}{
			"status": 8,
		})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	res = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_raw_float": gorm.Expr("balance_raw_float - ?", amountOrigin)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}

	res = ui.data.DB(ctx).Table("total").Where("id=?", 1).
		Updates(map[string]interface{}{"one": gorm.Expr("one - ?", amountOrigin)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}

	return nil
}

// UpdateUserRewardDailyLocation .
func (ui *UserInfoRepo) UpdateUserRewardDailyLocation(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool) error {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{
				"amount":     gorm.Expr("amount - ?", amountOrigin),
				"out_rate":   gorm.Expr("out_rate + ?", 1),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
			Updates(map[string]interface{}{
				"status":     2,
				"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	} else {
		if 0 < usdt {
			res := ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
				Updates(map[string]interface{}{
					"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				})
			if res.Error != nil {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}
		}
	}

	if 0 < usdt {
		if err = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{
				"balance_usdt_float":   gorm.Expr("balance_usdt_float + ?", usdt),
				"balance_raw_float":    gorm.Expr("balance_raw_float + ?", raw),
				"location_total_float": gorm.Expr("location_total_float + ?", usdtOrigin),
			}).Error; nil != err {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var reward Reward
		reward.UserId = userId
		reward.AmountNew = usdt
		reward.AmountNewTwo = raw
		reward.Reason = "location"
		err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateUserRewardAreaOne .
func (ui *UserInfoRepo) UpdateUserRewardAreaOne(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string, i, cl int64, two bool) error {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{
				"amount":     gorm.Expr("amount - ?", amountOrigin),
				"out_rate":   gorm.Expr("out_rate + ?", 1),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
			Updates(map[string]interface{}{
				"status":     2,
				"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	} else {
		if 0 < usdt {
			res := ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
				Updates(map[string]interface{}{
					"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				})
			if res.Error != nil {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}
		}
	}

	if 0 < usdt {
		if two {
			if err = ui.data.DB(ctx).Table("user_balance").
				Where("user_id=?", userId).
				Updates(map[string]interface{}{
					"balance_usdt_float":   gorm.Expr("balance_usdt_float + ?", usdt),
					"balance_raw_float":    gorm.Expr("balance_raw_float + ?", raw),
					"area_total_float_two": gorm.Expr("area_total_float_two + ?", usdtOrigin),
				}).Error; nil != err {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}

			var reward Reward
			reward.UserId = userId
			reward.AmountNew = usdt
			reward.AmountNewTwo = raw
			reward.Address = address
			reward.TypeRecordId = i
			reward.BalanceRecordId = cl
			reward.Reason = "area_two" // 直推
			err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
			if err != nil {
				return err
			}
		} else {
			if err = ui.data.DB(ctx).Table("user_balance").
				Where("user_id=?", userId).
				Updates(map[string]interface{}{
					"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", usdt),
					"balance_raw_float":  gorm.Expr("balance_raw_float + ?", raw),
					"area_total_float":   gorm.Expr("area_total_float + ?", usdtOrigin),
				}).Error; nil != err {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}

			var reward Reward
			reward.UserId = userId
			reward.AmountNew = usdt
			reward.AmountNewTwo = raw
			reward.Address = address
			reward.TypeRecordId = i
			reward.BalanceRecordId = cl
			reward.Reason = "area" // 直推
			err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// GetRewardYes .
func (u *UserRepo) GetRewardYes(ctx context.Context) ([]*biz.Reward, error) {
	var rewards []*Reward

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	} else {
		startDate = now.AddDate(0, 0, -2)
		endDate = now.AddDate(0, 0, -1)
	}

	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	fmt.Println(todayStart, todayEnd)
	res := make([]*biz.Reward, 0)
	if err := u.data.db.
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Where("reason=?", "buy").
		Table("reward").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:        reward.ID,
			UserId:    reward.UserId,
			AmountNew: reward.AmountNew,
		})
	}

	return res, nil
}

// UpdateUserRewardRecommendNewTwo .
func (ui *UserInfoRepo) UpdateUserRewardRecommendNewTwo(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string, i int64) error {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{
				"amount":     gorm.Expr("amount - ?", amountOrigin),
				"out_rate":   gorm.Expr("out_rate + ?", 1),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
			Updates(map[string]interface{}{
				"status":     2,
				"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	} else {
		if 0 < usdt {
			res := ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
				Updates(map[string]interface{}{
					"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				})
			if res.Error != nil {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}
		}
	}

	if 0 < usdt {
		if err = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{
				"balance_usdt_float":        gorm.Expr("balance_usdt_float + ?", usdt),
				"balance_raw_float":         gorm.Expr("balance_raw_float + ?", raw),
				"recommend_total_float_two": gorm.Expr("recommend_total_float_two + ?", usdtOrigin),
			}).Error; nil != err {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var reward Reward
		reward.UserId = userId
		reward.AmountNew = usdt
		reward.AmountNewTwo = raw
		reward.Address = address
		reward.TypeRecordId = i
		reward.Reason = "recommend_two" // 直推
		err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateUserRewardAllNew .
func (ui *UserInfoRepo) UpdateUserRewardAllNew(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool) error {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{
				"amount":     gorm.Expr("amount - ?", amountOrigin),
				"out_rate":   gorm.Expr("out_rate + ?", 1),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
			Updates(map[string]interface{}{
				"status":     2,
				"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	} else {
		if 0 < usdt {
			res := ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
				Updates(map[string]interface{}{
					"amount_get": gorm.Expr("amount_get + ?", usdtOrigin),
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				})
			if res.Error != nil {
				return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
			}
		}
	}

	if 0 < usdt {
		if err = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{
				"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", usdt),
				"balance_raw_float":  gorm.Expr("balance_raw_float + ?", raw),
				"all_float":          gorm.Expr("all_float + ?", usdtOrigin),
			}).Error; nil != err {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var reward Reward
		reward.UserId = userId
		reward.AmountNew = usdt
		reward.AmountNewTwo = raw
		reward.Reason = "all" // 直推
		err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// GetBuyRecord .
func (ui *UserInfoRepo) GetBuyRecord(ctx context.Context, day int) ([]*biz.BuyRecord, error) {
	res := make([]*biz.BuyRecord, 0)

	var buyRecord []*BuyRecord
	if err := ui.data.db.Table("buy_record").Where("status=?", 1).
		Order("id asc").Find(&buyRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "buy_record ERROR", err.Error())
	}

	for _, v := range buyRecord {
		res = append(res, &biz.BuyRecord{
			ID:          v.ID,
			UserId:      v.UserId,
			Status:      v.Status,
			Amount:      v.Amount,
			AmountGet:   v.AmountGet,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			LastUpdated: v.LastUpdated,
		})
	}

	return res, nil
}

// UpdateTotalOne .
func (ui *UserInfoRepo) UpdateTotalOne(ctx context.Context, amountUsdt float64) error {
	res := ui.data.DB(ctx).Table("total").Where("id=?", 1).
		Updates(map[string]interface{}{"one": gorm.Expr("one + ?", amountUsdt)})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}

	return nil
}

// UpdateUserRewardStake .
func (ui *UserInfoRepo) UpdateUserRewardStake(ctx context.Context, userId int64, amountUsdt float64, stakeId int64) (int64, error) {
	var err error

	//if err = ui.data.DB(ctx).Table("user_balance").
	//	Where("user_id=?", userId).
	//	Updates(map[string]interface{}{"balance_raw_float": gorm.Expr("balance_raw_float + ?", amountUsdt)}).Error; nil != err {
	//	return 0, errors.NotFound("user balance err", "user balance not found")
	//}

	if err = ui.data.DB(ctx).Table("stake").
		Where("id=?", stakeId).
		Updates(map[string]interface{}{"reward": gorm.Expr("reward + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = stakeId
	reward.Type = "system_reward_location_daily" // 本次分红的行为类型
	reward.Reason = "stake_reward"               // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// UpdateUserRewardStakeReomve .
func (ui *UserInfoRepo) UpdateUserRewardStakeReomve(ctx context.Context, userId int64, amountUsdt float64, stakeId int64) (int64, error) {
	var err error

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_raw_float": gorm.Expr("balance_raw_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	if err = ui.data.DB(ctx).Table("stake").
		Where("id=?", stakeId).
		Updates(map[string]interface{}{"status": 1}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = stakeId
	reward.Type = "system_reward_location_daily" // 本次分红的行为类型
	reward.Reason = "stake_reward_remove"        // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// UpdateUserRewardTotalOver .
func (ui *UserInfoRepo) UpdateUserRewardTotalOver(ctx context.Context) error {
	res := ui.data.DB(ctx).Table("total").Where("id=?", 1).
		Updates(map[string]interface{}{"one": 0})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}

	res = ui.data.DB(ctx).Table("user").Where("id>=?", 1).
		Updates(map[string]interface{}{"amount_biw": 0, "amount_recommend_usdt_get": 0})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "one信息修改失败")
	}
	return nil
}

// UpdateUserRewardTotalTwoNew .
func (ui *UserInfoRepo) UpdateUserRewardTotalTwoNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("user_id=?", userId).
			Updates(map[string]interface{}{"status": 2})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "area_total_float_three": gorm.Expr("area_total_float_three + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "USDT"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.ReasonLocationId = i
	reward.Type = "system_reward_total_two_daily" // 本次分红的行为类型
	reward.Reason = "total_two"

	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserRewardTotalOneNew .
func (ui *UserInfoRepo) UpdateUserRewardTotalOneNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("user_id=?", userId).
			Updates(map[string]interface{}{"status": 2})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "area_total_float_two": gorm.Expr("area_total_float_two + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "USDT"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.ReasonLocationId = i
	reward.Type = "system_reward_total_one_daily" // 本次分红的行为类型
	reward.Reason = "total_one"

	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserRewardAreaNew .
func (ui *UserInfoRepo) UpdateUserRewardAreaNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, level, i int64, address string) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("user_id=?", userId).
			Updates(map[string]interface{}{"status": 2})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "area_total_float": gorm.Expr("area_total_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "USDT"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = level
	reward.ReasonLocationId = i
	reward.Address = address
	reward.Type = "system_reward_area_daily" // 本次分红的行为类型
	reward.Reason = "area"

	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserRewardRecommendNew .
func (ui *UserInfoRepo) UpdateUserRewardRecommendNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64, address string) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "amount_recommend_usdt_get": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("user_id=?", userId).
			Updates(map[string]interface{}{"status": 2})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "recommend_level_float": gorm.Expr("recommend_level_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "USDT"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.Type = "system_reward_recommend_b_daily" // 本次分红的行为类型
	reward.Reason = "recommend_b"
	reward.ReasonLocationId = i
	reward.Address = address
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

func (ui *UserInfoRepo) UpdateUserRewardNewFour(ctx context.Context, userId int64, amountUsdt float64) (int64, error) {
	var err error

	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"amount_four_get": gorm.Expr("amount_four_get + ?", amountUsdt)})
	if res.Error != nil {
		return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.BalanceNew = userBalance.BalanceUsdtFloat
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "USDT"
	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_location_daily" // 本次分红的行为类型
	reward.Reason = "four_reward"                // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

func (ui *UserInfoRepo) UpdateUserRewardNew(ctx context.Context, id, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		res = ui.data.DB(ctx).Table("buy_record").Where("user_id=?", userId).
			Updates(map[string]interface{}{"status": 2})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		now := time.Now().UTC()
		var lasUpdated time.Time
		if 16 <= now.Hour() {
		} else {
			lasUpdated = now.AddDate(0, 0, -1)
		}
		todayStart := time.Date(lasUpdated.Year(), lasUpdated.Month(), lasUpdated.Day(), 16, 0, 0, 0, time.UTC)

		res = ui.data.DB(ctx).Table("buy_record").Where("id=?", id).
			Updates(map[string]interface{}{"last_updated": todayStart.Unix(), "amount_get": float64(1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "location_total_float": gorm.Expr("location_total_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.BalanceNew = userBalance.BalanceUsdtFloat
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "USDT"
	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_location_daily" // 本次分红的行为类型
	reward.Reason = "location"                   // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserReward .
func (ui *UserInfoRepo) UpdateUserReward(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "amount_recommend_usdt_get": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "location_total_float": gorm.Expr("location_total_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.BalanceNew = userBalance.BalanceUsdtFloat
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_location_daily" // 本次分红的行为类型
	reward.Reason = "location"                   // 给我分红的理由
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserRewardRecommendUserGet .
func (ui *UserInfoRepo) UpdateUserRewardRecommendUserGet(ctx context.Context, userId int64, amountUsdt float64, enough bool, amount float64) error {

	if enough {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_recommend_usdt_get": amount})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_recommend_usdt_get": gorm.Expr("amount_recommend_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	return nil
}

// UpdateUserRewardRecommend .
func (ui *UserInfoRepo) UpdateUserRewardRecommend(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, address string) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "amount_recommend_usdt_get": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "recommend_total_float": gorm.Expr("recommend_total_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.Type = "system_reward_recommend_daily" // 本次分红的行为类型
	reward.Reason = "recommend"
	reward.Address = address
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserRewardArea .
func (ui *UserInfoRepo) UpdateUserRewardArea(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, tmpLevel, stop bool, level, i int64, address string) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "amount_recommend_usdt_get": 0, "out_rate": gorm.Expr("out_rate + ?", 1)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdtTotal
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if tmpLevel {
		if err = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "area_total_float_three": gorm.Expr("area_total_float_three + ?", amountUsdt)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	} else {
		if err = ui.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "area_total_float": gorm.Expr("area_total_float + ?", amountUsdt)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = level
	reward.ReasonLocationId = i
	reward.Address = address
	if tmpLevel {
		reward.Type = "system_reward_area_daily_three" // 本次分红的行为类型
		reward.Reason = "area_three"
	} else {
		reward.Type = "system_reward_area_daily" // 本次分红的行为类型
		reward.Reason = "area"
	}
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserRewardAreaTwo .
func (ui *UserInfoRepo) UpdateUserRewardAreaTwo(ctx context.Context, userId int64, amountUsdt float64, stop bool) (int64, error) {
	var err error

	if stop {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": 0, "amount_usdt": 0, "last": 0, "amount_recommend_usdt_get": 0})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

		var rewardStop Reward
		rewardStop.UserId = userId
		rewardStop.AmountNew = amountUsdt
		rewardStop.Type = "out"   // 本次分红的行为类型
		rewardStop.Reason = "out" // 给我分红的理由
		err = ui.data.DB(ctx).Table("reward").Create(&rewardStop).Error
		if err != nil {
			return 0, err
		}
	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
			Updates(map[string]interface{}{"amount_usdt_get": gorm.Expr("amount_usdt_get + ?", amountUsdt)})
		if res.Error != nil {
			return 0, errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}
	}

	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": gorm.Expr("balance_usdt_float + ?", amountUsdt), "area_two_total_float": gorm.Expr("area_two_total_float + ?", amountUsdt)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ui.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"

	userBalanceRecode.AmountNew = amountUsdt
	err = ui.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.AmountNew = amountUsdt
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_area_two_daily" // 本次分红的行为类型
	reward.Reason = "area_two"
	err = ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateUserIspay .
func (ui *UserInfoRepo) UpdateUserIspay(ctx context.Context, userId int64, amount uint64) error {
	var err error
	if err = ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_raw_float": float64(amount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	return nil
}

// UpdateUserNewTwoNewThree .
func (ui *UserInfoRepo) UpdateUserNewTwoNewThree(ctx context.Context, userId int64, amount uint64, last uint64, coinType string) error {
	if err := ui.data.DB(ctx).Table("user").
		Where("id=?", userId).
		Updates(map[string]interface{}{"amount": amount}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.AmountNew = float64(amount)
	reward.AmountNewTwo = float64(last)
	reward.Type = "system_reward_area_two_daily" // 本次分红的行为类型
	reward.Reason = "deposit_two"
	err := ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserUsdtFloat .
func (ui *UserInfoRepo) UpdateUserUsdtFloat(ctx context.Context, userId int64, amount float64, last float64, coinType string) error {
	if err := ui.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_float": amount}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.AmountNew = amount
	reward.AmountNewTwo = last
	reward.Type = "system_reward_area_two_daily" // 本次分红的行为类型
	reward.Reason = "deposit_three"
	err := ui.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserRecommendLevel .
func (ui *UserInfoRepo) UpdateUserRecommendLevel(ctx context.Context, userId int64, level uint64) error {
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"recommend_level": level})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateUserPass .
func (ui *UserInfoRepo) UpdateUserPass(ctx context.Context, userId int64, pass string) error {
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"password": pass})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateUserRecommendLevel2 .
func (ui *UserInfoRepo) UpdateUserRecommendLevel2(ctx context.Context, userId int64, level uint64) error {
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"vip_admin": level})
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateUserAmountFour .
func (ui *UserInfoRepo) UpdateUserAmountFour(ctx context.Context, userId int64, amountFour float64, do bool) error {
	updateMap := map[string]interface{}{"amount_four": amountFour}
	if do {
		updateMap["amount_four_get"] = 0
	}
	res := ui.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(updateMap)
	if res.Error != nil {
		return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
	}

	return nil
}

// UpdateUserLast .
func (ui *UserInfoRepo) UpdateUserLast(ctx context.Context, userId int64, coinType string) error {
	if "USDT" == coinType {
		res := ui.data.DB(ctx).Table("user").Where("id=? and last>=?", userId, 0).
			Updates(map[string]interface{}{"last": 0})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

	} else {
		res := ui.data.DB(ctx).Table("user").Where("id=? and last_biw>=?", userId, 0).
			Updates(map[string]interface{}{"last_biw": 0})
		if res.Error != nil {
			return errors.New(500, "UPDATE_USER_ERROR", "用户信息修改失败")
		}

	}
	return nil
}

// GetUsersNewTwo .
func (u *UserRepo) GetUsersNewTwo(ctx context.Context) ([]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Order("id desc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, item := range users {
		res = append(res, &biz.User{
			ID:              item.ID,
			Address:         item.Address,
			AddressTwo:      item.AddressTwo,
			PrivateKey:      item.PrivateKey,
			PrivateKeyThree: item.PrivateKeyThree,
			WordThree:       item.WordThree,
			AddressThree:    item.AddressThree,
			Last:            item.Last,
			LastBiw:         item.LastBiw,
		})
	}

	return res, nil
}

// LocationRewardBiw .
func (ub *UserBalanceRepo) LocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, stop string, currentMaxNew int64, feeRate int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", rewardAmount), "location_total": gorm.Expr("location_total + ?", rewardAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	//if "stop" == stop {
	//if 0 < userBalance.BalanceDhb {
	//	tmp := currentMaxNew
	//	tmp -= tmp * feeRate / 1000
	//	if err = ub.data.DB(ctx).Table("user_balance").
	//		Where("user_id=?", userId).
	//		Updates(map[string]interface{}{"balance_dhb": 0, "balance_usdt": gorm.Expr("balance_usdt + ?", tmp)}).Error; nil != err {
	//		return 0, errors.NotFound("user balance err", "user balance not found")
	//	}
	//
	//	var userBalanceRecode UserBalanceRecord
	//	userBalanceRecode.Balance = userBalance.BalanceDhb
	//	userBalanceRecode.UserId = userBalance.UserId
	//	userBalanceRecode.Type = "exchange"
	//	userBalanceRecode.CoinType = "dhb"
	//	userBalanceRecode.Amount = tmp
	//	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	//	if err != nil {
	//		return 0, nil
	//	}
	//
	//	var (
	//		reward Reward
	//	)
	//
	//	reward.UserId = userId
	//	reward.Amount = userBalance.BalanceDhb
	//	reward.AmountB = tmp
	//	reward.Type = "exchange_system" // 本次分红的行为类型
	//	reward.TypeRecordId = userBalanceRecode.ID
	//	reward.Reason = "exchange_2" // 给我分红的理由
	//	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	//	if err != nil {
	//		return 0, nil
	//	}
	//}
	//}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_location_daily" // 本次分红的行为类型
	reward.Reason = "location"                   // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// PriceChange .
func (ub *UserBalanceRepo) PriceChange(ctx context.Context, userId int64, rewardAmount int64, up string) error {
	var err error

	var reward Reward
	reward.UserId = userId
	reward.Amount = rewardAmount
	reward.Type = up               // 本次分红的行为类型
	reward.Reason = "price_change" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// GetStake .
func (ub *UserBalanceRepo) GetStake(ctx context.Context) ([]*biz.Stake, error) {
	var stake []*Stake
	res := make([]*biz.Stake, 0)
	if err := ub.data.DB(ctx).Table("stake").
		Where("status=?", 0).Find(&stake).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "PRICE CHANGE ERROR", err.Error())
	}

	for _, v := range stake {
		res = append(res, &biz.Stake{
			ID:        v.ID,
			UserId:    v.UserId,
			Status:    v.Status,
			Day:       v.Day,
			Amount:    v.Amount,
			Reward:    v.Reward,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetPriceChangeConfig .
func (ub *UserBalanceRepo) GetPriceChangeConfig(ctx context.Context) (*biz.PriceChange, error) {
	var priceChange PriceChange
	if err := ub.data.DB(ctx).Table("price_change").Where("status=?", 0).Order("id asc").First(&priceChange).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "PRICE CHANGE ERROR", err.Error())
	}

	return &biz.PriceChange{
		ID:        priceChange.ID,
		Origin:    priceChange.Origin,
		Price:     priceChange.Price,
		Status:    priceChange.Price,
		CreatedAt: priceChange.CreatedAt,
	}, nil
}

// CreatePriceChangeConfig .
func (c *ConfigRepo) CreatePriceChangeConfig(ctx context.Context, origin int64, price int64) error {
	var err error

	var priceChange PriceChange
	priceChange.Origin = origin
	priceChange.Price = price
	err = c.data.DB(ctx).Table("price_change").Create(&priceChange).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdatePriceChangeStatus .
func (c *ConfigRepo) UpdatePriceChangeStatus(ctx context.Context, id int64, status int64) (bool, error) {
	var priceChange PriceChange
	priceChange.Status = status

	res := c.data.DB(ctx).Table("price_change").Where("id=?", id).Updates(&priceChange)
	if res.Error != nil {
		return false, errors.New(500, "UPDATE_PRICE_CHANGE_ERROR", "价格变更记录修改失败")
	}

	return true, nil
}

// AreaRewardBiw .
func (ub *UserBalanceRepo) AreaRewardBiw(ctx context.Context, userId int64, rewardAmount int64, tmpCurrentReward int64, areaType int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", tmpCurrentReward), "area_total": gorm.Expr("area_total + ?", tmpCurrentReward)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	//if "stop" == stop {
	//	if 0 < userBalance.BalanceDhb {
	//		tmp := tmpMaxNew
	//		tmp -= tmp * feeRate / 1000
	//		if err = ub.data.DB(ctx).Table("user_balance").
	//			Where("user_id=?", userId).
	//			Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb - ?", userBalance.BalanceDhb), "balance_usdt": gorm.Expr("balance_usdt + ?", tmp)}).Error; nil != err {
	//			return 0, errors.NotFound("user balance err", "user balance not found")
	//		}
	//
	//		var userBalanceRecode UserBalanceRecord
	//		userBalanceRecode.Balance = userBalance.BalanceDhb
	//		userBalanceRecode.UserId = userBalance.UserId
	//		userBalanceRecode.Type = "exchange"
	//		userBalanceRecode.CoinType = "dhb"
	//		userBalanceRecode.Amount = tmp
	//		err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	//		if err != nil {
	//			return 0, nil
	//		}
	//
	//		var (
	//			reward Reward
	//		)
	//
	//		reward.UserId = userId
	//		reward.Amount = userBalance.BalanceDhb
	//		reward.AmountB = tmp
	//		reward.Type = "exchange_system" // 本次分红的行为类型
	//		reward.TypeRecordId = userBalanceRecode.ID
	//		reward.Reason = "exchange_2" // 给我分红的理由
	//		err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	//		if err != nil {
	//			return 0, nil
	//		}
	//	}
	//}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = rewardAmount
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = tmpCurrentReward
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = tmpCurrentReward
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_area_daily" // 本次分红的行为类型
	reward.Reason = "area"                   // 给我分红的理由
	reward.ReasonLocationId = areaType
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// RecommendTeamReward .
func (ub *UserBalanceRepo) RecommendTeamReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, locationId int64, recommendNum int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountDhb)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_daily" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_team" // 给我分红的理由
	reward.ReasonLocationId = recommendNum
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// SystemReward .
func (ub *UserBalanceRepo) SystemReward(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "system_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// SystemDailyReward .
func (ub *UserBalanceRepo) SystemDailyReward(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "system_fee_daily"   // 本次分红的行为类型
	reward.Reason = "system_fee_daily" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// SystemWithdrawReward .
func (ub *UserBalanceRepo) SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "system_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// GetSystemYesterdayLocationReward .
func (ub *UserBalanceRepo) GetSystemYesterdayLocationReward(ctx context.Context, day int) ([]*biz.UserBalanceRecord, error) {
	var rewards []*UserBalanceRecord

	now := time.Now().UTC().AddDate(0, 0, day)
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Where("type=?", "exchange").
		Table("user_balance_record").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	res := make([]*biz.UserBalanceRecord, 0)
	for _, reward := range rewards {
		res = append(res, &biz.UserBalanceRecord{
			ID:           reward.ID,
			UserId:       reward.UserId,
			Balance:      reward.Balance,
			Amount:       reward.Amount,
			Type:         reward.Type,
			CoinType:     reward.CoinType,
			CreatedAt:    reward.CreatedAt,
			UpdatedAt:    reward.UpdatedAt,
			BalanceNew:   reward.BalanceNew,
			AmountNew:    reward.AmountNew,
			AmountNewTwo: reward.AmountNewTwo,
		})
	}

	return res, nil
}

// GetRewardYes .
func (ub *UserBalanceRepo) GetRewardYes(ctx context.Context) ([]*biz.Reward, error) {
	var rewards []*Reward

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	res := make([]*biz.Reward, 0)
	if err := ub.data.db.
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Table("reward").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			AmountNew:        reward.AmountNew,
		})
	}

	return res, nil
}

// GetTotal .
func (ub *UserBalanceRepo) GetTotal(ctx context.Context) (*biz.Total, error) {
	var total Total
	if err := ub.data.db.Table("total").First(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("REWARD_NOT_FOUND", "total not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}
	return &biz.Total{
		ID:    total.ID,
		One:   total.One,
		Two:   total.Two,
		Three: total.Three,
	}, nil
}

// GetSystemYesterdayDailyReward .
func (ub *UserBalanceRepo) GetSystemYesterdayDailyReward(ctx context.Context, day int) (*biz.Reward, error) {
	var reward Reward

	now := time.Now().UTC().AddDate(0, 0, day)
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.
		Where("user_id=?", 999999999).
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Where("type=?", "system_fee_daily").
		Where("reason=?", "system_fee_daily").
		Table("reward").First(&reward).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}
	return &biz.Reward{
		ID:               reward.ID,
		UserId:           reward.UserId,
		Amount:           reward.Amount,
		BalanceRecordId:  reward.BalanceRecordId,
		Type:             reward.Type,
		TypeRecordId:     reward.TypeRecordId,
		Reason:           reward.Reason,
		ReasonLocationId: reward.ReasonLocationId,
		LocationType:     reward.LocationType,
	}, nil
}

// SystemFee .
func (ub *UserBalanceRepo) SystemFee(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "system_fee" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// UserFee .
func (ub *UserBalanceRepo) UserFee(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward" // 本次分红的行为类型
	reward.Reason = "fee"         // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UserDailyFee .
func (ub *UserBalanceRepo) UserDailyFee(ctx context.Context, userId int64, amount int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}

	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_daily" // 本次分红的行为类型
	reward.Reason = "fee_daily"         // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UserDailyRecommendArea .
func (ub *UserBalanceRepo) UserDailyRecommendArea(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountDhb)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}

	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_daily"    // 本次分红的行为类型
	reward.Reason = "daily_recommend_area" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UserDailyBalanceReward .
func (ub *UserBalanceRepo) UserDailyBalanceReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountDhb)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}

	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward_daily"    // 本次分红的行为类型
	reward.Reason = "daily_balance_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// RecommendWithdrawReward .
func (ub *UserBalanceRepo) RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_vip" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// RecommendWithdrawTopReward .
func (ub *UserBalanceRepo) RecommendWithdrawTopReward(ctx context.Context, userId int64, amount int64, locationId int64, vip int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}

	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_vip_top" // 给我分红的理由
	reward.ReasonLocationId = vip
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalRecommendTopReward .
func (ub *UserBalanceRepo) NormalRecommendTopReward(ctx context.Context, userId int64, amount int64, locationId int64, reasonId int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_top" // 给我分红的理由
	reward.ReasonLocationId = reasonId
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NewNormalRecommendReward .
func (ub *UserBalanceRepo) NewNormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error) {
	var err error

	var location LocationNew
	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			return 0, errors.New(500, "LOCATION ERROR", err.Error())
		}

	}

	if 0 < location.ID {

		res := ub.data.DB(ctx).Table("location_new").
			Where("id=?", location.ID).
			Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
		if 0 == res.RowsAffected || res.Error != nil {
			return 0, res.Error
		}

		var reward Reward
		reward.UserId = userId
		reward.Amount = amount
		reward.BalanceRecordId = 0
		reward.Type = "location" // 本次分红的行为类型
		reward.TypeRecordId = locationId
		reward.Reason = "recommend" // 给我分红的理由
		err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
		if err != nil {
			return 0, err
		}
	}

	return 0, nil
}

// NewNormalRecommendReward .
//func (ub *UserBalanceRepo) NewNormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error) {
//	var err error
//
//	var location LocationNew
//	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//
//		} else {
//			return 0, errors.New(500, "LOCATION ERROR", err.Error())
//		}
//
//	}
//
//	if 0 < location.ID {
//		if len(tmpRecommendUserIdsInt) > 0 {
//			if err = ub.data.DB(ctx).Table("user_info").
//				Where("user_id in (?)", tmpRecommendUserIdsInt).
//				Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
//				return 0, errors.NotFound("user balance err", "user balance not found")
//			}
//		}
//
//		//res := ub.data.DB(ctx).Table("location_new").
//		//	Where("id=?", location.ID).
//		//	Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
//		//if 0 == res.RowsAffected || res.Error != nil {
//		//	return 0, res.Error
//		//}
//
//		if err = ub.data.DB(ctx).Table("user_balance").
//			Where("user_id=?", userId).
//			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
//			return 0, errors.NotFound("user balance err", "user balance not found")
//		}
//
//		var reward Reward
//		reward.UserId = userId
//		reward.Amount = amount
//		reward.BalanceRecordId = 0
//		reward.Type = "location" // 本次分红的行为类型
//		reward.TypeRecordId = locationId
//		reward.Reason = "recommend" // 给我分红的理由
//		err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
//		if err != nil {
//			return 0, err
//		}
//	}
//
//	return 0, nil
//}

// NormalRecommendReward .
func (ub *UserBalanceRepo) NormalRecommendReward(ctx context.Context, userId int64, rewardAmount int64, rewardAmount2 int64, locationId int64, status string, status2 string, type1 string, reason string) (int64, error) {
	var err error
	if "running" == status {
		amount := rewardAmount
		if "stop" == status2 {
			amount = rewardAmount2
		}

		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = reason // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = type1 // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalReward4 .
func (ub *UserBalanceRepo) NormalReward4(ctx context.Context, userId int64, rewardAmount int64, locationId int64) (int64, error) {
	var err error

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt_new": gorm.Expr("balance_usdt_new + ?", rewardAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdtNew
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward_withdraw"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "reward_withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "reward_withdraw" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalReward3 .
func (ub *UserBalanceRepo) NormalReward3(ctx context.Context, userId int64, rewardAmount int64, rewardAmount2 int64, locationId int64, status string, status2 string) (int64, error) {
	var err error

	if "running" == status {
		amount := rewardAmount
		if "stop" == status2 {
			amount = rewardAmount2
		}

		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt_new": gorm.Expr("balance_usdt_new + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdtNew
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward_withdraw"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "reward_withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "reward_withdraw" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalRecommendReward2 .
func (ub *UserBalanceRepo) NormalRecommendReward2(ctx context.Context, userId int64, rewardAmount int64, locationId int64, type1 string, reason string) (int64, error) {
	var err error

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", rewardAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward_token"
	userBalanceRecode.Amount = rewardAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = rewardAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = reason // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = type1 // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// LocationNewDailyReward .
func (ub *UserBalanceRepo) LocationNewDailyReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance_lock").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance_lock").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location_daily" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "location_daily" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// WithdrawNewRewardTeamRecommend .
func (ub *UserBalanceRepo) WithdrawNewRewardTeamRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error) {
	var err error
	var location LocationNew
	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.NotFound("LOCATION_NOT_FOUND", "location not found")
		}

		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	if 0 >= location.ID {
		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	//res := ub.data.DB(ctx).Table("location_new").
	//	Where("id=?", location.ID).
	//	Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
	//if 0 == res.RowsAffected || res.Error != nil {
	//	return 0, res.Error
	//}

	if len(tmpRecommendUserIdsInt) > 0 {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountB)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.Amount = amount
	reward.AmountB = amountB
	reward.BalanceRecordId = 0
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_team" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// WithdrawNewRewardRecommend .
func (ub *UserBalanceRepo) WithdrawNewRewardRecommend(ctx context.Context, userId int64, amount int64, amountB int64, withdrawId int64, tmpRecommendUserIdsInt []int64) (int64, error) {
	var err error

	//if err = ub.data.DB(ctx).Table("user_balance_lock").
	//	Where("user_id=?", userId).
	//	Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
	//	return 0, errors.NotFound("user balance err", "user balance not found")
	//}
	//
	//if len(tmpRecommendUserIdsInt) > 0 {
	//	if err = ub.data.DB(ctx).Table("user_info").
	//		Where("user_id in (?)", tmpRecommendUserIdsInt).
	//		Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
	//		return 0, errors.NotFound("user balance err", "user balance not found")
	//	}
	//}

	//var userBalance UserBalance
	//err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance_lock").First(&userBalance).Error
	//if err != nil {
	//	return 0, err
	//}

	//var userBalanceRecode UserBalanceRecord
	//userBalanceRecode.Balance = userBalance.BalanceUsdt
	//userBalanceRecode.UserId = userBalance.UserId
	//userBalanceRecode.Type = "reward"
	//userBalanceRecode.Amount = amount
	//err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	//if err != nil {
	//	return 0, err
	//}
	var location LocationNew
	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.NotFound("LOCATION_NOT_FOUND", "location not found")
		}

		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	if 0 >= location.ID {
		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	//res := ub.data.DB(ctx).Table("location_new").
	//	Where("id=?", location.ID).
	//	Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
	//if 0 == res.RowsAffected || res.Error != nil {
	//	return 0, res.Error
	//}

	if len(tmpRecommendUserIdsInt) > 0 {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountB)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.Amount = amount
	reward.AmountB = amountB
	reward.BalanceRecordId = 0
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = withdrawId
	reward.Reason = "recommend" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// UpdateLocationNewMax .
func (ub *UserBalanceRepo) UpdateLocationNewMax(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	var location LocationNew
	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.NotFound("LOCATION_NOT_FOUND", "location not found")
		}

		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	res := ub.data.DB(ctx).Table("location_new").
		Where("id=?", location.ID).
		Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
	if 0 == res.RowsAffected || res.Error != nil {
		return 0, res.Error
	}

	return 0, nil
}

// WithdrawNewRewardSecondRecommend .
func (ub *UserBalanceRepo) WithdrawNewRewardSecondRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error) {
	var err error
	var location LocationNew
	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.NotFound("LOCATION_NOT_FOUND", "location not found")
		}

		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	if 0 >= location.ID {
		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	//res := ub.data.DB(ctx).Table("location_new").
	//	Where("id=?", location.ID).
	//	Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
	//if 0 == res.RowsAffected || res.Error != nil {
	//	return 0, res.Error
	//}

	if len(tmpRecommendUserIdsInt) > 0 {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountB)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.Amount = amount
	reward.AmountB = amountB
	reward.BalanceRecordId = 0
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "second_recommend" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// WithdrawNewRewardLevelRecommend .
func (ub *UserBalanceRepo) WithdrawNewRewardLevelRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error) {
	var err error
	var location LocationNew
	if err = ub.data.db.Table("location_new").Where("user_id", userId).Order("id desc").First(&location).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.NotFound("LOCATION_NOT_FOUND", "location not found")
		}

		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	if 0 >= location.ID {
		return 0, errors.New(500, "LOCATION ERROR", err.Error())
	}

	//res := ub.data.DB(ctx).Table("location_new").
	//	Where("id=?", location.ID).
	//	Updates(map[string]interface{}{"current_max_new": gorm.Expr("current_max_new + ?", amount)})
	//if 0 == res.RowsAffected || res.Error != nil {
	//	return 0, res.Error
	//}

	if len(tmpRecommendUserIdsInt) > 0 {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount), "balance_dhb": gorm.Expr("balance_dhb + ?", amountB)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var reward Reward
	reward.UserId = userId
	reward.Amount = amount
	reward.AmountB = amountB
	reward.BalanceRecordId = 0
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "level_recommend" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// NormalWithdrawRecommendReward .
func (ub *UserBalanceRepo) NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalWithdrawRecommendTopReward .
func (ub *UserBalanceRepo) NormalWithdrawRecommendTopReward(ctx context.Context, userId int64, amount int64, locationId int64, reasonId int64, status string) (int64, error) {
	var err error
	if "running" == status {
		if err = ub.data.DB(ctx).Table("user_balance").
			Where("user_id=?", userId).
			Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
			return 0, errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_top" // 给我分红的理由
	reward.ReasonLocationId = reasonId
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// GetUserCurrentMonthRecommendByUserId .
func (uc *UserCurrentMonthRecommendRepo) GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*biz.UserCurrentMonthRecommend, error) {
	var userCurrentMonthRecommends []*UserCurrentMonthRecommend
	res := make([]*biz.UserCurrentMonthRecommend, 0)
	if err := uc.data.db.Where(&UserCurrentMonthRecommend{UserId: userId}).Table("user_current_month_recommend").Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found")
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error())
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		res = append(res, &biz.UserCurrentMonthRecommend{
			ID:              userCurrentMonthRecommend.ID,
			UserId:          userCurrentMonthRecommend.UserId,
			RecommendUserId: userCurrentMonthRecommend.RecommendUserId,
			Date:            userCurrentMonthRecommend.Date,
		})
	}
	return res, nil
}

// GetUserCurrentMonthRecommendGroupByUserId .
func (uc *UserCurrentMonthRecommendRepo) GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *biz.Pagination, userId int64) ([]*biz.UserCurrentMonthRecommend, error, int64) {
	var (
		count                      int64
		userCurrentMonthRecommends []*UserCurrentMonthRecommend
	)
	res := make([]*biz.UserCurrentMonthRecommend, 0)

	instance := uc.data.db.Table("user_current_month_recommend")
	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found"), 0
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error()), 0
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		res = append(res, &biz.UserCurrentMonthRecommend{
			ID:              userCurrentMonthRecommend.ID,
			UserId:          userCurrentMonthRecommend.UserId,
			RecommendUserId: userCurrentMonthRecommend.RecommendUserId,
			Date:            userCurrentMonthRecommend.Date,
		})
	}
	return res, nil, count
}

// GetUserRewardByUserId .
func (ub *UserBalanceRepo) GetUserRewardByUserId(ctx context.Context, userId int64) ([]*biz.Reward, error) {
	var rewards []*Reward
	res := make([]*biz.Reward, 0)
	if err := ub.data.db.Where("user_id", userId).Table("reward").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			CreatedAt:        reward.CreatedAt,
		})
	}
	return res, nil
}

// GetBalanceRewardCurrent .
func (ub *UserBalanceRepo) GetBalanceRewardCurrent(ctx context.Context, now time.Time) ([]*biz.BalanceReward, error) {
	var balanceRewards []*BalanceReward
	res := make([]*biz.BalanceReward, 0)

	if err := ub.data.db.
		Where("h=? and m<=? and m>=?", now.Hour(), now.Minute(), now.Minute()-5).
		Where("status=?", 1).Order("id asc").Table("balance_reward").Find(&balanceRewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, balanceReward := range balanceRewards {
		res = append(res, &biz.BalanceReward{
			ID:             balanceReward.ID,
			UserId:         balanceReward.UserId,
			Status:         balanceReward.Status,
			Amount:         balanceReward.Amount,
			SetDate:        balanceReward.SetDate,
			LastRewardDate: balanceReward.LastRewardDate,
		})
	}
	return res, nil
}

// GetUserTrades .
func (ub *UserBalanceRepo) GetUserTrades(ctx context.Context, b *biz.Pagination, userId int64) ([]*biz.Trade, error, int64) {
	var (
		trades []*Trade
		count  int64
	)
	res := make([]*biz.Trade, 0)

	instance := ub.data.db.Table("trade")

	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&trades).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, v := range trades {
		res = append(res, &biz.Trade{
			ID:           v.ID,
			UserId:       v.UserId,
			AmountCsd:    v.AmountCsd,
			RelAmountCsd: v.RelAmountCsd,
			AmountHbs:    v.AmountHbs,
			CsdReward:    v.CsdReward,
			RelAmountHbs: v.RelAmountHbs,
			Status:       v.Status,
			CreatedAt:    v.CreatedAt,
		})
	}

	return res, nil, count
}

// GetUserBuyById .
func (ub *UserBalanceRepo) GetUserBuyById(id int64) (*biz.BuyRecord, error) {
	var reward *BuyRecord

	instance := ub.data.db.Table("buy_record").Where("id=?", id)
	if err := instance.First(&reward).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	return &biz.BuyRecord{
		ID:          reward.ID,
		UserId:      reward.UserId,
		Status:      reward.Status,
		LastUpdated: reward.LastUpdated,
		Amount:      reward.Amount,
		AmountGet:   reward.AmountGet,
		CreatedAt:   reward.CreatedAt,
		UpdatedAt:   reward.UpdatedAt,
	}, nil
}

// GetUserBuyByUserId .
func (ub *UserBalanceRepo) GetUserBuyByUserId(ctx context.Context, userId int64) ([]*biz.BuyRecord, error) {
	var (
		rewards []*BuyRecord
	)
	res := make([]*biz.BuyRecord, 0)

	instance := ub.data.db.Table("buy_record").Where("user_id=?", userId)

	instance = instance.Where("status=?", 1)

	if err := instance.Order("id desc").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.BuyRecord{
			ID:          reward.ID,
			UserId:      reward.UserId,
			Status:      reward.Status,
			LastUpdated: reward.LastUpdated,
			Amount:      reward.Amount,
			AmountGet:   reward.AmountGet,
			CreatedAt:   reward.CreatedAt,
			UpdatedAt:   reward.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBuy .
func (ub *UserBalanceRepo) GetUserBuy(ctx context.Context, b *biz.Pagination, userId int64) ([]*biz.BuyRecord, error, int64) {
	var (
		rewards []*BuyRecord
		count   int64
	)
	res := make([]*biz.BuyRecord, 0)

	instance := ub.data.db.Table("buy_record")

	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	instance = instance.Where("status=?", 1)

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.BuyRecord{
			ID:          reward.ID,
			UserId:      reward.UserId,
			Status:      reward.Status,
			LastUpdated: reward.LastUpdated,
			Amount:      reward.Amount,
			AmountGet:   reward.AmountGet,
			CreatedAt:   reward.CreatedAt,
			UpdatedAt:   reward.UpdatedAt,
		})
	}

	return res, nil, count
}

// GetUserRewards .
func (ub *UserBalanceRepo) GetUserRewards(ctx context.Context, b *biz.Pagination, userId int64, reason string) ([]*biz.Reward, error, int64) {
	var (
		rewards []*Reward
		count   int64
	)
	res := make([]*biz.Reward, 0)

	instance := ub.data.db.Table("reward")

	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	if "" != reason {
		instance = instance.Where("reason=?", reason)
	} else {
		//tmpReason := make([]string, 0)
		//tmpReason = append(tmpReason, "area")
		//tmpReason = append(tmpReason, "total_two")
		//tmpReason = append(tmpReason, "total_one")
		//tmpReason = append(tmpReason, "recommend_b")
		//tmpReason = append(tmpReason, "location")
		//tmpReason = append(tmpReason, "recommend")
		//tmpReason = append(tmpReason, "recommend_two")
		//instance = instance.Where("reason in (?)", tmpReason)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			AmountB:          reward.AmountB,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			CreatedAt:        reward.CreatedAt,
			AmountNew:        reward.AmountNew,
			AmountNewTwo:     reward.AmountNewTwo,
			Status:           reward.Status,
			Address:          reward.Address,
		})
	}
	return res, nil, count
}

// GetUserRewardRecommendSort .
func (ub *UserBalanceRepo) GetUserRewardRecommendSort(ctx context.Context) ([]*biz.UserSortRecommendReward, error) {
	var total []*UserSortRecommendReward
	res := make([]*biz.UserSortRecommendReward, 0)

	now := time.Now().UTC().AddDate(0, 0, -1)
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("reward").
		Where("type=?", "location").
		Where("reason=?", "recommend").
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Group("user_id").
		Select("sum(amount) as total, user_id").
		Order("total desc").
		Scopes(Paginate(1, 4)).
		Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}
		return res, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, v := range total {
		res = append(res, &biz.UserSortRecommendReward{
			Total:  v.Total,
			UserId: v.UserId,
		})
	}

	return res, nil
}

// GetUserRewardsLastMonthFee .
func (ub *UserBalanceRepo) GetUserRewardsLastMonthFee(ctx context.Context) ([]*biz.Reward, error) {
	var (
		rewards []*Reward
	)
	res := make([]*biz.Reward, 0)

	instance := ub.data.db.Table("reward")

	now := time.Now().UTC().Add(8 * time.Hour)
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, time.UTC)
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, time.UTC)

	if err := instance.Where("created_at>=?", lastMonthStart).
		Where("created_at<=?", lastMonthEnd).
		Where("reason=?", "system_fee").
		Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			CreatedAt:        reward.CreatedAt,
		})
	}
	return res, nil
}

// CreateUserCurrentMonthRecommend .
func (uc *UserCurrentMonthRecommendRepo) CreateUserCurrentMonthRecommend(ctx context.Context, u *biz.UserCurrentMonthRecommend) (*biz.UserCurrentMonthRecommend, error) {
	var userCurrentMonthRecommend UserCurrentMonthRecommend
	userCurrentMonthRecommend.UserId = u.UserId
	userCurrentMonthRecommend.RecommendUserId = u.RecommendUserId
	userCurrentMonthRecommend.Date = u.Date
	res := uc.data.DB(ctx).Table("user_current_month_recommend").Create(&userCurrentMonthRecommend)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_CURRENT_MONTH_RECOMMEND_ERROR", "用户当月推荐人创建失败")
	}

	return &biz.UserCurrentMonthRecommend{
		ID:              userCurrentMonthRecommend.ID,
		UserId:          userCurrentMonthRecommend.UserId,
		RecommendUserId: userCurrentMonthRecommend.RecommendUserId,
		Date:            userCurrentMonthRecommend.Date,
	}, nil
}

// GetUserBalanceByUserIds .
func (ub UserBalanceRepo) GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.UserBalance, error) {
	var userBalances []*UserBalance
	res := make(map[int64]*biz.UserBalance)
	if err := ub.data.db.Where("user_id IN (?)", userIds).Table("user_balance").Find(&userBalances).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	for _, userBalance := range userBalances {
		res[userBalance.UserId] = &biz.UserBalance{
			ID:                     userBalance.ID,
			UserId:                 userBalance.UserId,
			BalanceUsdt:            userBalance.BalanceUsdt,
			BalanceDhb:             userBalance.BalanceDhb,
			BalanceC:               userBalance.BalanceC,
			BalanceUsdtFloat:       userBalance.BalanceUsdtFloat,
			BalanceRawFloat:        userBalance.BalanceRawFloat,
			BalanceKsdtFloat:       userBalance.BalanceKsdtFloat,
			AreaTotalFloat:         userBalance.AreaTotalFloat,
			AreaTotalFloatTwo:      userBalance.AreaTotalFloatTwo,
			RecommendTotalFloat:    userBalance.RecommendTotalFloat,
			LocationTotalFloat:     userBalance.LocationTotalFloat,
			AllFloat:               userBalance.AllFloat,
			RecommendLevelFloat:    userBalance.RecommendLevelFloat,
			RecommendTotalFloatTwo: userBalance.RecommendTotalFloatTwo,
		}
	}

	return res, nil
}

// GetUserBalanceLockByUserIds .
func (ub UserBalanceRepo) GetUserBalanceLockByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.UserBalance, error) {
	var userBalances []*UserBalance
	res := make(map[int64]*biz.UserBalance)
	if err := ub.data.db.Where("user_id IN (?)", userIds).Table("user_balance_lock").Find(&userBalances).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	for _, userBalance := range userBalances {
		res[userBalance.UserId] = &biz.UserBalance{
			ID:          userBalance.ID,
			UserId:      userBalance.UserId,
			BalanceUsdt: userBalance.BalanceUsdt,
			BalanceDhb:  userBalance.BalanceDhb,
		}
	}

	return res, nil
}

type UserBalanceTotal struct {
	Total int64
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

// GetUserBalanceUsdtTotal .
func (ub UserBalanceRepo) GetUserBalanceUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance").Select("sum(balance_usdt) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceLockUsdtTotal .
func (ub UserBalanceRepo) GetUserBalanceLockUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_lock").Select("sum(balance_usdt) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserLocationNewCurrentMaxNew .
func (ub UserBalanceRepo) GetUserLocationNewCurrentMaxNew(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("location_new").Select("sum(current_max_new) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserLocationNewCurrentMax .
func (ub UserBalanceRepo) GetUserLocationNewCurrentMax(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("location_new").Select("sum(current_max) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserLocationNewCurrent .
func (ub UserBalanceRepo) GetUserLocationNewCurrent(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("location_new").Select("sum(current) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceDHBTotal .
func (ub UserBalanceRepo) GetUserBalanceDHBTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance").Select("sum(balance_dhb) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordUsdtTotal .
func (ub UserBalanceRepo) GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "deposit").
		Where("coin_type=?", "usdt").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordUsdtTotalTwo .
func (ub UserBalanceRepo) GetUserBalanceRecordUsdtTotalTwo(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("eth_user_record").
		Where("coin_type=?", "USDT").
		Select("sum(amount_two) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

type UserBalanceTotalFloat struct {
	Total float64
}

// GetUserBalanceRecordUsdtTotalThree .
func (ub UserBalanceRepo) GetUserBalanceRecordUsdtTotalThree(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("eth_user_record").
		Select("sum(amount_two) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

func (ub UserBalanceRepo) GetUserBalanceRecordCsdTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "deposit").
		Where("coin_type=?", "CSD").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordHbsTotal .
func (ub UserBalanceRepo) GetUserBalanceRecordHbsTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "deposit").
		Where("coin_type=?", "HBS").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserRewardLocationTotalToday .
func (ub UserBalanceRepo) GetUserRewardLocationTotalToday(ctx context.Context, reason string) (int64, error) {
	var total UserBalanceTotal

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("reward").
		Where("reason=?", reason).
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, nil
		}

		return total.Total, errors.New(500, "REWARD RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordUsdtTotalToday .
func (ub UserBalanceRepo) GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error) {
	var total UserBalanceTotal

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "deposit").
		Where("coin_type=?", "usdt").
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetSystemWithdrawUsdtFeeTotalToday .
func (ub UserBalanceRepo) GetSystemWithdrawUsdtFeeTotalToday(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	startDate = now.AddDate(0, 0, -1)
	endDate = now

	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("reward").
		Where("user_id=?", 999999999).
		Where("type=?", "withdraw").
		Where("reason=?", "system_reward").
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserWithdrawUsdtTotalToday .
func (ub UserBalanceRepo) GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "withdraw").
		Where("coin_type=?", "usdt").
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserWithdrawDhbTotalToday .
func (ub UserBalanceRepo) GetUserWithdrawDhbTotalToday(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "withdraw").
		Where("coin_type=?", "dhb").
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserWithdrawUsdtTotalByUserIds .
func (ub UserBalanceRepo) GetUserWithdrawUsdtTotalByUserIds(ctx context.Context, userIds []int64) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("withdraw").
		Where("user_id in(?)", userIds).
		Where("status=?", "success").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserWithdrawUsdtTotal .
func (ub UserBalanceRepo) GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "withdraw").
		Where("coin_type=?", "usdt").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserWithdrawDhbTotal .
func (ub UserBalanceRepo) GetUserWithdrawDhbTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "withdraw").
		Where("coin_type=?", "dhb").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserRewardUsdtTotal .
func (ub UserBalanceRepo) GetUserRewardUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "reward").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserRewardBalanceRewardTotal .
func (ub UserBalanceRepo) GetUserRewardBalanceRewardTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("reason=?", "daily_balance_reward").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetBalanceRewardTotal .
func (ub UserBalanceRepo) GetBalanceRewardTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("balance_reward").
		Where("status=?", 1).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_REWARD_NOT_FOUND", "user balance reward not found")
		}

		return total.Total, errors.New(500, "USER BALANCE REWARD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetSystemRewardUsdtTotal .
func (ub UserBalanceRepo) GetSystemRewardUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("reward").
		Where("reason=? or reason=?", "system_reward", "system_fee").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserInfoByUserIds .
func (ui *UserInfoRepo) GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.UserInfo, error) {
	var userInfos []*UserInfo
	res := make(map[int64]*biz.UserInfo, 0)
	if err := ui.data.db.Where("user_id IN (?)", userIds).Table("user_info").Find(&userInfos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USERINFO_NOT_FOUND", "userinfo not found")
		}

		return nil, errors.New(500, "USERINFO ERROR", err.Error())
	}

	for _, userInfo := range userInfos {
		res[userInfo.UserId] = &biz.UserInfo{
			ID:               userInfo.ID,
			UserId:           userInfo.UserId,
			Vip:              userInfo.Vip,
			HistoryRecommend: userInfo.HistoryRecommend,
			TeamCsdBalance:   userInfo.TeamCsdBalance,
		}
	}

	return res, nil
}

// GetUserInfosByVipAndLockVip .
func (ui *UserInfoRepo) GetUserInfosByVipAndLockVip(ctx context.Context) ([]*biz.UserInfo, error) {
	var userInfos []*UserInfo
	res := make([]*biz.UserInfo, 0)
	if err := ui.data.db.Where("vip>=? and lock_vip=?", 1, 1).Table("user_info").Find(&userInfos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USERINFO_NOT_FOUND", "userinfo not found")
		}

		return nil, errors.New(500, "USERINFO ERROR", err.Error())
	}

	for _, userInfo := range userInfos {
		res = append(res, &biz.UserInfo{
			ID:               userInfo.ID,
			UserId:           userInfo.UserId,
			Vip:              userInfo.Vip,
			LockVip:          userInfo.LockVip,
			HistoryRecommend: userInfo.HistoryRecommend,
			TeamCsdBalance:   userInfo.TeamCsdBalance,
		})
	}

	return res, nil
}

// GetUserCurrentMonthRecommendCountByUserIds .
func (uc *UserCurrentMonthRecommendRepo) GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error) {
	var userCurrentMonthRecommends []*UserCurrentMonthRecommend
	res := make(map[int64]int64, 0)
	if err := uc.data.db.Where("user_id IN (?)", userIds).Table("user_current_month_recommend").Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found")
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error())
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		if _, ok := res[userCurrentMonthRecommend.UserId]; !ok {
			res[userCurrentMonthRecommend.UserId] = 1
		} else {
			res[userCurrentMonthRecommend.UserId]++
		}
	}
	return res, nil
}

// GetUserLastMonthRecommend .
func (uc *UserCurrentMonthRecommendRepo) GetUserLastMonthRecommend(ctx context.Context) ([]int64, error) {
	var userCurrentMonthRecommends []*UserCurrentMonthRecommend
	res := make([]int64, 0)

	now := time.Now().UTC().Add(8 * time.Hour)
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, time.UTC)
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, time.UTC)

	if err := uc.data.db.Table("user_current_month_recommend").
		Group("user_id").
		Having("count(id) >= 5").
		Where("date>=?", lastMonthStart).
		Where("date<=?", lastMonthEnd).
		Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found")
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error())
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		res = append(res, userCurrentMonthRecommend.UserId)
	}
	return res, nil
}
