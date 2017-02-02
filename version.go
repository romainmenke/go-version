package version

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type VersionList struct {
	Items []struct {
		Name string `json:"name"`
	} `json:"items"`
}

type Version struct {
	//Major version when you make incompatible API changes.
	Major       string
	MajorLatest bool
	//Minor version when you add functionality in a backwards-compatible manner.
	Minor       string
	MinorLatest bool
	//Patch version when you make backwards-compatible bug fixes.
	Patch       string
	PatchLatest bool
}

func (v Version) MajorInt() int {

	i, err := strconv.ParseInt(v.Major, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)

}

func (v Version) MinorInt() int {

	i, err := strconv.ParseInt(v.Minor, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)

}

func (v Version) PatchInt() int {

	i, err := strconv.ParseInt(v.Patch, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)

}

func (v Version) Version() string {

	if v.MajorLatest {
		return "master"
	}
	if v.Minor == "" && !v.MinorLatest {
		return "go" + v.Major
	}
	if v.MinorLatest {
		return "go" + v.Major + "." + "x"
	}
	if v.Patch == "" && !v.PatchLatest {
		return "go" + v.Major + "." + v.Minor
	}
	if v.PatchLatest {
		return "go" + v.Major + "." + v.Minor + "." + "x"
	}
	return "go" + v.Major + "." + v.Minor + "." + v.Patch

}

func (v Version) VersionWithoutWildcard() string {

	if v.MajorLatest {
		return "master"
	}
	if v.Minor == "" && !v.MinorLatest {
		return "go" + v.Major
	}
	if v.MinorLatest {
		return "go" + v.Major
	}
	if v.Patch == "" && !v.PatchLatest {
		return "go" + v.Major + "." + v.Minor
	}
	if v.PatchLatest {
		return "go" + v.Major + "." + v.Minor
	}
	return "go" + v.Major + "." + v.Minor + "." + v.Patch

}

func (v Version) MajorVersion() string {

	if v.MajorLatest {
		return "master"
	}
	return "go" + v.Major

}

func (v Version) MinorVersion() string {

	if v.MajorLatest {
		return "master"
	}
	if v.MinorLatest {
		return "go" + v.Major + "." + "x"
	}
	return "go" + v.Major + "." + v.Minor

}

func (v *Version) Before(v2 *Version) bool {

	if v.MajorInt() < v2.MajorInt() {
		return true
	}

	if v.MinorInt() < v2.MinorInt() {
		return true
	}

	if v.PatchInt() < v2.PatchInt() {
		return true
	}
	return false
}

func (v *Version) After(v2 *Version) bool {
	return v2.Before(v)
}

func GetVersionList() (*VersionList, error) {

	response, err := http.Get("https://www.googleapis.com/storage/v1/b/golang/o?fields=items%2Fname&maxResults=999999")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	out := &VersionList{}
	err = decoder.Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil

}

func generateVersionFromString(str string) *Version {

	str = strings.TrimPrefix(str, "go")
	v := &Version{}

	if str == "master" {
		v.MajorLatest = true
		return v
	}

	components := strings.Split(str, ".")
	for index, component := range components {

		switch {
		case index == 0:
			if component == "x" {
				v.MajorLatest = true
				return v
			}
			if !ANumberString(component) {
				return v
			}
			v.Major = component
		case index == 1:
			if component == "x" {
				v.MinorLatest = true
				return v
			}
			if !ANumberString(component) {
				return v
			}
			v.Minor = component
		case index == 2:
			if component == "x" {
				v.PatchLatest = true
				return v
			}
			if !ANumberString(component) {
				return v
			}
			v.Patch = component
		default:
			return v
		}
	}

	return v
}

func FindLatestVersion(list *VersionList, versionName string) string {

	v := generateVersionFromString(versionName)
	if v.MajorLatest {
		return "master"
	}

	var latest *Version
	for _, item := range list.Items {
		if strings.Contains(item.Name, v.VersionWithoutWildcard()) {

			otherV := generateVersionFromString(item.Name)

			if latest == nil {
				latest = otherV
			}

			if latest.Before(otherV) {
				latest = otherV
			}
		}
	}

	return latest.VersionWithoutWildcard()
}

func ANumberString(str string) bool {

	if str == "" {
		return false
	}

	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	}

	return true

}
