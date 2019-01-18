terraform-provider-bounds
=========================

This is a mock-up provider which only creates "resources" up to some
limit.  This is the failsafe behavior for the "Terraform with resource
bounds" system, in which Terraform itself would use this provider's
resource declarations to try to avoid going over the bound.
