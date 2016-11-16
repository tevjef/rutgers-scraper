package main

import (
	"strings"
	"time"
)

type (
	University struct {
		Name               string            `json:"name"`
		Abbr               string            `json:"abbr"`
		HomePage           string            `json:"home_page"`
		RegistrationPage   string            `json:"registration_page"`
		ResolvedSemesters  *ResolvedSemester `json:"resolved_semesters,omitempty"`
		Subjects           []*Subject        `json:"subjects,omitempty"`
		AvailableSemesters []*Semester       `json:"available_semesters,omitempty"`
		Registrations      []*Registration   `json:"registrations,omitempty"`
		Metadata           []*Metadata       `json:"metadata,omitempty"`
	}

	Subject struct {
		Name     string      `json:"name"`
		Number   string      `json:"number"`
		Season   string      `json:"season"`
		Year     string      `json:"year"`
		Courses  []*Course   `json:"courses,omitempty"`
		Metadata []*Metadata `json:"metadata,omitempty"`
	}

	Course struct {
		Name     string      `json:"name"`
		Number   string      `json:"number"`
		Synopsis string      `json:"synopsis,omitempty"`
		Sections []*Section  `json:"sections,omitempty"`
		Metadata []*Metadata `json:"metadata,omitempty"`
	}

	Section struct {
		Number      string        `json:"number"`
		IndexNumber string        `json:"index_number"`
		Status      string        `json:"status"`
		Credits     string        `json:"credits"`
		Meetings    []*Meeting    `json:"meetings,omitempty"`
		Instructors []*Instructor `json:"instructors,omitempty"`
		Books       []*Book       `json:"books,omitempty"`
		Metadata    []*Metadata   `json:"metadata,omitempty"`
	}

	Meeting struct {
		Room      string      `json:"room,omitempty"`
		Day       string      `json:"day,omitempty"`
		StartTime string      `json:"start_time,omitempty"`
		EndTime   string      `json:"end_time,omitempty"`
		ClassType string      `json:"class_type,omitempty"`
		Metadata  []*Metadata `json:"metadata,omitempty"`
	}

	Instructor struct {
		Name string `json:"name" db:"name"`
	}

	Book struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	Metadata struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	Registration struct {
		Period     string `json:"period"`
		PeriodDate int64  `json:"period_date"`
	}

	ResolvedSemester struct {
		Current *Semester `json:"current,omitempty"`
		Last    *Semester `json:"last,omitempty"`
		Next    *Semester `json:"next,omitempty"`
	}

	Semester struct {
		Year   int    `json:"year"`
		Season string `json:"season"`
	}
)

type (
	Period int
	Status int
)

const (
	InFall Period = iota
	InSpring
	InSummer
	InWinter
	StartFall
	StartSpring
	StartSummer
	StartWinter
	EndFall
	EndSpring
	EndSummer
	EndWinter
)

var period = [...]string{
	"fall",
	"spring",
	"summer",
	"winter",
	"start_fall",
	"start_spring",
	"start_summer",
	"start_winter",
	"end_fall",
	"end_spring",
	"end_summer",
	"end_winter",
}

func (s Period) String() string {
	return period[s]
}

const (
	Fall = "fall"
	Spring = "spring"
	Summer = "summer"
	Winter = "winter"
)

const (
	OPEN Status = 1 + iota
	CLOSED
)

var status = [...]string{
	"Open",
	"Closed",
}

func (s Status) String() string {
	return status[s-1]
}

func toTitle(str string) string {
	str = strings.Title(strings.ToLower(str))

	for i := len(str) - 1; i != 0; i-- {
		if strings.LastIndex(str, "i") == i {
			str = swapChar(str, "I", i)
		} else {
			break
		}
	}
	return str
}

func swapChar(s, char string, index int) string {
	left := s[:index]
	right := s[index+1:]
	return left + char + right
}

var trim = strings.TrimSpace

func (r Registration) month() time.Month {
	return time.Unix(r.PeriodDate, 0).UTC().Month()
}

func (r Registration) day() int {
	return time.Unix(r.PeriodDate, 0).UTC().Day()
}

func (r Registration) dayOfYear() int {
	return time.Unix(r.PeriodDate, 0).UTC().YearDay()
}

func (r Registration) season() string {
	switch r.Period {
	case InFall.String():
		return Fall
	case InSpring.String():
		return Spring
	case InSummer.String():
		return Summer
	case InWinter.String():
		return Winter
	default:
		return Summer
	}
}

func ResolveSemesters(t time.Time, registration []*Registration) *ResolvedSemester {
	month := t.Month()
	day := t.Day()
	yearDay := t.YearDay()

	var winterReg = registration[InWinter]
	var startFallReg = registration[StartFall]
	var startSpringReg = registration[StartSpring]
	var endSummerReg = registration[EndSummer]

	fall := &Semester{
		Year:   t.Year(),
		Season: Fall}

	winter := &Semester{
		Year:   t.Year(),
		Season: Winter}

	spring := &Semester{
		Year:   t.Year(),
		Season: Spring}

	summer := &Semester{
		Year:   t.Year(),
		Season: Summer}

	// Spring: Winter - StartFall
	if (month >= winterReg.month() && day >= winterReg.day()) || (month <= startFallReg.month() && day < startFallReg.day()) {
		if winterReg.month()-month <= 0 {
			spring.Year++
			summer.Year++
		} else {
			winter.Year--
			fall.Year--
		}
		return &ResolvedSemester{
			Last:    winter,
			Current: spring,
			Next:    summer}

		//StartFall: StartFall -- EndSummer
	} else if yearDay >= startFallReg.dayOfYear() && yearDay < endSummerReg.dayOfYear() {
		return &ResolvedSemester{
			Last:    spring,
			Current: summer,
			Next:    fall,
		}
		//Fall: EndSummer -- StartSpring
	} else if yearDay >= endSummerReg.dayOfYear() && yearDay < startSpringReg.dayOfYear() {
		return &ResolvedSemester{
			Last:    summer,
			Current: fall,
			Next:    winter,
		}
		//StartSpring: StartSpring -- Winter
	} else if yearDay >= startSpringReg.dayOfYear() && yearDay < winterReg.dayOfYear() {
		spring.Year++
		return &ResolvedSemester{
			Last:    fall,
			Current: winter,
			Next:    spring,
		}
	}

	return &ResolvedSemester{}
}
