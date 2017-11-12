package assert

import (
	"fmt"
	"testing"
	"unsafe"
	"sync"
	"strings"
	"bytes"
	"runtime"
)

type Asserts struct {
	t *testing.T
}

// NewAsserts，从T新建一个Asserts对象
func NewAsserts(tObj *testing.T) (a *Asserts){
	a = &Asserts{
		t: tObj,
	}
	return
}

// Fail，执行到这里自动fail
func (a *Asserts) Fail(msg ...string) {
	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = "Failed here."
	}
	a.LogFail(output)
}

// Equals，断言相等
func (a *Asserts) Equals(expect interface{}, actual interface{}, msg ...string) {
	if IsEquals(expect, actual) {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = fmt.Sprintf("Equals Assert Failed: expect %s, actual %s.", expect, actual)
	}
	a.LogFail(output)
}

// NotEquals，断言不相等
func (a *Asserts) NotEquals(expect interface{}, actual interface{}, msg ...string) {
	if !IsEquals(expect, actual) {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = fmt.Sprintf("NotEquals Assert Failed: expect %s, actual %s.", expect, actual)
	}
	a.LogFail(output)
}

// Nil，断言空
func (a *Asserts) Nil(actual interface{}, msg ...string) {
	if IsNil(actual) {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = "Nil Assert Failed."
	}
	a.LogFail(output)
}

// Nil，断言非空
func (a *Asserts) NotNil(actual interface{}, msg ...string) {
	if !IsNil(actual) {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = "NotNil Assert Failed."
	}
	a.LogFail(output)
}

// True，断言True
func (a *Asserts) True(actual bool, msg ...string) {
	if actual == true {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = "True Assert Failed."
	}
	a.LogFail(output)
}

// False，断言False
func (a *Asserts) False(actual bool, msg ...string) {
	if actual == false {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = "False Assert Failed."
	}
	a.LogFail(output)
}

// Error，断言有错误
func (a *Asserts) Error(err error, msg ...string) {
	if err != nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = "Error Assert Failed: expect error, but nil"
	}
	a.LogFail(output)
}

// NotError，断言无错误
func (a *Asserts) NotError(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = formatMessage(msg)
	} else {
		output = fmt.Sprintf("NotError Assert Failed: error is `%s`", err.Error())
	}
	a.LogFail(output)
}

// LogFail, 记录错误
// 如果需要封装类似Equals之类的函数，应当直接调用这个函数，否则堆栈层次可能有问题
func (a *Asserts) LogFail(s string) {
	a.log(s)
	a.t.Fail()
}

// log，改造自testing包下的log方法
func (a *Asserts) log(s string) {
	pt := unsafe.Pointer(a.t)
	pmu := unsafe.Pointer(uintptr(pt) + uintptr(0)) // mu的offset=0
	mu := (*sync.RWMutex)(pmu)

	// c.mu.Lock()
	// defer c.mu.Unlock()
	mu.Lock()
	defer mu.Unlock()

	// c.output = append(c.output, decorate(s)...)
	pOutput := unsafe.Pointer(uintptr(pt) + unsafe.Sizeof(sync.RWMutex{}))
	output := (*[]byte)(pOutput)
	*output = append(*output, decorate(s)...)
}

// decorate, 来自testing包下的decorate方法
// decorate prefixes the string with the file and line of the call site
// and inserts the final newline if needed and indentation tabs for formatting.
func decorate(s string) string {
	_, file, line, ok := runtime.Caller(4) // decorate + log + public function.
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	buf := new(bytes.Buffer)
	// Every line is indented at least one tab.
	buf.WriteByte('\t')
	fmt.Fprintf(buf, "%s:%d: ", file, line)
	lines := strings.Split(s, "\n")
	if l := len(lines); l > 1 && lines[l-1] == "" {
		lines = lines[:l-1]
	}
	for i, line := range lines {
		if i > 0 {
			// Second and subsequent lines are indented an extra tab.
			buf.WriteString("\n\t\t")
		}
		buf.WriteString(line)
	}
	buf.WriteByte('\n')
	return buf.String()
}