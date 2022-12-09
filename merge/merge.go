package merge

import (
	"OISA_2x_sistem/utility"
	"encoding/csv"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	fuzzy "github.com/paul-mannino/go-fuzzywuzzy"
	"log"
	"os"
	"strconv"
	"time"
)

type bookie struct {
	name string
	rows []*[8]string
}

func Merge(sportName string, data map[string][]*[8]string) [][]string {

	bookies := mapToSlice(data)

	if len(bookies) < 2 {
		fmt.Println("... nothing to merge -" + sportName)
		fmt.Println("Books passed in:", bookies)
		return nil
	}

	startTime := time.Now()
	fmt.Println("... merging scraped data - " + sportName)

	orderBooksByNumOfRecords(bookies)

	mergedRecords := [][]string{getMergedRecordsColumnNames(bookies)}
	mergedRecordsColIndxMap := getColumnIndexes(len(bookies))

	// Merge
	successfulMatches := 0
	for _, el1 := range bookies[0].rows {

		recordToMerge := initRecordWithEl1(el1, mergedRecordsColIndxMap, len(bookies))
		doAddRecordToMerged := false

		for bookieOrder := 1; bookieOrder < len(bookies); bookieOrder++ {
			for _, el2 := range bookies[bookieOrder].rows {

				// check if tip_names match
				if el1[utility.Tip1Name] != el2[utility.Tip1Name] &&
					el1[utility.Tip2Name] != el2[utility.Tip2Name] {
					continue
				}

				// check if kickoff times are similar
				t1, _ := strconv.ParseInt(el1[utility.Kickoff], 10, 64)
				t2, _ := strconv.ParseInt(el2[utility.Kickoff], 10, 64)
				oneHour := int64(3600)
				if utility.Abs(t1-t2) > oneHour {
					continue
				}

				// check if league numbers match
				if !isSameLeagueNum(el1[utility.League], el2[utility.League]) {
					continue
				}

				if sportName == utility.Soccer {
					if fuzzy.Ratio(el1[utility.Team1], el2[utility.Team1]) < 80 {
						continue
					}
					if fuzzy.Ratio(el1[utility.Team2], el2[utility.Team2]) < 80 {
						continue
					}

					addElToRecord(el2, bookieOrder, &recordToMerge, mergedRecordsColIndxMap, false)
					doAddRecordToMerged = true
					successfulMatches++
					continue
				}

				if fuzzy.Ratio(el1[utility.Team1], el2[utility.Team1]) >= 80 &&
					fuzzy.Ratio(el1[utility.Team2], el2[utility.Team2]) >= 80 {

					addElToRecord(el2, bookieOrder, &recordToMerge, mergedRecordsColIndxMap, false)
					doAddRecordToMerged = true
					successfulMatches++
					continue
				}
				if fuzzy.Ratio(el1[utility.Team1], el2[utility.Team2]) >= 80 &&
					fuzzy.Ratio(el1[utility.Team2], el2[utility.Team1]) >= 80 {

					addElToRecord(el2, bookieOrder, &recordToMerge, mergedRecordsColIndxMap,
						shouldSwitchTipVals(el1[utility.Tip1Name], sportName))
					doAddRecordToMerged = true
					successfulMatches++
					continue
				}

			}
		}
		if doAddRecordToMerged {
			mergedRecords = append(mergedRecords, recordToMerge)
		}
	}
	writeRecordsToCSV(mergedRecords, sportName)

	for _, bookie := range bookies {
		fmt.Println(bookie.name, ": ", len(bookie.rows))
	}
	fmt.Println("Successfully merged:", successfulMatches, "records")
	fmt.Println("--- ", time.Now().Sub(startTime), " ---")

	return mergedRecords
}

func writeRecordsToCSV(mergedRecords [][]string, sport string) {
	f, err := os.Create("C:\\Users\\Matija\\GolandProjects\\OISA_2x_sistem\\IO\\merged_records\\" + sport + ".csv")
	if err != nil {
		log.Println("Failed to open mergedRecords CSV file", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("Failed to close mergedRecords CSV file", err)
		}
	}(f)

	w := csv.NewWriter(f)
	for _, record := range mergedRecords {
		if err := w.Write(record); err != nil {
			log.Println("Failed to write record to mergedRecords CSV file", err)
		}
	}
	w.Flush()
}

func IsEmpty(data [][]string) bool {
	return data == nil || len(data) == 1
}

func PrintMergedData(data [][]string) {
	fmt.Println(dataframe.LoadRecords(data).Drop([]int{2, 3, 4}).String())
}
