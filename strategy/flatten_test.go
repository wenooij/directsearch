package strategy

import (
	"iter"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wenooij/directsearch"
)

func testStrategyIter() func() (directsearch.Strategy, bool) {
	var i int64
	return func() (directsearch.Strategy, bool) {
		if i < 3 {
			e := testStrategy{start: i}
			i++
			return e, true
		}
		return nil, false
	}
}

type testStrategy struct{ start int64 }

func (s testStrategy) Strategy() iter.Seq[directsearch.Action] {
	return func(yield func(directsearch.Action) bool) {
		for i := int64(0); i < 3 && yield(testAction(s.start+i)); i++ {
		}
	}
}

type testAction int64

func (i testAction) Action() string { return strconv.FormatInt(int64(i), 10) }

func TestFlatten(t *testing.T) {
	wantActions := []directsearch.Action{
		testAction(0),
		testAction(1),
		testAction(2),
		testAction(1),
		testAction(2),
		testAction(3),
		testAction(2),
		testAction(3),
		testAction(4),
	}
	var gotActions []directsearch.Action
	for a := range Flatten(testStrategyIter()).Strategy() {
		gotActions = append(gotActions, a)
	}

	if diff := cmp.Diff(wantActions, gotActions); diff != "" {
		t.Errorf("TestFlatten() got actions diff (-want, +got):\n%s", diff)
	}
}

func TestFlattenZero(t *testing.T) {
	wantActions := []directsearch.Action{}
	var gotActions []directsearch.Action
	for a := range Flatten(func() (directsearch.Strategy, bool) { return Zero{}, true }).Strategy() {
		gotActions = append(gotActions, a)
	}

	if diff := cmp.Diff(wantActions, gotActions); diff != "" {
		t.Errorf("TestFlattenZero() got actions diff (-want, +got):\n%s", diff)
	}
}

func TestFlattenForeverZero(t *testing.T) {
	wantActions := []directsearch.Action{}
	var gotActions []directsearch.Action
	for a := range FlattenForever(func() directsearch.Strategy { return Zero{} }).Strategy() {
		gotActions = append(gotActions, a)
	}

	if diff := cmp.Diff(wantActions, gotActions); diff != "" {
		t.Errorf("TestFlattenForeverZero() got actions diff (-want, +got):\n%s", diff)
	}
}
