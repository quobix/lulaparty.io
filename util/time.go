package util

import (
        "fmt"
        "time"
)

func Pluralize(val, pl string, num int) string {
        if(num>1) { return val + pl }
        return val
}

type TimeDiffClosure func(t time.Time) bool

func TimeDiffFormatter(h,m,s,d,mo,y int) string {
        var pl = "s"
        h1 := h-(d*24)
        if (y <= 0 && mo >= 1) {
                return fmt.Sprintf("%d " + Pluralize("Month", pl, mo), mo)
        }
        if (y <= 0 && d >= 1 && h1 <=0) {
                return fmt.Sprintf("%d " + Pluralize("Day", pl, d), d)
        }

        if (y <= 0 && d >= 1 && h1>=1) {
                return fmt.Sprintf("%d " + Pluralize("Day", pl, d) + ", %d " +
                Pluralize("Hour", pl,h1), d, h1)
        }
        if (y <= 0 && d <= 0 && h >= 1 && m <= 0) {
                return fmt.Sprintf("%d " + Pluralize("Hour", pl, h), h)
        }

        if (y <= 0 && d <= 0 && h >= 1 ) {
                return fmt.Sprintf("%d " + Pluralize("Hour", pl, h) + ", %d " +
                Pluralize("Min", pl, m), h, m)
        }

        if (y <= 0 && h <= 0 && m >= 1 && s >= 1) {
                return fmt.Sprintf("%d " + Pluralize("Minute", pl, m) + ", %d " +
                Pluralize("Second", pl, s), m, s)
        }

        if (y <= 0 && h <= 0 && m >= 1) {
                return fmt.Sprintf("%d " + Pluralize("Minute", pl, m), m)
        }

        if (y <= 0 && m <= 0 && s >= 1) {
                return fmt.Sprintf("%d " + Pluralize("Second", pl, s), s)
        }
        return fmt.Sprintf("%d " + Pluralize("Year", pl, y), y)
}

func TimeDiffHelper(diff time.Time, def string, m TimeDiffClosure) string {

        t := time.Now().UTC()
        d := t.Sub(diff)
        seconds := d.Seconds()
        hours := d.Hours()
        minutes := d.Minutes()

        if (m(diff)) {
                h := RoundPlusInt(hours * -1, 0)
                m := RoundPlusInt(minutes * -1, 0) - (h * 60)
                s := RoundPlusInt(seconds * -1, 0) - (m * 60)
                d := RoundPlusInt(hours * -1, 0) / 24
                mo := d / 30
                y := mo / 12
                return TimeDiffFormatter(h,m,s,d,mo,y)
        }
        return def
}
