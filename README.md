# Pulumi dependencies between resources & components

The normal way to define dependencies between resources is to use an output from one resource as
an input to another resource. See the documentation on 
[Inputs & Outputs](https://www.pulumi.com/docs/intro/concepts/inputs-outputs/) for more information.

Pulumi adds the `dependsOn` [resource option](https://www.pulumi.com/docs/intro/concepts/resources/options/dependson/)
in case you want to express a dependency between resources, but there is no output-to-input 
relation you can use.

Leveraging `dependsOn` does seem to have some subleties when needing a dependency **on** a component. 
This repository contains a set of use cases for depending on a component, implemented in any of the
supported programming languages and using:

* component resources in a single language
  * [C#](./single/cs/)
  * [Go](./single/go/)
  * [Python](./single/python/)
  * [Typescript](./single/ts/)
* component resources multi-language packages
  * [C#](./mlc/cs/)
  * [Go](./mlc/go/)
  * [Python](./mlc/python/)
  * [Typescript](./mlc/ts/)
  * [YAML](./mlc/yaml/)
