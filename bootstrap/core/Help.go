package core

// Debug 重新封装Debug方法
// @param args ...interface{}
func Debug(args ...interface{}) {
	Log.Debug(args)
}
