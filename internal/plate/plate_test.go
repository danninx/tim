package plate

import (
	"testing"

	"github.com/danninx/tim/internal/system"
)

func TestNewFilePlate(t *testing.T) {
	name := "TestName"
	origin := "TestOrigin"
	path := "timTest/files/TestName"

	plate, err := newFilePlate(name, origin, system.GetSystem())

	if err != nil {
		t.Errorf("Creating FilePlate: %v", err)
	}

	if plate.Name() != name || plate.Origin() != origin || plate.Path() != path || plate.Type() != "file" {
		t.Log("Mismatched Values: [actual --- expected]\n")
		t.Logf("NAME\t%v --- %v\n", plate.Name(), name)
		t.Logf("ORIGIN\t%v --- %v\n", plate.Origin(), origin)
		t.Logf("PATH\t%v --- %v\n", plate.Path(), path)
		t.Fail()
	}
}

func TestNewDirPlate(t *testing.T) {
	name := "TestName"
	origin := "TestOrigin"
	path := "timTest/dir/TestName"

	plate, err := newDirPlate(name, origin, system.GetSystem())

	if err != nil {
		t.Errorf("Creating DirPlate: %v", err)
	}

	if plate.Name() != name || plate.Origin() != origin || plate.Path() != path || plate.Type() != "dir" {
		t.Log("Mismatched Values: [actual --- expected]\n")
		t.Logf("NAME\t%v --- %v\n", plate.Name(), name)
		t.Logf("ORIGIN\t%v --- %v\n", plate.Origin(), origin)
		t.Logf("PATH\t%v --- %v\n", plate.Path(), path)
		t.Fail()
	}
}

func TestNewGitPlate(t *testing.T) {
	name := "TestName"
	origin := "TestOrigin"
	path := "timTest/git/TestName"

	plate, err := newGitPlate(name, origin, system.GetSystem())

	if err != nil {
		t.Errorf("Creating GitPlate: %v", err)
	}

	if plate.Name() != name || plate.Origin() != origin || plate.Path() != path || plate.Type() != "git" {
		t.Log("Mismatched Values: [actual --- expected]\n")
		t.Logf("NAME\t%v --- %v\n", plate.Name(), name)
		t.Logf("ORIGIN\t%v --- %v\n", plate.Origin(), origin)
		t.Logf("PATH\t%v --- %v\n", plate.Path(), path)
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	unloadedDir := UnloadedPlate{
		"dir",
		"dirOrigin",
		"dirPath",
	}
	unloadedFile := UnloadedPlate{
		"file",
		"fileOrigin",
		"filePath",
	}
	unloadedGit := UnloadedPlate{
		"git",
		"gitOrigin",
		"gitPath",
	}
	unloaded := []UnloadedPlate{unloadedDir, unloadedFile, unloadedGit}

	for _, u := range unloaded {
		p, err := Load(u.Type, u, system.GetSystem())

		if err != nil {
			t.Errorf("Error loading plate: %v", err)
		}

		if p.Type() != u.Type || p.Name() != u.Type || p.Origin() != u.Origin || p.Path() != u.Path {
			t.Log("Mismatched Values: [actual --- expected]\n")
			t.Logf("NAME\t%v --- %v\n", p.Name(), u.Type)
			t.Logf("TYPE\t%v --- %v\n", p.Type(), u.Type)
			t.Logf("ORIGIN\t%v --- %v\n", p.Origin(), u.Origin)
			t.Logf("PATH\t%v --- %v\n", p.Path(), u.Path)
			t.Fail()
		}
	}
}

func TestLoadInvalid(t *testing.T) {
	unloaded := UnloadedPlate{
		"invalid",
		"something",
		"something",
	}

	_, err := Load("invalid", unloaded, system.GetSystem())
	if err == nil {
		t.Logf("Invalid plate should've thrown an error!")
	}
}

func TestUnload(t *testing.T) {
	loaded, err := NewPlate("git", "loaded", "loadedOrigin", system.GetSystem())
	if err != nil {
		t.Errorf("Error creating plate for unload: %v", err)
	}

	u := Unload(loaded)
	
	if u.Type != "git" || u.Origin != loaded.Origin() || u.Path != loaded.Path() {
		t.Log("Mismatched Values: [actual --- expected]\n")
		t.Logf("TYPE\t%v --- %v\n", u.Type, loaded.Type())
		t.Logf("ORIGIN\t%v --- %v\n", u.Origin, loaded.Origin())
		t.Logf("PATH\t%v --- %v\n", u.Path, loaded.Path())
		t.Fail()
	}
}
