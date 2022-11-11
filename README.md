# xml of Go supporting namespace

The standard for namespaces is https://www.w3.org/TR/xml-names/ is incompletely supported by Go lang.
Main questions are listed in [#13400](https://github.com/golang/go/issues/13400).

A general [fix](https://go-review.googlesource.com/c/go/+/109855) is now abandoned for several reasons.
It is based on Go 1.11 and cannot be rebased correctly.
It should be considered as an archive.

Issues handled by this commit are below:
  - golang/go#7113
  - golang/go#7535
  - golang/go#8068, merged in Go 1.20 cycle 
  - golang/go#8535, merged in Go 1.20 cycle 
  - golang/go#10538 
  - golang/go#11431 
  - golang/go#13185 
  - golang/go#16497 
  - golang/go#20396, merged in Go 1.20 cycle 
  - golang/go#20614 
  - golang/go#20685, merged in Go 1.20 cycle
