package security

import (
        "time"
        "github.com/dgrijalva/jwt-go"
        "fmt"
        "lulaparty.io/model"
        "lulaparty.io/util"
)

const (
        defaultTTL = 3600 * 24 * 7 // 1 week
        JWTSECRET = "giantkingslugthing|chickie,maggie,fox"
)

// I have heavily modified this. Originally taken from github.com/dghubble/jwts
type Provider interface {
        CreateToken() *jwt.Token
        SignByte(token *jwt.Token) ([]byte, error)
        SignString(token *jwt.Token) (string, error)
        GetClaims(token string) (*jwt.Token, error)
}

type Config struct {
        Method jwt.SigningMethod
        TTL int64
}

type Manager struct {
        key    []byte
        method jwt.SigningMethod
        ttl    int64
}

type LuluClaims struct {
        Lpu string `json:"lpu"`
        Lph string `json:"lph"`
        jwt.StandardClaims
}

func CreateNewManager(key []byte, configs ...Config) *Manager {
        var c Config
        if len(configs) == 0 {
                c = Config{}
        } else {
                c = configs[0]
        }
        m := &Manager{
                key:    key,
                method: c.Method,
                ttl:    c.TTL,
        }
        m.setDefaults()
        return m
}

func (m *Manager) setDefaults() {
        if m.method == nil { m.method = jwt.SigningMethodHS256 }
        if m.ttl == 0 { m.ttl = defaultTTL }
}

func (m *Manager) CreateToken(u *model.User) *jwt.Token {
        d := time.Duration(m.ttl) * time.Second
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "iat": time.Now().Unix(),
                "lpu": u.Id,
                "lph": util.GetMD5Hash(u.Id.String()+u.Email+u.Created.String()),
                "exp": time.Now().Add(d).Unix(),
        })
        return token
}

func (m *Manager) SignByte(token *jwt.Token) ([]byte, error) {
        jwtString, err := token.SignedString(m.key)
        return []byte(jwtString), err
}

func (m *Manager) SignString(token *jwt.Token) (string, error) {
        jwtString, err := token.SignedString(m.key)
        return jwtString, err
}

func (m *Manager) GetClaims(jwtString  string) (*LuluClaims, error) {
        token, err := jwt.ParseWithClaims(jwtString, &LuluClaims{}, m.getKey)
        if err == nil && token.Valid {
                if claims, ok := token.Claims.(*LuluClaims); ok && token.Valid {
                        return claims, nil
                } else {
                        return nil, err
                }
        }
        return nil, fmt.Errorf("no token found in request")
}

// getKey accepts an unverified JWT and returns the signing/verification key.
// Also ensures tha the token's algorithm matches the signing method expected
// by the manager.
func (m *Manager) getKey(unverified *jwt.Token) (interface{}, error) {
        // require token alg to match the set signing method, do not allow none
        if meth := unverified.Method; meth == nil || meth.Alg() != m.method.Alg() {
                return nil, jwt.ErrHashUnavailable
        }
        return m.key, nil
}