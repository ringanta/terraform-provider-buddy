---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "buddy_workspace_member Resource - terraform-provider-buddy"
subcategory: ""
description: |-
  buddy_workspace_member manages member on a Buddy workspace.
  A new member can be invited into a workspace using their email address.
---

# buddy_workspace_member (Resource)

`buddy_workspace_member` manages member on a Buddy workspace.

A new member can be invited into a workspace using their email address.

## Example Usage

```terraform
resource "buddy_workspace_member" "self" {
  email = "example@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **email** (String) Email address to be invited into Buddy workspace

### Optional

- **admin** (Boolean) Flag to indicate whether member has admin right
- **id** (String) The ID of this resource.

### Read-Only

- **name** (String) Member name

## Import

Import is supported using the following syntax:

```shell
# import existing workspace member using its ID
# You can get a member ID via user profile page under the People menu (last part of the URL).
terraform import buddy_workspace_member.self  12345
```
