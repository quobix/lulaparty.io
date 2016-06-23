package model

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "time"
)

var party = &Party{}

func TestSetCreated_Party(t *testing.T) {

        Convey("Given we have a party, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(party)
        })
}

func TestUpdate_Party(t *testing.T) {

        Convey("Given we party, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(party)
        })
}

func TestGetId_Party(t *testing.T) {

        Convey("Given we have a party there should be an appropriate ID", t, func () {
                helperGetId(party)
        })
}

func TestSetId_Party(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Party{})
        })
}

var partyinv = &PartyInventory{}

func TestSetCreated_PartyInventory(t *testing.T) {

        Convey("Given we have a party inventory, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(partyinv)
        })
}

func TestUpdate_PartyInventory(t *testing.T) {

        Convey("Given we party inventory, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(partyinv)
        })
}

func TestGetId_PartyInventory(t *testing.T) {

        Convey("Given we have a party inventory there should be an appropriate ID", t, func () {
                helperGetId(partyinv)
        })
}

func TestSetId_PartyInventory(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&PartyInventory{})
        })
}

var pinvite= &PartyInvite{}

func TestSetCreated_PartyInvite(t *testing.T) {

        Convey("Given we have a party invite, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(pinvite)
        })
}

func TestUpdate_PartyInvite(t *testing.T) {

        Convey("Given have a party invite, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(pinvite)
        })
}

func TestGetId_PartyInvite(t *testing.T) {

        Convey("Given we have a party invite there should be an appropriate ID", t, func () {
                helperGetId(pinvite)
        })
}

func TestSetId_PartyInvite(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&PartyInvite{})
        })
}

func TestSetId_Reward(t *testing.T) {

        Convey("Given we have a persisted entity, we should be able to set an ID ", t, func () {
                helperSetId(&Reward{})
        })
}

var reward = &Reward{}

func TestSetCreated_Reward(t *testing.T) {

        Convey("Given we have a reward, we should be able to test the creation date is correct", t, func() {
                helperSetCreated(reward)
        })
}

func TestUpdate_Reward(t *testing.T) {

        Convey("Given have a reward, we should be able to test the last updated date is correct", t, func () {
                helperUpdated(reward)
        })
}

func TestGetId_Reward(t *testing.T) {

        Convey("Given we have a reward there should be an appropriate ID", t, func () {
                helperGetId(reward)
        })
}


func TestParty_IsOpen(t *testing.T) {

        Convey("Given we have a party, we should be able to validate that the IsOpen method is correctly working", t, func () {

                p1 := &Party {
                        Starts: time.Date(2019, time.January, 10, 10, 0, 0, 0, time.UTC),
                        Ends: time.Date(2019, time.January, 10, 12, 0, 0, 0, time.UTC),
                        Enabled: true }



                p2 := &Party {
                        Starts: time.Date(2009, time.January, 10, 10, 0, 0, 0, time.UTC),
                        Ends: time.Date(2009, time.March, 10, 12, 0, 0, 0, time.UTC),
                        Enabled: true }

                t := time.Now()
                d := time.Hour * 24 * 2

                p3 := &Party {
                        Starts: t,
                        Ends: t.Add(d),
                        Enabled: true }


                Convey("given a non reached start date, the party should be closed", func () {
                        So(p1.IsOpen(), ShouldBeFalse)
                })

                Convey("given a past date, the party should be closed", func () {
                        So(p2.IsOpen(), ShouldBeFalse)
                })

                Convey("given a passed start and unreached end datethe party should be open", func () {
                        So(p3.IsOpen(), ShouldBeTrue)
                })

                Convey("given a party is open, if we disable it, it should be closed regardless", func () {
                        p3.Enabled = false;
                        So(p3.IsOpen(), ShouldBeFalse)
                })


        })
}

