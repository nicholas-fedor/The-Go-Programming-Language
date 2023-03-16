package main

import "testing"

func TestAreAnagram(t *testing.T) {
	if checkAnagram("listen", "silent") != true {
		t.Error(`"listen", "silent"`)
		}
	if checkAnagram("test", "ttew") != false {
		t.Error(`"test", "ttew"`)
		}
	if checkAnagram("geeksforgeeks", "forgeeksgeeks") != true {
		t.Error(`"geeksforgeeks", "forgeeksgeeks"`)
		}
	if checkAnagram("triangle", "integral") != true {
		t.Error(`"triangle", "integral"`)
		}
	if checkAnagram("abd", "acb") != false {
		t.Error(`"abd", "acb"`)
		}
}