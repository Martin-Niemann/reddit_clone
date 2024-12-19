import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime)

export function getShortenedDateString(date: Date) {
    if (dayjs(date).year() == dayjs(Date.now()).year()) {
        if (dayjs(date).month() == dayjs(Date.now()).month()) {
            if (dayjs(date).day() == dayjs(Date.now()).day()) {
                if (dayjs(date).hour() == dayjs(Date.now()).hour()) {
                    if (dayjs(date).minute() == dayjs(Date.now()).minute()) {
                        return dayjs(Date.now()).second() - dayjs(date).second() + "sec ago"
                    } else {
                        return dayjs(Date.now()).minute() - dayjs(date).minute() + "min ago"
                    }
                } else {
                    return dayjs(Date.now()).hour() - dayjs(date).hour() + "hr ago"
                }
            } else {
                return dayjs(Date.now()).day() - dayjs(date).day() + "d ago"
            }
        } else {
            return dayjs(Date.now()).month() - dayjs(date).month() + "mon ago"
        }
    } else {
        return dayjs(Date.now()).year() - dayjs(date).year() + "yr ago"
    }
}