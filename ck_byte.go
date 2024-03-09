/*
   Golang test helper library: sztest.
   Copyright (C) 2023, 2024 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package sztest

// Bytef compares the wanted byte against the gotten byte invoking an
// error should they not match.
func (chk *Chk) Bytef(got, want byte, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChkf(got, want, "byte", msgFmt, msgArgs...)
}

// Byte compares the wanted byte against the gotten byte invoking an
// error should they not match.
func (chk *Chk) Byte(got, want byte, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChk(got, want, "byte", msg...)
}

// ByteSlicef checks two byte slices for equality.
func (chk *Chk) ByteSlicef(
	got, want []byte, msgFmt string, msgArgs ...any,
) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()

	return errSlicef(chk,
		got, want, "byte", defaultCmpFunc[byte], msgFmt, msgArgs...,
	)
}

// ByteSlice checks two byte slices for equality.
func (chk *Chk) ByteSlice(got, want []byte, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()

	return errSlice(chk, got, want, "byte", defaultCmpFunc[byte], msg...)
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges.
////////////////////////////////////////////////////////////////

// ByteBoundedf checks value is within specified bounded range.
func (chk *Chk) ByteBoundedf(
	got byte, option BoundedOption, min, max byte, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "byte"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// ByteBounded checks value is within specified bounded range.
func (chk *Chk) ByteBounded(
	got byte, option BoundedOption, min, max byte, msg ...any,
) bool {
	const typeName = "byte"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWnt(typeName, got, want, msg...)
}

// ByteUnboundedf checks value is within specified unbounded range.
func (chk *Chk) ByteUnboundedf(
	got byte, option UnboundedOption, bound byte, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "byte"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// ByteUnbounded checks value is within specified unbounded range.
func (chk *Chk) ByteUnbounded(
	got byte, option UnboundedOption, bound byte, msg ...any,
) bool {
	const typeName = "byte"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWnt(typeName, got, want, msg...)
}
