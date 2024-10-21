variable "db_url" {
    type = string
    default = getenv("DATABASE_URL")
}

data "composite_schema" "app" {
  schema "public" {
    url = "file://model/atlas/enums.hcl"
  }

  schema "public" {
    url = data.external_schema.gorm.url
  }
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./model",
    "--dialect", "postgres",
  ]
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/16/dev?search_path=public"
  url = var.db_url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

