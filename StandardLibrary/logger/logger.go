package logger

import (
	"log/slog"
	"os"
)

// InitLogger logerni loyiha rejimiga qarab sozlaydi
func InitLogger(isProd bool) {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo, // Faqat Info va undan yuqori loglarni ko'rsatish
		AddSource: !isProd,        // Dev rejimda log qaysi fayl va qatordan chiqqanini ko'rsatadi
	}

	if isProd {
		// Production rejim: JSON format (Kiberxavfsizlik va monitoring tizimlari uchun qulay)
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		// Development rejim: Oddiy matn format (Dasturchi o'qishi uchun qulay)
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	// Sozlangan logerni butun loyiha uchun standart (global) qilib o'rnatamiz
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
