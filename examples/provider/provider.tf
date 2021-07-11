provider "buddy" {
  buddy_url  = "https://buddy.example.com/api/workspaces/my-workspace" # Alternatively use BUDDY_URL env variable
  token      = "dummyrandomtoken"                                      # Alternatively use BUDDY_TOKEN env variable
  verify_ssl = true                                                    # Alternatively use BUDDY_VERIFY_SSL env variable
}