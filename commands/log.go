package commands

import "fmt"

// LogSink sets a standard basic logging interface by which a discarding logsink or a forwarding
// logsink (unittests, splunk, whatever) can be supplied.  Yep. Go needs a flexible log sink API.
type LogSink interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// QuacksLikeATesting is a mimic of a testing.T to avoid a dependency.  Using Duck-typing we check
// whether it has the few functions we need, and run with it.  ...or Quack with it.  I'm not sure
// if it makes a huge difference, but this avoids a "is it testing or not?" sort of question, and
// allows a user to layer in a Testing sink that isn't called "testing.T".
type QuacksLikeATesting interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
}

// testLogger is a logsink that passes messages to T.Error, or a thing that quacks like a t.Error.
// Private symbol, let's try not to actually need to refer to is aside from how it Quacks.
type testLogger struct {
	T QuacksLikeATesting
}

func (l testLogger) Print(v ...interface{})                 { l.T.Log(v...) }
func (l testLogger) Printf(format string, v ...interface{}) { l.T.Logf(format, v...) }
func (l testLogger) Println(v ...interface{})               { l.T.Logf("%s", fmt.Sprintln(v...)) }

var logsink LogSink = &DiscardingLogger{} // default to running quietly

// SetLogSink allows a caller to override (ie in testing) the logger from a dicard to whatever.  Should be called with things that offer either a LogSink or a testing.T interface (literally something that "quacks like a testing.T", ie has the member functions we need)
//
// ie  SetLogSink(&MyBetterLogSink{...})
//
// ie func TestAThing(t *testing.T) {
//        SetLogSink(t)
//        ...
//    }
//
// Even better:
//
// import "github.com/stretchr/testify/assert"
// ...
// ie func TestAThing(t *testing.T) {
//        assert.Nil(t, SetLogSink(t))
//        ...
//    }
//
// TypeSwitch not used here; we don't it to *be* a LogSink, just act like one
func SetLogSink(newLS interface{}) error {
	if nl, ok := interface{}(newLS).(LogSink); ok {
		logsink = nl
	} else if ql, ok := interface{}(newLS).(QuacksLikeATesting); ok {
		logsink = &testLogger{T: ql}
	} else {
		return fmt.Errorf("Type Error: given type %T should offer interface of LogSink or QuacksLikeATesting (ie testing.T)", newLS)
	}

	return nil // found a matching if/type
}

// DiscardingLogger defines the default logsink that discards all content by doing nothing with it
type DiscardingLogger struct{}

func (d DiscardingLogger) Print(v ...interface{})                 { /* discard */ }
func (d DiscardingLogger) Printf(format string, v ...interface{}) { /* discard */ }
func (d DiscardingLogger) Println(v ...interface{})               { /* discard */ }
