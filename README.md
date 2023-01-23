
## Examples

```terraform

resource "spotinstadmin_account" "custdep" {
  name            = "${var.custdep}"
}

resource "spotinstadmin_account_external_id" "custdep" {
  account_id            = "${spotinstadmin_account.custdep.id}"
}

resource "spotinstadmin_account_aws_link" "custdep" {
  account_id      = "${spotinstadmin_account.custdep.id}"
  aws_role_arn    = "${aws_iam_role.spotinst.arn}"
}

resource "spotinstadmin_programmatic_user" "custdep" {
  name        = "${var.custdep}"
  account_id  = "${spotinstadmin_account.custdep.id}"
  description = "Account for ${var.custdep}"
}
```

## Building the Provider
```bash
go get && go build .
```
