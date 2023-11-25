package env

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	environment = "devnet"
	databaseurl string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// house, err = solana.PrivateKeyFromBase58(os.Getenv("HOUSE_PRIVATE_KEY"))
	// if err != nil {
	// 	panic(err)
	// }
	// towerHouse, err = solana.PrivateKeyFromBase58(os.Getenv("TOWER_HOUSE_PRIVATE_KEY"))
	// if err != nil {
	// 	panic(err)
	// }
	// fee, err = solana.PublicKeyFromBase58(os.Getenv("FEE_PUBLIC_KEY"))
	// if err != nil {
	// 	panic(err)
	// }
	// feeBasisPoints, err = strconv.Atoi(os.Getenv("HOUSE_FEE_BASIS_POINTS"))
	// if err != nil {
	// 	panic(err)
	// }
	// towerFeeBasisPoints, err = strconv.Atoi(os.Getenv("TOWER_HOUSE_FEE_BASIS_POINTS"))
	// if err != nil {
	// 	panic(err)
	// }
	// towerHouseEdgeBasisPoints, err = strconv.Atoi(os.Getenv("TOWER_HOUSE_EDGE_BASIS_POINTS"))
	// if err != nil {
	// 	panic(err)
	// }
	// towerMaxPayout, err = strconv.Atoi(os.Getenv("TOWER_MAX_PAYOUT"))
	// if err != nil {
	// 	panic(err)
	// }
	// rpcUrl = os.Getenv("MAINNET_RPC_URL")
	// if rpcUrl == "" {
	// 	panic(err)
	// }
	// wsUrl = os.Getenv("MAINNET_WS_URL")
	// if wsUrl == "" {
	// 	panic(err)
	// }
	// webhook = os.Getenv("WEBHOOK")
	// if webhook == "" {
	// 	panic(errors.New("webook is nil"))
	// }

	databaseurl = os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		panic(errors.New("databaseurl is nil"))
	}
}

func Set(e string) {
	environment = e
}

func Get() string {
	return environment
}

func GetDatabaseURL() string {
	return databaseurl
}

func GetPort() string {
	switch environment {
	case "mainnet-beta":
		return ":4242"
	case "devnet":
		return ":4343"
	default:
		return ":4343"
	}
}

func GetHomePagePort() string {
	return ":4040"
}
