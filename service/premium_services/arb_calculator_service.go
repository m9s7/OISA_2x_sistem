package premium_services

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ProvideArbCalculatorService(arbString string, userReply string) (string, error) {

	capitalStr := strings.TrimLeft(userReply, "/UuLlOoGg ")

	capital, err := strconv.Atoi(capitalStr)
	if err != nil {
		return "", err
	}

	reply, err := generateArbCalculatorReply(arbString, capital)
	if err != nil {
		return "", err
	}

	return reply, nil
}

func arbCalc(capital int, tip1StakePercentage float64, tip2StakePercentage float64) (string, string) {
	tip1Stake := fmt.Sprintf("%.0f", tip1StakePercentage*float64(capital))
	tip2Stake := fmt.Sprintf("%.0f", tip2StakePercentage*float64(capital))
	return tip1Stake, tip2Stake
}

func generateArbCalculatorReply(arbString string, capital int) (string, error) {

	bookie1, bookie2, err := extractBookies(arbString)
	if err != nil {
		return "", err
	}

	roi, err := extractROI(arbString)
	if err != nil {
		return "", err
	}

	stakePercentage1, stakePercentage2, err := extractStakePercentages(arbString)
	if err != nil {
		return "", err
	}

	t1Investment, t2Investment := arbCalc(capital, stakePercentage1, stakePercentage2)
	if err != nil {
		return "", err
	}

	responseMsg := " ulog: \n" +
		t1Investment + " @" + strings.ToLower(bookie1) + "\n" +
		t2Investment + " @" + strings.ToLower(bookie2) + "\n" +
		"\nPROFIT: " + fmt.Sprintf("%.0f\n_sigurica ca ca..._", (roi/100)*float64(capital))

	return responseMsg, nil

}

func extractStakePercentages(arbString string) (float64, float64, error) {

	re := regexp.MustCompile(`ukupno \* (\d+\.\d+)\n`)
	stakePercentages := re.FindAllStringSubmatch(arbString, -1)
	if stakePercentages == nil {
		return 0, 0, errors.New("failed to extract stake percentages")
	}
	stakePercentage1, err := strconv.ParseFloat(stakePercentages[0][1], 64)
	if err != nil {
		return 0, 0, errors.New("failed to parse stake percentage 1")
	}
	stakePercentage2, err := strconv.ParseFloat(stakePercentages[1][1], 64)
	if err != nil {
		return 0, 0, errors.New("failed to parse stake percentage 2")
	}

	return stakePercentage1, stakePercentage2, nil
}

func extractROI(arbString string) (float64, error) {

	re := regexp.MustCompile(`ROI: (\d+\.\d+)%`)
	roi := re.FindStringSubmatch(arbString)
	if roi == nil {
		return 0, errors.New("failed to extract roi")
	}
	roiValue, err := strconv.ParseFloat(roi[1], 64)
	if err != nil {
		return 0, errors.New("failed to parse roi")
	}

	return roiValue, nil
}

func extractBookies(arbString string) (string, string, error) {

	re := regexp.MustCompile(`@ ([A-Z]+?)\n`)
	extractedBookies := re.FindAllStringSubmatch(arbString, -1)
	if extractedBookies == nil {
		return "", "", errors.New("failed to extract bookies")
	}

	var bookies []string
	for _, value := range extractedBookies {
		bookies = append(bookies, value[1])
	}

	if len(bookies) != 2 {
		return "", "", errors.New("failed to extract bookies")
	}

	return bookies[0], bookies[1], nil
}
