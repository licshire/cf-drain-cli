// Code generated by go-bindata.
// sources:
// test-certs/syslog-ca.crl
// test-certs/syslog-ca.crt
// test-certs/syslog-ca.key
// test-certs/syslog.crt
// test-certs/syslog.csr
// test-certs/syslog.key
// DO NOT EDIT!

package testhelper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _syslogCaCrl = []byte(`-----BEGIN X509 CRL-----
MIICgjBsAgEBMA0GCSqGSIb3DQEBCwUAMBMxETAPBgNVBAMTCHN5c2xvZ0NBFw0x
ODA3MTIyMjE4NDdaFw0yMDAxMTIyMjE4NDZaMACgIzAhMB8GA1UdIwQYMBaAFCQY
ykgBbJ+xvUTvFD5ro5ApHdwZMA0GCSqGSIb3DQEBCwUAA4ICAQAj7fDdFAgvOGjQ
3rfnYLVEY/KjJrdGYNrRnqDD9OeSGi5bqvI/pfpx0m8tEBpan/W7smrehGdDvPzK
xwraFU07rIzNGl+YYnrnO0ruJisJbH8MLqLAPomYMcNcF8sVE1pRx6sLaAYQJ+aQ
BSUwYKFG9C7oB9UufqOjPr1mvLJK3QPILRWgq9ynuQkmXGS0GLvLJT9vsHFvwFBN
Dz7l1EpH8PITagilmXsHPlUsiZXq32BrTRAUjClegWiw+mgXQEEX6xf9slSkujda
bJIGbtM/XzcSnFBhPD3/keIEfdr/aIcR9KDqpo+hxRO/ATFieXxPt/ht+oqEK8yL
BFEQIrpBYn9aOskVO4uo9IoOO9L+oWHaQIHfPTbgFUm3htaBmoKWiwTcYpwyDQob
lzDwuAp87+8ei2HRIWEClN+zryP6+1iLQ3UBqmBu5nGIj8eanXPHA3SRZVY9ZLMJ
fg/pABwjwah7CIyR+oSjZhi6thT3HZBLmHt5igD1+BP6iC2uMjsu+yQyRXL7lvvr
G6eMFCJdDB3VI/cUJycDjmfP40gjiUpphm9i4csU+FBILlLmxX2s/JaftF4xprRr
/S23Znbk/fS5D2+BV6a8H8Z16Pq+Qdu2glJSpt1N8FNxQMMh9Hd7A5MkimL2dTC6
5qUSg2gnGkXIteBu6XUdCf6POW2Ucw==
-----END X509 CRL-----
`)

func syslogCaCrlBytes() ([]byte, error) {
	return _syslogCaCrl, nil
}

