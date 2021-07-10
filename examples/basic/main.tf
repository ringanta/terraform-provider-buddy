resource "buddy_workspace_variable" "self" {
  key       = "TF_TEST_VAR"
  value     = "dummy"
  encrypted = true
}

resource "buddy_project_variable" "self" {
  key       = "TF_TEST_PROJECT_VAR"
  value     = "dummy"
  project   = "unique-payment-code-generator"
  encrypted = true
}
