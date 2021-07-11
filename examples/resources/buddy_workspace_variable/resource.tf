resource "buddy_workspace_variable" "self" {
  key       = "TEST_WORKSPACE_VAR"
  value     = "dummy"
  encrypted = true
}