func TestParty_StartsIn(t *testing.T) {


        Convey("Given we have a party, we should have freindly ways to determining the diff between now and party start", t, func () {

                ti :=time.Now().UTC()

                p1 := &Party {
                        Starts: time.Date(ti.Year()+3, ti.Month(), ti.Day(),
                                ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("There should be 3 Years Left", func () {
                        So(p1.StartsIn(), ShouldEqual, "3 Years")
                })



                p1 = &Party {
                        Starts: time.Date(ti.Year()+2, ti.Month(), ti.Day(),
                                ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 2 years left", func () {
                        So(p1.StartsIn(), ShouldEqual, "2 Years")
                })

                p1 = &Party {
                        Starts: time.Date(ti.Year()+1, ti.Month(), ti.Day(),
                                ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 1 year left", func () {
                        So(p1.StartsIn(), ShouldEqual, "1 Year")
                })

                p1 = &Party {
                        Starts: time.Date(ti.Year(), ti.Month()+2, ti.Day(),
                                ti.Hour(), ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 2 Months left", func () {
                        So(p1.StartsIn(), ShouldEqual, "2 Months")
                })

                to:=time.Now().UTC()

                p2 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day()+4,
                                to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 4 Days left", func () {
                        So(p2.StartsIn(), ShouldEqual, "4 Days")
                })

                p3 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day(),
                                to.Hour()+1, to.Minute()+10, to.Second(), to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 1 Hour and 1 Mins Left", func () {
                        So(p3.StartsIn(), ShouldEqual, "1 Hour, 10 Mins")
                })

                p4 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day(),
                                to.Hour()+10, to.Minute()+22, to.Second(), to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 10 Hours and 22 Mins Left", func () {
                        So(p4.StartsIn(), ShouldEqual, "10 Hours, 22 Mins")
                })

                p5 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day(),
                                to.Hour(), to.Minute()+1, to.Second()+23, to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 1 Minute and 23 Seconds left", func () {
                        So(p5.StartsIn(), ShouldEqual, "1 Minute, 23 Seconds")
                })


                p6 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day(),
                                to.Hour(), to.Minute(), to.Second()+18, to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 18 Seconds left", func () {
                        So(p6.StartsIn(), ShouldEqual, "18 Seconds")
                })

                p7 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day(),
                                to.Hour(), to.Minute(), to.Second()+1, to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should be 1 Second left", func () {
                        So(p7.StartsIn(), ShouldEqual, "1 Second")
                })

                p8 := &Party {
                        Starts: time.Date(to.Year(), to.Month(), to.Day(),
                                to.Hour(), to.Minute(), to.Second(), to.Nanosecond(), time.UTC),
                        Enabled: true }

                Convey("There should have just started", func () {
                        So(p8.StartsIn(), ShouldEqual, PARTY_STARTED)
                })

        })
}