func syslogCaCrl() (*asset, error) {
	bytes, err := syslogCaCrlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslog-ca.crl", size: 926, mode: os.FileMode(292), modTime: time.Unix(1531433927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogCaCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIE5jCCAs6gAwIBAgIBATANBgkqhkiG9w0BAQsFADATMREwDwYDVQQDEwhzeXNs
b2dDQTAeFw0xODA3MTIyMjE4NDZaFw0yMDAxMTIyMjE4NDZaMBMxETAPBgNVBAMT
CHN5c2xvZ0NBMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAv6aHDwpx
ZqsgzXhW0tvOoylenocQg6UHQJKGZ4kYHMzVPDS+bt+Vw+qvnOeMnIZypMX3kaUr
y+BjgzjUtMb5fRrWoIO5pZOq/C5b7iMNe6xK8+NMeBtiOrDB79XxfGeythJZDySA
L/vyrUcBZGBRkLHNSdsjtKuOrPro5MNBWItzvwGlAR+Tqz8A2ulECvq0I9s+waGB
WlsHYmVNUSX+rsyEltdfeRPvVhM+sh5x6XRwc3WRIo2VsudTXJcWFsulJNJveV4h
9sayL7C0XNyGhwZkCyDnwH7hivmhTJYaLbp8NAE16jJfFwavnR2zN1eqXGU64oVB
Dz1w5UIvUVUpcn1Wabq6m1Hfsmhi3zdWf9Shx5tGOj9bB8+b2MhdwuBnSQzIgqrT
g2g8BJ3WVmD0GSbxwxjb4An3hScQ7bf7lsPIiO6M80XqwJDrD2aOLOaH1QJNwGOG
QG7OobGywiwuH3vRxiQWj95kfcwp7+JepHvFEdn5dMYehMAMy2lsNII5hZZ9sG42
cVFxzSpnKW0ngF6P+nBE56Pf6ZMuCbfvlTX6URlD4ruKXhLfjfnuc0EGvGqUaFXP
e/f1dvrixRT8c1AFj90PuGU0euc6PravunhLc+1vY/fqlvvrvejffhTSXCW9Jh/6
8ULngYjs9N6zNwE/KPOE+Kugh2YY5krQM8MCAwEAAaNFMEMwDgYDVR0PAQH/BAQD
AgEGMBIGA1UdEwEB/wQIMAYBAf8CAQAwHQYDVR0OBBYEFCQYykgBbJ+xvUTvFD5r
o5ApHdwZMA0GCSqGSIb3DQEBCwUAA4ICAQC/bCH1Py44kRO10QkjFKWChaYqtT42
xqzI7k7nE0AP+4fro978nmtarOUnst/9eOTvYp4kkv4Qi82nto/m/H/21ZOMFHz3
3omxLtgdoXsnfxkA3cmkwGVcUxDbzDgd4LoroKKtniSFrXMPOJnbF0oUfldtSvH/
UrvEYn8IDDFb5YWJ/iNbo1b908MZiGNVOC//wkPzrj3VLg13fy1fcahLFqTla4ra
X7EPwaaKdUqikWTI1VAVOGP2dUH7thzHpeYqhpZ8zQWT3NrKck0JoeCsBG9wZIan
iG4GQRxPgXOw1Jx0YOMuvlO1C5ohlKQ8wr2DxobVLHHQXHxNK6eHPKC1BZLsInNW
C+O/Zg+/R2myv9bIzLXZNKXeSt/Q2KayvYaGigJHYcucnSIJyWWowY033J8xh/cO
/zveAS1cNveXfMhUX7sSFnfd/z5OT36xdpBwqmoR6Bx3yRZ6Ia0PwmrxT5JWsXA7
Y0l0SlxnslAKFxPXLQq+scr8ZYwZF6ORhNjYI+bqOGOhUGdOu9w67JKwrvAJa/VW
widZMmwxB3qUD2FPlFQp+jbaS6w0LhG47lIqvsEn6eK4VByH9A1gzsukWTvrPTDT
1vynDEkll1TopseeO6Sf2go11XIl+E++op++nSZSCH7r4g1S/C3/sqRyU1WoFkNr
0CFIQB3RuK1X+w==
-----END CERTIFICATE-----
`)

func syslogCaCrtBytes() ([]byte, error) {
	return _syslogCaCrt, nil
}

func syslogCaCrt() (*asset, error) {
	bytes, err := syslogCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslog-ca.crt", size: 1761, mode: os.FileMode(292), modTime: time.Unix(1531433927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogCaKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

func syslogCaKeyBytes() ([]byte, error) {
	return _syslogCaKey, nil
}

func syslogCaKey() (*asset, error) {
	bytes, err := syslogCaKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslog-ca.key", size: 3243, mode: os.FileMode(288), modTime: time.Unix(1531433927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIEIDCCAgigAwIBAgIRAIFumfRKipgHzq+xeJ38dZIwDQYJKoZIhvcNAQELBQAw
EzERMA8GA1UEAxMIc3lzbG9nQ0EwHhcNMTgwNzEyMjIxODQ3WhcNMjAwMTEyMjIx
ODQ1WjARMQ8wDQYDVQQDEwZzeXNsb2cwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQDK4vy91/uUVxFVZnwSZL7hah3mdeb2zKqt5nIncvzlSiMSK9qnzOAO
CbwlIuueVOcjMKSK3eh/LOTBKdDrpszaRjrOU2g0+KPNeJqeJSQtFrQ9UGKOuYxD
fHV8TOrbuHwOyWb74TkyZhWXcLlJou1cvTBpkafkQLAZgEBDqHd/HaAe1444kpSn
bIumcxNYKlgu9TIbcKa6su234jGQboNvqqCGD3rhZvQFDH3viXQg507FiKHBCxxv
UojOytWuc6CwpzjSGlbBdgdNNo9nV7OKmZtX9ZhWCY3/RrT8buv5sBt57ZrFCe0w
759kLlUuzmQfwET5ylC5o09zRbZYA2o/AgMBAAGjcTBvMA4GA1UdDwEB/wQEAwID
uDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHQYDVR0OBBYEFJ7WiO9E
EbafRpdubhyaNxFpVNCDMB8GA1UdIwQYMBaAFCQYykgBbJ+xvUTvFD5ro5ApHdwZ
MA0GCSqGSIb3DQEBCwUAA4ICAQATkEcRKk2Y5Gxgiakbw9vEhfWlNnm4S93xVDpc
FOEufBSV89HJnJYHNKbhotOb0x6p9vZBH5tPtS0QnCYGX4e9gdttkM0CFTobIdhs
mMOGYRRRRf0AW/Q4rR1oarQ94oh9sBYec6it0cZIQJD1GzHfJEuBTnTCgzKIR10b
n6r61XZlC0qbgorITbRcFl1nBSNO4JS38Q4nX+c9yyaOc93dsttGd1lJ/7K7tEWC
X8JCK5PUdoAq55OjtBSiwwiwrrTC+W6T3Hls0IRlo5ABXu/qds+wgtr9gdVfee8+
cyz/5dosZisuJc7uFtBPjj7QwFawICs99czevqK3UMirtVlPPrR5Mup+JLX6pYPq
T20Dk0TrLFWbiaDYvwpP0plEL9K4ywqkiiaTMNHWki7oK01gbBS1BDe6CMBbb7M2
MWOjtUbBah20vT0fCgrYqnHH0YgMnokrJlgWnrHrJHHzm/42fAaV6M5j21p0rp2V
pvJKL/YXEZBN3Y0ODW9YjK0hvnW+nX+QNXsP98zkoOqyapB0wYqMfhrnd7vY+tgV
jY9F+4XkjNOYARJ/ASpn8RMcx/YaGI4Jub0LzyXNBl1Ej4q/Jy8UWVYVWNwqs4x/
Ese5d/t9QT6F6AqX1RjTKWHLpT8vI5Rc90b+OqgKdXDjNsrS78I57aczRW+mkTYd
/INCzw==
-----END CERTIFICATE-----
`)

