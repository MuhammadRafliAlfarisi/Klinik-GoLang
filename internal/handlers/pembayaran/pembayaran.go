// // payment_handler.go
package belajarHandler

//
//import (
//	"fmt"
//	"github.com/midtrans/midtrans-go"
//	"github.com/veritrans/go-midtrans"
//	"net/http"
//
//	"github.com/gofiber/fiber/v2"
//	"sistem-informasi-klinik/database"
//	"sistem-informasi-klinik/internal/model"
//)
//
//// HandleReservationPayment menangani pembayaran reservasi
//func Bayar(c *fiber.Ctx) error {
//	// Dapatkan data pembayaran dari request
//	var bayar model.Bayar
//	if err := c.BodyParser(&bayar); err != nil {
//		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses data pembayaran reservasi"})
//	}
//
//	// Inisialisasi klien Midtrans
//	midtrans.ClientKey = "Mid-client-mW8rTmTz_DJlDssb"
//	midtrans.ServerKey = "Mid-server-rj9KyfZ9dSgCEqlOcRxbmNR-"
//	midtrans.Environment = midtrans.Sandbox
//
//	// Lakukan pembayaran dengan Midtrans
//	chargeReq := &midtrans.chargerReq{
//		PaymentType: midtrans.PaymentTypeEwallet,
//		TransactionDetails: midtrans.TransactionDetails{
//			OrderID:  fmt.Sprintf("reservation_%d", reservationID),
//			GrossAmt: int64(totalAmount),
//		},
//		EWallet: &midtrans.EWalletDetail{
//			Provider: "DANA", // Ganti dengan provider e-wallet yang sesuai
//		},
//	}
//
//	payment, err := midtrans.Charge(chargeReq)
//	if err != nil {
//		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Gagal melakukan pembayaran: %v", err)})
//	}
//
//	// Simpan data pembayaran ke database
//	err = SavePaymentData(Bayar.TotalBayar)
//	if err != nil {
//		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan data pembayaran"})
//	}
//
//	// Respon berhasil pembayaran
//	return c.JSON(fiber.Map{"message": "Pembayaran berhasil", "payment_status": payment.Status})
//}
//
//// SavePaymentData menyimpan data pembayaran ke dalam database
//// SavePaymentData menyimpan data pembayaran ke dalam database
//func SavePaymentData(totalBayar int64) error {
//	// Buat objek Bayar
//	paymentData := model.Bayar{
//		TotalBayar: totalBayar,
//	}
//
//	// Simpan data pembayaran ke dalam database
//	result := database.DB.Create(&paymentData)
//	if result.Error != nil {
//		return result.Error
//	}
//
//	return nil
//}
