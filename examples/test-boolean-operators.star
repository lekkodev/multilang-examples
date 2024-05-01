export(
    Config(
        description = "test boolean operators",
        default = "default",
        overrides = [
            ("is_test == true", "true"),
            ("is_test == false", "false"),
        ],
    ),
)
