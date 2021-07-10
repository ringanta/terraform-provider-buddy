resource "buddy_workspace_variable" "self" {
  key       = "TF_TEST_VAR"
  value     = "dummy"
  encrypted = true
}
