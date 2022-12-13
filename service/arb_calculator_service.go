package service

import (
	"OISA_2x_sistem/telegram"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// TODO: make this take the message and the source msg, so it doesnt have to import telegram and bot api, just works the numbers

func ProvideArbCalculatorService(updateMessage *tgbotapi.Message) error {

	capitalStr := strings.TrimPrefix(updateMessage.Text, "/ulog ")
	capital, err := strconv.Atoi(capitalStr)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error converting capital to int")
	}

	reply, err := generateArbCalculatorReply(updateMessage.ReplyToMessage.Text, capital)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error generating message reply")
	}

	response := telegram.ReplyToMsg(reply, updateMessage.MessageID, updateMessage.Chat.ID)
	telegram.CheckIfSent(response,
		"Error back sending ulog, calculation's good just sending failed",
		strconv.FormatInt(updateMessage.Chat.ID, 10),
	)

	return nil
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
		"\nPROFIT: " + fmt.Sprintf("%.0f\n_sigurica_", (roi/100)*float64(capital))

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
