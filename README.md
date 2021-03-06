# gopwt

[![Travis-CI Build Status](https://travis-ci.org/ToQoz/gopwt.svg?branch=master)](https://travis-ci.org/ToQoz/gopwt)
[![AppVeyor Build Status](https://ci.appveyor.com/api/projects/status/5m2c6xiwgs8yamj1/branch/master?svg=true&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending)](https://ci.appveyor.com/project/ToQoz/gopwt/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ToQoz/gopwt)](https://goreportcard.com/report/github.com/ToQoz/gopwt)
[![codecov](https://codecov.io/gh/ToQoz/gopwt/branch/master/graph/badge.svg)](https://codecov.io/gh/ToQoz/gopwt)

PowerAssert library for golang.

<img src="https://i.gyazo.com/fde9f5c049a94b02019a578d4b7e19c5.png" width="713" alt="screenshot">

![logo](http://toqoz.net/art/images/gopwt.svg)

## Supported go versions

See [.travis.yml](/.travis.yml)

## Getting Started

Check [online demo](http://gopwt.toqoz.net) at first!

## Install

```
$ go get -u github.com/ToQoz/gopwt/...
```

## Try

go to project:

```
$ mkdir -p $GOPATH/src/github.com/$(whoami)/gopwtexample
$ cd $GOPATH/src/github.com/$(whoami)/gopwtexample
```

write main_test.go

```
$ cat <<EOF > main_test.go
package main

import (
	"flag"
	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower()
	os.Exit(m.Run())
}

func TestFoo(t *testing.T) {
	a := "xxx"
	b := "yyy"
	assert.OK(t, a == b, "a should equal to b")
}
EOF
```

run tests:

```
$ go test
--- FAIL: TestFoo (0.00s)
	assert.go:85: FAIL main_test.go:20
		assert.OK(t, a == b, "a should equal to b")
		             | |  |
		             | |  "yyy"
		             | false
		             "xxx"
		
		Assertion messages:
			- a should equal to b
		
		--- [string] b
		+++ [string] a
		@@ -1,1 +1,1@@
		-yyy
		+xxx
		
		
FAIL
exit status 1
FAIL	github.com/gopwter/gopwtexample	0.057s
exit status 1
exit status 1
FAIL	github.com/gopwter/gopwtexample	1.330s
```

## Example
```go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/ToQoz/gopwt"
	"github.com/ToQoz/gopwt/assert"
)

func TestMain(m *testing.M) {
	flag.Parse()
	gopwt.Empower() // magic happens
	os.Exit(m.Run())
}

func TestWithMessage(t *testing.T) {
	var receiver *struct{}
	receiver = nil
	assert.OK(t, receiver != nil, "receiver should not be nil")
}

func TestWithTestingRun(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		assert.OK(t, 1 == 2)
	})
	t.Run("B", func(t *testing.T) {
		assert.OK(t, 3 == 4)
	})
}

func TestBasicLit(t *testing.T) {
	assert.OK(t, "a" == "b")
	assert.OK(t, 1 == 2)

	a := 1
	b := 2
	c := 3
	assert.OK(t, a+c == b)
	assert.OK(t, (a+c)+a == b)
	assert.OK(t, `foo
bar` == "bar")
}

func TestUnaryExpr(t *testing.T) {
	assert.OK(t, +1 == -1)

	a := "a"
	b := "a"
	assert.OK(t, &a == &b)

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func() {
		assert.OK(t, <-ch == 0)
		wg.Done()
	}()
	ch <- 1
	wg.Wait()
	close(ch)
}

func TestStringDiff(t *testing.T) {
	assert.OK(t, "supersoper" == "supersuper")

	expected := `foo
baz
bar2
bar`
	got := `foo
baz
bar
bar`
	assert.OK(t, got == expected)

	expected = `<div>
<div>
foo
</div>
</div>`
	got = `<div>
bar
</div>`
	assert.OK(t, got == expected)
}

func TestMapType(t *testing.T) {
	k := "b--------key"
	v := "b------value"
	assert.OK(t, reflect.DeepEqual(map[string]string{}, map[string]string{
		"a": "a",
		k:   v,
	}))
}

func TestArrayType(t *testing.T) {
	index := 1
	assert.OK(t, []int{
		1,
		2,
	}[index] == 3)
}

func TestStructType(t *testing.T) {
	foox := "foo------x"
	assert.OK(t, reflect.DeepEqual(struct{ Name string }{foox}, struct{ Name string }{"foo"}))
	assert.OK(t, reflect.DeepEqual(struct{ Name string }{Name: foox}, struct{ Name string }{Name: "foo"}))
}

func TestNestedCallExpr(t *testing.T) {
	rev := func(a bool) bool {
		return !a
	}

	assert.OK(t, rev(rev(rev(true))))
}

func TestCallWithNonIdempotentFunc(t *testing.T) {
	i := 0
	incl := func() int {
		i++
		return i
	}

	assert.OK(t, incl()+incl() == incl()+incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, (incl() == incl()) != (incl() == incl()))

	i2 := 0
	incl2 := func(i3 int) int {
		i2 += i3
		return i2
	}

	assert.OK(t, incl2(incl2(2)) == 10)
}

func TestPkgValue(t *testing.T) {
	assert.OK(t, sql.ErrNoRows == fmt.Errorf("error"))
}

func TestSelectorExpr(t *testing.T) {
	type user struct{ name string }
	u := user{"gopwt"}
	assert.OK(t, u.name == "no-name")
}
```

```
$ go test
--- FAIL: TestWithMessage (0.00s)
	assert.go:85: FAIL main_test.go:25
		assert.OK(t, receiver != nil, "receiver should not be nil")
		             |        |  |
		             |        |  <nil>
		             |        false
		             (*struct {})(nil)
		
		Assertion messages:
			- receiver should not be nil
		
--- FAIL: TestWithTestingRun (0.00s)
    --- FAIL: TestWithTestingRun/A (0.00s)
    	assert.go:85: FAIL main_test.go:30
    		assert.OK(t, 1 == 2)
    		               |
    		               false
    		
    		--- [int] 2
    		+++ [int] 1
    		@@ -1,1 +1,1@@
    		-2
    		+1
    		
    		
    --- FAIL: TestWithTestingRun/B (0.00s)
    	assert.go:85: FAIL main_test.go:33
    		assert.OK(t, 3 == 4)
    		               |
    		               false
    		
    		--- [int] 4
    		+++ [int] 3
    		@@ -1,1 +1,1@@
    		-4
    		+3
    		
    		
--- FAIL: TestBasicLit (0.00s)
	assert.go:85: FAIL main_test.go:38
		assert.OK(t, "a" == "b")
		                 |
		                 false
		
		--- [string] "b"
		+++ [string] "a"
		@@ -1,1 +1,1@@
		-b
		+a
		
		
	assert.go:85: FAIL main_test.go:39
		assert.OK(t, 1 == 2)
		               |
		               false
		
		--- [int] 2
		+++ [int] 1
		@@ -1,1 +1,1@@
		-2
		+1
		
		
	assert.go:85: FAIL main_test.go:44
		assert.OK(t, a+c == b)
		             ||| |  |
		             ||| |  2
		             ||| false
		             ||3
		             |4
		             1
		
		--- [int] b
		+++ [int] a + c
		@@ -1,1 +1,1@@
		-2
		+4
		
		
	assert.go:85: FAIL main_test.go:45
		assert.OK(t, (a+c)+a == b)
		              ||| || |  |
		              ||| || |  2
		              ||| || false
		              ||| |1
		              ||| 5
		              ||3
		              |4
		              1
		
		--- [int] b
		+++ [int] (a + c) + a
		@@ -1,1 +1,1@@
		-2
		+5
		
		
	assert.go:85: FAIL main_test.go:46
		assert.OK(t, "foo\nbar" == "bar")
		             |          |
		             |          false
		             "foo
		             bar"
		
		--- [string] "bar"
		+++ [string] "foo\nbar"
		@@ -1,1 +1,2@@
		+foo
		 bar
		
		
--- FAIL: TestUnaryExpr (0.00s)
	assert.go:85: FAIL main_test.go:51
		assert.OK(t, +1 == -1)
		             |  |  |
		             |  |  -1
		             |  false
		             1
		
		--- [int] -1
		+++ [int] +1
		@@ -1,1 +1,1@@
		--1
		+1
		
		
	assert.go:85: FAIL main_test.go:55
		assert.OK(t, &a == &b)
		             || |  ||
		             || |  |"a"
		             || |  (*string)(0xc42008f450)
		             || false
		             |"a"
		             (*string)(0xc42008f440)
		
		--- [*string] &b
		+++ [*string] &a
		@@ -1,1 +1,1@@
		&"a"
		
	assert.go:85: FAIL main_test.go:62
		assert.OK(t, <-ch == 0)
		             | |  |
		             | |  false
		             | (chan int)(0xc42008c8a0)
		             1
		
		--- [int] 0
		+++ [int] <-ch
		@@ -1,1 +1,1@@
		-0
		+1
		
		
--- FAIL: TestStringDiff (0.00s)
	assert.go:85: FAIL main_test.go:71
		assert.OK(t, "supersoper" == "supersuper")
		                          |
		                          false
		
		--- [string] "supersuper"
		+++ [string] "supersoper"
		@@ -1,1 +1,1@@
		supers[-u-]{+o+}per
		
	assert.go:85: FAIL main_test.go:81
		assert.OK(t, got == expected)
		             |   |  |
		             |   |  "foo
		             |   |  baz
		             |   |  bar2
		             |   |  bar"
		             |   false
		             "foo
		             baz
		             bar
		             bar"
		
		--- [string] expected
		+++ [string] got
		@@ -1,4 +1,4@@
		foo
		baz
		bar[-2-]
		bar
		
	assert.go:85: FAIL main_test.go:91
		assert.OK(t, got == expected)
		             |   |  |
		             |   |  "<div>
		             |   |  <div>
		             |   |  foo
		             |   |  </div>
		             |   |  </div>"
		             |   false
		             "<div>
		             bar
		             </div>"
		
		--- [string] expected
		+++ [string] got
		@@ -1,5 +1,3@@
		 <div>
		-<div>
		-foo
		-</div>
		+bar
		 </div>
		
		
--- FAIL: TestMapType (0.00s)
	assert.go:85: FAIL main_test.go:97
		assert.OK(t, reflect.DeepEqual(map[string]string{}, map[string]string{"a": "a", k: v}))
		             |                                                                  |  |
		             |                                                                  |  "b------value"
		             |                                                                  "b--------key"
		             false
		
		--- [map[string]string] map[string]string{"a": "a", k: v}
		+++ [map[string]string] map[string]string{}
		@@ -1,3 +1,1@@
		-map[string]string{ "a":            "a",
		-                   "b--------key": "b------value",
		-}
		+map[string]string{}
		
		
--- FAIL: TestArrayType (0.00s)
	assert.go:85: FAIL main_test.go:105
		assert.OK(t, []int{1, 2}[index] == 3)
		                        ||      |
		                        ||      false
		                        |1
		                        2
		
		--- [int] 3
		+++ [int] []int{1, 2}[index]
		@@ -1,1 +1,1@@
		-3
		+2
		
		
--- FAIL: TestStructType (0.00s)
	assert.go:85: FAIL main_test.go:113
		assert.OK(t, reflect.DeepEqual(struct{ Name string }{foox}, struct{ Name string }{"foo"}))
		             |                                       |
		             |                                       "foo------x"
		             false
		
		--- [struct { Name string }] struct{ Name string }{"foo"}
		+++ [struct { Name string }] struct{ Name string }{foox}
		@@ -1,3 +1,3@@
		struct { Name string }{
		  Name: "foo{+------x+}",
		}
		
	assert.go:85: FAIL main_test.go:114
		assert.OK(t, reflect.DeepEqual(struct{ Name string }{Name: foox}, struct{ Name string }{Name: "foo"}))
		             |                                             |
		             |                                             "foo------x"
		             false
		
		--- [struct { Name string }] struct{ Name string }{Name: "foo"}
		+++ [struct { Name string }] struct{ Name string }{Name: foox}
		@@ -1,3 +1,3@@
		struct { Name string }{
		  Name: "foo{+------x+}",
		}
		
--- FAIL: TestNestedCallExpr (0.00s)
	assert.go:85: FAIL main_test.go:122
		assert.OK(t, rev(rev(rev(true))))
		             |   |   |   |
		             |   |   |   true
		             |   |   false
		             |   true
		             false
		
--- FAIL: TestCallWithNonIdempotentFunc (0.00s)
	assert.go:85: FAIL main_test.go:132
		assert.OK(t, incl()+incl() == incl()+incl())
		             |     ||      |  |     ||
		             |     ||      |  |     |4
		             |     ||      |  |     7
		             |     ||      |  3
		             |     ||      false
		             |     |2
		             |     3
		             1
		
		--- [int] incl() + incl()
		+++ [int] incl() + incl()
		@@ -1,1 +1,1@@
		-7
		+3
		
		
	assert.go:85: FAIL main_test.go:133
		assert.OK(t, incl() == incl())
		             |      |  |
		             |      |  6
		             |      false
		             5
		
		--- [int] incl()
		+++ [int] incl()
		@@ -1,1 +1,1@@
		-6
		+5
		
		
	assert.go:85: FAIL main_test.go:134
		assert.OK(t, incl() == incl())
		             |      |  |
		             |      |  8
		             |      false
		             7
		
		--- [int] incl()
		+++ [int] incl()
		@@ -1,1 +1,1@@
		-8
		+7
		
		
	assert.go:85: FAIL main_test.go:135
		assert.OK(t, incl() == incl())
		             |      |  |
		             |      |  10
		             |      false
		             9
		
		--- [int] incl()
		+++ [int] incl()
		@@ -1,1 +1,1@@
		-10
		+9
		
		
	assert.go:85: FAIL main_test.go:136
		assert.OK(t, incl() == incl())
		             |      |  |
		             |      |  12
		             |      false
		             11
		
		--- [int] incl()
		+++ [int] incl()
		@@ -1,1 +1,1@@
		-12
		+11
		
		
	assert.go:85: FAIL main_test.go:137
		assert.OK(t, (incl() == incl()) != (incl() == incl()))
		              |      |  |       |   |      |  |
		              |      |  |       |   |      |  16
		              |      |  |       |   |      false
		              |      |  |       |   15
		              |      |  |       false
		              |      |  14
		              |      false
		              13
		
	assert.go:85: FAIL main_test.go:145
		assert.OK(t, incl2(incl2(2)) == 10)
		             |     |         |
		             |     |         false
		             |     2
		             4
		
		--- [int] 10
		+++ [int] incl2(incl2(2))
		@@ -1,1 +1,1@@
		-10
		+4
		
		
--- FAIL: TestPkgValue (0.00s)
	assert.go:85: FAIL main_test.go:149
		assert.OK(t, sql.ErrNoRows == fmt.Errorf("error"))
		                 |         |  |
		                 |         |  &errors.errorString{s:"error"}
		                 |         false
		                 &errors.errorString{s:"sql: no rows in result set"}
		
		--- [*errors.errorString] fmt.Errorf("error")
		+++ [*errors.errorString] sql.ErrNoRows
		@@ -1,3 +1,3@@
		 &errors.errorString{
		-  s: "error",
		+  s: "sql: no rows in result set",
		 }
		
		
--- FAIL: TestSelectorExpr (0.00s)
	assert.go:85: FAIL main_test.go:155
		assert.OK(t, u.name == "no-name")
		             | |    |
		             | |    false
		             | "gopwt"
		             main.user{name:"gopwt"}
		
		--- [string] "no-name"
		+++ [string] u.name
		@@ -1,1 +1,1@@
		-no-name
		+gopwt
		
		
FAIL
exit status 1
FAIL	github.com/ToQoz/gopwt/_example	0.074s
exit status 1
exit status 1
FAIL	github.com/ToQoz/gopwt/_example	1.277s
```

## Tips

### Off gopwt.Empoer()

If you don't need gopwt or need to stop gopwt, you can!

```
GOPWT_OFF=1 go test -covermode count
PASS
coverage: 60.8% of statements
ok      github.com/ToQoz/gopwt  1.008s
```

## See Also

- [Why does Go not have assertions?](http://golang.org/doc/faq#assertions)
- http://dontmindthelanguage.wordpress.com/2009/12/11/groovy-1-7-power-assert/
- https://github.com/twada/power-assert
