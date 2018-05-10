# encoding/xml package of Go

The standard for namespaces is https://www.w3.org/TR/xml-names/ is incompletely supported by Go lang.
Related issues are mainly listed in [#13400](https://github.com/golang/go/issues/13400).

Purpose of the repository is to ease development and testing of the complete set of patches against the real world.
History of work is out of the scope. Commits will be rewritten with updates of tip of package from [Go source code](https://github.com/golang/go)
and patches as see fit. Merged patches will not appear anymore.

## Issues

Patches are under review (:hammer:) or work in progress (:construction:) (currently empty).  

### :hammer:
  - [disallow empty namespace when prefix is set](https://github.com/golang/go/issues/7535)
  - [recognize xmlns as reserved](https://github.com/golang/go/issues/8068), merged and reverted in Go 1.20 cycle
  - [unexpected behavior of encoder.Indent("", "")](https://github.com/golang/go/issues/13185)
  - [add whitespace normalization from spec](https://github.com/golang/go/issues/20614)
  - [embedded struct with XMLName field not being ignored](https://github.com/golang/go/issues/10538)  
  - [panic when unmarshaling specially crafted structs](https://github.com/golang/go/issues/16497)  

## How to use the repository

Recommended use is to download the patch file of interest from this repository.  
Something like `git am 0001....patch` should apply the patch to your package.  

Other suggested methods are:
    - Create patch files for the desired commits using something like: `git format-patch -1 <commit>`
    - Ignoring history is not recommended but may help
        - Copy and paste files from `src/encoding/xml` in a Go repository.  
        - Use something like `git show <commit>` for the commit of interest and re-do the changes.  

### Maintenance

To ease the display of the fix, history of the folder will be rewritten when a patch is updated.
The baseline commit will be the tip of the `encoding/xml` package without history.

Changes for *.patch files (update, removal, add) will show in history as feasible.
Commits for one issue will be updated (rewritten) or removed when merged.

## History of repository

A general [fix](https://go-review.googlesource.com/c/go/+/109855) based on Go 1.11 is abandoned for several reasons.
Branch `go111` has an archive of this global fix.
