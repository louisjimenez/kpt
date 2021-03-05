---
title: "Annotate"
linkTitle: "annotate"
weight: 4
type: docs
description: >
  Set an annotation on one or more resources
---

<!--mdtogo:Short
    Set an annotation on one or more resources
-->

{{< asciinema key="cfg-annotate" rows="10" preload="1" >}}

Annotate sets annotations on resources.

Annotate can be useful when combined with other tools or commands that
read annotations to configure their behavior.

### Examples

{{% hide %}}

<!-- @makeWorkplace @verifyExamples-->
```
# Set up workspace for the test.
TEST_HOME=$(mktemp -d)
cd $TEST_HOME
```

<!-- @fetchPackage @verifyExamples-->
```sh
export SRC_REPO=https://github.com/GoogleContainerTools/kpt.git
kpt pkg get $SRC_REPO/package-examples/helloworld-set@v0.5.0 DIR
```

{{% /hide %}}

<!--mdtogo:Examples-->

<!-- @cfgAnnotate @verifyExamples-->
```sh
# set an annotation on all Resources: 'key: value'
kpt cfg annotate DIR --kv key=value
```

<!-- @cfgAnnotate @verifyExamples-->
```sh
# set an annotation on all Service Resources
kpt cfg annotate DIR --kv key=value --kind Service
```

<!-- @cfgAnnotate @verifyExamples-->
```sh
# set an annotation on the foo Service Resource only
kpt cfg annotate DIR --kv key=value --kind Service --name foo
```

<!-- @cfgAnnotate @verifyExamples-->
```sh
# set multiple annotations
kpt cfg annotate DIR --kv key1=value1 --kv key2=value2
```

<!--mdtogo-->

### Synopsis

<!--mdtogo:Long-->

```
kpt cfg annotate DIR --kv KEY=VALUE...
```

#### Args

```
DIR:
  Path to a package directory
```

<!--mdtogo-->

#### Flags

```sh
--apiVersion
  Only set annotations on resources with this apiVersion.

--kind
  Only set annotations on resources of this kind.

--kv
  The annotation key and value to set.  May be specified multiple times
  to set multiple annotations at once.

--namespace
  Only set annotations on resources in this namespace.

--name
  Only set annotations on resources with this name.

--recurse-subpackages, -R
  Add annotations recursively in all the nested subpackages
```