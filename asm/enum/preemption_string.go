// Code generated by "stringer -linecomment -type Preemption /home/u/Desktop/go/src/github.com/llir/l/ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/l/ir/enum"

const _Preemption_name = "nonedso_localdso_preemptable"

var _Preemption_index = [...]uint8{0, 4, 13, 28}

func PreemptionFromString(s string) enum.Preemption {
	if len(s) == 0 {
		return 0
	}
	for i := range _Preemption_index[:len(_Preemption_index)-1] {
		if s == _Preemption_name[_Preemption_index[i]:_Preemption_index[i+1]] {
			return enum.Preemption(i)
		}
	}
	panic(fmt.Errorf("unable to locate Preemption enum corresponding to %q", s))
}
