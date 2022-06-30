package main

import "testing"

type lineStripTest struct {
	input, expected string
}

var lineStripTests = []lineStripTest{
	{input: "15:41:44 [rundeck@zt-ops-0 2][NORMAL] TASK [APPLY SQL - FAIL if SQL failed] ******************************************", expected: "TASK [APPLY SQL - FAIL if SQL failed] ******************************************"},
	{input: "15:48:00 [rundeck@zt-ops-0 2][NORMAL] fatal: [la-grant-so-prod-1]: FAILED! => {\"changed\": true, \"rc\": 4}", expected: "fatal: [la-grant-so-prod-1]: FAILED! => {\"changed\": true, \"rc\": 4}"},
}

func TestStripLineInfo(t *testing.T) {
	for test_num, test := range lineStripTests {
		if output := stripLineInfo(test.input); output != test.expected {
			t.Errorf("Test %d failed\nResult: %s\nExpected: %s", test_num, output, test.expected)
		}
	}
}
