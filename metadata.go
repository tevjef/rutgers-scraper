package main

import "time"

func getRutgers(campus string) University {
	university := University{
		Name:             "Rutgers University–New Brunswick",
		Abbr:             "RU-NB",
		HomePage:         "http://newbrunswick.edu/",
		RegistrationPage: "https://sims.edu/webreg/",
		Registrations: []*Registration{
			{
				Period:     InFall.String(),
				PeriodDate: time.Date(2000, time.September, 6, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     InSpring.String(),
				PeriodDate: time.Date(2000, time.January, 17, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     InSummer.String(),
				PeriodDate: time.Date(2000, time.May, 30, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     InWinter.String(),
				PeriodDate: time.Date(2000, time.December, 23, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     StartFall.String(),
				PeriodDate: time.Date(2000, time.March, 20, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     StartSpring.String(),
				PeriodDate: time.Date(2000, time.October, 5, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     StartSummer.String(),
				PeriodDate: time.Date(2000, time.January, 14, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     StartWinter.String(),
				PeriodDate: time.Date(2000, time.September, 21, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     EndFall.String(),
				PeriodDate: time.Date(2000, time.September, 13, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     EndSpring.String(),
				PeriodDate: time.Date(2000, time.January, 27, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     EndSummer.String(),
				PeriodDate: time.Date(2000, time.June, 15, 0, 0, 0, 0, time.UTC).Unix(),
			},
			{
				Period:     EndWinter.String(),
				PeriodDate: time.Date(2000, time.December, 22, 0, 0, 0, 0, time.UTC).Unix(),
			},
		},
		Metadata: []*Metadata{{
			Title: "About", Content: aboutNewBrunswick,
		},
		},
	}

	if campus == "NK" {
		university.Name = "Rutgers University–Newark"
		university.Abbr = "RU-NK"
		university.HomePage = "http://www.newark.edu/"
		university.Metadata = []*Metadata{
			{
				Title: "About", Content: aboutNewark,
			},
		}
	}

	if campus == "CM" {
		university.Name = "Rutgers University–Camden"
		university.Abbr = "RU-CAM"
		university.HomePage = "http://www.camden.edu/"
		university.Metadata = []*Metadata{
			{
				Title: "About", Content: aboutCamden,
			},
		}
	}

	return university
}

const (
	aboutNewBrunswick = ``

	aboutNewark = ``

	aboutCamden = ``
)
