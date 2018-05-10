# encoding/xml package of Go

The standard for namespaces is https://www.w3.org/TR/xml-names/ is incompletely supported by Go lang.
Related issues are mainly listed in [#13400](https://github.com/golang/go/issues/13400).

Purpose of the repository is to ease development and testing of the complete set of patches against the real world.
History of work is out of the scope and updates of the package occur for language changes and other reasons.
Commits will be rewritten with updates of tip of package from [Go source code](https://github.com/golang/go) and patches when updated.

Changes for patches can be followed on [Gerrit](https://go-review.googlesource.com/q/encoding/xml+status:open).

- [recognize xmlns as reserved](https://go.dev/issue/7535) ([gerrit-patch](https://go-review.googlesource.com/c/go/+/107755))  
- [disallow empty namespace when prefix is set](https://go.dev/issue/8068)  ([gerrit-patch](https://go-review.googlesource.com/c/go/+/105636))  
- [support xml namespace prefix](https://go.dev/issue/9519) ([gerrit-patch](https://go-review.googlesource.com/c/go/+/116056))  
- [embedded struct with XMLName field not being ignored](https://go.dev/issue/10538) ([gerrit-patch](https://go-review.googlesource.com/c/go/+/108616))  
- [indent even when both indent and prefix are empty](https://go.dev/issue/13185) ([gerrit-patch](https://go-review.googlesource.com/c/go/+/108797))  
- [add whitespace normalization from spec](https://go.dev/issue/20614) ([gerrit-patch](https://go-review.googlesource.com/c/go/+/104655))  

Some comments in related issues might also be of value.

## How to

### Using patch files

Download the patch file of interest from this repository by accessing [Gerrit](https://go-review.googlesource.com/c/go/).  
Something like `git am 0001....patch` should apply the patch to your go source code.  

### Local development and evaluation

Git clone as usual. The repository is an ordinary module which can be imported as usual.

## History of repository

- A folder using the src/ directory to match patches was cumbersome to use since the availability of Go workspace.

- A general [fix](https://go-review.googlesource.com/c/go/+/109855) based on Go 1.11 is abandoned for several reasons.
Branch `go111` has an archive of this global fix.
