Guide: https://medium.com/@krishnan.srm/graphql-with-golang-331de956d956

go run github.com/99designs/gqlgen init


Guide running GraphQL
# mutation removeUser($removeUser:DeleteUser!){
#   removeUser(input: $removeUser){
#     id
#     firstName
#     lastName
#   }
# }
# query {
# 	users{
#      id
#     firstName
#     lastName
#   }
# }
# find a user
# query {
# 	user(id:"2"){
#     id
#     firstName
#     lastName
#   }
# }
#Create a User
# mutation createUser($user:NewUser!){
#   createUser(input:$user){
#     id
#     firstName
#     lastName
#     dob
#   }
# }
----createdUser variable---{
  "user": {
    "userId": "4",
    "firstName": "ngo",
    "lastName": "nguyen",
    "dob": "11/7/1888"
  }
}