/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strcase

import (
	"testing"
)

func TestToSnake(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "test_case"},
		[]string{"TestCase", "test_case"},
		[]string{"Test Case", "test_case"},
		[]string{" Test Case", "test_case"},
		[]string{"Test Case ", "test_case"},
		[]string{" Test Case ", "test_case"},
		[]string{"test", "test"},
		[]string{"test_case", "test_case"},
		[]string{"Test", "test"},
		[]string{"", ""},
		[]string{"ManyManyWords", "many_many_words"},
		[]string{"manyManyWords", "many_many_words"},
		[]string{"AnyKind of_string", "any_kind_of_string"},
		[]string{"numbers2and55with000", "numbers_2_and_55_with_000"},
		[]string{"JSONData", "json_data"},
		[]string{"userID", "user_id"},
		[]string{"AAAbbb", "aa_abbb"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToSnake(in)
		if result != out {
			t.Error("'" + in + "'('" + result + "' != '" + out + "')")
		}
	}
}

func TestToDelimited(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "test@case"},
		[]string{"TestCase", "test@case"},
		[]string{"Test Case", "test@case"},
		[]string{" Test Case", "test@case"},
		[]string{"Test Case ", "test@case"},
		[]string{" Test Case ", "test@case"},
		[]string{"test", "test"},
		[]string{"test_case", "test@case"},
		[]string{"Test", "test"},
		[]string{"", ""},
		[]string{"ManyManyWords", "many@many@words"},
		[]string{"manyManyWords", "many@many@words"},
		[]string{"AnyKind of_string", "any@kind@of@string"},
		[]string{"numbers2and55with000", "numbers@2@and@55@with@000"},
		[]string{"JSONData", "json@data"},
		[]string{"userID", "user@id"},
		[]string{"AAAbbb", "aa@abbb"},
		[]string{"test-case", "test@case"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToDelimited(in, '@')
		if result != out {
			t.Error("'" + in + "' ('" + result + "' != '" + out + "')")
		}
	}
}

func TestToScreamingSnake(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "TEST_CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToScreamingSnake(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestToKebab(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "test-case"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToKebab(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestToScreamingKebab(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "TEST-CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToScreamingKebab(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestToScreamingDelimited(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "TEST.CASE"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToScreamingDelimited(in, '.', true)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
