# generics
Quick experiments with Go generics

- `algebra`, a generic square root function for float, complex and and rational.
- `future`, a concurrent cache ("future cache"). 
- `mapreduce`, parallel Map, Reduce, and ForEach utilities
- `maps`, a map with sorted keys based on a binary tree.
- `metric`, a streamz-style multidimensional variable for production monitoring.
- `number`, generic functions related to numbers (min, max, abs) and a user-defined complex type.
- `oddities`, bugs and quirks.
- `pq`, a priority queue
- `slices`, generic slice utilities, and a user-defined Slice type.
- `stream`, a streams library.
- `striped`, a concurrency-safe map using lock striping, and also a custom hash/eq relation.

First impression:

This is really nice. It addresses the main things I miss about generics:
- being able to change the equivalence relation of map;
- a better API for a priority queue;
- and the ability to generate specialized code for a range of data types
  (though I understand that's not guaranteed).

By comparison, C++'s templates are extremely powerful, but syntactically and
semantically quite complex, and historically the error messages have
been both confusing and too late. (Templates also lead to considerable bloat
in the text segment, but that may be a risk of the Go approach too.)
Java's generics are limited to reference types, and thus are no use for
efficient algorithms on (say) arrays of integers.
Somehow the Go approach seems to do most of what I need while still
feeling simple and easy to use.

Observations:

- Slices can now be implemented in the language, but not without `unsafe` pointer arithmetic.
  The generic slice algorithms work nicely.
  Sorting with a custom order works nicely.

- `unsafe.Sizeof` is disallowed on type parameters, yet it can be simulated using pointer arithmetic
  (though not as a constant expression). Why?

- I often need a hash table with an alternative hash function, for
    - comparing non-canonical pointers (e.g. *big.Int, go/types.Type) by their referent;
    - comparing non-comparable values (such as slices) under the obvious relation;
    - using an alternative comparator (e.g. case insensitive, absolute value) for simple types.
  However, the custom hash function often wants to be at least partly defined in terms of
  the standard hash function, so the latter needs to be exposed somehow; see hacks.RuntimeHash.
  I imagine that could be problematic.
  
- min, max, abs work nicely.

- Go's built-in complex numbers could be satisfactorily replaced by a library.

- The abstract algebraic ring generates pretty good code.
  One can imagine writing some numerical analysis routines this way
  when the algorithm is sufficiently complex that it is best not duplicated.

- I couldn't find a way to achieve ad-hoc polymorphism, that is, defining a generic
  function by cases specialized to each possible type. For example, I don't know
  how to write a generic version of all the math/bits.OnesCount functions that uses
  the most efficient implementation.
  (Typeswitch doesn't handle all possible named types, and using reflect is cheating.)

Users will no doubt build generic libraries of collections, of stream processing functions,
and of numeric analysis routines. The design space for each is large, and I imagine arriving
at simple, efficient, and coherent APIs worthy of the standard library will be an arduous task.
But there is no need to hurry.

