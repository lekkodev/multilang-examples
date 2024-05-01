export(
    Config(
        description = "test string operators",
        default = "default",
        overrides = [
            ("env == \"prod\"", "equals"),
            ("env != \"dev\"", "not equals"),
            ("env co \"test\"", "test"),
            ("env sw \"staging\"", "staging"),
            ("env ew \"beta\"", "beta"),
            ("env in [\"special1\",\"special2\"]", "special"),
        ],
    ),
)
