# SunSpec tools for Go

This package contains Go data types representing the [SunSpec][] information model.
Subpackages implement particular use-cases that rely on this domain model:

 * [/impl](./impl) contains the types used by the remainder of the library
 * [/types](./types) holds the canonical, encoding-neutral SunSpec model definition
   (`Model`, `Block`, `Point`, `Symbol`) plus the global registry. Loaders for the
   two on-disk spec formats live in [/types/xml](./types/xml) (legacy SMDX) and
   [/types/json](./types/json) (the JSON spec maintained at
   [github.com/sunspec/models][SMDX]).
 * [/generators](./generators) contains code generators that emit a Go package per
   SunSpec model: [/generators/json](./generators/json) reads from the JSON spec
   (default), [/generators/xml](./generators/xml) reads from the legacy SMDX XML.
 * [/models](./models) contains one generated package for each SunSpec model
 * [/xml](./xml) contains utilities for exchanging SunSpec data using XML format
 * [/modbus](./modbus) contains utilities for talking to SunSpec devices via Modbus TCP

[SunSpec]: http://sunspec.org/
[SMDX]: https://github.com/sunspec/models

This implementation is based on the original work from [crabmusket/gosunspec](https://github.com/crabmusket/gosunspec/issues/38).

## Generated code

This package uses typesafe representations of each model type. To avoid
excessive manual maintenance, structs for each type of model defined by SunSpec
are generated from the upstream spec. To regenerate, first initialise the
spec submodule:

    git submodule update --init spec

Then regenerate from the JSON spec (default — covers all current models including
the DER 7xx series and 64410-64415 that have no SMDX counterpart):

    make spec

To regenerate from the legacy SMDX/XML spec instead (covers only the older subset
of models):

    make spec-xml

Both generators produce the same `types.Model` shape via `types.RegisterModel`,
so downstream code does not care which source format was used.
