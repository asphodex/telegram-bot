package raketa

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	 months = map[string]string{
		 "Января":   "January",
		 "Февраля":  "February",
		 "Марта":    "March",
		 "Апреля":   "April",
		 "Мая":      "May",
		 "Июня":     "June",
		 "Июля":     "July",
		 "Августа":  "August",
		 "Сентября": "September",
		 "Октября":  "October",
		 "Ноября":   "November",
		 "Декабря":  "December",
	}
)

var (
	currentYear = fmt.Sprintf("%d", time.Now().Year())
	client = http.Client{}
)

func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
}

// formatDate принимает на вход дату в формате '12 сентября 2024 11:04'
func formatDate(date string) (*time.Time, error) {
	dateParts := strings.Split(date, " ")
	if len(dateParts) != 4 {
		return nil, fmt.Errorf("invalid date format: %s", date)
	}

	month, ok := months[dateParts[1]]
	if !ok {
		return nil, fmt.Errorf("can't parse month: %s", date)
	}

	normalDate, err := time.Parse("2 January 2006 15:04",fmt.Sprintf("%s %s %s %s", dateParts[0], month, dateParts[2], dateParts[3]))
	if err != nil {
		return nil, err
	}

	return &normalDate, nil
}

func GetDeliveryTrace(track string) (*DeliveryTrace, error) {
	params := url.Values{}
	params.Add("option", comLkOption)
	params.Add("task", trackTask)
	params.Add("track", track)

	fullURL := fmt.Sprintf("%s?%s", raketaLk, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("raketa request creating failed: %v", err)
	}

	setHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("raketa request doing failed: %v", err)
	}

	defer res.Body.Close()


	body, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	var (
		deliveryTrace DeliveryTrace
		date *time.Time
	)

	body.Find(".delivery-step").Each(func(i int, s *goquery.Selection) {
		status := s.Find(".status-text").Text()
		dateText := s.Find(".date-text").Text()

		dateParts := strings.Split(dateText, currentYear)
		dateStr := strings.TrimSpace(dateParts[0])
		timeStr := strings.TrimSpace(dateParts[1])

		// Объединяем дату и время в одну строку
		dateTimeStr := fmt.Sprintf("%s %s %s", dateStr, currentYear, timeStr)

		date, err = formatDate(dateTimeStr)
		if err != nil {
			return
		}

		step := Step{
			Status: status,
			Date:   *date,
		}

		deliveryTrace.Step = append(deliveryTrace.Step, step)
	})

	if deliveryTrace.Step == nil {
		return nil, nil
	}

	return &deliveryTrace, nil
}

