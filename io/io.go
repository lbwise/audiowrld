package io

type OutputDevice interface {
	Write([]byte) (int, error)
}
