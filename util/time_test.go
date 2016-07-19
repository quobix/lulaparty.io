package util

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "time"
)

func TestTimeDiffHelper(t *testing.T) {

        Convey("Given we know the time difference, we should be able to test the time difference", t, func() {

                ti :=time.Now().UTC()

                diff1 := time.Date(ti.Year()+3, ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff2 := time.Date(ti.Year()+2, ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff3 := time.Date(ti.Year()+1, ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff4 := time.Date(ti.Year(), ti.Month()+1, ti.Day()+2,
                        ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff5 := time.Date(ti.Year(), ti.Month(), ti.Day()+2,
                        ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff6 := time.Date(ti.Year(), ti.Month(), ti.Day(),
                        ti.Hour()+8, ti.Minute()+22, ti.Second(), ti.Nanosecond(), time.UTC)

                diff7 := time.Date(ti.Year(), ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute()+4, ti.Second()+13, ti.Nanosecond(), time.UTC)

                Convey("There should be 3 Years Left", func () {
                        So(TimeDiffHelper(diff1,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "3 Years")
                })

                Convey("There should be 2 Years Left", func () {
                        So(TimeDiffHelper(diff2,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "2 Years")
                })

                Convey("There should be 1 Year Left", func () {
                        So(TimeDiffHelper(diff3,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "1 Year")
                })

                Convey("There should be 1 Month Left", func () {
                        So(TimeDiffHelper(diff4,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "1 Month")
                })

                Convey("There should be 2 Days Left", func () {
                        So(TimeDiffHelper(diff5,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "2 Days")
                })

                Convey("There should be 8 Hours, 22 mins left", func () {
                        So(TimeDiffHelper(diff6,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "8 Hours, 22 Mins")
                })

                Convey("There should be 4 Mins and 13 Seconds left", func () {
                        So(TimeDiffHelper(diff7,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "4 Minutes, 13 Seconds")
                })
        })
}

func TestTimeDiffFormatter(t *testing.T) {
}

func TestPluralize(t *testing.T) {
}