func TestParty_EndsIn(t *testing.T) {


        Convey("Given we have a party, we should have freindly ways to determining the diff between now and party end", t, func () {

                ti :=time.Now().UTC()

                p1 := &Party {
                        Starts: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()-5, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Ends: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("There should be 2 Hours Left", func () {
                        So(p1.IsOpen(), ShouldBeTrue)
                        So(p1.EndsIn(), ShouldEqual, "2 Hours")
                })

                p2 := &Party {
                        Starts: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()-6, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Ends: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+3, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("There should be 3 hours, 17 minutes left", func () {
                        So(p2.IsOpen(), ShouldBeTrue)
                        So(p2.StartsIn(), ShouldEqual, PARTY_STARTED)
                        So(p2.EndsIn(), ShouldEqual, "3 Hours, 17 Mins")
                })

                p3 := &Party {
                        Starts: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()-2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Ends: time.Date(ti.Year(), ti.Month(), ti.Day()+2,
                                ti.Hour()+12, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("There should be 2 day and 12 hours left", func () {
                        So(p3.IsOpen(), ShouldBeTrue)
                        So(p3.StartsIn(), ShouldEqual, PARTY_STARTED)
                        So(p3.EndsIn(), ShouldEqual, "2 Days, 12 Hours")
                })

                p4 := &Party {
                        Starts: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()-2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        Ends: time.Date(ti.Year(), ti.Month(), ti.Day()+7,
                                ti.Hour()+22, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("There should be 7 day and 22 hours left", func () {
                        So(p4.IsOpen(), ShouldBeTrue)
                        So(p4.StartsIn(), ShouldEqual, PARTY_STARTED)
                        So(p4.EndsIn(), ShouldEqual, "7 Days, 22 Hours")
                })


                p5 := &Party {
                        Starts: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+1, ti.Minute()+2, ti.Second(), ti.Nanosecond(), time.UTC),
                        Ends: time.Date(ti.Year(), ti.Month(), ti.Day()+7,
                                ti.Hour()+22, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("Party should not have started and shgould be closed", func () {
                        So(p5.IsOpen(), ShouldBeFalse)
                        So(p5.StartsIn(), ShouldEqual, "1 Hour, 2 Mins")
                        So(p5.EndsIn(), ShouldEqual, PARTY_CLOSED)
                })


                p6 := &Party {
                        Starts: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour(), ti.Minute()+22, ti.Second(), ti.Nanosecond(), time.UTC),
                        Ends: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour(), ti.Minute()+50, ti.Second(), ti.Nanosecond(), time.UTC),
                        Enabled: true }


                Convey("There should be 22 minutes left", func () {
                        So(p6.IsOpen(), ShouldBeFalse)
                        So(p6.StartsIn(), ShouldEqual, "22 Minutes")
                        So(p6.EndsIn(), ShouldEqual, PARTY_CLOSED)
                })





        })
}

func TestPartyInvite_HasExpired(t *testing.T) {

        Convey("Given we have a party with and end date, we should be able to test if it's expired", t, func () {

                Convey("The party invite should have expired two hours ago", func () {
                        ti :=time.Now().UTC()
                        p1 := &PartyInvite {
                                Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                        ti.Hour()-2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        }

                        So(p1.HasExpired(), ShouldBeTrue)
                })

                Convey("The party invite should expire in two hours from now", func () {
                        ti :=time.Now().UTC()
                        p1 := &PartyInvite {
                                Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                        ti.Hour()+2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        }

                        So(p1.HasExpired(), ShouldBeFalse)
                })


                Convey("The party invite should should expire in 22 minutes from now", func () {
                        ti :=time.Now().UTC()
                        p1 := &PartyInvite {
                                Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                        ti.Hour(), ti.Minute()+22, ti.Second(), ti.Nanosecond(), time.UTC),
                        }

                        So(p1.HasExpired(), ShouldBeFalse)
                })

                Convey("The party invite should should expire in in an hour, but has been accepted already", func () {
                        ti :=time.Now().UTC()
                        p1 := &PartyInvite {
                                Accepted: true,
                                Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                        ti.Hour()+1, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC),
                        }

                        So(p1.HasExpired(), ShouldBeTrue)
                })


        })
}


func TestPartyInvite_ExpiresIn(t *testing.T) {


        Convey("Given we have a party invite, we should have freindly ways to determining the diff between now and expiration", t, func () {

                ti :=time.Now().UTC()

                p1 := &PartyInvite {
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC) }

                Convey("There should be 2 Hours Left", func () {
                        So(p1.ExpiresIn(), ShouldEqual, "2 Hours")
                })

                p2 := &PartyInvite {
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()+3, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC) }



                Convey("There should be 3 hours, 17 minutes left", func () {
                        So(p2.ExpiresIn(), ShouldEqual, "3 Hours, 17 Mins")

                })

                p3 := &PartyInvite {
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day()+2,
                                ti.Hour()+12, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC) }


                Convey("There should be 2 day and 12 hours left", func () {
                        So(p3.ExpiresIn(), ShouldEqual, "2 Days, 12 Hours")
                })

                p4 := &PartyInvite {
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day()+7,
                                ti.Hour()+22, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC) }


                Convey("There should be 7 day and 22 hours left", func () {
                        So(p4.ExpiresIn(), ShouldEqual, "7 Days, 22 Hours")
                })


                p5 := &PartyInvite {
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day(),
                                ti.Hour()-2, ti.Minute(), ti.Second(), ti.Nanosecond(), time.UTC) }


                Convey("Party Invite should not expired", func () {
                        So(p5.ExpiresIn(), ShouldEqual, PARTY_INVITE_EXPIRED)
                })

                p6 := &PartyInvite {
                        Accepted: true,
                        Expires: time.Date(ti.Year(), ti.Month(), ti.Day()+7,
                                ti.Hour()+22, ti.Minute()+17, ti.Second(), ti.Nanosecond(), time.UTC) }

                Convey("Party invite should be unexpired but already accepted", func () {
                        So(p6.ExpiresIn(), ShouldEqual, PARTY_INVITE_ACCEPTED)
                })



        })
}