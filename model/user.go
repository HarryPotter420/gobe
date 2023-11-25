package model

type User struct {
	Id            string `json:"id" bson:"_id,omitempty"`
	Username      string `json:"username" bson:"username"`
	Password      string `json:"password,omitempty" bson:"password"`
	WalletAddress string `json:"walletAddress" bson:"walletAddress"`
	GameWon       int    `json:"gameWon"`
	GameLost      int    `json:"gameLost"`
	GamePlayed    int    `json:"gamePlayed"`
	TotalWagered  int    `json:"totalWagered"`
	TotalWon      int    `json:"totalWon"`
}
