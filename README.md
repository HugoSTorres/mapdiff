# mapdiff
Pretty print a diff of two maps

Recently when running some integration tests in go I came
across the need to see a diff of two objects. In these integration tests,
expected responses are read in from json files and unmarshalled
into maps of type map[string]interface{}. The actual responses are unmarshalled
from the response body into maps as well, and both are finally compared with
reflect.DeepEqual. When the test fails, the objects get logged out so they can
be inspected. This code really just fills the need two compare the two objects
and see their differences more easily than by just logging the maps themselves.
