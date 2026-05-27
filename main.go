package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Генерируем случайное число от 1 до 100
	randomNum := rand.Intn(100) + 1

	// Текущее время для отображения
	currentTime := time.Now().Format("02.01.2006 15:04:05")

	// Создаём HTML-контент
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate">
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Expires" content="0">
    <title>Случайное число | DevOps практика</title>
    <style>
        body {
            background-color: #81D8D6;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            text-align: center;
            padding: 50px;
            margin: 0;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background: white;
            border-radius: 20px;
            padding: 40px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.2);
        }
        h1 {
            color: #2C3E50;
            margin-bottom: 20px;
        }
        .number {
            font-size: 120px;
            font-weight: bold;
            color: #81D8D6;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.1);
            margin: 20px;
            padding: 20px;
            background: #f0f0f0;
            border-radius: 20px;
            display: inline-block;
            min-width: 200px;
        }
        .info {
            color: #666;
            font-size: 14px;
            margin-top: 30px;
            padding-top: 20px;
            border-top: 1px solid #eee;
        }
        .badge {
            display: inline-block;
            background: #2C3E50;
            color: white;
            padding: 5px 10px;
            border-radius: 5px;
            font-size: 12px;
            margin-top: 10px;
        }
        button {
            background: #2C3E50;
            color: white;
            border: none;
            padding: 12px 24px;
            font-size: 16px;
            border-radius: 8px;
            cursor: pointer;
            margin-top: 20px;
            transition: background 0.3s;
        }
        button:hover {
            background: #1a2a3a;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎲 Случайное число</h1>
        <div class="number">%d</div>
        <p>✨ Сгенерировано автоматически ✨</p>
        <button onclick="location.reload()">🔄 Обновить страницу</button>
        <div class="info">
            <div>⚡ Число генерируется при каждой сборке GitHub Actions</div>
            <div>📅 Последняя генерация: %s</div>
            <div class="badge">🚀 Автоматический деплой через GitHub Actions</div>
        </div>
    </div>
</body>
</html>`, randomNum, currentTime)

	// Записываем HTML в файл index.html
	err := os.WriteFile("index.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("❌ Ошибка при записи файла:", err)
		os.Exit(1)
	}

	fmt.Println("✅ index.html успешно сгенерирован!")
	fmt.Printf("🎲 Случайное число: %d\n", randomNum)
	fmt.Printf("🕐 Время генерации: %s\n", currentTime)
}
