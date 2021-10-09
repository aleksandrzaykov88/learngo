// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var IsPalindrome func(string) bool

func init() {
	if len(os.Args) > 1 &&
		(os.Args[1] == "-a" || os.Args[1] == "--ascii") {
		os.Args = append(os.Args[:1], os.Args[2:]...)
		IsPalindrome = func(s string) bool {
			if len(s) <= 1 {
				return true
			}
			j := len(s) - 1
			for i := 0; i < len(s)/2; i++ {
				if s[i] != s[j] {
					return false
				}
				j--
			}
			return true
		}
	} else {
		IsPalindrome = func(s string) bool {
			if len(s) <= 1 {
				return true
			}
			rs := []rune(s)
			j := len(rs) - 1
			for i := 0; i < len(rs)/2; i++ {
				if rs[i] != rs[j] {
					return false
				}
				j--
			}
			return true
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s [-a|--ascii] word1 [word2 [... wordN]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	words := os.Args[1:]
	for _, word := range words {
		fmt.Printf("%5t %q\n", IsPalindrome(word), word)
	}
}
