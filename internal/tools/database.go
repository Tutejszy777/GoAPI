package tools

// Database return
type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinsDetails struct {
	Username string
	Coins    int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinsDetails
	SetupDatabase() error
}


func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if er != nil {
		log.error(err)
		return nil, err
	}
	return &database, nil 
}
