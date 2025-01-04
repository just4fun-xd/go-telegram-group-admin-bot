package config

import "time"

var (
	SubscriptionWeek  = 7 * 24 * time.Hour
	SubscriptionMonth = 30 * 24 * time.Hour
	SubscriptionYear  = 365 * 24 * time.Hour
)
