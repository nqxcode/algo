package interview1

// Условие задачи
// Мы в любим проводить соревнования, — недавно мы устроили чемпионат по шагам. И вот настало время подводить итоги!

// Необходимо определить userIds участников, которые прошли наибольшее количество шагов steps за все дни, не пропустив ни одного дня соревнований.

// Пример
// # Пример 1
// # ввод
// statistics = [
//         [{ userId: 1, steps: 1000 }, { userId: 2, steps: 1500 }],
//         [{ userId: 2, steps: 1000 }]
// ]

// # вывод
// champions = { userIds: [2], steps: 2500 }

// # Пример 2
// statistics = [
//         [{ userId: 1, steps: 2000 }, { userId: 2, steps: 1500 }],
//         [{ userId: 2, steps: 4000 }, { userId: 1, steps: 3500 }]
// ]

// # вывод
// champions = { userIds: [1, 2], steps: 5500 }

type Result struct {
	UserIds []int
	Steps   int
}

type UserInfo struct {
	UserID int
	Days   int
	Steps  int
}

func getChampions(statistics [][]map[string]int) Result {
	var result Result

	userInfoMap := make(map[int]*UserInfo)

	for _, userDataList := range statistics {
		for _, userData := range userDataList {
			userID := userData["userId"]
			steps := userData["steps"]

			if _, ok := userInfoMap[userID]; !ok {
				userInfoMap[userID] = &UserInfo{UserID: userID, Days: 1, Steps: steps}
			} else {
				userInfo := userInfoMap[userID]

				userInfo.Days = userInfo.Days + 1
				userInfo.Steps = userInfo.Steps + steps
			}
		}
	}

	userList := make([]UserInfo, 0, len(userInfoMap))

	validDays := len(statistics)

	for _, userInfo := range userInfoMap {
		if userInfo.Days == validDays {
			userList = append(userList, *userInfo)
		}
	}

	if len(userList) == 0 {
		return result
	}

	var maxSteps int
	for _, user := range userList {
		if user.Steps > maxSteps {
			maxSteps = user.Steps
		}
	}

	topSteps := maxSteps

	for _, user := range userList {
		if user.Steps == topSteps {
			result.UserIds = append(result.UserIds, user.UserID)
		}
	}

	result.Steps = topSteps

	return result
}
