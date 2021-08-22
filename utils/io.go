package utils

type NilWriter struct{}

func (*NilWriter) Write(b []byte) (int, error) { return len(b), nil }

func (*NilWriter) Close() error { return nil }
