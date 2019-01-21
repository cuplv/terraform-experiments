terraform-provider-bounds
=========================

This is a mock-up provider which only creates "resources" up to some
limit.  This is the failsafe behavior for the "Terraform with resource
bounds" system, in which Terraform itself would use this provider's
resource declarations to try to avoid going over the bound.

See [`./test.tf`](./test.tf) for an example configuration that will
fail unless the provider's "allowance" is increased to 2.  See
[`./provider.go`](./provider.go) for all implementation details.

This approach, in which the provider itself manages the resource
counter logic, does not *really* work because (as far as I can tell)
the provider has no way to check the list of existing resources before
taking some create/destroy action.  This means that on a second
`terraform apply`, the resource allowance is set back to full even if
existing resource instances are already consuming it.


Next step
---------

The next prototyping step involves modifying Terraform to manage the
resource counter itself.  The provider will be responsible for
answering two sorts of queries:

1. `Allowance(*ResourceConfig) int`, which gives the total allowed
   resource bounds defined by a provider configuration.
2. `ResourceCost(*ResourceConfig) int`, which gives the cost of a
   configured resource.

In order to avoid a change to the `terraform.ResourceProvider`
interface, I think these answers can be encoded as warning strings
returned by `Validate` and `ValidateResource`.

Resources in this prototype are represented as `int`; in the full
form, providers should be able to define their own domain-specific
resource datatypes.
