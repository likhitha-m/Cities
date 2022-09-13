package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	logger "github.com/sirupsen/logrus"
)




func GenerateContNumber(n int, length int) (string, error) {
	s := strconv.Itoa(n)
	l := len(s)
	if len(s) < length {
		s = ("00000000" + s)

		s = s[l:]
	}
	fmt.Println(s)
	return s, nil
}
func GenerateRandomUniqueContNumber(n int, length int) (string, error) {
	s := strconv.Itoa(n)
	l := len(s)
	if len(s) < length {
		s = ("000000000000000" + s)

		s = s[l:]
	}
	return s, nil
}

// Change decimal with fixed .00 precision
func Decimal64p2(val float64) float64 {
	twoDigit, err := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	if err != nil {
		logger.Error("func_Decimal64p2: Error in change decimal with fixed: ", err)
	}
	return twoDigit
}

func DecimalPrecision2(no float64) float64 {
	return math.Round(no*100) / 100
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func BoolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

func GetStringify(data interface{}) (string, error) {
	out, err := json.Marshal(data)
	if err != nil {
		logger.Error("func_GetStringify: Error in marshal ", err)
		return "", err
	}
	return string(out), nil
}

func ToUpperCase(str string) string {
	return strings.ToUpper(str)
}

func CapFirstChar(s string) string {
	loweredVal := ToLowerCase(s)
	for index, value := range loweredVal {
		return string(unicode.ToUpper(value)) + loweredVal[index+1:]
	}
	return ""
}

func ToLowerCase(str string) string {
	return strings.ToLower(str)
}






func GetUTCDateTime(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		logger.Error("func_GetUTCDateTime: Error loading location: ", err)
		return t, err
	}
	utcDateTime := t.In(loc)
	return utcDateTime, nil
}

func GetDateTimeUsingLocation(locationName string) time.Time {
	// Start - Get time zone
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		logger.Error("func_GetDateTimeUsingLocation: Error while loading location timezone , error: ", err)
		// return err
	}
	// End - Get time zone

	_, offset := time.Now().In(loc).Zone()
	offSetInHour := offset / 3600
	log.Println(offset)
	log.Println(offSetInHour)

	estDuration, _ := time.ParseDuration(strconv.Itoa(offSetInHour) + "h")
	estDateTime := time.Now().Add(estDuration).UTC()
	return estDateTime
}

func GetEstDateTimePR(t time.Time) (time.Time, error) {

	loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		logger.Error("func_GetEstDateTimePR: Error while loading location timezone , error: ", err)
		// return err
	}
	// End - Get time zone

	_, offset := t.In(loc).Zone()
	//	fmt.Println("_,offset>>>>>>>>>>>>>>>>>", n, offset)
	offSetInHour := offset / 3600
	log.Println(offset)
	log.Println(offSetInHour)

	estDuration, err := time.ParseDuration(strconv.Itoa(offSetInHour) + "h")
	if err != nil {
		logger.Error("func_GetEstDateTimePR: Error loading location: ", err)
		return t, err
	}
	estDateTime := t.Add(estDuration)
	//	fmt.Println("est>>>>>>>>>>>>>>>>>", estDateTime)
	//	fmt.Println("nor>>>>>>>>>>>>>>>>>", t)
	return estDateTime, nil
}

func GetEstDateTimeAddPR(t time.Time) (time.Time, error) {

	loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		logger.Error("func_GetEstDateTimePR: Error while loading location timezone , error: ", err)
		// return err
	}
	// End - Get time zone

	_, offset := t.In(loc).Zone()
	//	fmt.Println("_,offset>>>>>>>>>>>>>>>>>", n, offset)
	offSetInHour := offset / 3600
	log.Println(offset)
	log.Println(offSetInHour)

	estDuration, err := time.ParseDuration(strconv.Itoa(int(math.Abs(float64(offSetInHour)))) + "h")
	if err != nil {
		logger.Error("func_GetEstDateTimePR: Error loading location: ", err)
		return t, err
	}
	estDateTime := t.Add(estDuration)
	//	fmt.Println("est>>>>>>>>>>>>>>>>>", estDateTime)
	//	fmt.Println("nor>>>>>>>>>>>>>>>>>", t)
	return estDateTime, nil
}

func GetDuration(m string) (time.Duration, error) {
	var t time.Duration
	min, err := strconv.Atoi(m)
	if err != nil {
		logger.Error("func_GetDuration: ", err)
		return t, err
	}
	t = time.Duration(min)
	return t, nil
}


func encodeBase64cr(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func decodeBase64cr(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func AppendZeros(n int, length int) (string, error) {
	s := strconv.Itoa(n)
	l := len(s)
	if len(s) < length {
		s = ("000000000" + s)

		s = s[l:]
	}
	fmt.Println(s)
	return s, nil
}

func SliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// func ESTstringUTCTime(value string) (*time.Time, error) {
// 	location, err := time.LoadLocation("EST")
// 	if err != nil {
// 		return nil, err
// 	}
// 	timeFromRequest := value // Fare Quote , EST

// 	timeLayout := "2006-01-02T15:04:05"

// 	estTime, err := time.ParseInLocation(timeLayout, timeFromRequest, location)
// 	if err != nil {
// 		return nil, err
// 	}
// 	loc, err := time.LoadLocation("UTC")
// 	if err != nil {
// 		return nil, err
// 	}
// 	UTC := estTime.In(loc)

// 	return &UTC, nil
// }

// This function accept this EST value with format = "2022-03-18T13:24:45"
// Convert est string to Est time
// using time_zone fetching Offset
// Adding offset to est time
// thne finally converting to UTC
// func ESTstringUTCTime(estTimeString string) (*time.Time, error) {
// 	location, err := time.LoadLocation("EST")
// 	if err != nil {
// 		return nil, err
// 	}

// 	estTime, err := time.ParseInLocation("2006-01-02T15:04:05", estTimeString, location)
// 	if err != nil {
// 		return nil, err
// 	}

// 	timeZone, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, offset := estTime.In(timeZone).Zone()
// 	offSetInHour := offset / 3600
// 	estDuration, err := time.ParseDuration(strconv.Itoa(offSetInHour*(-1)) + "h")
// 	if err != nil {
// 		return nil, err
// 	}
// 	utcDate := estTime.Add(estDuration)

// 	loc, err := time.LoadLocation("UTC")
// 	if err != nil {
// 		return nil, err
// 	}
// 	UTC := utcDate.In(loc)

// 	return &UTC, nil
// }

func ESTstringUTCTime(estTimeString string) (*time.Time, error) {

	timeZone, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		return nil, err
	}

	estTime, err := time.ParseInLocation("2006-01-02T15:04:05", estTimeString, timeZone)
	if err != nil {
		return nil, err
	}
	UTC := estTime.UTC()

	return &UTC, nil
}
