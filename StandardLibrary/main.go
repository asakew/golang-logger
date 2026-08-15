package main

import (
	"StandardLibrary/logger"
	"log/slog"
)

func main() {
	// 1. Logerni sozlash (Development rejimi uchun false bering)
	isProduction := false
	logger.InitLogger(isProduction)

	slog.Info("Dastur muvaffaqiyatli ishga tushdi")

	// Simulyatsiya: Tizimga kirayotgan foydalanuvchi ma'lumotlari
	userID := "usr_9981"
	userIP := "192.168.1.50"

	// 2. Strukturali log yozish (Kalit va Qiymat juftligi orqali)
	slog.Info("Foydalanuvchi tizimga kirdi",
		slog.String("user_id", userID),
		slog.String("ip_address", userIP),
	)

	// 3. Xatolik logini yozish misoli
	errReason := "not_found"
	slog.Error("Ma'lumotlar bazasidan xatolik",
		slog.String("reason", errReason),
		slog.Int("attempt_count", 3),
	)
}
