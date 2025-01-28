package main

import (
    "fmt"
    "image/gif"
    "log"
    "net/http"
    "os"
    "os/exec"
    "time"
    "github.com/hajimehoshi/ebiten/v2"
)

// Game - структура для игры
type Game struct {
    gifImage     *gif.GIF
    currentFrame int
    delay        time.Duration
    nextFrameAt  time.Time
    screenWidth  int
    screenHeight int
}

// loadGif загружает GIF-файл
func loadGif(filename string) (*gif.GIF, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return gif.DecodeAll(file)
}

// openBrowser открывает браузер в киоск-режиме
func openBrowser(url string) error {
    cmd := exec.Command("chromium", "--kiosk", url) // Используем флаг "--kiosk" для киоск-режима
    err := cmd.Start()
    if err != nil {
        return fmt.Errorf("failed to open browser: %v", err)
    }
    return nil
}

// Update обновляет состояние игры
func (g *Game) Update() error {
    // Проверка состояния Home Assistant
    if isHomeAssistantAvailable() {
        // Если Home Assistant доступен, открываем браузер
        err := openBrowser("http://192.168.11.12:8123")
        if err != nil {
            log.Println("Failed to open browser:", err)
        }
        os.Exit(0) // Завершаем программу после открытия браузера
    }

    // Обновление текущего кадра GIF-анимации
    if g.gifImage != nil && time.Now().After(g.nextFrameAt) {
        g.currentFrame++
        if g.currentFrame >= len(g.gifImage.Image) {
            g.currentFrame = 0
        }

        // Устанавливаем задержку до следующего кадра
        g.nextFrameAt = time.Now().Add(g.delay)
    }

    return nil
}

// Draw рисует кадр
func (g *Game) Draw(screen *ebiten.Image) {
    if g.gifImage != nil {
        img := g.gifImage.Image[g.currentFrame]
        ebitenImage := ebiten.NewImageFromImage(img)

        // Масштабируем изображение на весь экран
        op := &ebiten.DrawImageOptions{}
        screenWidth, screenHeight := float64(g.screenWidth), float64(g.screenHeight)
        imgWidth, imgHeight := float64(img.Bounds().Dx()), float64(img.Bounds().Dy())

        // Рассчитываем коэффициенты масштабирования
        scaleX := screenWidth / imgWidth
        scaleY := screenHeight / imgHeight
        op.GeoM.Scale(scaleX, scaleY)

        screen.DrawImage(ebitenImage, op)
    }
}

// Layout сообщает размеры окна
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return g.screenWidth, g.screenHeight
}

// isHomeAssistantAvailable проверяет доступность Home Assistant
func isHomeAssistantAvailable() bool {
    resp, err := http.Get("http://192.168.11.12:8123")
    if err != nil {
        return false
    }
    return resp.StatusCode == http.StatusOK
}

// main запускает игру
func main() {
    // Загрузка анимации GIF
    gifImage, err := loadGif("animation.gif")
    if err != nil {
        log.Fatal("Failed to load GIF:", err)
    }

    // Устанавливаем задержку между кадрами (например, 100 мс по умолчанию)
    delay := time.Duration(gifImage.Delay[0]) * 10 * time.Millisecond

    // Получаем текущие размеры экрана
    screenWidth, screenHeight := ebiten.ScreenSizeInFullscreen()

    // Создание экземпляра игры с загруженной анимацией
    game := &Game{
        gifImage:     gifImage,
        delay:        delay,
        screenWidth:  screenWidth,
        screenHeight: screenHeight,
    }

    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Home Assistant Checker")
    ebiten.SetFullscreen(true) // Включаем полноэкранный режим
    ebiten.SetWindowDecorated(false) // Отключаем рамки окна

    // Запуск игры
    go func() {
        // Запуск цикла проверки доступности Home Assistant каждые 5 секунд
        for {
            if isHomeAssistantAvailable() {
                break
            }
            time.Sleep(5 * time.Second)
        }
    }()

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
