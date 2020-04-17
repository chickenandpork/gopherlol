package commands

// LogSink sets a standard basic logging interface by which a discarding logsink or a forwarding
// logsink (unittests, splunk, whatever) can be supplied.  Yep. Go needs a flexible log sink API.
type LogSink interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// DiscardingLogger defines the default logsink that discards all content by doing nothing with it
type DiscardingLogger struct{}

func (d DiscardingLogger) Print(v ...interface{})                 { /* discard */ }
func (d DiscardingLogger) Printf(format string, v ...interface{}) { /* discard */ }
func (d DiscardingLogger) Println(v ...interface{})               { /* discard */ }
