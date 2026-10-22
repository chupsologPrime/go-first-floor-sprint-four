package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	elements := strings.Split(data, ",")
	if len(elements) != 2 {
		return 0, 0, fmt.Errorf("Неверный ввод с клавиатуры")
	}
	steps, err := strconv.Atoi(elements[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("Число шагов должно быть больше 0")
	}
	period, err := time.ParseDuration(elements[1])
	if err != nil {
		return 0, 0, err
	}
	if period <= 0 {
		return 0, 0, fmt.Errorf("Время должно быть больше 0")
	}
	return steps, period, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, period, err := parsePackage(data)
	if err != nil {
		fmt.Println("Возникла ошибка:", err.Error())
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceMeter := float64(steps) * stepLength
	distanceKmeter := distanceMeter / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, period)
	if err != nil {
		fmt.Println("Возникла ошибка:", err.Error())
		return ""
	}
	resultStr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKmeter, calories)
	return resultStr

}
