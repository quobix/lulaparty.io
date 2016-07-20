package util

import (
        "testing"
        . "github.com/smartystreets/goconvey/convey"
        "strings"
        "gopkg.in/mgo.v2/bson"
        "net/http"
)

func TestGenerateUUID(t *testing.T) {
        Convey("Given we have access to the OS, we should be able to call uuidgen and get an ID", t, func() {
                uuid := GenerateUUID()
                So(uuid, ShouldNotBeNil)
                So(len(uuid), ShouldBeGreaterThan, 10)

                Convey("We should also be able to check the structure for accurate lengths", func() {

                        s := strings.Split(uuid, "-");
                        i := 0
                        for _, n := range s {
                                switch {
                                case i == 0:
                                        So(len(n), ShouldEqual, 8)
                                case i == 1:
                                        So(len(n), ShouldEqual, 4)
                                case i == 2:
                                        So(len(n), ShouldEqual, 4)
                                case i == 3:
                                        So(len(n), ShouldEqual, 4)
                                case i == 4:
                                        So(len(n), ShouldEqual, 12)
                                }
                                i++;
                        }
                })

                Convey("It should also return a valid UUID", func() {

                        So(len(uuid), ShouldEqual, 36)
                        So(ValidateUUID(uuid), ShouldBeTrue)
                })
        })
}

func TestRound(t *testing.T) {
        Convey("We should be able to validate rounding up",t,  func() {
                So(Round(1.6), ShouldEqual, 2.0)
                So(Round(1.0), ShouldEqual, 1.0)
                So(Round(1.2), ShouldEqual, 1.0)
                So(Round(1.4), ShouldEqual, 1.0)
                So(Round(1.5), ShouldEqual, 2.0)
                So(Round(1.6), ShouldEqual, 2.0)
        })
}

func TestRoundPlus(t *testing.T) {
        Convey("We should be able to validate rounding up, with a decimal place!",t, func() {
                So(RoundPlus(1.6233, 2), ShouldEqual, 1.62)
                So(RoundPlus(1.11111, 3), ShouldEqual, 1.111)
                So(RoundPlus(22.2332, 3), ShouldEqual, 22.233)
                So(RoundPlus(2211.421112, 0), ShouldEqual, 2211.0)
                So(RoundPlus(1.00000000, 1), ShouldEqual, 1.00)
                So(RoundPlus(99.9021, 10), ShouldEqual, 99.9021)
        })
}

func TestRandStringRunes(t *testing.T) {
        Convey("I want to create some random runes, they should not be equal",t, func() {
                r1 := RandStringRunes(5)
                r2 := RandStringRunes(5)
                r3 := RandStringRunes(5)
                r4 := RandStringRunes(5)
                r5 := RandStringRunes(2)
                r6 := RandStringRunes(4)
                r7 := RandStringRunes(22)
                r8 := RandStringRunes(100)

                So(r1, ShouldNotEqual, r2)
                So(len(r1), ShouldEqual, len(r2))
                So(r2, ShouldNotEqual, r3)
                So(len(r2), ShouldEqual, len(r3))
                So(r3, ShouldNotEqual, r4)
                So(len(r1), ShouldEqual, len(r4))

                So(r5, ShouldNotEqual, r6)
                So(len(r5), ShouldNotEqual, len(r6))
                So(r6, ShouldNotEqual, r7)
                So(len(r6), ShouldNotEqual, len(r7))
                So(r7, ShouldNotEqual, r8)
                So(len(r7), ShouldNotEqual, len(r8))
        })
}

func TestSliceHelper(t *testing.T) {
        Convey("I want to remove an item from a slice:",t,  func() {

                id1 := bson.NewObjectId();
                id2 := bson.NewObjectId();
                id3 := bson.NewObjectId();
                vals := []bson.ObjectId {id1, id2, id3}

                i:=-1
                for idx, itms := range vals {
                        if(itms.Hex() == id1.Hex()) {
                                i = idx
                        }
                }

                So(i, ShouldBeGreaterThanOrEqualTo, 0)

                // slice trick to delete from collection, don't care about
                newvals := SliceHelper(vals,i)
                So(len(newvals), ShouldBeLessThan, len(vals))
                So(len(newvals), ShouldEqual, 2)

                i=-1
                for idx, itms := range vals {
                        if(itms.Hex() == id3.Hex()) {
                                i = idx
                        }
                }
        })
}

func TestGetMD5Hash(t *testing.T) {
        Convey("We should be able to validate that our MD5 Hash generator matches the output of a verified one",t,  func() {
                // generated with an external tool
                h := "8dedcde9e72b1cdbcfe341b7306d0c8f"
                k := "pup|shop|trick|shot|cotton|fox|maggiepop"
                v := GetMD5Hash(k)
                So(v, ShouldNotBeNil)
                So(v, ShouldEqual, h)
        })
}

func TestReadBody(t *testing.T) {
        Convey("Given we can connect to a remote resource, we should be able to parse the body",t,  func() {
                resp, err := http.Get("http://httpbin.org/robots.txt")

                So(err, ShouldBeNil)
                So(resp, ShouldNotBeNil)
                So(resp.StatusCode, ShouldNotEqual, 401)
                So(resp.StatusCode, ShouldNotEqual, 404)
                So(resp.StatusCode, ShouldEqual, 200)

                str := ReadBody(resp);

                So(strings.Contains(str, "User-agent: *"), ShouldBeTrue)
                So(strings.Contains(str, "Disallow: /deny"), ShouldBeTrue)
        })
}
