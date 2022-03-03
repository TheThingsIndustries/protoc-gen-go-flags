# protoc-gen-go-flags

> Protoc plugin for generating CLI flags for command-line interfaces written in Go.

## Background

This plugin generates cli flags from protobuf message definitions. The generated flags are POSIX/GNU-style --flags based on the `github.com/spf13/pflag` project.  This allows us to replace the current reflect based flag generation for proto structs in The Things Stack V3.

## Usage in Proto Code

```proto
syntax = "proto3";

import "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations.proto";

package thethings.flags.example;

option go_package = "github.com/TheThingsIndustries/protoc-gen-go-flags/example";

message MessageWithFieldsWithCustomFlags {
  option (thethings.flags.message) = { select: true, set: true };

  double double_value = 1;
  repeated double double_values = 2;

  bytes hex_bytes_value = 31 [
    (thethings.flags.field) = {
      set_flag_new_func: "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin.NewHexBytesFlag",
      set_flag_getter_func: "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin.GetBytes"
    }
  ];

  repeated bytes hex_bytes_values = 32 [
    (thethings.flags.field) = {
      set_flag_new_func: "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin.NewHexBytesSliceFlag",
      set_flag_getter_func: "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin.GetBytesSlice"
    }
  ];
}

message MessageWithEmbeddedMessage {
  option (thethings.flags.message) = { select: true, set: true };

  message EmbeddedMessage {
      option (thethings.flags.message) = { select: true, set:true};

      double double_value = 1;
  }

  EmbeddedMessage embedded = 1;
}


```

## Generating Go Code

```bash
$ protoc -I ./path/to -I . \
  --go_opt=paths=source_relative --go_out=./path/to \
  --go-flags_opt=paths=source_relative --go-flags_out=./path/to \
  ./path/to/*.proto
```

## Usage in Go Code

```go
  selectFlags := &pflag.FlagSet{}
  setFlags := &pflag.FlagSet{}
  // Add generated flags to a flag set.
  AddSelectFlagsForMessageWithEmbeddedMessage(selectFlags, "")
  AddSetFlagsForMessageWithEmbeddedMessage(setFlags, "")
```

```go
  // Get paths that are selected from select flags.
  paths, err := PathsFromSelectFlagsForMessageWithEmbeddedMessage(selectFlags, "")

  // Get paths and set the struct fields from flag values
  var embeddedMessage MessageWithEmbeddedMessage
  paths, err := embeddedMessage.SetFromFlags(setFlags, "")
```

## Contributing

We do not accept external issues with feature requests. This plugin only supports what we actually use ourselves at The Things Industries.

We do not accept external pull requests with new features, but everyone is free to fork it and add features in their own fork.

We do accept issues with bug reports and pull requests with bug fixes.
