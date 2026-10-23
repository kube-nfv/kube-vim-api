# hack/python

Inputs for the Python client that `make generate-python` builds with openapi-generator (pinned Docker image, see the Makefile) from `generated/openapi/kubevim.openapi.json` into `generated/python`.

- `openapi-generator-ignore`: generated files the client does not ship (CI templates, test stubs; client tests live in `test/python`).
- `templates/model_enum.mustache`: the generator's enum template with a `_missing_` hook, so a value added by a newer server becomes `UNKNOWN_DEFAULT_OPEN_API` instead of failing validation. `enumUnknownDefaultCase` alone only adds that member. The template is copied from the pinned generator version; re-export it (`author template -g python`) when the version changes.
