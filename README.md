# xml namespace encoder of Go

The standard for namespaces is https://www.w3.org/TR/xml-names/

go lang encoding/xml contains several issues related to namespace standards summarized in https://github.com/golang/go/issues/13400

A general fix has been submitted, cf. https://go-review.googlesource.com/c/go/+/109855 for many of them.
It won't be included in Go 1.11.
If you need to find what is fixed, please read the CL details.
