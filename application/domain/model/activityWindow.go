package model

import (
	"errors"
	"sort"
	"time"
)

type ActivityWindow struct {
	Activities []Activity
}

func (aw ActivityWindow) GetStartTimeStamp() (time.Time, error) {
	if len(aw.Activities) == 0 {
		return time.Time{}, errors.New("no activities found")
	}

	sort.Slice(aw.Activities, func(i, j int) bool {
		return aw.Activities[i].Timestamp.Before(aw.Activities[j].Timestamp)
	})

	return aw.Activities[0].Timestamp, nil
}

func (aw ActivityWindow) GetEndTimeStamp() (time.Time, error) {
	if len(aw.Activities) == 0 {
		return time.Time{}, errors.New("no activities found")
	}

	sort.Slice(aw.Activities, func(i, j int) bool {
		return aw.Activities[i].Timestamp.After(aw.Activities[j].Timestamp)
	})

	return aw.Activities[0].Timestamp, nil
}

func (aw ActivityWindow) CalculateBalance(accountId AccountId) Money {
	depositBalance := ZERO
	withdrawalBalance := ZERO

	for _, activity := range aw.Activities {
		if activity.TargetAccountId == accountId {
			depositBalance = depositBalance.Plus(activity.Money)
		}

		if activity.SourceAccountId == accountId {
			withdrawalBalance = withdrawalBalance.Plus(activity.Money)
		}
	}

	return depositBalance.Plus(withdrawalBalance.Negate())
}

func NewActivityWindow(activities []Activity) ActivityWindow {
	return ActivityWindow{Activities: activities}
}

func NewActivityWindowFromSlice(activities ...Activity) ActivityWindow {
	return ActivityWindow{Activities: activities}
}

func (aw ActivityWindow) GetActivities() []Activity {
	activitiesCopy := make([]Activity, len(aw.Activities))
	copy(activitiesCopy, aw.Activities)
	return activitiesCopy
}

func (aw *ActivityWindow) AddActivity(activity Activity) {
	aw.Activities = append(aw.Activities, activity)
}
