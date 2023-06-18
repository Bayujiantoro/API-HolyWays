package handlers

import (
	"fmt"
	donationdto "holyways/dto/donation"
	dto "holyways/dto/result"
	"holyways/models"
	"holyways/repository"
	"log"
	"net/http"
	"os"
	"time"

	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerDonation struct {
	DonationRepo repository.DonationRepo
}

func DonationHandler(Donation repository.DonationRepo) *handlerDonation {
	return &handlerDonation{Donation}
}

func (h *handlerDonation) CreateDonation(c echo.Context) error  {
	// Money, _ := strconv.Atoi(c.FormValue("Money"))
	// FundID, _ := strconv.Atoi(c.FormValue("FundID"))
	request := new(donationdto.CreateDonation)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var donationIsMatch = false
	var donationId int
	for !donationIsMatch {
		donationId = int(time.Now().Unix())
		donationData, _ := h.DonationRepo.GetDonation(donationId)
		if donationData.ID == 0 {
			donationIsMatch = true
		}
	}

	userLogin := c.Get("userLogin")
	fmt.Println("userLogin : ", userLogin)
	userID := userLogin.(jwt.MapClaims)["id"].(float64)
	user, _ := h.DonationRepo.GetUserById(int(userID))

	Fund, _ := h.DonationRepo.GetFundById(request.FundID)
	fmt.Println(user)

	donate := models.Donation{
		ID: donationId,
		Date: time.Now(),
		Money: request.Money,
		FundID: Fund.ID,
		UserID: int(userID),
		User: user,
		Fund: Fund,
	}
	fmt.Println("ini request : ", donate)

	data, err := h.DonationRepo.CreateDonation(donate)
	fmt.Println("ini data : ", data)
	

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(data.ID),
			GrossAmt: int64(data.Money),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.User.Name,
			Email: data.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: snapResp,
	})
}

func (h *handlerDonation) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)
	donation, _ := h.DonationRepo.GetDonation(order_id)
	fmt.Print("ini payloadnya", notificationPayload)

	fmt.Println("transaction satus handler : ",transactionStatus)
	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			h.DonationRepo.UpdateDonateMidtrans("pending",order_id)
		} else if fraudStatus == "accept" {
			SendMail("success", donation)
			h.DonationRepo.UpdateDonateMidtrans("success", order_id)
		}
	} else if transactionStatus == "settlement" {
		SendMail("success", donation)
		h.DonationRepo.UpdateDonateMidtrans("success", order_id)
	} else if transactionStatus == "deny" {
		h.DonationRepo.UpdateDonateMidtrans("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.DonationRepo.UpdateDonateMidtrans("failed", order_id)

	} else if transactionStatus == "pending" {
		h.DonationRepo.UpdateDonateMidtrans("pending", order_id)

	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: notificationPayload})
}

func (h *handlerDonation) UpdateDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	Money, _ := strconv.Atoi(c.FormValue("Money"))
	request := models.Donation{
		Date: time.Now(),
		Money: Money,
	}
	fmt.Println("ini request : ", request)
	donation, err := h.DonationRepo.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	
	if request.Money > 0 {
		donation.Date = request.Date
	}
	data, err := h.DonationRepo.UpdateDonation(donation)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}

func (h *handlerDonation) FindDonation(c echo.Context) error {
	donation, err := h.DonationRepo.FindDonation() 

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: donation,
	})
}

func (h *handlerDonation) GetDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.DonationRepo.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}

func (h *handlerDonation) DeleteDonation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	donation, err := h.DonationRepo.GetDonation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	data, err := h.DonationRepo.DeleteDonation(donation, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{
		Code: http.StatusOK,
		Data: convertResponseDonation(data),
	})
}


func convertResponseDonation(donation models.Donation) donationdto.DonationResponse{
	return donationdto.DonationResponse{
		ID: donation.ID,
		Date: donation.Date,
		Money: donation.Money,
		UserID: donation.UserID,
		User: donation.User,
		FundID: donation.FundID,
		Fund: donation.Fund,
	}
}

func SendMail(status string, Donation models.Donation) {

	if status != Donation.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var productName = Donation.Fund.Title
		var price = strconv.Itoa(Donation.Money)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", Donation.User.Email)
		mailer.SetHeader("Subject", "Donation Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, productName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + Donation.User.Email)
	}
}