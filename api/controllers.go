package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var weekDays = [7]string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

// *************** Controllers ***************

// /user POST
func UserRegistration(c *gin.Context) {
	user, err := ParseUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	// If provided username already exists, the
	// new user's creation is ignored.
	err = InsertUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.String(http.StatusCreated, "User created")
}

// /user/:name GET
func UserRetrieval(c *gin.Context) {
	user, err := GetUser(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// /user/ PUT
func UserModification(c *gin.Context) {
	user, err := ParseUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	err = UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.String(http.StatusOK, "User updated")
}

// /user/:name DELETE
func UserDeletion(c *gin.Context) {
	err := DeleteUser(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.String(http.StatusOK, "User deleted")
}

// /slots POST
func SlotCreation(c *gin.Context) {
	slot, err := ParseSlot(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	err = InsertSlot(slot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.String(http.StatusCreated, "Slot(s) created")
}

// /slots/:name GET
func SlotRetrieval(c *gin.Context) {
	slot, err := GetSlotUser(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"slot": slot})
}

// /slots/ GET
func SlotRetrievalAll(c *gin.Context) {
	slots, err := GetAllSlots()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"slots": slots})
}

// /slots/:name DELETE
// Deletes all slots of a user, could be improved
func SlotDeletion(c *gin.Context) {
	err := DeleteSlots(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Slots deleted")
}

// /schedule/ GET
func ScheduleRetrieval(c *gin.Context) {
	targets, err := ParseSchedule(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	slots, err := GetSchedule(targets)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"possible_slots": slots})
}

// *************** Auxiliary/Parsing ***************

func IsValidDay(day string) error {
	for _, d := range weekDays {
		if d == day {
			return nil
		}
	}
	return errors.New("not a valid day of the week")
}

func IsValidPeriod(start, end uint) error {
	if start > 24 || end > 24 {
		return errors.New("hours must be between 1 and 24")
	}
	if end-start < 1 {
		return errors.New("invalid time period")
	}
	return nil
}

// Validates and "sanitizes" the User Model
func ParseUser(c *gin.Context) (User, error) {
	var input User
	var user User
	// Initial sanitization by Gin
	if err := c.ShouldBindJSON(&input); err != nil {
		return user, err
	}
	// Business/BD specific
	if input.Type != "candidate" && input.Type != "interviewer" {
		return user, errors.New("type must be either candidate or interviewer")
	}
	user = User{Name: input.Name, Mail: strings.ToLower(input.Mail), Type: strings.ToLower(input.Type)}
	return user, nil
}

// Validates and "sanitizes" the Slot Model
func ParseSlot(c *gin.Context) (Slot, error) {
	var input Slot
	var slot Slot
	// Initial sanitization by Gin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return slot, err
	}
	// Business/BD specific
	if 1 > input.WeekNumber || input.WeekNumber > 53 {
		return slot, errors.New("week number must be between 1 and 53")
	}
	if err := IsValidDay(strings.ToLower(input.Day)); err != nil {
		return slot, err
	}
	if err := IsValidPeriod(input.HourStart, input.HourEnd); err != nil {
		return slot, err
	}
	slot = Slot{Name: input.Name, WeekNumber: input.WeekNumber,
		Day: strings.ToLower(input.Day), HourStart: input.HourStart, HourEnd: input.HourEnd}
	return slot, nil
}

func ParseSchedule(c *gin.Context) ([]string, error) {
	var input ScheduleInput
	// Initial sanitization by Gin
	if err := c.ShouldBindJSON(&input); err != nil {
		return nil, err
	}
	// Business/BD specific
	if t, err := GetUserType(input.Candidate); t != "candidate" || err != nil {
		if t != "candidate" {
			err = errors.New(input.Candidate + " is not a candidate")
		}
		return nil, err
	}
	s := strings.Split(input.Interviewers, ",")
	for _, i := range s {
		if t, err := GetUserType(i); t != "interviewer" || err != nil {
			if t != "interviewer" {
				err = errors.New(i + " is not an interviewer")
			}
			return nil, err
		}
	}
	targets := make([]string, len(s)+1)
	targets[0] = input.Candidate
	copy(targets[1:], s)

	return targets, nil
}
