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
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Получаем текущее время
	now := time.Now().Format("15:04:05 02.01.2006")
	
	// Создаём HTML с красным X на чёрном фоне
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate">
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Expires" content="0">
    <title>░ X ░</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            background-color: #000000;
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            font-family: 'Courier New', monospace;
        }
        
        .container {
            text-align: center;
        }
        
        .x {
            font-size: 40vw;
            font-weight: bold;
            color: #8B0000;
            text-shadow: 0 0 20px rgba(139, 0, 0, 0.5);
            animation: pulse 2s ease-in-out infinite;
            cursor: pointer;
            user-select: none;
            transition: transform 0.1s;
        }
        
        .x:hover {
            transform: scale(1.05);
            color: #FF0000;
            text-shadow: 0 0 40px rgba(255, 0, 0, 0.8);
        }
        
        .x:active {
            transform: scale(0.95);
        }
        
        @keyframes pulse {
            0% { opacity: 0.8; text-shadow: 0 0 10px rgba(139, 0, 0, 0.3); }
            50% { opacity: 1; text-shadow: 0 0 30px rgba(139, 0, 0, 0.8); }
            100% { opacity: 0.8; text-shadow: 0 0 10px rgba(139, 0, 0, 0.3); }
        }
        
        .timestamp {
            position: fixed;
            bottom: 20px;
            right: 20px;
            color: #333;
            font-size: 10px;
            font-family: monospace;
        }
        
        .instruction {
            position: fixed;
            bottom: 20px;
            left: 20px;
            color: #222;
            font-size: 10px;
            font-family: monospace;
        }
        
        @media (max-width: 768px) {
            .x {
                font-size: 60vw;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="x" onclick="location.reload()">✘</div>
    </div>
    <div class="timestamp">%s</div>
    <div class="instruction">click to refresh</div>
</body>
</html>`, now)
	
	// Записываем в файл
	err := os.WriteFile("index.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("❌ Ошибка при записи файла:", err)
		os.Exit(1)
	}
	
	fmt.Println("✅ index.html успешно сгенерирован!")
	fmt.Println("🗻 Дизайн: красный X на чёрном фоне")
}
