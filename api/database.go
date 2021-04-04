package api

import (
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDataBase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("Database connection failed: + " + err.Error())
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Slot{})

	db = database
	return db
}

func InsertUser(user User) error {
	return db.Create(&user).Error
}

// Fetches user by username. Could be improved by
// fetching through primary key, but some conversion mechanism
// client side would be needed.
func GetUser(name string) (User, error) {
	var user User
	err := db.Where("name = ?", name).First(&user).Error
	return user, err
}

func GetUserType(name string) (string, error) {
	var target User
	err := db.Where("name = ?", name).First(&target).Error
	return target.Type, err
}

func UpdateUser(new User) error {
	return db.Where("name = ?", new.Name).Updates(&new).Error
}

func DeleteUser(name string) error {
	return db.Where("name = ?", name).Delete(User{}).Error
}

func InsertSlot(slot Slot) error {
	return db.Create(&slot).Error
}

func GetSlotUser(name string) ([]Slot, error) {
	var slots []Slot
	err := db.Where("name = ?", name).Find(&slots).Error
	return slots, err
}

func GetAllSlots() ([]Slot, error) {
	var slots []Slot
	err := db.Find(&slots).Error
	return slots, err
}

func DeleteSlots(name string) error {
	return db.Where("name = ?", name).Delete(Slot{}).Error
}

// Sorts slots chronologically, via manual bubblesort
func SortSlots(slots []Slot) []Slot {
	// Translates string weekday to int
	getdayint := func(day string) int {
		for i, d := range weekDays {
			if d == day {
				return i
			}
		}
		return 0
	}
	// Swaps two consecutive slots
	swap := func(f []Slot, i int) {
		f[i], f[i+1] = f[i+1], f[i]
	}
	// Order Weeks
	for {
		swapped := false
		for i, e := range slots {
			if i < len(slots)-1 {
				if int(slots[i+1].WeekNumber) < int(e.WeekNumber) {
					swap(slots, i)
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	// Days
	for {
		swapped := false
		for i, e := range slots {
			if i < len(slots)-1 {
				if getdayint(slots[i+1].Day) < getdayint(e.Day) &&
					slots[i+1].WeekNumber == e.WeekNumber {
					swap(slots, i)
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	// Hours
	for {
		swapped := false
		for i, e := range slots {
			if i < len(slots)-1 {
				if getdayint(slots[i+1].Day) == getdayint(e.Day) &&
					slots[i+1].WeekNumber == e.WeekNumber &&
					slots[i+1].HourStart < e.HourStart {
					swap(slots, i)
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	return slots
}

// Splits a slot into 1-h slots
func SplitSlots(slot Slot) []Slot {
	n := int(slot.HourEnd - slot.HourStart)
	slots := make([]Slot, n)
	for i := 0; i < n; i++ {
		s := Slot{
			Name:       slot.Name,
			WeekNumber: slot.WeekNumber,
			Day:        slot.Day,
			HourStart:  uint(int(slot.HourStart) + i), // Quick hack since 0 isn't uint
			HourEnd:    uint(int(slot.HourStart) + i + 1)}
		slots[i] = s
	}
	return slots
}

// Lists all possible slots between a candidate and
// one or more interviewers.
func GetSchedule(targets []string) ([]Slot, error) {
	var err error
	schedule := make([]Slot, 100) // List of 1 hour slots, hardcoded len needed
	slotsCand, _ := GetSlotUser(targets[0])
	pos := 0 // To keep track where to store the new slots
	for i := 0; i < len(slotsCand); i++ {
		var s []Slot
		// Get fitting slots
		db.Where("day = ? AND hour_start <= ? AND hour_end >= ?", slotsCand[i].Day,
			slotsCand[i].HourStart,
			slotsCand[i].HourEnd).Find(&s, Slot{WeekNumber: slotsCand[i].WeekNumber})
		// Create new 1-h slots based on our candidate
		if len(s) == len(targets) {
			meetingSlot := Slot{
				Name:       strings.Join(targets, "-"),
				WeekNumber: slotsCand[i].WeekNumber,
				Day:        slotsCand[i].Day,
				HourStart:  slotsCand[i].HourStart,
				HourEnd:    slotsCand[i].HourEnd,
			}
			meetings := SplitSlots(meetingSlot)
			// Schedule is a fixed length array
			for ii := 0; ii < 100; ii++ {
				schedule[pos] = meetings[ii]
				pos++
				if ii == len(meetings)-1 {
					break
				}
			}
		}
	}
	return schedule[:pos], err
}
