# pghash
Go clone of Postgres's builtin hash_any, which in turn is an clone of http://burtleburtle.net/bob/hash/doobs.html

Postgres version implements 4 versions
    
	Optimized-for-word-aligned input -----| X  |----  Big-Endian
	Non-optimized / generic input    -----| X  |----  Little-endian

This repo only implements the generic+little-endian combo.

# Benchmark

	BenchmarkPghash-8   	20000000	        88.7 ns/op	       0 B/op	       0 allocs/op


