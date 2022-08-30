package reports

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"regexp"
	"strings"
)

const (
	sectionStartDenom = "=====> "
	infoIndent        = "       "

	dokkuTagName = "dokku"
)

var (
	ErrInvalidReport = errors.New("invalid report")

	appNameRe = regexp.MustCompile(`^=====> (\S*)\s`)
	rowRe     = regexp.MustCompile(`^\s+([\s\w]*):(.*)$`)
)

type Report map[string]string
type ReportMap map[string]map[string]string

func rowPair(row string) (string, string) {
	matches := rowRe.FindStringSubmatch(row)
	if matches == nil || len(matches) < 3 {
		return "", ""
	}
	key := matches[1]
	val := matches[2]
	return key, strings.Trim(val, " \t")
}

func ParseSingle(sReport string) (map[string]string, error) {
	report := Report{}
	lines := strings.Split(sReport, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, sectionStartDenom) {
			appNameMatches := appNameRe.FindStringSubmatch(line)
			if len(appNameMatches) < 2 {
				return nil, errors.New("invalid report line: '" + line)
			}
			if len(lines) > i+1 && !strings.HasPrefix(lines[i+1], infoIndent) {
				continue
			}
		} else if strings.HasPrefix(line, infoIndent) {
			k, v := rowPair(line)
			if k != "" {
				report[k] = v
			}
		}
	}

	return report, nil
}

func ParseMultiple(rawReport string) (ReportMap, error) {
	report := ReportMap{}

	var currentAppName string
	var currentReportMap map[string]string

	lines := strings.Split(rawReport, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, sectionStartDenom) {
			appNameMatches := appNameRe.FindStringSubmatch(line)
			if len(appNameMatches) < 2 {
				return nil, errors.New("invalid report line: '" + line)
			}
			if len(lines) > i+1 && !strings.HasPrefix(lines[i+1], infoIndent) {
				continue
			}
			if currentAppName != "" {
				report[currentAppName] = currentReportMap
			}

			currentAppName = appNameMatches[1]
			currentReportMap = map[string]string{}
		} else if strings.HasPrefix(line, infoIndent) {
			k, v := rowPair(line)
			if k != "" {
				currentReportMap[k] = v
			}
		}
	}

	if currentAppName != "" {
		report[currentAppName] = currentReportMap
	}

	return report, nil
}

func ParseIntoMap(rawReport string, reportPtr interface{}) error {
	reportMaps, err := ParseMultiple(rawReport)
	if err != nil {
		fmt.Println("failed to parse report: ", err.Error())
		return err
	}

	// TODO: check reportPtr is a map, str->report
	// we expect "reports" to be a mapping of 'app name' -> report
	// we should export a struct with mapped fields, and automatically convert it

	// this gets kinda gross, but so be it
	reportVal := reflect.ValueOf(reportPtr)

	// Since reports are a map, we want elements of the type held by
	// the real element passed in reportPtr
	elemValType := reportVal.Elem().Type().Elem()

	for appName, reportMap := range reportMaps {
		// indirect to get the value held by a pointer to new map value
		// as an interface, so we can actually pass the data pointer
		appReport := reflect.Indirect(reflect.New(elemValType)).Interface()

		decoderCfg := &mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			Result:           &appReport,
			TagName:          dokkuTagName,
			ErrorUnused:      false,
			ErrorUnset:       false,
		}
		decoder, err := mapstructure.NewDecoder(decoderCfg)
		if err != nil {
			return errors.New("failed to create decoder: " + err.Error())
		}

		if err := decoder.Decode(reportMap); err != nil {
			return errors.New("failed to decode report map: " + err.Error())
		}

		k := reflect.ValueOf(appName)
		v := reflect.ValueOf(appReport)
		reflect.Indirect(reportVal).SetMapIndex(k, v)
	}

	return nil
}

func ParseInto(singleReport string, reportPtr interface{}) error {
	report, err := ParseSingle(singleReport)
	if err != nil {
		fmt.Println("failed to parse report: ", err.Error())
		return err
	}

	decoderCfg := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           reportPtr,
		TagName:          dokkuTagName,
		ErrorUnused:      true,
		ErrorUnset:       true,
	}
	decoder, err := mapstructure.NewDecoder(decoderCfg)
	if err != nil {
		return errors.New("failed to create decoder: " + err.Error())
	}

	if err := decoder.Decode(report); err != nil {
		return errors.New("failed to decode report map: " + err.Error())
	}

	return nil
}
