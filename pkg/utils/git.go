/*
Package utils includes helpful utilities for the Create Go App CLI.

Copyright © 2020 Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import "github.com/go-git/go-git/v5"

// GitClone function for `git clone` defined project template
func GitClone(rootFolder, templateName string) error {
	// Clone project template
	_, err := git.PlainClone(rootFolder, false, &git.CloneOptions{
		URL: "https://" + templateName,
	})
	if err != nil {
		return ThrowError("Repository was not cloned!")
	}

	return nil
}
