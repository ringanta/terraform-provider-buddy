resource "buddy_workspace_member" "self" {
  email = "example@example.com"
}

resource "buddy_project_member" "self" {
  project_name      = "my-project"
  member_id         = buddy_workspace_member.self.id
  permission_set_id = 12345 # Developer
}