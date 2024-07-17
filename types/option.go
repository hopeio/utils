package types

import (
	"encoding/json"
)

// 返回option 返回时会有两次复制value,后续使用还有可能更多次,自行选择用不用
type Option[T any] struct {
	value T
	ok    bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{value: v, ok: true}
}

func None[T any]() Option[T] {
	return Option[T]{ok: false}
}
func Nil[T any]() Option[T] {
	return Option[T]{ok: false}
}

func (opt *Option[T]) Val() (T, bool) {
	return opt.value, opt.ok
}

func (opt *Option[T]) Get() (T, bool) {
	return opt.value, opt.ok
}

func (opt *Option[T]) IsNone() bool {
	return !opt.ok
}

func (opt *Option[T]) IsSome() bool {
	return opt.ok
}

func (opt *Option[T]) Unwrap() T {
	if opt.IsNone() {
		panic("Attempted to unwrap an empty Option.")
	}
	return opt.value
}

func (opt *Option[T]) UnwrapOr(def T) T {
	if opt.IsSome() {
		return opt.Unwrap()
	}
	return def
}

func (opt *Option[T]) UnwrapOrElse(fn func() T) T {
	if opt.IsSome() {
		return opt.Unwrap()
	}
	return fn()
}

func MapOption[T any, R any](opt Option[T], fn func(T) R) Option[R] {
	if !opt.IsSome() {
		return None[R]()
	}
	return Some(fn(opt.Unwrap()))
}

func (opt *Option[T]) IfSome(action func(value T)) {
	if opt.ok {
		action(opt.value)
	}
}

func (opt *Option[T]) IfNone(action func()) {
	if !opt.ok {
		action()
	}
}

func (opt *Option[T]) Next() *Option[T] {
	return opt
}

func (opt *Option[T]) MarshalJSON() ([]byte, error) {
	if opt.ok {
		return json.Marshal(opt.value)
	}
	return []byte("null"), nil
}

func (opt *Option[T]) UnmarshalJSON(data []byte) error {
	if len(data) < 5 && string(data) == "null" {
		opt.ok = false
		return nil
	}
	opt.ok = true
	return json.Unmarshal(data, &opt.value)
}

type OptionP[T any] struct {
	value T
}

func SomeP[T any](v T) *OptionP[T] {
	return &OptionP[T]{value: v}
}

func NoneP[T any]() *OptionP[T] {
	return nil
}
func NilP[T any]() *OptionP[T] {
	return nil
}

func (opt *OptionP[T]) Val() (T, bool) {
	if opt == nil {
		return *new(T), false
	}
	return opt.value, true
}

func (opt *OptionP[T]) Get() (T, bool) {
	if opt == nil {
		return *new(T), false
	}
	return opt.value, true
}

func (opt *OptionP[T]) IsNone() bool {
	return opt == nil
}

func (opt *OptionP[T]) IsSome() bool {
	return opt != nil
}

func (opt *OptionP[T]) Unwrap() T {
	if opt.IsNone() {
		panic("Attempted to unwrap an empty OptionP.")
	}
	return opt.value
}

func (opt *OptionP[T]) UnwrapOr(def T) T {
	if opt.IsSome() {
		return opt.Unwrap()
	}
	return def
}

func (opt *OptionP[T]) UnwrapOrElse(fn func() T) T {
	if opt.IsSome() {
		return opt.Unwrap()
	}
	return fn()
}

func MapOptionP[T any, R any](opt *OptionP[T], fn func(T) R) *OptionP[R] {
	if !opt.IsSome() {
		return NoneP[R]()
	}
	return SomeP(fn(opt.Unwrap()))
}

func (opt *OptionP[T]) IfSome(action func(value T)) {
	if opt.IsSome() {
		action(opt.value)
	}
}

func (opt *OptionP[T]) IfNone(action func()) {
	if opt.IsNone() {
		action()
	}
}

func (opt *OptionP[T]) Next() *OptionP[T] {
	return opt
}

func (opt *OptionP[T]) MarshalJSON() ([]byte, error) {
	if opt.IsSome() {
		return json.Marshal(opt.value)
	}
	return []byte("null"), nil
}

func (opt *OptionP[T]) UnmarshalJSON(data []byte) error {
	if len(data) < 5 && string(data) == "null" {
		return nil
	}
	return json.Unmarshal(data, &opt.value)
}
