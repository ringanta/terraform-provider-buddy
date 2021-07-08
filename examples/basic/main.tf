resource "buddy_global_variable" "self" {
  key       = "TF_TEST_VAR"
  value     = "dummy"
  encrypted = true
}
