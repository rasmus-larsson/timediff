package timediff_test

import (
	"fmt"
	"testing"

	"github.com/mergestat/timediff"
)

var fixtures_sv_SE = map[string]string{
	"-10s":                            "några sekunder sedan",
	"-44s":                            "några sekunder sedan",
	"-45s":                            "en minut sedan",
	"-89s":                            "en minut sedan",
	"-90s":                            "2 minuter sedan",
	"-91s":                            "2 minuter sedan",
	"-2m":                             "2 minuter sedan",
	"-10m":                            "10 minuter sedan",
	"-44m":                            "44 minuter sedan",
	"-45m":                            "en timme sedan",
	"-60m":                            "en timme sedan",
	"-1h":                             "en timme sedan",
	"-80m":                            "en timme sedan",
	"-89m":                            "en timme sedan",
	"-90m":                            "2 timmar sedan",
	"-2h":                             "2 timmar sedan",
	"-20h":                            "20 timmar sedan",
	"-21h":                            "21 timmar sedan",
	"-21h30m":                         "en dag sedan",
	"-22h":                            "en dag sedan",
	"-24h":                            "en dag sedan",
	"-24h30m":                         "en dag sedan",
	"-34h59m":                         "en dag sedan",
	"-36h":                            "2 dagar sedan",
	fmt.Sprintf("-%dh", 10*24):        "10 dagar sedan",
	fmt.Sprintf("-%dh", 25*24):        "25 dagar sedan",
	fmt.Sprintf("-%dh", 26*24):        "en månad sedan",
	fmt.Sprintf("-%dh", 45*24):        "en månad sedan",
	fmt.Sprintf("-%dh2m", 45*24):      "2 månader sedan",
	fmt.Sprintf("-%dh", 46*24+1):      "2 månader sedan",
	fmt.Sprintf("-%dh", 80*24):        "3 månader sedan",
	fmt.Sprintf("-%dh", 9*24*30):      "9 månader sedan",
	fmt.Sprintf("-%dh", 10*24*30):     "10 månader sedan",
	fmt.Sprintf("-%dh1m", 10*24*30):   "ett år sedan",
	fmt.Sprintf("-%dh", 12*24*30):     "ett år sedan",
	fmt.Sprintf("-%dh", 17*24*30+1):   "2 år sedan",
	fmt.Sprintf("-%dh", 24*24*30):     "2 år sedan",
	fmt.Sprintf("-%dh", 20*24*30*12):  "20 år sedan",
	fmt.Sprintf("-%dh", 100*24*30*12): "100 år sedan",

	"10s":                            "om några sekunder",
	"44s":                            "om några sekunder",
	"45s":                            "om en minut",
	"89s":                            "om en minut",
	"90s":                            "om 2 minuter",
	"2m":                             "om 2 minuter",
	"10m":                            "om 10 minuter",
	"44m":                            "om 44 minuter",
	"45m":                            "om en timme",
	"60m":                            "om en timme",
	"1h":                             "om en timme",
	"80m":                            "om en timme",
	"89m":                            "om en timme",
	"89m10s":                         "om 2 timmar",
	"90m":                            "om 2 timmar",
	"2h":                             "om 2 timmar",
	"20h":                            "om 20 timmar",
	"21h":                            "om 21 timmar",
	"21h30m":                         "om en dag",
	"22h":                            "om en dag",
	"24h":                            "om en dag",
	"35h10m":                         "om 2 dagar",
	"36h":                            "om 2 dagar",
	fmt.Sprintf("%dh", 10*24):        "om 10 dagar",
	fmt.Sprintf("%dh", 25*24):        "om 25 dagar",
	fmt.Sprintf("%dh", 26*24):        "om en månad",
	fmt.Sprintf("%dh", 45*24):        "om en månad",
	fmt.Sprintf("%dh1m", 45*24):      "om 2 månader",
	fmt.Sprintf("%dh", 46*24):        "om 2 månader",
	fmt.Sprintf("%dh", 80*24):        "om 3 månader",
	fmt.Sprintf("%dh", 9*24*30):      "om 9 månader",
	fmt.Sprintf("%dh", 10*24*30):     "om 10 månader",
	fmt.Sprintf("%dh1m", 10*24*30):   "om ett år",
	fmt.Sprintf("%dh", 12*24*30):     "om ett år",
	fmt.Sprintf("%dh", 24*24*30):     "om 2 år",
	fmt.Sprintf("%dh", 20*24*30*12):  "om 20 år",
	fmt.Sprintf("%dh", 100*24*30*12): "om 100 år",
}

// execFixtures is a helper function for executing tests against fixtures
func TestTimeDiffSvSE(t *testing.T) {
	execFixtures(t, fixtures_sv_SE, timediff.WithLocale("sv-SE"))
}