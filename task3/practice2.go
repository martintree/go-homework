package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Accounts struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Balance float32
}

type Transactions struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	FromAccountId int
	ToAccountId   int
	Amount        float32
}

func insertAccounts(db *gorm.DB, balance float32) {
	account := &Accounts{Balance: balance}
	result := db.Create(account)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("Account id=%d RowsAffected=%d\n", account.ID, result.RowsAffected)
}

func transferAccounts(db *gorm.DB, fromAccount *Accounts, toAccount *Accounts, transferAmount float32) error {
	if fromAccount == nil || toAccount == nil {
		return errors.New("fromAccount or toAccount is nil")
	}

	if fromAccount.Balance < transferAmount {
		return errors.New("余额不足")
	}
	fromAccount.Balance -= transferAmount
	toAccount.Balance += transferAmount
	db.Debug().Save(fromAccount)
	db.Debug().Save(toAccount)
	result := db.Debug().Create(&Transactions{FromAccountId: fromAccount.ID,
		ToAccountId: toAccount.ID, Amount: transferAmount})
	return result.Error
}

func findAccount(db *gorm.DB, id int) Accounts {
	var account Accounts
	db.Debug().First(&account, "ID = ?", id)
	return account
}

func main() {
	db, err := gorm.Open(mysql.Open("root:admin123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	//建表
	db.AutoMigrate(&Accounts{}, &Transactions{})

	//插入数据
	//insertAccounts(db, 1000)
	//insertAccounts(db, 100)

	//查询数据
	fromAccount := findAccount(db, 1)
	toAccount := findAccount(db, 2)

	//转账100
	db.Transaction(func(tx *gorm.DB) error {
		if err := transferAccounts(db, &fromAccount, &toAccount, 100); err != nil {
			fmt.Println("转账失败：", err)
			return err
		}

		fmt.Println("转账成功")
		return nil
	})

}
