package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleOutput = `=====> garbage field
=====> APP_NAME blah
       key:                      	  value
       int key:                       3
       boolean key:                   true
       really long key wow it is:     value
       empty value:       			  `

const exampleOutputWithTwoSections = `=====> APP_NAME blah
       key:                      	  value
       int key:                       3
       boolean key:                   true
       really long key wow it is:     value
       empty value:       			  
=====> SECOND_APP blah
       key:                      	  value
       int key:                       3
       boolean key:                   true
       really long key wow it is:     value
       empty value:       			  `

const exampleOutputWithMissingKeys = `=====> APP_NAME blah
       int key:                       3
       boolean key:                   true
       really long key wow it is:     value
       empty value:       			  `

type ExampleIndividualReport struct {
	Key      string `dokku:"key"`
	IntKey   int    `dokku:"int key"`
	BoolKey  bool   `dokku:"boolean key"`
	LongKey  string `dokku:"really long key wow it is"`
	EmptyVal string `dokku:"empty value"`
}

func TestParseIndividualReport(t *testing.T) {
	report := ExampleIndividualReport{}
	assert.NoError(t, ParseInto(exampleOutput, &report))
	assert.Equal(t, "value", report.Key)
	assert.Equal(t, 3, report.IntKey)
	assert.Equal(t, true, report.BoolKey)
	assert.Equal(t, "value", report.LongKey)
	assert.Empty(t, report.EmptyVal)
}

type ExampleReport map[string]ExampleIndividualReport

func TestParseReport(t *testing.T) {
	report := ExampleReport{}
	assert.NoError(t, ParseIntoMap(exampleOutputWithTwoSections, &report))
	assert.Contains(t, report, "APP_NAME")

	appReport, _ := report["APP_NAME"]
	assert.Equal(t, "value", appReport.Key)
	assert.Equal(t, 3, appReport.IntKey)
	assert.Equal(t, true, appReport.BoolKey)
	assert.Equal(t, "value", appReport.LongKey)
	assert.Empty(t, appReport.EmptyVal)

	assert.Error(t, ParseInto(exampleOutputWithMissingKeys, &report))
}
