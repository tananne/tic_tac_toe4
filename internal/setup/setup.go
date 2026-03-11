package setup

import (
	"fmt"
	"tic-tac-toe/internal/model"
	"tic-tac-toe/internal/player"
)

func ChooseMode() (string, error) {
	fmt.Println("1 - Игрок против игрока")
	fmt.Println("2 - Игрок против бота")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		return "human", nil
	case 2:
		return "bot", nil
	default:
		return "", fmt.Errorf("Неверный выбор. Пожалуйста, выберите 1 или 2.")
	}
}

func ChooseSide() (model.Mark, error) {
	fmt.Println("Вы играете против бота. Выберите сторону:")
	fmt.Println("1 - " + string(model.X))
	fmt.Println("2 - " + string(model.O))

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		return model.X, nil
	case 2:
		return model.O, nil
	default:
		return model.Empty, fmt.Errorf("Неверный выбор. Пожалуйста, выберите 1 или 2.")
	}
}

func getPlayerName() string {
	var name string
	fmt.Print("Введите имя игрока: ")
	fmt.Scanln(&name)

	if name == "" {
		name = "Игрок"
	}

	return name
}

func GetPlayers(mode string) (player.Player, player.Player) {
	switch mode {
	case "human":
		fmt.Println("Введите имя игрока " + string(model.X) + ":")
		nameX := getPlayerName()
		fmt.Println("Введите имя игрока " + string(model.O) + ":")
		nameO := getPlayerName()
		return player.NewHuman(nameX, model.X), player.NewHuman(nameO, model.O)
	case "bot":
		humanSide, err := ChooseSide()
		for err != nil {
			fmt.Println(err)
			humanSide, err = ChooseSide()
		}
		// bot should take opposite mark
		var botSide model.Mark
		if humanSide == model.X {
			botSide = model.O
		} else {
			botSide = model.X
		}
		fmt.Println("Бот играет за " + string(botSide) + ".")
		return player.NewHuman("Игрок", humanSide), player.NewBot(botSide)

	default:
		return nil, nil
	}
}
