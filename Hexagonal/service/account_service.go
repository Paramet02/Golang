package service

import (
	"paramet/logs"
	"paramet/repository"
	"time"
	"strings"
	"github.com/gofiber/fiber/v2"
)

type accountService struct {
	accRepo repository.AccountRpeository
}

func NewaccountService(accRepo repository.AccountRpeository) AccountService {
	return &accountService{accRepo : accRepo}
}

func (s accountService) NewAccount(id int , Request NewAccountRequest) (*AccountResponse, error) {
	// Validate
	// เช็ดว่า เงินน้อยกว่า 5000
	if Request.Amount < 5000 {
		return nil , fiber.NewError(fiber.StatusUnprocessableEntity , "Amount at least 5,000")
	}

	// เช็ดว่า ถ้าไม่ใช่ saveing และ checking ให้ return err
	if strings.ToLower(Request.AccountType) != "saving" && strings.ToLower(Request.AccountType) != "checking" {
		return nil , fiber.NewError(fiber.StatusBadRequest, "account type should be saving or checking")
	}

	// Request
	account := repository.Account{
		CustomerID: id,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"), // วันที่ ปัจจุบัน
		AccountType: Request.AccountType , // รับค่า request 
		Amount:      Request.Amount, // รับค่า request 
		Status:      1,
	}

	// ส่งข้อมูล account ไปยัง repository create  
	newAcc , err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)

		return nil , fiber.NewError(fiber.StatusInternalServerError)
	}

	// ส่งค่าให้ ผู้รับ(respones)
	respones := AccountResponse{
		AccountID: newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount: newAcc.Amount,
		Status: newAcc.Status,
	}
	
	return &respones , nil
}

func (s accountService) GetAccount(id int) ([]AccountResponse , error) {

	// รับ id จาก Params 
	accounts , err := s.accRepo.GetAll(id)
	if err != nil {
		logs.Error(err)

		return nil , fiber.NewError(fiber.StatusInternalServerError)
	}

	// ปั้นข้อมูลส่งให้ response
	response := []AccountResponse{}
	for _ , acccount := range accounts {
		response = append(response, AccountResponse{
			AccountID: acccount.AccountID,
			OpeningDate: acccount.OpeningDate,
			AccountType: acccount.AccountType,
			Amount: acccount.Amount,
			Status: acccount.Status,
		})
	}

	return response , nil
}