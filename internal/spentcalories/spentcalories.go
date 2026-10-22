package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	element := strings.Split(data, ",")
	if len(element) != 3 {
		return 0, "", 0, fmt.Errorf("Неверный ввод")
	}
	steps, err := strconv.Atoi(element[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("Шагов должно быть больше 0")
	}
	period, err := time.ParseDuration(element[2])
	if err != nil {
		return 0, "", 0, err
	}
	if period <= 0 {
		return 0, "", 0, fmt.Errorf("Время должно быть больше 0")
	}
	return steps, element[1], period, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	meters := stepLength * float64(steps)
	kilometers := meters / mInKm
	return kilometers
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	kilometers := distance(steps, height)
	hours := duration.Hours()
	return kilometers / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, name, period, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var calories float64
	var resultDistance float64
	var averageSpeed float64

	switch name {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, period)
		if err != nil {
			return "", err
		}
		resultDistance = distance(steps, height)
		averageSpeed = meanSpeed(steps, height, period)

	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, period)
		if err != nil {
			return "", err
		}
		resultDistance = distance(steps, height)
		averageSpeed = meanSpeed(steps, height, period)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	resultStr := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", name, period.Hours(), resultDistance, averageSpeed, calories)
	return resultStr, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некорректные параметры")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * averageSpeed * durationInMinutes) / minInH
	return calories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некорректные параметры")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * averageSpeed * durationInMinutes) / minInH
	calories = calories * walkingCaloriesCoefficient
	return calories, nil
}
