provider "bounds" {
  budget = 2
}

resource "bounds_thing" "example1" {
  name = "asdf1"
}

resource "bounds_thing" "example2" {
  name = "asdf2"
}

resource "bounds_thing" "example3" {
  name = "asdf3"
}
