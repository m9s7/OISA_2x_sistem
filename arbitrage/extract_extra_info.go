package arbitrage

import (
	"fmt"
	"strings"
)

func (a Arb) getTip1DeviationTable() string {

	header := make([]string, len(a.TipValueLabels))
	body := make([]string, len(a.Tip1Vals))

	for i := range a.TipValueLabels {
		header[i] = " " + a.TipValueLabels[i] + " "
		body[i] = fmt.Sprintf(" %4.2f ", a.Tip1Vals[i])
	}
	labelsString := strings.Join(header, "|")
	valuesString := strings.Join(body, "|")
	arrowString := strings.Repeat(" ", a.Tip1MaxIdx*7) + " ^^^^"

	return strings.Join([]string{labelsString, valuesString, arrowString}, "\n")
}

func (a Arb) getTip2DeviationTable() string {

	header := make([]string, len(a.TipValueLabels))
	body := make([]string, len(a.Tip2Vals))

	for i := range a.TipValueLabels {
		header[i] = " " + a.TipValueLabels[i] + " "
		body[i] = fmt.Sprintf(" %4.2f ", a.Tip2Vals[i])
	}
	labelsString := strings.Join(header, "|")
	valuesString := strings.Join(body, "|")
	arrowString := strings.Repeat(" ", a.Tip2MaxIdx*7) + " ^^^^"

	return strings.Join([]string{labelsString, valuesString, arrowString}, "\n")
}
