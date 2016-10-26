package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var rutgers = []*Registration{
	{
		Period:     InFall.String(),
		PeriodDate: time.Date(0000, time.September, 6, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     InSpring.String(),
		PeriodDate: time.Date(0000, time.January, 17, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     InSummer.String(),
		PeriodDate: time.Date(0000, time.May, 30, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     InWinter.String(),
		PeriodDate: time.Date(0000, time.December, 23, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     StartFall.String(),
		PeriodDate: time.Date(0000, time.March, 20, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     StartSpring.String(),
		PeriodDate: time.Date(0000, time.October, 18, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     StartSummer.String(),
		PeriodDate: time.Date(0000, time.January, 14, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     StartWinter.String(),
		PeriodDate: time.Date(0000, time.September, 21, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     EndFall.String(),
		PeriodDate: time.Date(0000, time.September, 13, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     EndSpring.String(),
		PeriodDate: time.Date(0000, time.January, 27, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     EndSummer.String(),
		PeriodDate: time.Date(0000, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(),
	},
	{
		Period:     EndWinter.String(),
		PeriodDate: time.Date(0000, time.December, 22, 0, 0, 0, 0, time.UTC).Unix(),
	}}

func TestResolveSemesters(t *testing.T) {
	var semesters *ResolvedSemester

	//semesters := ResolveSemesters(time.Now(), rutgers)
	//assert.Equal(t, 2016, int(semesters.Last.Year))
	//assert.Equal(t, SUMMER, semesters.Last.Season)
	//assert.Equal(t, 2016, int(semesters.Current.Year))
	//assert.Equal(t, FALL, semesters.Current.Season)
	//assert.Equal(t, 2016, int(semesters.Next.Year))
	//assert.Equal(t, WINTER, semesters.Next.Season)

	semesters = ResolveSemesters(time.Date(2015, time.December, 24, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2015, int(semesters.Last.Year))
	assert.Equal(t, WINTER, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, SPRING, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, SUMMER, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2015, int(semesters.Last.Year))
	assert.Equal(t, WINTER, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, SPRING, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, SUMMER, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.March, 19, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2015, int(semesters.Last.Year))
	assert.Equal(t, WINTER, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, SPRING, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, SUMMER, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.March, 20, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, SPRING, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, SUMMER, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, FALL, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.April, 30, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, SPRING, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, SUMMER, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, FALL, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.June, 14, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, SPRING, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, SUMMER, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, FALL, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.June, 15, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, SUMMER, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, FALL, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, WINTER, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.September, 15, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, SUMMER, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, FALL, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, WINTER, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.October, 17, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, SUMMER, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, FALL, semesters.Current.Season)
	assert.Equal(t, 2016, int(semesters.Next.Year))
	assert.Equal(t, WINTER, semesters.Next.Season)

	semesters = ResolveSemesters(time.Date(2016, time.October, 18, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, FALL, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, WINTER, semesters.Current.Season)
	assert.Equal(t, 2017, int(semesters.Next.Year))
	assert.Equal(t, SPRING, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.November, 1, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, FALL, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, WINTER, semesters.Current.Season)
	assert.Equal(t, 2017, int(semesters.Next.Year))
	assert.Equal(t, SPRING, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

	semesters = ResolveSemesters(time.Date(2016, time.December, 22, 0, 0, 0, 0, time.UTC), rutgers)
	assert.Equal(t, 2016, int(semesters.Last.Year))
	assert.Equal(t, FALL, semesters.Last.Season)
	assert.Equal(t, 2016, int(semesters.Current.Year))
	assert.Equal(t, WINTER, semesters.Current.Season)
	assert.Equal(t, 2017, int(semesters.Next.Year))
	assert.Equal(t, SPRING, semesters.Next.Season)
	//fmt.Printf("%#v\n", semesters)

}
