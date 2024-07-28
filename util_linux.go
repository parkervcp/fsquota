package fsquota

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/hashicorp/go-version"
)

type Project struct {
	ID   string
	Name string
}

var expectedKernelVersion = version.Must(version.NewVersion("4.6.0"))

func isKernel46OrLater() (is46OrLater bool, err error) {
	var utsname syscall.Utsname
	if err = syscall.Uname(&utsname); err != nil {
		return
	}

	// get kernel release string
	kernelReleaseBytes := make([]byte, 0, cap(utsname.Release))
	for _, b := range utsname.Release {
		if b == '-' {
			// Replace the first dash with a null-byte
			b = 0x0
		}

		if b == 0x0 {
			break
		}

		kernelReleaseBytes = append(kernelReleaseBytes, byte(b))
	}

	var kernelRelease *version.Version
	if kernelRelease, err = version.NewVersion(string(kernelReleaseBytes)); err != nil {
		return
	}

	is46OrLater = kernelRelease.GreaterThan(expectedKernelVersion) || kernelRelease.Equal(expectedKernelVersion)
	return
}

const passwdFile = "/etc/passwd"
const groupFile = "/etc/group"
const projectFile = "/etc/projid"

func getIDsFromUserOrGroupFile(path string) (ids []uint32, err error) {
	var f *os.File

	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, ":") {
			// Skip malformed line
			continue
		}
		lineParts := strings.SplitN(line, ":", 4)
		if len(lineParts) != 4 {
			continue
		}

		var id uint64
		var parseErr error
		if id, parseErr = strconv.ParseUint(lineParts[2], 10, 32); parseErr == nil {
			ids = append(ids, uint32(id))
		}
	}

	return
}

func getIDsFromProjectFile(path string) (ids []uint32, err error) {
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), ":")
		var id uint64
		if id, err = strconv.ParseUint(lineParts[1], 10, 32); err == nil {
			ids = append(ids, uint32(id))
		}
	}

	// return if unable to close file
	if err = file.Close(); err != nil {
		return
	}

	return
}

func getProjects() (projects []Project, err error) {
	var file *os.File
	if file, err = os.Open(projectFile); err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newProject Project
		lineParts := strings.Split(scanner.Text(), ":")
		newProject.Name = lineParts[0]
		newProject.ID = lineParts[1]

		projects = append(projects, newProject)
	}

	return
}

func LookupProject(name string) (project *Project, err error) {
	var projects []Project
	if projects, err = getProjects(); err != nil {
		return
	}

	for _, proj := range projects {
		if proj.Name == name {
			return &proj, nil
		}
	}

	return project, errors.New("no project found with that name")
}

func LookupProjectID(id string) (project *Project, err error) {
	var projects []Project
	if projects, err = getProjects(); err != nil {
		return
	}

	for _, proj := range projects {
		if proj.ID == id {
			return &proj, nil
		}
	}

	return project, errors.New("no project found with that name")
}
