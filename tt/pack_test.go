package tt_test

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/ttcc/tt/tt"
	"testing"
)

type Item struct {
	Name        string
	privateName string
	NameOther   string
}
type Item2 struct {
	Name        string
	privateName string
	Name2       string
	Name3       string
}

func Test_Pack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oo := Item{
			Name:        tt.RandomString(10),
			privateName: "privateName",
			NameOther:   "NameOther",
		}
		b, _ := tt.Pack(oo)
		var aa Item2
		_ = tt.UnpackTo(b, &aa)
		t.Assert(aa.Name, oo.Name)
	})
}

func Test_UnPack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oo := Item{
			Name:        tt.RandomString(10),
			privateName: "privateName",
			NameOther:   "NameOther",
		}
		b, _ := tt.Pack(oo)
		value, _ := tt.Unpack(b)
		t.Assert(value, g.Map{
			"Name":      oo.Name,
			"NameOther": oo.NameOther,
		})
	})
}

func Test_UnPackTo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oo := Item{
			Name:        tt.RandomString(10),
			privateName: "privateName",
			NameOther:   "NameOther",
		}
		b, _ := tt.Pack(oo)
		var aa Item2
		_ = tt.UnpackTo(b, &aa)
		t.Assert(aa.Name, oo.Name)
	})
}
