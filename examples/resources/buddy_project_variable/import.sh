# import existing project variable using its ID
# Variable ID can be retrieve via Buddy API https://buddy.works/docs/api/general/environment-variables/list-environment-variables
# Use this jq command to filter the result by a certain variable
#   jq '.variables[] | select(.key == "<VARIABLE_NAME>")'
terraform import buddy_project_variable.self 12345