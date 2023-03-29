package main

import (
	"log"
	"strings"
	"time"
	"math"
)



type Offers []struct {
	ID              ID                `json:"_id"`
	Title           string            `json:"title"`
	City            string            `json:"city"`
	MarkerIcon      string            `json:"marker_icon"`
	ExperienceLevel string            `json:"experience_level"`
	EmploymentTypes []EmploymentTypes `json:"employment_types"`
	Remote          bool              `json:"remote"`
	InternalID      string            `json:"internal_id"`
	Skills          []Skills          `json:"skills"`
	PublishedAt     PublishedAt       `json:"published_at"`
	ValidTo         ValidTo           `json:"valid_to"`
}
type ID struct {
	Oid string `json:"$oid"`
}
type Salary struct {
	From     int64    `json:"from_"`
	To       int64    `json:"to"`
	Currency string `json:"currency"`
}
type EmploymentTypes struct {
	Type   string `json:"type"`
	Salary Salary `json:"salary"`
}
type Skills struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}
type PublishedAt struct {
	Date time.Time `json:"$date"`
}
type ValidTo struct {
	Date time.Time `json:"$date"`
}

func (o Offers) Len(seniority string) (int) {
	var offers_len int = 0
	for i := 0; i < len(o); i++ {
		res := strings.EqualFold(o[i].ExperienceLevel, seniority) 
		if res == true {
			offers_len ++
		} else {
			log.Printf("seniority %s in offer is not %s", o[i].ExperienceLevel, seniority)
		}
	}
	return offers_len
}

func (o Offers) GetSalaryBySeniorityLvl(seniority string) (int64, int64, float32) {
	// var b2bSalaries []Salary
	var minSalaries []int64
	var maxSalaries []int64
	var min int64
	var max int64
	var avg float64

	for i := 0; i < len(o); i++ {
		res := strings.EqualFold(o[i].ExperienceLevel, seniority) 
		if res == true && o[i].EmploymentTypes != nil {
			var employmentTypes []EmploymentTypes = o[i].EmploymentTypes
			for i := 0; i < len(employmentTypes); i++ {
				res := strings.EqualFold(employmentTypes[i].Type, "b2b") 
				if res == true {
					minSalaries = append(minSalaries, employmentTypes[i].Salary.From)
					maxSalaries = append(maxSalaries, employmentTypes[i].Salary.To)
				}

			}
			
		}
	}
	log.Println(minSalaries)
	log.Println(minSalaries)
	log.Printf("liczba min: %v", len(minSalaries))
	log.Printf("liczba max: %v", len(maxSalaries))
	min = findMin(minSalaries)
	max = findMax(maxSalaries)
	avg =  float64(min+max) / 2.0
	res := math.Round(avg*100) / 100
	return min, max, float32(res)

}

func findMin(numbers []int64) int64 {
    if len(numbers) == 0 {
        return 0
    }

    min := numbers[0]
    for _, number := range numbers {
        if number < min {
            min = number
        }
    }

    return min
}

func findMax(numbers []int64) int64 {
    if len(numbers) == 0 {
        return 0
    }

    max := numbers[0]
    for _, number := range numbers {
        if number > max {
            max = number
        }
    }

    return max
}






