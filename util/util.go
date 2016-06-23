package util

import (
        "os/exec"
        "fmt"
        "regexp"
        "strings"
        "github.com/quobix/lulaparty.io/model"
        "os"
        "path"
        "gopkg.in/mgo.v2/bson"
        "math"
        "time"
        "math/rand"
        "crypto/md5"
        "encoding/hex"
)

const (
        FILE_UUID_FSSEP = "/"
        FILE_UUID_EXT = "_"
)

func GenerateUUID() string {
        out, err := exec.Command("uuidgen").Output()
        if err != nil {
                panic("unable to generate uuid via uuidgen")
        }
        return strings.ToLower(strings.TrimSpace(fmt.Sprintf("%s",out)))
}

func ValidateUUID(text string) bool {
        r := regexp.MustCompile("(?i)^[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}$")
        return r.MatchString(text)
}


func GenerateGalleryItemUUID(u *model.User, g *model.Gallery, i *model.GalleryItem, f *os.File) string {
       return GenerateRawGalleryItemUUID(u.Id, g.Id, i.Id, f)
}

func GenerateRawGalleryItemUUID(u bson.ObjectId, g bson.ObjectId, i bson.ObjectId, f *os.File) string {
        return u.Hex() + FILE_UUID_FSSEP + g.Hex() +
                FILE_UUID_FSSEP + i.Hex() + FILE_UUID_EXT + path.Base(f.Name())
}


func Round(f float64) float64 {
        return math.Floor(f + .5)
}
func RoundPlus(f float64, places int) (float64) {
        shift := math.Pow(10, float64(places))
        return Round(f * shift) / shift;
}

func SliceHelper(sl []bson.ObjectId, i int) []bson.ObjectId {
        return append(sl[:i], sl[i+1:]...)
}

//http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func init() {
        rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
        b := make([]rune, n)
        for i := range b {
                b[i] = letterRunes[rand.Intn(len(letterRunes))]
        }
        return string(b)
}

func GetMD5Hash(text string) string {
        hasher := md5.New()
        hasher.Write([]byte(text))
        return hex.EncodeToString(hasher.Sum(nil))
}