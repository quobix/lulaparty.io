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

                diff7 := time.Date(ti.Year(), ti.Month(), ti.Day()+1,
                        ti.Hour()+8, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff8 := time.Date(ti.Year(), ti.Month(), ti.Day(),
                        ti.Hour()+9, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC)

                diff9 := time.Date(ti.Year(), ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute()+28, ti.Second()+17, ti.Nanosecond(), time.UTC)

                diff10 := time.Date(ti.Year(), ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute()+5, ti.Second(), ti.Nanosecond(), time.UTC)

                diff11 := time.Date(ti.Year(), ti.Month(), ti.Day(),
                        ti.Hour(), ti.Minute(), ti.Second()+2, ti.Nanosecond(), time.UTC)

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

                Convey("There should be 8 Hours 22 Minutes Left", func () {
                        So(TimeDiffHelper(diff6,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "8 Hours, 22 Minutes")
                })

                Convey("There should be 1 Day and 8 Hours left", func () {
                        So(TimeDiffHelper(diff7,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "1 Day, 8 Hours")
                })

                Convey("There should be 9 Hours Left", func () {
                        So(TimeDiffHelper(diff8,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "9 Hours")
                })

                Convey("There should be 28 Minutes and 17 seconds Left", func () {
                        So(TimeDiffHelper(diff9,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "28 Minutes, 17 Seconds")
                })

                Convey("There should be 5 Minutes Left", func () {
                        So(TimeDiffHelper(diff10,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "5 Minutes")
                })

                Convey("There should be 2 Seconds Left", func () {
                        So(TimeDiffHelper(diff11,"", func(t time.Time) bool {
                                return true
                        }), ShouldEqual, "2 Seconds")
                })
        })
}

func TestTimeDiffFormatter(t *testing.T) {
}

func TestPluralize(t *testing.T) {
}