func syslogCrtBytes() ([]byte, error) {
	return _syslogCrt, nil
}

func syslogCrt() (*asset, error) {
	bytes, err := syslogCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslog.crt", size: 1493, mode: os.FileMode(292), modTime: time.Unix(1531433927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogCsr = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIICVjCCAT4CAQAwETEPMA0GA1UEAxMGc3lzbG9nMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAyuL8vdf7lFcRVWZ8EmS+4Wod5nXm9syqreZyJ3L85Uoj
Eivap8zgDgm8JSLrnlTnIzCkit3ofyzkwSnQ66bM2kY6zlNoNPijzXianiUkLRa0
PVBijrmMQ3x1fEzq27h8Dslm++E5MmYVl3C5SaLtXL0waZGn5ECwGYBAQ6h3fx2g
HteOOJKUp2yLpnMTWCpYLvUyG3CmurLtt+IxkG6Db6qghg964Wb0BQx974l0IOdO
xYihwQscb1KIzsrVrnOgsKc40hpWwXYHTTaPZ1ezipmbV/WYVgmN/0a0/G7r+bAb
ee2axQntMO+fZC5VLs5kH8BE+cpQuaNPc0W2WANqPwIDAQABoAAwDQYJKoZIhvcN
AQELBQADggEBALHjMW+7g78cY3nW6TNxyqujF3T/0DmBbtVDotCQBYSEmb4yk6Bx
9Py/aNvu8knzUdFhQxjDAKva44pqzxPRfmX4Hv1hoOEgfnEIxl7Oghhq7GweZwvv
KKR1KZl2ktAPP2n6u6q1axtNJSGtfX59+QS0S9g5UWh4z6CTegpsErUJTRgHZPKC
ObdahORi3dh7rr5NwaUDrien/gWFAw5m2GoM1tR/Nfd0AvEIINHVUYOLfEmVg/ku
TTqFC7AslWt2WzMmbAW81eGGHhfvurNztAjNxxgBUZpjsSvSqPElbMrck6KJaAcA
spXhUXvZrDFemjiW3IDu3FtXJrOvoyqF3sY=
-----END CERTIFICATE REQUEST-----
`)

func syslogCsrBytes() ([]byte, error) {
	return _syslogCsr, nil
}

func syslogCsr() (*asset, error) {
	bytes, err := syslogCsrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslog.csr", size: 887, mode: os.FileMode(292), modTime: time.Unix(1531433927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAyuL8vdf7lFcRVWZ8EmS+4Wod5nXm9syqreZyJ3L85UojEiva
p8zgDgm8JSLrnlTnIzCkit3ofyzkwSnQ66bM2kY6zlNoNPijzXianiUkLRa0PVBi
jrmMQ3x1fEzq27h8Dslm++E5MmYVl3C5SaLtXL0waZGn5ECwGYBAQ6h3fx2gHteO
OJKUp2yLpnMTWCpYLvUyG3CmurLtt+IxkG6Db6qghg964Wb0BQx974l0IOdOxYih
wQscb1KIzsrVrnOgsKc40hpWwXYHTTaPZ1ezipmbV/WYVgmN/0a0/G7r+bAbee2a
xQntMO+fZC5VLs5kH8BE+cpQuaNPc0W2WANqPwIDAQABAoIBAQCvO3XIEm+8I6dY
93dZ+6HR13GTx2mA6CzSniMiZ3J+J5Y6752fKJisB0R1Xr/atMId5rl3J0rT7NZ2
78L+hcaRD+9inS938PipUu1YggPI3b825eL6GtBGyYw/m7+Nnr2nl9RW4KQiH4Gu
lzGxpAMXAmB6Dc/QmP6ASAE5bpwnMMiWfvCv8B3l36nfdFtuRdqV+qaz2k9dX41W
zT954SuV7GCu9xu3AlmoJB9VwvfAXRFp184NSUoBDxMZmfzsZIF6O/7nRhG08Kmk
WMrYBypq8VQhPoBFNEyE4qGQHKR1fzA7qO3Xlws1tFTOewvrplJ5sxYRg+BK1U/k
1VNLgarhAoGBANQl57tdUjyRQa244qYxFrgm0f4vDmE+6jYPRH6tNN/ey6oQPd1K
Jny9ow1Eg5Iu5pT4yAoGmWOLpFeyfW7t0zblEPq5mbu6yXCLONK4oR5i1dzZY2aA
QyBMRWtZ4EloLyoNI8PmZvuglFsR0drTPpazdynKsiEwU/G2exphl1ylAoGBAPTT
AMKZG8ZSWQkGFb2yuvDvBa0HUPG1GZaUEkzqNIffkLDnpVFUU0uUMspxcsNWzMeW
+0dv8zxqXCsZ9rfJnHYHu0AykQcDaDsf1YscTASaGSwOuR5oN0UUm42bZsLCjINi
lIBgTk2a/vBw5hmv+vrGqLvvK+60nJUjpif2QUITAoGAQjovFFu2r3e9JClif/U1
HXno9kvkVlFYtlWf0Vfq+LRJqiZ12Z4rU6ezvhGUbU885maftBmJSv0t6b6rz7Ro
ymtKGKtWJbfS6NJdg9LlYWIDV3V5xWbnDa8hwLiG/wOKLbt6Xc6QeZ8QkTmH5KMr
l8oguigv4ZJ/siaAWaT7po0CgYACvfI+O7TB7d4tabIIo93QIP0xSOmxK1QUCaiC
0ASymOBn+tAxLv6X1BsVhVvBsFEu0xQRYsGYpaBOo0XOHSXxdHi2aVqSd4lNUDf2
DOUVEBzED44nNweXcHmlhl4KRF+KgSokne+Ckv/T9Y++LehjsqHlKPKXA9LVIMcV
n35vnQKBgAJzH1cL5+cLIXpbCfMnhlBJI1+1UPVLqrE824veuNBlBqMhUaRr1V9f
gQtCjtsTvJ7CxnUbBuNg2WwN4wN+lQ8otSCBHuub+naZimvQjt8SnpOc6Qmurbe+
nsbPeV0nkcgYlZ5TegvRB0HxP+AK+WiHqxvWxcfkTZuxtyunlb93
-----END RSA PRIVATE KEY-----
`)

func syslogKeyBytes() ([]byte, error) {
	return _syslogKey, nil
}

func syslogKey() (*asset, error) {
	bytes, err := syslogKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslog.key", size: 1675, mode: os.FileMode(288), modTime: time.Unix(1531433927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"syslog-ca.crl": syslogCaCrl,
	"syslog-ca.crt": syslogCaCrt,
	"syslog-ca.key": syslogCaKey,
	"syslog.crt": syslogCrt,
	"syslog.csr": syslogCsr,
	"syslog.key": syslogKey,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"syslog-ca.crl": &bintree{syslogCaCrl, map[string]*bintree{}},
	"syslog-ca.crt": &bintree{syslogCaCrt, map[string]*bintree{}},
	"syslog-ca.key": &bintree{syslogCaKey, map[string]*bintree{}},
	"syslog.crt": &bintree{syslogCrt, map[string]*bintree{}},
	"syslog.csr": &bintree{syslogCsr, map[string]*bintree{}},
	"syslog.key": &bintree{syslogKey, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

