// +build windows

package supervisord

func Deamonize(proc func()) {
	proc()
}
