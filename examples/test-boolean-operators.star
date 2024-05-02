export(
    Config(
        description = "test boolean operators",
        default = "default",
        overrides = [
            ("isTest == true", "true"),
            ("isTest == false", "false"),
        ],
    ),
)
