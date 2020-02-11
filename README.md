# gomodulesdemo
Go Modules: Creating projects independent of $GOPATH

For this Week's reflection learning I demonstrated a project to show my learning. Here are the details:

Go Modules: Creating projects independent of $GOPATH

All projects had to be created inside the $GOPATH. It is used by the Go compiler to search for dependencies when building your Go application. $GOPATH contains source code and binaries.

Shortcomings before Go Modules?
    You have to create all your projects inside a single folder where $GOPATHis defined. (Some feel this is a good thing)
    Versioning Go packages were not supported. It doesn’t allow you to specify a particular version for a Go package like you do in package.json. Also, you couldn’t use two different versions of the same package.
    All external packages were kept in a vendor folder and pushed to the server. Dependency Hell can occur.


What exactly is a Go Module?
    A module is a collection of related Go packages that are versioned together as a single unit. 
   Modules record precise dependency requirements and create reproducible builds.

Working Demo of Go Modules
