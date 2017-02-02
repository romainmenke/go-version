package version

import (
	"fmt"
	"testing"
)

func TestGetVersionList(t *testing.T) {

	v, err := GetVersionList()
	if err != nil {
		t.Fatal(err)
	}

	_ = v

}

func TestVersion(t *testing.T) {

	if generateVersionFromString("master").Version() != "master" {
		t.Log(generateVersionFromString("master").Version())
		t.Fail()
	}

	if generateVersionFromString("go1.x").Version() != "go1.x" {
		t.Log(generateVersionFromString("go1.x").Version())
		t.Fail()
	}

	if generateVersionFromString("go1").Version() != "go1" {
		t.Log(generateVersionFromString("go1").Version())
		t.Fail()
	}

	if generateVersionFromString("go1.5.x").Version() != "go1.5.x" {
		t.Log(generateVersionFromString("go1.5.x").Version())
		t.Fail()
	}

	if generateVersionFromString("go1.5").Version() != "go1.5" {
		t.Log(generateVersionFromString("go1.5").Version())
		t.Fail()
	}

	if generateVersionFromString("go1.5.4").Version() != "go1.5.4" {
		t.Log(generateVersionFromString("go1.5.4").Version())
		t.Fail()
	}

}

func TestFindLatestVersion(t *testing.T) {

	list, err := GetVersionList()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(FindLatestVersion(list, "go1.2.x"))
	t.Log(FindLatestVersion(list, "go1.3.x"))
	t.Log(FindLatestVersion(list, "go1.4.x"))
	t.Log(FindLatestVersion(list, "go1.5.x"))
	t.Log(FindLatestVersion(list, "go1.6.x"))
	t.Log(FindLatestVersion(list, "go1.7.x"))
	t.Log(FindLatestVersion(list, "go1.x"))
	t.Log(FindLatestVersion(list, "go1"))
	t.Log(FindLatestVersion(list, "master"))

}

func ExampleFindLatestVersion() {

	list, err := GetVersionList()
	if err != nil {
		panic(err)
	}

	fmt.Println(FindLatestVersion(list, "go1.5.x"))
	// Output: go1.5.4

}
