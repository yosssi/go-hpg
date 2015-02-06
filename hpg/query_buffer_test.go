package hpg

import "testing"

func Test_queryBuffer_append(t *testing.T) {
	bf := new(queryBuffer)

	bf.append("testKey", "testValue")

	want := "&testKey=testValue"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendString_zero(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendString("testKey", "")

	want := ""

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendString(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendString("testKey", "testValue")

	want := "&testKey=testValue"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendStringSlice_zero(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendStringSlice("testKey", nil)

	want := ""

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendStringSlice(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendStringSlice("testKey", []string{"testValue"})

	want := "&testKey=testValue"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendInt_zero(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendInt("testKey", 0)

	want := ""

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendInt(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendInt("testKey", 1)

	want := "&testKey=1"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendFloatPtr_zero(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendFloatPtr("testKey", nil)

	want := ""

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendFloatPtr(t *testing.T) {
	bf := new(queryBuffer)

	f := 1.0

	bf.appendFloatPtr("testKey", &f)

	want := "&testKey=1.0000000000"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendBool_zero(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendBool("testKey", false)

	want := ""

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func Test_queryBuffer_appendBool(t *testing.T) {
	bf := new(queryBuffer)

	bf.appendBool("testKey", true)

	want := "&testKey=1"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
