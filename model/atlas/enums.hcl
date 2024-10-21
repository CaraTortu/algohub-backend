schema "public" {}

enum "user_type" {
  schema = schema.public
  values = [
    "user",
    "staff",
  ]
}

