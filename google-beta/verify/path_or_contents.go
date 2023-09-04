<<<<<<<< HEAD:google-beta/services/compute/path_or_contents.go
package google
========
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package verify
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/verify/path_or_contents.go

import (
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

// If the argument is a path, pathOrContents loads it and returns the contents,
// otherwise the argument is assumed to be the desired contents and is simply
// returned.
//
// The boolean second return value can be called `wasPath` - it indicates if a
// path was detected and a file loaded.
<<<<<<<< HEAD:google-beta/services/compute/path_or_contents.go
func pathOrContents(poc string) (string, bool, error) {
========
func PathOrContents(poc string) (string, bool, error) {
>>>>>>>> e214aac40503b2e28c5bcc73b7c91726014c7e35:google-beta/verify/path_or_contents.go
	if len(poc) == 0 {
		return poc, false, nil
	}

	path := poc
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return path, true, err
		}
	}

	if _, err := os.Stat(path); err == nil {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			return string(contents), true, err
		}
		return string(contents), true, nil
	}

	return poc, false, nil
}
