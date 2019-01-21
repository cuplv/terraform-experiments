provider "bounds" {
  allowance = "1"
}

resource "bounds_thing" "example1" {
  name = "asdf1"
}

resource "bounds_thing" "example2" {
  name = "asdf2"
}
