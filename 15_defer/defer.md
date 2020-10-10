# defer
A `defer` statement defers the execution of a function until the surrounding function returns

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. (see `defer_counting.go`)