package misc

import "JVM-GO/ch11/native"
import "JVM-GO/ch11/rtda"

func init() {
	native.Register("sun/misc/URLClassPath", "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}
