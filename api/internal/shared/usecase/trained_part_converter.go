package usecase

func ConvertTrainedPartToTinyInt(trainedPart string) int {
	switch trainedPart {
	case "体の中心":
		return 1
	case "背骨":
		return 2
	case "顔":
		return 3
	case "首":
		return 4
	case "背中":
		return 5
	case "お腹":
		return 6
	case "肩":
		return 7
	case "二の腕":
		return 8
	case "腕（ひじ下）":
		return 9
	case "手":
		return 10
	case "おしり":
		return 11
	case "太もも":
		return 12
	case "内もも":
		return 13
	case "ふくらはぎ":
		return 14
	case "足":
		return 15
	default:
		return 0
	}
}

func ConvertTrainedPartToString(trainedPart int) string {
	switch trainedPart {
	case 1:
		return "体の中心"
	case 2:
		return "背骨"
	case 3:
		return "顔"
	case 4:
		return "首"
	case 5:
		return "背中"
	case 6:
		return "お腹"
	case 7:
		return "肩"
	case 8:
		return "二の腕"
	case 9:
		return "腕（ひじ下）"
	case 10:
		return "手"
	case 11:
		return "おしり"
	case 12:
		return "太もも"
	case 13:
		return "内もも"
	case 14:
		return "ふくらはぎ"
	case 15:
		return "足"
	default:
		return "不明"
	}

}
