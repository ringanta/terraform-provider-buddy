# import existing workspace variable using its ID
# Variable ID can be retrieve via Buddy the following API https://buddy.works/docs/api/general/environment-variables/list-environment-variables
# Use this jq command to filter the result by a certain variable
#   jq '.variables[] | select(.key == "<VARIABLE_NAME>")'
terraform import buddy_workspace_variable.self 12345