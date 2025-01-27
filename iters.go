package iters

import (
    "iter"
    "slices"
)

func IterValue[T any](v ...T) iter.Seq[T] {
    return slices.Values(v)
}

func IterValue2[T1, T2 any](v1 T1, v2 T2) iter.Seq2[T1, T2] {
    return func(yield func(T1, T2) bool) {
        yield(v1, v2)
    }
}

func IterUnique[T comparable](iter ...iter.Seq[T]) iter.Seq[T] {
    return func(yield func(T) bool) {
        s := map[T]struct{}{}
        for _, i := range iter {
            for v := range i {
                if _, ok := s[v]; !ok {
                    s[v] = struct{}{}
                    if !yield(v) {
                        return
                    }
                }
            }
        }
    }
}

func IterJoin[T any](iter ...iter.Seq[T]) iter.Seq[T] {
    return func(yield func(T) bool) {
        for _, i := range iter {
            for v := range i {
                if !yield(v) {
                    return
                }
            }
        }
    }
}

func IterJoin2[T1, T2 any](iter ...iter.Seq2[T1, T2]) iter.Seq2[T1, T2] {
    return func(yield func(T1, T2) bool) {
        for _, i := range iter {
            for v1, v2 := range i {
                if !yield(v1, v2) {
                    return
                }
            }
        }
    }
}

func IterCollect[T any](iter ...iter.Seq[T]) []T {
    return slices.Collect(IterJoin(iter...))
}
