resource "buddy_project_variable" "self" {
  key       = "TEST_PROJECT_VAR"
  value     = "dummy"
  project   = "example-project"
  encrypted = true
}