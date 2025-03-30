package formatter

import "time"

func EpochToKitchen(epoch int) string {
	return time.Unix(int64(epoch), 0).Format(time.Kitchen)
}
