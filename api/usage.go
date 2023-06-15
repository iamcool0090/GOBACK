package api

import (
	"encoding/json"
	"veda-backend/db"
)

func AddLoginTime(FromTime int) error {
	dataBase, err := db.NewUsageRepository()
	if err != nil {
		return err
	}

	use := db.Usage{
		FromTime: FromTime,
	}

	err = dataBase.CreateUsage(&use)
	if err != nil {
		return err
	}

	err = dataBase.Close()
	if err != nil {
		return err
	}

	return nil
}

func AddLogoutTime(ToTime int, SessionID int) error {
	dataBase, err := db.NewUsageRepository()
	if err != nil {
		return err
	}

	use := db.Usage{
		ToTime:    ToTime,
		SessionID: SessionID,
	}

	err = dataBase.UpdateToTime(&use)
	if err != nil {
		return err
	}

	return nil
}

func ReadUsage() ([]db.Usage, error) {
	dataBase, err := db.NewUsageRepository()
	if err != nil {
		return nil, err
	}

	usageBytes, err := dataBase.ReadUsage()
	if err != nil {
		return nil, err
	}

	var usages []db.Usage
	err = json.Unmarshal(usageBytes, &usages)
	if err != nil {
		return nil, err
	}

	return usages, nil
}
